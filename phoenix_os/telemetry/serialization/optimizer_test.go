package serialization

import (
	"testing"
)

type TestStruct struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

func TestOptimizer(t *testing.T) {
	data := TestStruct{ID: 1, Value: "test"}
	encoded, err := OptimizedMarshaler(data)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}
	if string(encoded) == "" {
		t.Error("Marshaled data is empty")
	}
}
