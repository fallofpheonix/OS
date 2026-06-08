/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package decision_runtime

import "github.com/fallofpheonix/phoenixmind-validator/truth/evidence"

// Decision represents an action to be taken based on the truth state.
type Decision struct {
	Action     string // e.g., "CONTAIN", "ESCALATE", "LOG"
	EntityID   string
	TruthState evidence.TruthState
	Confidence float64
	Reason     string
}

// MakeDecision evaluates a truth state and makes a corresponding decision.
func MakeDecision(entityID string, currentState evidence.TruthState, confidence float64) Decision {
	switch currentState {
	case evidence.BLOCKED, evidence.REJECTED:
		return Decision{
			Action:     "CONTAIN",
			EntityID:   entityID,
			TruthState: currentState,
			Confidence: confidence,
			Reason:     "Critical threat detected, requires immediate containment.",
		}
	case evidence.ESCALATED:
		return Decision{
			Action:     "ESCALATE",
			EntityID:   entityID,
			TruthState: currentState,
			Confidence: confidence,
			Reason:     "Threat requires human intervention or higher-level review.",
		}
	case evidence.WARNING:
		return Decision{
			Action:     "LOG_AND_MONITOR",
			EntityID:   entityID,
			TruthState: currentState,
			Confidence: confidence,
			Reason:     "Potential issue detected, monitoring advised.",
		}
	case evidence.VALIDATED, evidence.OBSERVED:
		return Decision{
			Action:     "ALLOW",
			EntityID:   entityID,
			TruthState: currentState,
			Confidence: confidence,
			Reason:     "No immediate threats, entity is behaving as expected.",
		}
	default:
		return Decision{
			Action:     "INVESTIGATE",
			EntityID:   entityID,
			TruthState: currentState,
			Confidence: confidence,
			Reason:     "Unknown state, requires further investigation.",
		}
	}
}
