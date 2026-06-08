package security

import (
	ledger "github.com/fallofpheonix/phoenix/foundation/ledger/src"
	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
	"testing"
)

func TestWarden_TransitionMatrix(t *testing.T) {
	// Initialize Warden with minimal substrate
	w := NewWarden(&bus.Bus{}, &ledger.Ledger{})

	type testCase struct {
		name     string
		from     SystemState
		to       SystemState
		hasProof bool
		want     bool
	}

	cases := []testCase{
		// 1. ESCALATION (Fast-Up)
		{"SafeToWatch", StateSafe, StateWatch, false, true},
		{"SafeToCritical", StateSafe, StateCritical, false, true},
		{"SafeToCompromised", StateSafe, StateCompromised, false, true},
		{"WatchToCritical", StateWatch, StateCritical, false, true},

		// 2. DE-ESCALATION (Slow-Down - 1 step only)
		{"CriticalToSuspicious", StateCritical, StateSuspicious, false, true},
		{"SuspiciousToWatch", StateSuspicious, StateWatch, false, true},
		{"WatchToSafe", StateWatch, StateSafe, false, true},

		// 3. INVALID DE-ESCALATION (Ladder Violation)
		{"CriticalToWatch", StateCritical, StateWatch, false, false},
		{"CriticalToSafe", StateCritical, StateSafe, false, false},
		{"SuspiciousToSafe", StateSuspicious, StateSafe, false, false},

		// 4. COMPROMISED EXIT (ROOT-005)
		{"CompromisedToCritical_NoProof", StateCompromised, StateCritical, false, false},
		{"CompromisedToCritical_WithProof", StateCompromised, StateCritical, true, true},
		{"CompromisedToSafe_WithProof", StateCompromised, StateSafe, true, false},     // Ladder still enforced
		{"CompromisedSelf_NoProof", StateCompromised, StateCompromised, false, false}, // Fix for Q128
		{"CompromisedSelf_WithProof", StateCompromised, StateCompromised, true, true},

		// 5. ZERO-VALUE TRAP
		{"InvalidToSafe", StateInvalid, StateSafe, false, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w.state.Store(uint32(tc.from))
			snap := PostureSnapshot{State: tc.from}
			err := w.validateTransition(snap, tc.to, tc.hasProof)
			got := (err == nil)
			if got != tc.want {
				t.Errorf("%v -> %v (Proof:%v): got %v, want %v. Err: %v",
					tc.from, tc.to, tc.hasProof, got, tc.want, err)
			}
		})
	}
}

func TestWarden_Lockout(t *testing.T) {
	w := NewWarden(&bus.Bus{}, &ledger.Ledger{})
	w.Lock()

	snap := PostureSnapshot{State: StateSafe}
	err := w.validateTransition(snap, StateWatch, false)
	if err == nil {
		t.Error("Expected failure for transition on locked Warden")
	}
}
