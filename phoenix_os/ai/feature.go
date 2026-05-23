package ai

import (
	"fmt"
	"time"

	"phoenix/arbiter"
	"phoenix/bus"
	"phoenix/common/serialization"
	"phoenix/ledger/src"
	"phoenix/monitor"
	"phoenix/tcs"
	"phoenix/trace"
	"phoenix/warden"
)

// Feature represents a modular subsystem of PhoenixOS acting under the AI Orchestrator.
type Feature interface {
	Name() string
}

// TraceFeature wraps the L4 Graph Intelligence / Trace storage.
type TraceFeature struct {
	Store *trace.TraceStorage
}

func (tf *TraceFeature) Name() string {
	return "trace"
}

func (tf *TraceFeature) Write(event bus.TelemetryEvent) error {
	return tf.Store.Write(event)
}

// MonitorFeature wraps L3 Telemetry Math and Signal processing.
type MonitorFeature struct {
	Service *monitor.MonitorService
}

func (mf *MonitorFeature) Name() string {
	return "monitor"
}

func (mf *MonitorFeature) Process(event bus.TelemetryEvent) monitor.DriftScore {
	return mf.Service.Process(event)
}

// TCSFeature wraps sliding window and degradation monitors.
type TCSFeature struct {
	Window *tcs.SlidingWindow
	DegMon *tcs.DegradationMonitor
}

func (tf *TCSFeature) Name() string {
	return "tcs"
}

func (tf *TCSFeature) AddAndEvaluate(event bus.TelemetryEvent) float64 {
	if event.SeqID > 0 {
		tf.Window.AddEvent(tcs.TelemetryEvent{
			Timestamp:  time.Unix(event.WallTimeUnix, 0),
			SequenceID: uint64(event.SeqID),
			Payload:    event.Payload,
			JitterNS:   0,
		})
	}
	return tf.Window.Evaluate()
}

func (tf *TCSFeature) EvaluateDegradation(tcsScore float64, evLedger *ledger.Ledger) {
	oldDegraded := tf.DegMon.IsDegraded()
	tf.DegMon.Evaluate(tcsScore)
	if tf.DegMon.IsDegraded() != oldDegraded {
		action := "ENTER_NORMAL_MODE"
		if tf.DegMon.IsDegraded() {
			action = "ENTER_DEGRADED_MODE"
		}
		payload, _ := serialization.CanonicalJSON(map[string]interface{}{
			"action": action,
			"score":  tcsScore,
		})
		evLedger.AddEntry("STATE-TRANSITION", "TCS-VIOLATION", payload)
	}
}

// ArbiterFeature wraps the L5.5 Strategic Policy layer.
type ArbiterFeature struct {
	Arb *arbiter.Arbiter
}

func (af *ArbiterFeature) Name() string {
	return "arbiter"
}

func (af *ArbiterFeature) Evaluate(score monitor.DriftScore, tcsScore float64) (warden.SystemState, warden.ActuationClass, bool) {
	return af.Arb.Evaluate(score, tcsScore)
}

// WardenFeature wraps the L5 Tactical Actuation FSM.
type WardenFeature struct {
	Warden *warden.Warden
}

func (wf *WardenFeature) Name() string {
	return "warden"
}

func (wf *WardenFeature) Actuate(targetState warden.SystemState, class warden.ActuationClass, tcsScore float64, seqID int64, wallTime int64, tick uint64, evLedger *ledger.Ledger) {
	stateBefore := string(wf.Warden.State)
	transitioned := wf.Warden.Actuate(targetState, class, tcsScore, seqID, wallTime, tick)
	if transitioned {
		payload, _ := serialization.CanonicalJSON(map[string]interface{}{
			"state": string(wf.Warden.State),
		})
		evLedger.AddEntryV2(fmt.Sprintf("WARDEN-ACTION-%d", seqID), "POLICY-ACTUATION", payload, stateBefore, string(wf.Warden.State), "1.0.0")
	}
}

// LedgerFeature wraps the cryptographic evidence ledger.
type LedgerFeature struct {
	Ledger *ledger.Ledger
}

func (lf *LedgerFeature) Name() string {
	return "ledger"
}
