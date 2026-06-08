// Package serialization provides utilities for deterministic data representation.
// Core Domain Logic: Implements canonical serialization (Stable JSON) to ensure that logically equivalent
// data structures always produce identical byte streams, a prerequisite for distributed consensus and hashing.
package serialization

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
)

// LABEL: [PURE] [PUBLIC_API] [STABLE]
// StableMarshal produces a deterministic byte representation of the input.
// It handles recursive map sorting, deterministic slice ordering, and
// explicitly distinguishes between nil and empty slices.
// I/O: None.
// Side Effects: None.
// Complexity: O(N log N) where N is the total number of elements/nodes in the structure, driven by recursive sorting.
func StableMarshal(v interface{}) ([]byte, error) {
	// First, perform a standard JSON marshal.
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	// Then, recursively process and re-marshal to ensure determinism.
	var raw interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return data, nil
	}

	return marshalRecursive(raw)
}

// CanonicalJSON is an alias for StableMarshal, used for deterministic JSON serialization.
func CanonicalJSON(v interface{}) ([]byte, error) {
	return StableMarshal(v)
}

// LABEL: [PURE] [INTERNAL_ONLY] [STABLE]
// marshalRecursive handles the dispatch of serialization logic based on type.
// I/O: None.
// Complexity: O(1) + complexity of the specific type handler.
func marshalRecursive(v interface{}) ([]byte, error) {
	switch val := v.(type) {
	case map[string]interface{}:
		return marshalSortedMap(val)
	case []interface{}:
		return marshalSortedSlice(val)
	default:
		return json.Marshal(val)
	}
}

// LABEL: [PURE] [INTERNAL_ONLY] [STABLE]
// marshalSortedMap serializes a map with keys in lexicographical order.
// I/O: None.
// Complexity: O(K log K) where K is the number of keys in the map.
func marshalSortedMap(m map[string]interface{}) ([]byte, error) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, k := range keys {
		buf.WriteString(fmt.Sprintf("%q:", k))
		valBytes, err := marshalRecursive(m[k])
		if err != nil {
			return nil, err
		}
		buf.Write(valBytes)
		if i < len(keys)-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

// LABEL: [PURE] [INTERNAL_ONLY] [STABLE]
// marshalSortedSlice serializes a slice by sorting its elements by their canonical byte representation.
// I/O: None.
// Complexity: O(L log L) where L is the number of elements in the slice, due to byte-comparison sorting.
func marshalSortedSlice(s []interface{}) ([]byte, error) {
	// If nil, explicitly mark as null or handled distinctly
	if s == nil {
		return []byte("null"), nil
	}

	// Create a slice of serialized elements to sort by their canonical representation
	serializedElements := make([][]byte, len(s))
	for i, element := range s {
		b, err := marshalRecursive(element)
		if err != nil {
			return nil, err
		}
		serializedElements[i] = b
	}

	// Sort based on the serialized byte representation
	sort.Slice(serializedElements, func(i, j int) bool {
		return bytes.Compare(serializedElements[i], serializedElements[j]) < 0
	})

	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, b := range serializedElements {
		buf.Write(b)
		if i < len(serializedElements)-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteByte(']')
	return buf.Bytes(), nil
}
