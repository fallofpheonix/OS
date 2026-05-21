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
