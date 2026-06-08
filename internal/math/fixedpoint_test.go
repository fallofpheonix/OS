package phxmath

import (
	"math"
	"testing"
)

func TestFixedPoint_Determinism(t *testing.T) {
	// 1. NewFixedPoint (Scaling check)
	f1 := NewFixedPoint(100)
	if f1.V != 100000000 {
		t.Errorf("Scaling mismatch: expected 100000000, got %d", f1.V)
	}

	// 2. Addition
	f2 := NewFixedPoint(50)
	sum := f1.SaturatingAdd(f2)
	if sum.V != 150000000 {
		t.Errorf("Addition mismatch: expected 150000000, got %d", sum.V)
	}

	// 3. Multiplication (A * B) / Precision
	// 100.0 * 0.5 = 50.0
	f3 := NewFixedPointRaw(500000) // 0.5
	prod := f1.Mul(f3)
	if prod.V != 50000000 {
		t.Errorf("Multiplication mismatch: expected 50000000, got %d", prod.V)
	}

	// 4. Division (A * Precision) / B
	// 100.0 / 2.0 = 50.0
	f4 := NewFixedPoint(2)
	div := f1.Div(f4)
	if div.V != 50000000 {
		t.Errorf("Division mismatch: expected 50000000, got %d", div.V)
	}
}

func TestFixedPoint_Saturation(t *testing.T) {
	max := FixedPoint{V: math.MaxInt64}
	one := NewFixedPoint(1)

	// Overflow
	sum := max.SaturatingAdd(one)
	if sum.V != math.MaxInt64 {
		t.Error("SaturatingAdd failed to saturate at MaxInt64")
	}

	min := FixedPoint{V: math.MinInt64}
	// Underflow
	diff := min.SaturatingSub(one)
	if diff.V != math.MinInt64 {
		t.Error("SaturatingSub failed to saturate at MinInt64")
	}
}

func TestFixedPoint_Rounding(t *testing.T) {
	// 1 / 3 = 0.333333 (Truncated)
	f1 := NewFixedPoint(1)
	f3 := NewFixedPoint(3)
	res := f1.Div(f3)
	if res.V != 333333 {
		t.Errorf("Rounding mismatch: expected 333333, got %d", res.V)
	}
}

func TestFixedPointPrecisionContract(t *testing.T) {
	// 1.5 * 1.5 = 2.25
	a := NewFixedPointRaw(1500000) // 1.5
	b := NewFixedPointRaw(1500000) // 1.5
	got := a.Mul(b)
	want := int64(2250000) // 2.25
	if got.V != want {
		t.Errorf("Mul(1.5, 1.5) = %v; want %v", got.V, want)
	}

	// 1.234567 * 1.0 = 1.234567 (no change)
	c := NewFixedPointRaw(1234567)
	d := NewFixedPoint(1)
	if c.Mul(d).V != 1234567 {
		t.Errorf("Mul(1.234567, 1.0) = %v; want 1234567", c.Mul(d).V)
	}

	// Verify Truncation: 1.999999 / 2.0 = 0.9999995 -> 0.999999 (truncate toward zero)
	e := NewFixedPointRaw(1999999)
	f := NewFixedPoint(2)
	gotDiv := e.Div(f)
	wantDiv := int64(999999)
	if gotDiv.V != wantDiv {
		t.Errorf("Div(1.999999, 2.0) = %v; want %v (truncation check)", gotDiv.V, wantDiv)
	}

	// Precision Loss Contract: Div(Mul(x, y), y) != x
	// (1.5 * 0.000001) / 0.000001
	x := NewFixedPointRaw(1500000) // 1.5
	y := NewFixedPointRaw(1)       // 0.000001
	intermediate := x.Mul(y)       // (1500000 * 1) / 1000000 = 1 (Truncated from 1.5)
	if intermediate.V != 1 {
		t.Errorf("Intermediate Mul(1.5, 0.000001) = %v; want 1", intermediate.V)
	}
	final := intermediate.Div(y) // (1 * 1000000) / 1 = 1000000 (1.0)
	if final.V == x.V {
		t.Errorf("Precision Contract Violated: Div(Mul(x, y), y) should not equal x when precision is lost")
	}
}
