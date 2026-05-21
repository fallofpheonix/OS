package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"phoenix/ai/agents"
	"phoenix/kernel/sandbox"
	"phoenix/ledger/src"
	"phoenix/security/control"
	"phoenix/security/game/stackelberg"
	"phoenix/security/physics"
	"phoenix/telemetry/detector"
	"phoenix/telemetry/events"
	"phoenix/telemetry/process_lineage"
)

type ReplayResult struct {
	TotalEvents  int     `json:"total_events"`
	GraphSize    int     `json:"graph_size"`
	Duration     string  `json:"duration"`
	GraphHash    string  `json:"graph_hash"`
	TP           int     `json:"tp"`
	FP           int     `json:"fp"`
	TN           int     `json:"tn"`
	FN           int     `json:"fn"`
	Precision    float64 `json:"precision"`
	Recall       float64 `json:"recall"`
	AvgReaction  string  `json:"avg_reaction"`
	Transitions  int     `json:"transitions"`
	MaxDisorder  float64 `json:"max_disorder"`
	FinalTemp    float64 `json:"final_temp"`
	GameUtility  float64 `json:"game_utility"`
	OptimalMoves int     `json:"optimal_moves"`
	ExplainCount int     `json:"explain_count"`
	KernelErrors int     `json:"kernel_errors"`
	MTTD         string  `json:"mttd"`
}

func main() {
	inputFile := flag.String("input", "../../../09_telemetry/replay_events_large.jsonl", "Path to JSONL events file")
	outputFile := flag.String("output", "replay_result.json", "Path to output result file")
	flag.Parse()

	log.Printf("Starting replay of %s", *inputFile)

	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Failed to open input: %v", err)
	}
	defer file.Close()

	graph := lineage.NewLineageGraph()
	det := detector.NewDetector()
	ctrl := control.NewController()
	gameMatrix := stackelberg.NewDefaultMatrix()
	explainer := agents.NewExplainer()
	ks := sandbox.NewKernelSimulator()
	
	reader := bufio.NewReader(file)
	count := 0
	start := time.Now()

	var tp, fp, tn, fn int
	var totalReaction time.Duration
	var transitions int
	var maxDisorder float64
	var totalUtility float64
	var optimalMoves int
	var explainCount int
	var kernelErrors int
	
	var firstAttackTime time.Time
	var firstDetectionTime time.Time
	
	categoryCounts := make(map[string]float64)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error reading line: %v", err)
		}

		var evt events.Event
		if err := json.Unmarshal(line, &evt); err != nil {
			log.Printf("Skip invalid line: %v", err)
			continue
		}

		// Ground Truth
		isGTThreat := false
		if entropy, ok := evt.Payload["entropy_score"].(float64); ok {
			if entropy > 7.0 {
				isGTThreat = true
				if firstAttackTime.IsZero() {
					firstAttackTime = evt.Timestamp
				}
			}
		}

		// Update Graph
		switch evt.EventType {
		case "execve", "fork":
			graph.AddProcess(evt.PID, evt.PPID, evt.Comm, evt.ExePath, evt.Timestamp)
		case "exit":
			graph.ExitProcess(evt.PID, evt.Timestamp)
		}

		// Update Physics Distributions
		categoryCounts[evt.Category]++

		// Detect
		res := det.Analyze(evt, graph)
		if res.IsThreat && isGTThreat && firstDetectionTime.IsZero() {
			firstDetectionTime = evt.Timestamp
		}
		
		// Kernel Simulation (Stage 10)
		if err := ks.UpdateMap(evt.EventID, res.ImportanceScore); err != nil {
			kernelErrors++
			ks.CurrentEntries = 0
		}

		// Game Theory (Stage 8)
		move, utility := stackelberg.Solve(gameMatrix, res.ImportanceScore)
		totalUtility += utility
		if (move == stackelberg.MoveDefend && res.IsThreat) || (move == stackelberg.MoveMonitor && !res.IsThreat) {
			optimalMoves++
		}

		// Physics
		physState := physics.ComputeState(categoryCounts, res.ImportanceScore)
		if physState.Disorder > maxDisorder {
			maxDisorder = physState.Disorder
		}

		// Control
		prevState := ctrl.CurrentState
		newState, action := ctrl.UpdateState(res.ImportanceScore)
		if newState != prevState {
			transitions++
			totalReaction += ctrl.ReactionTime
			if action != "NONE" {
				evidence := ledger.Evidence{
					TraceHash:  evt.EventID,
					SDI:        physState.Disorder,
					Action:     action,
					Confidence: res.ImportanceScore,
				}
				_, _ = explainer.Explain(evidence)
				explainCount++
			}
		}

		if res.IsThreat && isGTThreat {
			tp++
		} else if res.IsThreat && !isGTThreat {
			fp++
		} else if !res.IsThreat && isGTThreat {
			fn++
		} else {
			tn++
		}

		count++
	}

	elapsed := time.Since(start)
	hash := calculateGraphHash(graph)

	precision := 0.0
	if (tp + fp) > 0 {
		precision = float64(tp) / float64(tp+fp)
	}
	recall := 0.0
	if (tp + fn) > 0 {
		recall = float64(tp) / float64(tp+fn)
	}

	avgReaction := time.Duration(0)
	if transitions > 0 {
		avgReaction = totalReaction / time.Duration(transitions)
	}

	mttd := time.Duration(0)
	if !firstDetectionTime.IsZero() && !firstAttackTime.IsZero() {
		mttd = firstDetectionTime.Sub(firstAttackTime)
	}

	result := ReplayResult{
		TotalEvents: count,
		GraphSize:   graph.Size(),
		Duration:    elapsed.String(),
		GraphHash:   hash,
		TP:          tp,
		FP:          fp,
		TN:          tn,
		FN:          fn,
		Precision:   precision,
		Recall:      recall,
		AvgReaction: avgReaction.String(),
		Transitions: transitions,
		MaxDisorder: maxDisorder,
		FinalTemp:   physics.ComputeState(categoryCounts, 0).Temperature,
		GameUtility: totalUtility / float64(count),
		OptimalMoves: optimalMoves,
		ExplainCount: explainCount,
		KernelErrors: kernelErrors,
		MTTD:         mttd.String(),
	}

	resData, _ := json.MarshalIndent(result, "", "  ")
	_ = os.WriteFile(*outputFile, resData, 0644)

	fmt.Printf("Replay complete. MTTD: %v, Precision: %.2f%%\n", mttd, precision*100)
}

func calculateGraphHash(g *lineage.LineageGraph) string {
	data, _ := json.Marshal(g.Nodes)
	h := sha256.New()
	h.Write(data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
