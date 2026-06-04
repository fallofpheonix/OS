// Package math provides specialized arithmetic for the Phoenix substrate.
// Core Domain Logic: Implements fixed-point arithmetic with saturating semantics to ensure
// deterministic and safe calculations for system resources and state transitions.
package math

import (
	"fmt"
)

// FixedPoint represents a decimal number using a fixed-point integer representation.
// Internal State: int64 underlying value representing (real_value * divisor).
// API Scope: Public utility for resource accounting.
// Concurrency: Thread-safe as an immutable value type.
type FixedPoint int64

// LABEL: [PURE] [PUBLIC_API] [STABLE]
// SaturatingAdd adds two FixedPoint numbers, saturating at MaxInt64 or MinInt64 on overflow/underflow.
// I/O: None.
// Complexity: O(1).
func (a FixedPoint) SaturatingAdd(b FixedPoint) FixedPoint {
	if a > 0 && b > (1<<63-1)-a {
		return FixedPoint(1<<63 - 1)
	}
	if a < 0 && b < (-1<<63)-a {
		return FixedPoint(-1 << 63)
	}
	return a + b
}

// LABEL: [PURE] [PUBLIC_API] [STABLE]
// SaturatingSub subtracts two FixedPoint numbers, saturating at Min/MaxInt64 on overflow/underflow.
// I/O: None.
// Complexity: O(1).
func (a FixedPoint) SaturatingSub(b FixedPoint) FixedPoint {
	if b > 0 && a < (-1<<63)+b {
		return FixedPoint(-1 << 63)
	}
	if b < 0 && a > (1<<63-1)+b {
		return FixedPoint(1<<63 - 1)
	}
	return a - b
}

// LABEL: [PURE] [PUBLIC_API] [STABLE]
// NewFixedPoint creates a new FixedPoint value from an integer and a divisor.
// I/O: None.
// Complexity: O(1).
func NewFixedPoint(value int64, divisor int64) (FixedPoint, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("divisor cannot be zero")
	}
	return FixedPoint(value * divisor), nil
}
