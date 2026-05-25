package serialization

import (
	"encoding/json"
	"fmt"
	"sort"
)

// CanonicalJSON marshals an interface into a byte slice with sorted keys.
func CanonicalJSON(v interface{}) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return data, nil // Return as is if not a map
	}

	return marshalSortedMap(m)
}

func marshalSortedMap(m map[string]interface{}) ([]byte, error) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	result := "{"
	for i, k := range keys {
		val, _ := json.Marshal(m[k])
		result += fmt.Sprintf("%q:%s", k, string(val))
		if i < len(keys)-1 {
			result += ","
		}
	}
	result += "}"
	return []byte(result), nil
}
