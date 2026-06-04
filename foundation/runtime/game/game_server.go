/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: CYCLE 12 — HTTP API / OPERATOR INTERFACE (Layer 6)
//
// The GameServer is the HTTP API for operator interaction and training.
// It provides endpoints for:
//   - Viewing telemetry events (/events)
//   - Viewing the causal DAG (/graph)
//   - Triggering Warden state transitions (/warden/policy)
//   - Managing game score and challenges (/game/score, /game/action)
//   - Loading ecosystem configuration (/game/config)
//
// WORKFLOW: Operator → HTTP request → GameServer handler → Warden/Bus/Ledger
//   → Response with current state
//
// SECURITY: This server runs with CORS: * (all origins allowed).
// The /warden/policy endpoint directly triggers Warden state transitions
// WITHOUT authentication. Any HTTP client can change the system state.
//
// TRUST BOUNDARY: This is the UNTRUSTED ZONE — any network client can
// access all endpoints. The server should be behind an auth proxy.
// =========================================================================
package game

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	securityv1 "github.com/fallofpheonix/phoenix/foundation/contracts/security/v1"
	ledger "github.com/fallofpheonix/phoenix/foundation/ledger/src"
)

type GameServer struct {
	Score      *ScoreState
	Warden     securityv1.Actuator
	Ledger     *ledger.Ledger
	EventsFile string
	mu         sync.Mutex
}

func NewGameServer(score *ScoreState, w securityv1.Actuator, l *ledger.Ledger, eventsFile string) *GameServer {
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
		server := &http.Server{
			Addr:              addr,
			Handler:           mux,
			ReadHeaderTimeout: 3 * time.Second,
		}
		if err := server.ListenAndServe(); err != nil {
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
	defer func() { _ = file.Close() }()

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
	defer func() { _ = file.Close() }()

	// Process DAG structures
	type Node struct {
		ID      string  `json:"id"`
		Label   string  `json:"label"`
		Title   string  `json:"title"`
		Group   string  `json:"group"`
		Entropy float64 `json:"entropy"`
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
				ID:      strconv.Itoa(ev.PID),
				Label:   fmt.Sprintf("%s (PID: %d)", ev.Comm, ev.PID),
				Title:   fmt.Sprintf("EXE: %s\nEntropy: %.2f\nEvent: %s", ev.ExePath, entropy, ev.EventType),
				Group:   group,
				Entropy: entropy,
			}

			// Link parent-child edges
			if ev.PPID > 0 && ev.PPID != ev.PID {
				// Ensure parent node exists (defaults if not seen yet)
				if _, exists := nodeMap[ev.PPID]; !exists {
					nodeMap[ev.PPID] = Node{
						ID:      strconv.Itoa(ev.PPID),
						Label:   fmt.Sprintf("PID: %d", ev.PPID),
						Title:   "Parent Process (Implicit)",
						Group:   "normal",
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

type simpleContainment struct {
	target string
	level  securityv1.ContainmentLevel
	reason string
}

func (c *simpleContainment) Target() string                      { return c.target }
func (c *simpleContainment) Level() securityv1.ContainmentLevel  { return c.level }
func (c *simpleContainment) Reason() string                      { return c.reason }

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

	var level securityv1.ContainmentLevel
	switch req.TargetState {
	case "SAFE":
		level = securityv1.LevelNone
	case "WATCH":
		level = securityv1.LevelMonitor
	case "SUSPICIOUS":
		level = securityv1.LevelSandbox
	case "CRITICAL":
		level = securityv1.LevelIsolate
	case "COMPROMISED":
		level = securityv1.LevelQuench
	default:
		http.Error(w, fmt.Sprintf("invalid target state: %s", req.TargetState), http.StatusBadRequest)
		return
	}

	action := &simpleContainment{
		target: "system",
		level:  level,
		reason: "Operator request",
	}

	var ctx context.Context = r.Context()
	err := gs.Warden.Actuate(ctx, action)
	transitioned := (err == nil)

	currentLevel, _ := gs.Warden.GetCurrentLevel()
	var stateStr string
	switch currentLevel {
	case securityv1.LevelNone:
		stateStr = "SAFE"
	case securityv1.LevelMonitor:
		stateStr = "WATCH"
	case securityv1.LevelSandbox:
		stateStr = "SUSPICIOUS"
	case securityv1.LevelIsolate:
		stateStr = "CRITICAL"
	case securityv1.LevelQuench:
		stateStr = "COMPROMISED"
	default:
		stateStr = "UNKNOWN"
	}

	res := map[string]interface{}{
		"status":       "ok",
		"state":        stateStr,
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
