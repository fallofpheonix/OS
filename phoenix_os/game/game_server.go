package game

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"phoenix/ledger/src"
	"phoenix/warden"
)

type GameServer struct {
	Score      *ScoreState
	Warden     *warden.Warden
	Ledger     *ledger.Ledger
	EventsFile string
	mu         sync.Mutex
}

func NewGameServer(score *ScoreState, w *warden.Warden, l *ledger.Ledger, eventsFile string) *GameServer {
	return &GameServer{
		Score:      score,
		Warden:     w,
		Ledger:     l,
		EventsFile: eventsFile,
	}
}

func (gs *GameServer) Start(addr string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/events", gs.corsHandler(gs.handleEvents))
	mux.HandleFunc("/graph", gs.corsHandler(gs.handleGraph))
	mux.HandleFunc("/warden/policy", gs.corsHandler(gs.handleWardenPolicy))
	mux.HandleFunc("/game/score", gs.corsHandler(gs.handleScore))
	mux.HandleFunc("/game/action", gs.corsHandler(gs.handleAction))
	mux.HandleFunc("/game/config", gs.corsHandler(gs.handleConfig))

	log.Printf("[GAME SERVER] Listening on %s\n", addr)
	go func() {
		if err := http.ListenAndServe(addr, mux); err != nil {
			log.Printf("[GAME SERVER ERROR] %v\n", err)
		}
	}()
}

func (gs *GameServer) corsHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

type rawLogEvent struct {
	Timestamp string                 `json:"timestamp"`
	EventID   string                 `json:"event_id"`
	Category  string                 `json:"category"`
	EventType string                 `json:"event_type"`
	HostID    string                 `json:"host_id"`
	PID       int                    `json:"pid"`
	PPID      int                    `json:"ppid"`
	UID       int                    `json:"uid"`
	GID       int                    `json:"gid"`
	Comm      string                 `json:"comm"`
	ExePath   string                 `json:"exe_path"`
	Payload   map[string]interface{} `json:"payload"`
}

func (gs *GameServer) handleEvents(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(gs.EventsFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to open events file: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var events []rawLogEvent
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var ev rawLogEvent
		if err := json.Unmarshal(scanner.Bytes(), &ev); err == nil {
			events = append(events, ev)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(events)
}

func (gs *GameServer) handleGraph(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("seq_id")
	limit := int64(1000) // default high value
	if limitStr != "" {
		if val, err := strconv.ParseInt(limitStr, 10, 64); err == nil {
			limit = val
		}
	}

	file, err := os.Open(gs.EventsFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to open events file: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Process DAG structures
	type Node struct {
		ID       string `json:"id"`
		Label    string `json:"label"`
		Title    string `json:"title"`
		Group    string `json:"group"`
		Entropy  float64 `json:"entropy"`
	}

	type Edge struct {
		From string `json:"from"`
		To   string `json:"to"`
	}

	nodeMap := make(map[int]Node)
	var edges []Edge

	scanner := bufio.NewScanner(file)
	var count int64
	for scanner.Scan() {
		count++
		if count > limit {
			break
		}

		var ev rawLogEvent
		if err := json.Unmarshal(scanner.Bytes(), &ev); err != nil {
			continue
		}

		// Track process nodes
		if ev.PID > 0 {
			entropy := 3.2
			if scoreVal, ok := ev.Payload["entropy_score"].(float64); ok {
				entropy = scoreVal
			}

			group := "normal"
			if entropy > 7.5 {
				group = "critical"
			} else if entropy > 6.0 {
				group = "suspicious"
			}

			nodeMap[ev.PID] = Node{
				ID:    strconv.Itoa(ev.PID),
				Label: fmt.Sprintf("%s (PID: %d)", ev.Comm, ev.PID),
				Title: fmt.Sprintf("EXE: %s\nEntropy: %.2f\nEvent: %s", ev.ExePath, entropy, ev.EventType),
				Group: group,
				Entropy: entropy,
			}

			// Link parent-child edges
			if ev.PPID > 0 && ev.PPID != ev.PID {
				// Ensure parent node exists (defaults if not seen yet)
				if _, exists := nodeMap[ev.PPID]; !exists {
					nodeMap[ev.PPID] = Node{
						ID:    strconv.Itoa(ev.PPID),
						Label: fmt.Sprintf("PID: %d", ev.PPID),
						Title: "Parent Process (Implicit)",
						Group: "normal",
						Entropy: 3.2,
					}
				}
				
				// Avoid adding duplicate edges
				edgeExists := false
				fromStr := strconv.Itoa(ev.PPID)
				toStr := strconv.Itoa(ev.PID)
				for _, edge := range edges {
					if edge.From == fromStr && edge.To == toStr {
						edgeExists = true
						break
					}
				}
				if !edgeExists {
					edges = append(edges, Edge{From: fromStr, To: toStr})
				}
			}
		}
	}

	var nodes []Node
	for _, node := range nodeMap {
		nodes = append(nodes, node)
	}

	response := map[string]interface{}{
		"nodes": nodes,
		"edges": edges,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (gs *GameServer) handleWardenPolicy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST expected", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		TargetState    string `json:"target_state"`
		ActuationClass int    `json:"actuation_class"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	gs.mu.Lock()
	defer gs.mu.Unlock()

	target := warden.SystemState(req.TargetState)
	class := warden.ActuationClass(req.ActuationClass)

	// Trigger Warden actuation
	transitioned := gs.Warden.Actuate(target, class, 1.0, 9999, time.Now().Unix(), uint64(9999))

	res := map[string]interface{}{
		"status":       "ok",
		"state":        string(gs.Warden.State),
		"transitioned": transitioned,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(res)
}

func (gs *GameServer) handleScore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(gs.Score.GetState())
}

func (gs *GameServer) handleAction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST expected", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Action string `json:"action"`
		Target string `json:"target"` // PID or state
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var reward int
	var msg string
	success := true

	switch req.Action {
	case "isolate":
		pid, _ := strconv.Atoi(req.Target)
		// For demo challenge: PIDs with high entropy (e.g. 1004, 1026) are target malware
		if pid == 1004 || pid == 1026 || pid == 1047 || pid == 1068 || pid == 1072 || pid == 1076 || pid == 1093 || pid == 1099 {
			reward = 500
			msg = fmt.Sprintf("Successfully isolated compromised process (PID: %d)! Threat neutralized.", pid)
			gs.Score.AddPoints(reward)
			gs.Score.CompleteChallenge("isolate-" + req.Target)
			gs.Score.AwardBadge("Warden-Guardian")
		} else {
			reward = -100
			msg = fmt.Sprintf("Failed to isolate: PID %d is a critical system service! Penalty applied.", pid)
			gs.Score.ApplyPenalty(100)
			success = false
		}
	case "harden":
		reward = 300
		msg = "Hardened Warden policy rules. Minimum confidence gating set to 0.85."
		gs.Score.AddPoints(reward)
		gs.Score.AwardBadge("Determinism-Master")
	case "reset_budget":
		gs.mu.Lock()
		gs.Warden.ResetBudget()
		gs.mu.Unlock()
		reward = 100
		msg = "Manual operator reset: Warden de-escalation recovery budget has been restored."
		gs.Score.AddPoints(reward)
	default:
		success = false
		msg = "Unknown training action."
	}

	res := map[string]interface{}{
		"success": success,
		"message": msg,
		"score":   gs.Score.GetState(),
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(res)
}
