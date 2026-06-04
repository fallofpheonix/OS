/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/* =========================================================================
 * WORKFLOW POSITION: TEST — ENTROPY AND KL DIVERGENCE VERIFICATION
 *
 * This test file verifies the correctness of the entropy and KL divergence
 * calculations used by the Monitor for anomaly detection.
 *
 * TEST CASES:
 *   1. ZeroData: Empty input → zero entropy
 *   2. LowEntropy: Uniform data → zero entropy
 *   3. HighEntropy: Unique bytes → 8.0 entropy
 *   4. MixedEntropy: Real text → entropy between 0 and 8
 *   5. BufferUnderflow: Short input → small positive entropy
 *   6. KLDivergenceWithBaseline: Uniform vs skewed distributions
 *   7. AnomalyThresholdBoundary: Tests anomaly detection thresholds
 *   8. SchemaValidation: JSON serialization correctness
 *
 * INVARIANT: The entropy calculation must be deterministic and correct.
 * Any deviation would cause false positives/negatives in threat detection.
 * ========================================================================= */
package monitor

import (
	"encoding/json"
	"math"
	"testing"
)

func TestCalculate(t *testing.T) {
	t.Run("ZeroData", func(t *testing.T) {
		res := Calculate([]byte{}, nil)
		if res.Entropy != 0 || res.KLDivergence != 0 || res.IsAnomaly {
			t.Errorf("Expected zeroed result for empty input, got %+v", res)
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

	t.Run("MixedEntropy", func(t *testing.T) {
		data := []byte("The quick brown fox jumps over the lazy dog. 1234567890!")
		res := Calculate(data, nil)
		if res.Entropy <= 0.0 || res.Entropy >= 8.0 {
			t.Errorf("Expected entropy between 0 and 8, got %f", res.Entropy)
		}
	})

	t.Run("BufferUnderflow", func(t *testing.T) {
		// Short buffer: testing that short input doesn't crash and has some expected range
		data := []byte("hello")
		res := Calculate(data, nil)
		if res.Entropy <= 0.0 || res.Entropy > 3.0 {
			t.Errorf("Expected small positive entropy for hello, got %f", res.Entropy)
		}
	})

	t.Run("KLDivergenceWithBaseline", func(t *testing.T) {
		// Setup baseline as uniform distribution
		baseline := make([]float64, 256)
		for i := 0; i < 256; i++ {
			baseline[i] = 1.0 / 256.0
		}

		// Uniform data should have KL divergence near 0
		data := make([]byte, 256)
		for i := 0; i < 256; i++ {
			data[i] = byte(i)
		}
		res := Calculate(data, baseline)
		if res.KLDivergence > 0.001 {
			t.Errorf("Expected KL Divergence near 0 for uniform data matching baseline, got %f", res.KLDivergence)
		}

		// Skewed data should have higher KL divergence
		skewedData := make([]byte, 256)
		for i := range skewedData {
			skewedData[i] = 0 // all zeros
		}
		resSkewed := Calculate(skewedData, baseline)
		if resSkewed.KLDivergence < 1.0 {
			t.Errorf("Expected high KL Divergence for skewed data, got %f", resSkewed.KLDivergence)
		}
	})

	t.Run("AnomalyThresholdBoundary", func(t *testing.T) {
		// Anomaly triggers: entropy > 7.5 OR klDiv > 4.0

		// 1. High entropy triggers anomaly
		highEntropyData := make([]byte, 256)
		for i := 0; i < 256; i++ {
			highEntropyData[i] = byte(i)
		}
		res1 := Calculate(highEntropyData, nil)
		if !res1.IsAnomaly {
			t.Errorf("Expected high entropy (8.0 > 7.5) to be flagged as anomaly")
		}

		// 2. KL Divergence triggers anomaly
		// We want a case where entropy is low but KL divergence is high.
		// If baseline is uniform (1/256), and data is all 0, KL divergence is:
		// P(0) = 1.0, P(i) = 0 for i > 0.
		// D_KL = P(0) * log2(P(0)/Q(0)) = 1.0 * log2(1.0 / (1/256)) = log2(256) = 8.0
		// 8.0 > 4.0, so it should flag an anomaly.
		baseline := make([]float64, 256)
		for i := 0; i < 256; i++ {
			baseline[i] = 1.0 / 256.0
		}
		zeroData := make([]byte, 100)
		res2 := Calculate(zeroData, baseline)
		if !res2.IsAnomaly {
			t.Errorf("Expected high KL divergence (8.0 > 4.0) to be flagged as anomaly, got KL %f, entropy %f", res2.KLDivergence, res2.Entropy)
		}
	})

	t.Run("SchemaValidation", func(t *testing.T) {
		res := Result{Entropy: 4.5, KLDivergence: 1.2, IsAnomaly: false}
		data, err := json.Marshal(res)
		if err != nil {
			t.Fatalf("Failed to marshal Result to JSON: %v", err)
		}

		var parsed map[string]interface{}
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatalf("Failed to unmarshal JSON: %v", err)
		}

		// Verify keys exist and have correct types
		if _, ok := parsed["entropy"].(float64); !ok {
			t.Errorf("entropy missing or not a float64 in JSON: %s", string(data))
		}
		if _, ok := parsed["kl_divergence"].(float64); !ok {
			t.Errorf("kl_divergence missing or not a float64 in JSON: %s", string(data))
		}
		if _, ok := parsed["is_anomaly"].(bool); !ok {
			t.Errorf("is_anomaly missing or not a bool in JSON: %s", string(data))
		}
	})
}

func BenchmarkCalculate4KB(b *testing.B) {
	data := make([]byte, 4096)
	for i := range data {
		data[i] = byte(i % 256)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Calculate(data, nil)
	}
}

func BenchmarkCalculate64KB(b *testing.B) {
	data := make([]byte, 65536)
	for i := range data {
		data[i] = byte(i % 256)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Calculate(data, nil)
	}
}
