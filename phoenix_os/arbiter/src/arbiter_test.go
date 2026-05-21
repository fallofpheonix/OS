package main

import "testing"

func TestMiniMax(t *testing.T) {
	matrix := PayoffMatrix{
		{{10, -10}, {0, 0}},
		{{0, 0}, {5, -5}},
	}
	action := SolveMiniMax(matrix)
	if action != 1 {
		t.Errorf("Expected action 1, got %d", action)
	}
}
