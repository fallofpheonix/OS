package ai

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/fallofpheonix/phoenix-os/phoenix_os/arbiter"
	"github.com/fallofpheonix/PheonixGuard"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/monitor"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/tcs"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/trace"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/common/serialization"
	ledger "github.com/fallofpheonix/PheonixTruth/src"
	"github.com/fallofpheonix/phoenix-os/phoenixmind-core/contracts"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/telemetry/process_graphs"
)

type Feature interface {
	Name() string
}

type GraphFeature struct {
	Graph  *process_graphs.Graph
	mu     sync.RWMutex
	Paused bool
}

func (gf *GraphFeature) Name() string { return "graph" }

func (gf *GraphFeature) SetPaused(paused bool) {
	gf.mu.Lock()
	defer gf.mu.Unlock()
	gf.Paused = paused
}

func (gf *GraphFeature) RebuildFromLedger(l contracts.ILedger) {
	gf.mu.Lock()
	defer gf.mu.Unlock()

	// Clear existing graph
	gf.Graph = process_graphs.NewGraph()

	if local, ok := l.(*ledger.Ledger); ok {
		entries := local.SortedEntries()
		for _, entry := range entries {
			nodeID := fmt.Sprintf("evt-%s", entry.EventID)
			gf.Graph.AddNode(nodeID, process_graphs.Process, int64(entry.LogicalTick))
			if entry.CauseID != "" {
				parentID := fmt.Sprintf("evt-%s", entry.CauseID)
				gf.Graph.AddEdge(parentID, nodeID)
			}
		}
	}
}

func (gf *GraphFeature) AddEvent(event bus.TelemetryEvent) {
	gf.mu.RLock()
	if gf.Paused {
		gf.mu.RUnlock()
		return
	}
	gf.mu.RUnlock()

	nodeID := fmt.Sprintf("evt-%d", event.SeqID)
	gf.Graph.AddNode(nodeID, process_graphs.Process, int64(event.LamportClock))
	if event.CausalID != "" {
		gf.Graph.AddEdge(event.CausalID, nodeID)
	}
}

func (gf *GraphFeature) VerifyPath(path []string) (bool, error) {
	gf.Graph.Mu.RLock()
	defer gf.Graph.Mu.RUnlock()

	for i := 0; i < len(path)-1; i++ {
		from := path[i]
		to := path[i+1]

		// Check if nodes exist
		if _, ok := gf.Graph.Nodes[from]; !ok {
			return false, fmt.Errorf("node %s not found", from)
		}
		if _, ok := gf.Graph.Nodes[to]; !ok {
			return false, fmt.Errorf("node %s not found", to)
		}

		// Check if edge exists
		found := false
		for _, neighbor := range gf.Graph.Edges[from] {
			if neighbor == to {
				found = true
				break
			}
		}
		if !found {
			return false, fmt.Errorf("edge %s -> %s not found", from, to)
		}
	}
	return true, nil
}

type TraceFeature struct {
	Store *trace.TraceStorage
}

func (tf *TraceFeature) Name() string { return "trace" }
func (tf *TraceFeature) Write(event bus.TelemetryEvent) error { return tf.Store.Write(event) }

type MonitorFeature struct {
	Service *monitor.MonitorService
}

func (mf *MonitorFeature) Name() string { return "monitor" }
func (mf *MonitorFeature) Process(event bus.TelemetryEvent) monitor.DriftScore { return mf.Service.Process(event) }

type TCSFeature struct {
	Window *tcs.SlidingWindow
	DegMon *tcs.DegradationMonitor
}

func (tf *TCSFeature) Name() string { return "tcs" }
func (tf *TCSFeature) AddAndEvaluate(event bus.TelemetryEvent) float64 {
	tf.Window.AddEvent(tcs.TelemetryEvent{
		Timestamp:  time.Unix(int64(event.LamportClock), 0),
		SequenceID: uint64(event.SeqID),
		Payload:    event.Payload,
	})
	return tf.Window.Evaluate()
}

func (tf *TCSFeature) EvaluateDegradation(tcsScore float64, seqID int64, evLedger contracts.ILedger) {
	oldDegraded := tf.DegMon.IsDegraded()
	tf.DegMon.Evaluate(tcsScore)
	if tf.DegMon.IsDegraded() != oldDegraded {
		action := "NORMAL"
		if tf.DegMon.IsDegraded() { action = "DEGRADED" }
		payload, _ := serialization.CanonicalJSON(map[string]interface{}{"action": action, "score": tcsScore})
		evLedger.AddEntry(fmt.Sprintf("TCS-ANOMALY-%d", seqID), "TCS-VIOLATION", payload)
	}
}

type ArbiterFeature struct {
	Arb *arbiter.Arbiter
}

func (af *ArbiterFeature) Name() string { return "arbiter" }
func (af *ArbiterFeature) Evaluate(score monitor.DriftScore, tcsScore float64) (warden.SystemState, warden.ActuationClass, bool) {
	return af.Arb.Evaluate(score, tcsScore)
}

type WardenFeature struct {
	Warden *warden.Warden
}

func (wf *WardenFeature) Name() string { return "warden" }
func (wf *WardenFeature) Actuate(targetState warden.SystemState, class warden.ActuationClass, tcsScore float64, seqID int64, targetPID int, targetTgid uint32, lamportClock uint64, evLedger contracts.ILedger, proof *warden.GraphProof) {
	stateBefore := string(wf.Warden.State)
	eventID := fmt.Sprintf("TCS-ANOMALY-%d", seqID)
	cert, _ := evLedger.GenerateCertificate(eventID, tcsScore)
	req := warden.AuthorityEscalationRequest{
		EventID:        eventID,
		TargetPID:      targetPID,
		TargetTgid:     int(targetTgid),
		TargetNsproxy:  proof.ExpectedNsproxy,
		TargetState:    targetState,
		ActuationClass: class,
		EvidenceWeight: tcsScore,
		Certificate:    cert,
		GraphProof:     proof,
	}
	if transitioned := wf.Warden.Actuate(req, int(seqID), lamportClock); transitioned {
		payload, _ := serialization.CanonicalJSON(map[string]interface{}{"state": string(wf.Warden.State)})
		evLedger.AddEntryV2(fmt.Sprintf("WARDEN-ACTION-%d", seqID), "POLICY-ACTUATION", payload, "", stateBefore, string(wf.Warden.State), "1.0.0")
	}
}


type RealityFeature struct {
	AuditPath string
}

func (rf *RealityFeature) Name() string { return "reality" }
func (rf *RealityFeature) ReadAudit() string {
	content, err := os.ReadFile(rf.AuditPath)
	if err != nil { return "Audit report missing" }
	return string(content)
}

type LedgerFeature struct {
	Ledger contracts.ILedger
}

func (lf *LedgerFeature) Name() string { return "ledger" }
