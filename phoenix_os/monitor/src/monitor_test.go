package main

import "testing"

func TestEntropy(t *testing.T) {
	data := make([]byte, 256)
	for i := 0; i < 256; i++ {
		data[i] = byte(i)
	}
	res := CalculateEntropy(data)
	if res.Entropy < 7.9 {
		t.Errorf("Expected max entropy, got %f", res.Entropy)
	}
}

func TestKalman(t *testing.T) {
	f := NewKalmanFilter(0.1, 0.1, 1.0, 0.0)
	v := f.Filter(10.0)
	if v == 0 || v == 10.0 {
		t.Errorf("Filter did not smooth signal: %f", v)
	}
}
