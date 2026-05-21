package main

import (
	"fmt"
)

// Payoff represents (Defender, Attacker) utilities
type Payoff struct {
	D, A float64
}

// PayoffMatrix for a 2x2 game
type PayoffMatrix [][]Payoff

// Node represents a participant in the swarm
type Node struct {
	ID         string
	Reputation float64
	Authorized bool
}

// Arbiter handles quorum and reputation
type Arbiter struct {
	Nodes           []Node
	QuorumThreshold float64
}

// CalculateQuorum validates quorum based on reputation
func (a *Arbiter) CalculateQuorum(votes map[string]bool) bool {
	var totalReputation float64
	var positiveReputation float64
	for _, node := range a.Nodes {
		if !node.Authorized {
			continue
		}
		totalReputation += node.Reputation
		if vote, ok := votes[node.ID]; ok && vote {
			positiveReputation += node.Reputation
		}
	}
	if totalReputation == 0 {
		return false
	}
	return (positiveReputation / totalReputation) >= a.QuorumThreshold
}

func SolveMiniMax(m PayoffMatrix) int {
	// Simple pure strategy minimax for demonstration
	bestDefenderAction := 0
	maxMinVal := -1e9

	for i := range m {
		minAttackerVal := 1e9
		for j := range m[i] {
			if m[i][j].A < minAttackerVal {
				minAttackerVal = m[i][j].A
			}
		}
		if minAttackerVal > maxMinVal {
			maxMinVal = minAttackerVal
			bestDefenderAction = i
		}
	}
	return bestDefenderAction
}

func main() {
	// Boot check
	matrix := PayoffMatrix{
		{{1, -1}, {-1, 1}},
		{{-1, 1}, {1, -1}},
	}
	fmt.Printf("Best Action: %d\n", SolveMiniMax(matrix))
}
