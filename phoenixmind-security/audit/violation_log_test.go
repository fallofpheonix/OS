package audit

import (
	"os"
	"testing"
)

func TestViolationLog(t *testing.T) {
	logFile := "test_audit.jsonl"
	defer os.Remove(logFile)

	v := Violation{
		ID:       "SEC-TEST-001",
		Actor:    "test_actor",
		Action:   "test_action",
		Severity: "HIGH",
		Result:   "BLOCKED",
	}

	err := WriteJSONL(logFile, v)
	if err != nil {
		t.Fatalf("Failed to write to audit log: %v", err)
	}

	// In a real scenario, you'd read the file back and verify contents
	// For this test, we just check that the file was created.
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		t.Errorf("Audit log file was not created")
	}
}
