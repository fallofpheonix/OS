/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package stackelberg

// Move represents a strategic choice by an agent.
type Move string

const (
	MoveAttack  Move = "ATTACK"
	MoveWait    Move = "WAIT"
	MoveDefend  Move = "DEFEND"
	MoveMonitor Move = "MONITOR"
)

// Payoff represents the utility for an agent.
type Payoff struct {
	Leader   float64 // Defender (Leader in Stackelberg for PhoenixOS)
	Follower float64 // Attacker
}

// Matrix represents the payoff matrix for the game.
type Matrix map[Move]map[Move]Payoff

// NewDefaultMatrix creates the standard Stackelberg matrix for PhoenixOS.
func NewDefaultMatrix() Matrix {
	m := make(Matrix)
	m[MoveDefend] = make(map[Move]Payoff)
	m[MoveMonitor] = make(map[Move]Payoff)

	// Defender Defends, Attacker Attacks: High cost for attacker, high protection for defender
	m[MoveDefend][MoveAttack] = Payoff{Leader: 5.0, Follower: -10.0}

	// Defender Defends, Attacker Waits: Waste of resources for defender
	m[MoveDefend][MoveWait] = Payoff{Leader: -2.0, Follower: 0.0}

	// Defender Monitors, Attacker Attacks: High cost for defender, high reward for attacker
	m[MoveMonitor][MoveAttack] = Payoff{Leader: -10.0, Follower: 10.0}

	// Defender Monitors, Attacker Waits: Efficient for defender
	m[MoveMonitor][MoveWait] = Payoff{Leader: 2.0, Follower: 0.0}

	return m
}

// Solve finds the optimal leader move given the follower's likely rational response.
func Solve(m Matrix, attackerProb float64) (move Move, utility float64) {
	// Calculate expected utility for each leader move

	uDefend := attackerProb*m[MoveDefend][MoveAttack].Leader + (1-attackerProb)*m[MoveDefend][MoveWait].Leader

	uMonitor := attackerProb*m[MoveMonitor][MoveAttack].Leader + (1-attackerProb)*m[MoveMonitor][MoveWait].Leader

	if uDefend > uMonitor {
		return MoveDefend, uDefend
	}
	return MoveMonitor, uMonitor
}
