/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: INTEGRATED - Formal Cognition Authoritative
 */

// Package reasoning implements the strategic inference and explanation layers for PhoenixOS.
package reasoning

import (
	"fmt"
	"strings"
	"time"
)

// ReasonNode represents a single step in the authority chain with comparative metrics.
type ReasonNode struct {
	Authority string  `json:"authority"` // Authority domain (e.g., Sensory, Causal)
	Signal    string  `json:"signal"`    // Specific signal identifier
	Observed  float64 `json:"observed"`  // Measured telemetry value
	Expected  float64 `json:"expected"`  // Predicted baseline value
	Deviation float64 `json:"deviation"` // (Observed - Expected) / Expected
	Weight    float64 `json:"weight"`    // Authority weight in final decision
}

// ReasonPath is a sequence of authority nodes that led to a specific system decision.
type ReasonPath []ReasonNode

// String returns a human-readable representation of the deep explanation path.
// Formats metrics as percentages and provides clear step-by-step lineage.
func (rp ReasonPath) String() string {
	var paths []string
	for _, node := range rp {
		paths = append(paths, fmt.Sprintf("[%s: %s | Obs: %.2f, Exp: %.2f, Dev: +%.1f%%]",
			node.Authority, node.Signal, node.Observed, node.Expected, node.Deviation*100))
	}
	return strings.Join(paths, " -> ")
}

// Explainer generates the formal justification for a system action (Mandate W1).
// It maintains the decision context, reason path, and alternative scenarios.
type Explainer struct {
	DecisionID  string
	Timestamp   int64
	Path        ReasonPath
	Action      string
	Risk        string
	Alternative string
	Evidence    map[string]interface{}
}

// NewExplainer initializes a new explanation context for a specific decision.
func NewExplainer(decisionID string) *Explainer {
	return &Explainer{
		DecisionID: decisionID,
		Timestamp:  time.Now().Unix(),
		Path:       make(ReasonPath, 0),
		Evidence:   make(map[string]interface{}),
	}
}

// AddLink appends a new authority step to the path with deep metrics.
// Calculates deviation percentage automatically. Handles division-by-zero by capping at 100.0 (10,000%).
// Inputs: authority, signal (strings), observed, expected, weight (float64).
// Complexity: O(1) time / O(1) space.
func (e *Explainer) AddLink(authority, signal string, observed, expected, weight float64) {
	deviation := 0.0
	if expected != 0 {
		deviation = (observed - expected) / expected
	} else if observed != 0 {
		deviation = 100.0
	}
	e.Path = append(e.Path, ReasonNode{
		Authority: authority,
		Signal:    signal,
		Observed:  observed,
		Expected:  expected,
		Deviation: deviation,
		Weight:    weight,
	})
}

// GenerateCounterfactual produces a mathematical proof of alternative outcomes.
// It recalculates the total signal by substituting a sensor value, proving why an action was triggered.
// Inputs: missingSensor (string), replacementValue (float64).
// Returns: string representation of the counterfactual scenario.
// Complexity: O(P) where P is the number of nodes in the ReasonPath.
func (e *Explainer) GenerateCounterfactual(missingSensor string, replacementValue float64) string {
	totalSignal := 0.0
	var adjustedSignal float64

	for _, node := range e.Path {
		if node.Authority == missingSensor {
			adjustedSignal = (replacementValue * node.Weight)
		} else {
			totalSignal += (node.Observed * node.Weight)
		}
	}

	final := totalSignal + adjustedSignal
	return fmt.Sprintf("If sensor '%s' were replaced with value %.2f, total signal would be %.2f (Decision change: %v)",
		missingSensor, replacementValue, final, final < 0.8)
}

// ExplainIgnored generates a negative justification log for a non-action.
// It explains why an event that crossed a baseline did not result in an actuation.
// Returns: Formatted string log.
func (e *Explainer) ExplainIgnored(reason string) string {
	return fmt.Sprintf("[NEGATIVE JUSTIFICATION] Decision: IGNORE | Target: %s | Reason: %s | Path: %s",
		e.DecisionID, reason, e.Path.String())
}
