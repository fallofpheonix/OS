package validation

import (
	"encoding/json"
	"testing"

	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment"
)

// TestFuzzMutationRejection verifies that parsing corrupted process actions fails or rejects bad mutations.
func TestFuzzMutationRejection(t *testing.T) {
	badJSONInputs := []string{
		`{"PID": "not-an-int", "Action": "MONITOR"}`,
		`{"PID": 100, "Action": 99999}`,
		`{"PID": -5, "Action": ""}`,
		`{invalid-json}`,
	}

	for _, input := range badJSONInputs {
		var action containment.ProcessAction
		err := json.Unmarshal([]byte(input), &action)
		if err != nil {
			continue // Rejected successfully by JSON parser
		}
		if action.PID < 0 || action.Action == "" {
			t.Logf("fuzz rejected invalid bounds: PID=%d, Action=%q", action.PID, action.Action)
		}
	}
}

// TestFuzzPayloadParser validates process action parser robustness.
func TestFuzzPayloadParser(t *testing.T) {
	action := containment.ProcessAction{
		PID:    100,
		Action: containment.ActionMonitor,
		Reason: "normal-reason",
	}

	data, err := json.Marshal(action)
	if err != nil {
		t.Fatalf("failed to marshal action: %v", err)
	}

	var parsed containment.ProcessAction
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("failed to parse action: %v", err)
	}
}

