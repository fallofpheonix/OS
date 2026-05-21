package main

import (
    "bytes"
    "testing"
)

func BenchmarkShannonEntropy_Random512(b *testing.B) {
    data := make([]byte, 512)
    for i := 0; i < 512; i++ {
        data[i] = byte(i % 256)
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = ShannonEntropy(data)
    }
}

func TestShannonEntropy_Known(t *testing.T) {
    data := bytes.Repeat([]byte{0x00}, 1024)
    h := ShannonEntropy(data)
    if h != 0 {
        t.Fatalf("expected entropy 0, got %v", h)
    }
}
