package phxmath

import (
	"math/big"
	"testing"
)

func BenchmarkFixedPoint_Mul_Native(b *testing.B) {
	f1 := FixedPoint{V: 4000000}
	f2 := FixedPoint{V: 2000000}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FixedPoint{V: (f1.V * f2.V) / Precision}
	}
}

func BenchmarkFixedPoint_Mul_BigInt(b *testing.B) {
	f1 := FixedPoint{V: 4000000000000000000}
	f2 := FixedPoint{V: 4000000000000000000}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f1.Mul(f2)
	}
}

func BenchmarkFixedPoint_Div_BigInt(b *testing.B) {
	f1 := FixedPoint{V: 4000000000000000000}
	f2 := FixedPoint{V: 2000000}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f1.Div(f2)
	}
}

func BenchmarkFixedPoint_Mul_BigInt_Alloc(b *testing.B) {
	f1 := big.NewInt(4000000000000000000)
	f2 := big.NewInt(4000000000000000000)
	p := big.NewInt(Precision)
	res := new(big.Int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res.Mul(f1, f2)
		res.Quo(res, p)
	}
}
