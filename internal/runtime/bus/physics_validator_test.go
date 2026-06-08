package bus

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/fallofpheonix/phoenix/foundation/ledger/src"
)

func TestPhysicsValidatorRejectsInvalidEvents(t *testing.T) {
	// 1. Setup
	eventBus := NewBus()
	substrateLedger := ledger.NewLedger(nil)
	physicsValidator := ledger.NewPhysicsValidator()

	applier := NewApplier(ApplierConfig{
		BufferSize: 100,
		Ledger:     substrateLedger,
		Bus:        eventBus,
		Validator:  physicsValidator,
		Topics:     []string{"TEST_TOPIC"},
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	applier.Start(ctx)

	// 2. Test Empty EventID (should be rejected)
	eventEmpty := TelemetryEvent{
		EventID: "",
		Payload: json.RawMessage(`{"coherence": 500000}`),
	}
	eventBus.Publish("TEST_TOPIC", eventEmpty)
	time.Sleep(50 * time.Millisecond)
	if len(substrateLedger.SortedEntries()) != 0 {
		t.Error("Expected event with empty EventID to be rejected, but it was ledgered")
	}

	// 3. Test Coherence Out of Range (should be rejected)
	eventBadRange := TelemetryEvent{
		EventID:     "EVT-BAD",
		LogicalTick: 1,
		Payload:     json.RawMessage(`{"coherence": 2000000}`), // > 1,000,000
	}
	eventBus.Publish("TEST_TOPIC", eventBadRange)
	time.Sleep(50 * time.Millisecond)
	if len(substrateLedger.SortedEntries()) != 0 {
		t.Error("Expected event with out-of-range coherence to be rejected, but it was ledgered")
	}

	// 4. Test Valid Event (should be accepted)
	eventValid := TelemetryEvent{
		EventID:     "EVT-VALID",
		LogicalTick: 2,
		Payload:     json.RawMessage(`{"coherence": 800000, "entropy": 100000}`),
	}
	eventBus.Publish("TEST_TOPIC", eventValid)

	// Wait for ingestion
	deadline := time.After(500 * time.Millisecond)
	found := false
	for !found {
		select {
		case <-deadline:
			t.Fatal("Timeout waiting for valid event to reach ledger")
		default:
			for _, entry := range substrateLedger.SortedEntries() {
				if entry.EventID == "EVT-VALID" {
					found = true
					break
				}
			}
			if !found {
				time.Sleep(10 * time.Millisecond)
			}
		}
	}
}
