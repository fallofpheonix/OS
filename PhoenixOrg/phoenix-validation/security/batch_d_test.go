package security

import (
	"testing"
	"time"

	"github.com/fallofpheonix/phoenix-control/warden"
	"github.com/fallofpheonix/phoenix-contracts"
	"github.com/fallofpheonix/phoenix-control/arbiter"
	"github.com/fallofpheonix/phoenix-logic/monitor"
)

// Mock components for security validation
func setupEnforcement() (*warden.Warden, *arbiter.Arbiter) {
	w := warden.NewWarden(nil)
	arb := arbiter.NewArbiter(nil)
	
	// Ensure UID 1234 is highly critical
	arb.Policy.NodeCriticality[1234] = 1.0
	
	return w, arb
}

func simulateAttack(t *testing.T, name string, sdi float64) {
	t.Run(name, func(t *testing.T) {
		w, arb := setupEnforcement()
		
		// 1. Initial State
		if w.State != contracts.StateSafe {
			t.Errorf("expected SAFE initial state, got %v", w.State)
		}

		// 2. Telemetry Ingestion (Simulated via DriftScore)
		score := monitor.DriftScore{
			UID:       1234,
			Severity:  sdi,
			Frequency: 0.1, // Rare event = high importance
			ZScore:    sdi * 25, // High drift
		}

		// 3. Arbiter Evaluation
		target, class, authorized := arb.Evaluate(score, 1.0)
		
		// 4. Warden Actuation
		w.Actuate(target, class, 1.0, 1, time.Now().Unix(), 1)

		// 5. Verification
		if sdi > 0.8 && w.State == contracts.StateSafe {
			t.Errorf("Warden failed to transition for critical attack %s (SDI: %.2f, Auth: %v, State: %v)", name, sdi, authorized, w.State)
		}
		
		t.Logf("[BLUE] Attack %s processed. Authorized: %v, Warden State: %v", name, authorized, w.State)
	})
}

func TestAttackSimulations_BatchD(t *testing.T) {
	attacks := []struct {
		name string
		sdi  float64
	}{
		{"Fork Bomb", 0.95},
		{"Reverse Shell", 0.85},
		{"Beacon", 0.75},
		{"Port Scan", 0.65},
		{"File Exfiltration", 0.88},
		{"Ransomware Simulation", 0.99},
		{"SSH Abuse", 0.70},
		{"Persistence", 0.78},
		{"CPU Exhaustion", 0.92},
		{"Memory Exhaustion", 0.91},
		{"Timeline Divergence", 0.82},
		{"Hash Tampering", 0.96},
	}

	for _, a := range attacks {
		simulateAttack(t, a.name, a.sdi)
	}
}
