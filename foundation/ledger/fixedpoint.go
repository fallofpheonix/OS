/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * 
 * FILE: fixedpoint.go
 * PATH: Phoenix.Nucleus/ledger/fixedpoint.go
 *
 * PURPOSE:
 * Provides a deterministic, architecture-independent fixed-point number 
 * representation. This eliminates IEEE-754 float64 rounding discrepancies 
 * to ensure 100% deterministic state reconstruction across the mesh.
 */

package ledger

import "fmt"

// FixedPointDivisor is set to 10^9 to provide billionth-level precision.
const FixedPointDivisor uint64 = 1_000_000_000

// FixedPoint replaces float64 for all Simulation, Confidence, and Drift scores.
type FixedPoint struct {
	Value int64 `json:"value"` // Backed by int64 to allow negative metrics (e.g. antagonistic weights, negative drift)
}

// NewFixedPoint creates a FixedPoint from an integer numerator.
func NewFixedPoint(numerator int64) FixedPoint {
	return FixedPoint{
		Value: numerator,
	}
}

// FromFloat64 converts a float64 to FixedPoint (useful for bridging external nondeterministic inputs).
func FromFloat64(f float64) FixedPoint {
	return FixedPoint{
		Value: int64(f * float64(FixedPointDivisor)),
	}
}

// Float64 converts the fixed point back to a float64 for display or non-consensus logging.
func (f FixedPoint) Float64() float64 {
	return float64(f.Value) / float64(FixedPointDivisor)
}

// String provides a human-readable representation.
func (f FixedPoint) String() string {
	return fmt.Sprintf("%f", f.Float64())
}
