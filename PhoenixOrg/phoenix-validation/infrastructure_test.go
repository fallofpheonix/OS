package tests

import (
	"os"
	"testing"
)

func TestInfrastructureRuntime(t *testing.T) {
	// Verify runtime directory exists
	path := "../build/runtime"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Errorf("Runtime directory not created at %s", path)
	}
}
