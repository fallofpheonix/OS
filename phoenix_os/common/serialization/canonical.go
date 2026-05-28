package serialization

import (
	"encoding/json"
)

// CanonicalJSON returns a deterministic JSON representation.
func CanonicalJSON(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
