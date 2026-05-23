package entropy_engine

import (
	"fmt"
	"time"
)

func BenchmarkMain() {
	// Construct a 4096-byte buffer with pseudo-random content
	data := make([]byte, 4096)
	for i := 0; i < len(data); i++ {
		data[i] = byte((i*37 + 17) % 256)
	}

	N := 20000
	start := time.Now()
	for i := 0; i < N; i++ {
		_ = ShannonEntropy(data)
	}
	elapsed := time.Since(start)
	per := elapsed.Seconds() * 1e6 / float64(N)
	fmt.Printf("Go entropy bench: events=%d total_s=%.6f per_event_us=%.3f\n", N, elapsed.Seconds(), per)
}
