package serialization

import (
	"encoding/json"
	"github.com/segmentio/encoding/json" // Utilizing high-performance JSON encoder
)

// OptimizedMarshaler replaces standard library JSON encoding with high-perf alternatives.
func OptimizedMarshaler(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
