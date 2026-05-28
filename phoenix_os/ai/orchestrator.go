package ai

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/fallofpheonix/PheonixDistributed/discovery"
	distLedger "github.com/fallofpheonix/PheonixDistributed/ledger"
	"github.com/fallofpheonix/PheonixGuard"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/monitor"
	"github.com/fallofpheonix/phoenix-os/phoenixmind-core/contracts"
)

type AIOrchestrator struct {
	features map[string]Feature
	Oracle   *NexusBridge
	Wg       sync.WaitGroup
	mu       sync.Mutex
	Advisor  *PredictiveAdvisor
	
	// Distributed Networking & Consensus
	Discovery discovery.PeerDiscovery
	Consensus distLedger.ConsensusLedger
}

func NewAIOrchestrator() *AIOrchestrator {
	advisor, _ := NewPredictiveAdvisor(0.0, 0.045)
	return &AIOrchestrator{
		features: make(map[string]Feature),
		Oracle:   NewNexusBridge("http://127.0.0.1:7860", ""), // G0DM0D3 Live Port
		Wg:       sync.WaitGroup{},
		mu:       sync.Mutex{},
		Advisor:  advisor,
	}
}

func (o *AIOrchestrator) RegisterFeature(f Feature) {
	o.features[f.Name()] = f
}

func (o *AIOrchestrator) GetFeature(name string) Feature {
	return o.features[name]
}

// StartNetworking initiates the P2P discovery loop and consensus layer.
func (o *AIOrchestrator) StartNetworking(ctx context.Context, disc discovery.PeerDiscovery, cons distLedger.ConsensusLedger) error {
	o.Discovery = disc
	o.Consensus = cons

	if err := o.Discovery.Start(ctx); err != nil {
		return fmt.Errorf("failed to start peer discovery: %v", err)
	}

	o.Wg.Add(1)
	go func() {
		defer o.Wg.Done()
		log.Println("[NETWORKING] Pheonix Beacon Discovery Active.")
		for {
			select {
			case <-ctx.Done():
				return
			default:
				// Simplified watch loop for now
				peers := o.Discovery.Peers()
				if len(peers) > 0 {
					log.Printf("[NETWORKING] Cluster Census: %d verified peers.", len(peers))
				}
				time.Sleep(10 * time.Second)
			}
		}
	}()

	return nil
}

// StopNetworking halts all networking activities.
func (o *AIOrchestrator) StopNetworking() {
	if o.Discovery != nil {
		o.Discovery.Stop()
	}
}

const DRY_RUN = false

func (o *AIOrchestrator) ReconcileGraph() {
	o.mu.Lock()
	defer o.mu.Unlock()
	
	gf, ok := o.features["graph"].(*GraphFeature)
	lf, hasLedger := o.features["ledger"].(*LedgerFeature)
	
	if ok && hasLedger {
		gf.SetPaused(true)
		gf.RebuildFromLedger(lf.Ledger)
		gf.SetPaused(false)
		log.Println("[RECONCILIATION] DAG Rebuilt from Ledger.")
	}
}

func (o *AIOrchestrator) OrchestrateTick(event bus.TelemetryEvent, logicalTick uint64) {
	ledgerFeature, hasLedger := o.features["ledger"].(*LedgerFeature)
	
	// 1. Sensory (Monitor)
	var score monitor.DriftScore
	if monitorFeature, ok := o.features["monitor"].(*MonitorFeature); ok {
		score = monitorFeature.Process(event)
		log.Printf("[DEBUG] Event %d ZScore: %f", event.SeqID, score.ZScore)
		if o.Advisor != nil && score.ZScore > o.Advisor.Threshold {
			o.Advisor.LogPrediction(event.EventID, score.ZScore, o.Advisor.CalculateWeight(score.ZScore))
		}
	}

	// 2. Causal (Graph/Trace)
	graphContext := "No causal data available."
	if graphFeature, ok := o.features["graph"].(*GraphFeature); ok {
		graphFeature.AddEvent(event)
		graphContext = fmt.Sprintf("Nodes: %d, Lineage tracking active.", len(graphFeature.Graph.Nodes))
	}
	if traceFeature, ok := o.features["trace"].(*TraceFeature); ok {
		_ = traceFeature.Write(event)
	}

	// TCS Confidence
	tcsScore := 1.0
	if tcsFeature, ok := o.features["tcs"].(*TCSFeature); ok {
		tcsScore = tcsFeature.AddAndEvaluate(event)
		if hasLedger && ledgerFeature != nil {
			tcsFeature.EvaluateDegradation(tcsScore, int64(event.SeqID), ledgerFeature.Ledger)
		}
	}

	// 3. Strategic (Arbiter)
	authorized := false
	var targetState warden.SystemState
	var class warden.ActuationClass
	if arbiterFeature, ok := o.features["arbiter"].(*ArbiterFeature); ok {
		targetState, class, authorized = arbiterFeature.Evaluate(score, tcsScore)
	}

	// 4. Cognitive (Nexus Oracle - G0DM0D3)
	if authorized && o.Oracle != nil {
		o.Wg.Add(1)
		go func() {
			defer o.Wg.Done()
			audit := ""
			if rf, ok := o.features["reality"].(*RealityFeature); ok { audit = rf.ReadAudit() }
			
			directive, err := o.Oracle.RequestDirective(event, tcsScore, audit, graphContext)
			if err != nil {
				log.Printf("[Nexus Oracle Error] %v", err)
				return
			}
			
			log.Printf("[Nexus Oracle Directive] Command: %s, Confidence: %.2f, Reasoning: %s", directive.Command, directive.ConfidenceScore, directive.Reasoning)
			
			if DRY_RUN {
				log.Printf("[SHADOW-MODE] AI proposes: %s for PID %d with confidence %.2f", directive.Command, event.PID, directive.ConfidenceScore)
				return
			}

			// If Oracle agrees, actuate via Warden
			if directive.ConfidenceScore > 0.7 {
				o.mu.Lock()
				if wf, ok := o.features["warden"].(*WardenFeature); ok {
					var evLedger contracts.ILedger
					if hasLedger && ledgerFeature != nil {
						evLedger = ledgerFeature.Ledger
					}
					
					// Map GraphProof to warden structure
					proof := &warden.GraphProof{
						Path:            directive.GraphProof,
						ExpectedNsproxy: uint32(event.Nsproxy),
					}
					
					wf.Actuate(targetState, class, tcsScore, event.SeqID, event.PID, event.Tgid, logicalTick, evLedger, proof)
				}
				o.mu.Unlock()
			}

		}()
	}
}
