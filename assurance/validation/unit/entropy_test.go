/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package unit

import (
	"bytes"
	"testing"

	"github.com/fallofpheonix/phoenix/foundation/runtime/telemetry/entropy_engine_go"
)

func BenchmarkShannonEntropy_Random512(b *testing.B) {
	data := make([]byte, 512)
	for i := 0; i < 512; i++ {
		data[i] = byte(i % 256)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = entropy_engine.ShannonEntropy(data)
	}
}

func TestShannonEntropy_Known(t *testing.T) {
	data := bytes.Repeat([]byte{0x00}, 1024)
	h := entropy_engine.ShannonEntropy(data)
	if h != 0 {
		t.Fatalf("expected entropy 0, got %v", h)
	}
}
