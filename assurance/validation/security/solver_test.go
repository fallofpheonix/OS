/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package security

import (
	"github.com/fallofpheonix/phoenix/foundation/runtime/security/game/stackelberg"
	"testing"
)

func TestSolver(t *testing.T) {
	m := stackelberg.NewDefaultMatrix()

	// High probability of attack -> Should Defend
	move1, _ := stackelberg.Solve(m, 0.9)
	if move1 != stackelberg.MoveDefend {
		t.Errorf("Expected DEFEND for high attack prob, got %s", move1)
	}

	// Low probability of attack -> Should Monitor
	move2, _ := stackelberg.Solve(m, 0.1)
	if move2 != stackelberg.MoveMonitor {
		t.Errorf("Expected MONITOR for low attack prob, got %s", move2)
	}
}
