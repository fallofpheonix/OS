package main

import (
	"math"
	"testing"
)

func TestCalculate(t *testing.T) {
	t.Run("ZeroData", func(t *testing.T) {
		res := Calculate([]byte{}, nil)
		if res.Entropy != 0 {
			t.Errorf("Expected 0 entropy, got %f", res.Entropy)
		}
	})

	t.Run("LowEntropy", func(t *testing.T) {
		data := make([]byte, 100)
		for i := range data {
			data[i] = 'A'
		}
		res := Calculate(data, nil)
		if res.Entropy != 0 {
			t.Errorf("Expected 0 entropy for uniform data, got %f", res.Entropy)
		}
	})

	t.Run("HighEntropy", func(t *testing.T) {
		data := make([]byte, 256)
		for i := 0; i < 256; i++ {
			data[i] = byte(i)
		}
		res := Calculate(data, nil)
		if math.Abs(res.Entropy-8.0) > 0.0001 {
			t.Errorf("Expected 8.0 entropy for unique byte distribution, got %f", res.Entropy)
		}
	})
}

func BenchmarkCalculate(b *testing.B) {
	data := make([]byte, 4096)
	for i := range data {
		data[i] = byte(i % 256)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Calculate(data, nil)
	}
}
