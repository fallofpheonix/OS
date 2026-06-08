package bus

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/fallofpheonix/phoenix/foundation/ledger/src"
)

// PURPOSE: Verifies the full vertical slice from Bus to Ledger.
// CONTRACT: A TelemetryEvent published to the Bus must appear in the Ledger
//           after being processed by the Applier.
// FAILURE: Fails if the event is lost, reordered, or corrupted.

func TestVerticalSlice_BusToLedger(t *testing.T) {
	// 1. Setup
	eventBus := NewBus()
	substrateLedger := ledger.NewLedger(nil) // No allocator needed for test

	applier := NewApplier(ApplierConfig{
		BufferSize: 100,
		Ledger:     substrateLedger,
		Bus:        eventBus,
		Topics:     []string{"TEST_TOPIC"},
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := applier.Start(ctx); err != nil {
		t.Fatalf("Failed to start applier: %v", err)
	}

	// 2. Act: Publish an event
	testPayload := map[string]string{"key": "value"}
	payloadBytes, _ := json.Marshal(testPayload)

	event := TelemetryEvent{
		EventID:     "EVT-001",
		CausalID:    "ROOT",
		LogicalTick: 123,
		Payload:     json.RawMessage(payloadBytes),
	}

	eventBus.Publish("TEST_TOPIC", event)

	// 3. Assert: Wait for ingestion
	// WHY: Channels are async. We need a small poll loop to check the ledger.
	deadline := time.After(1 * time.Second)
	found := false
	for !found {
		select {
		case <-deadline:
			t.Fatal("Timeout waiting for event to reach ledger")
		default:
			entries := substrateLedger.SortedEntries()
			for _, entry := range entries {
				if entry.EventID == "EVT-001" {
					if entry.LogicalTick != 123 {
						t.Errorf("Expected tick 123, got %d", entry.LogicalTick)
					}
					// Severity and Raw payload are wrapped by event_adapter.ToLedgerParams
					expectedWrapped := fmt.Sprintf(`{"severity":{"v":0},"raw":%s}`, string(payloadBytes))
					if string(entry.Payload) != expectedWrapped {
						t.Errorf("Payload corruption: expected %s, got %s", expectedWrapped, string(entry.Payload))
					}
					found = true
					break
				}
			}
			if !found {
				time.Sleep(10 * time.Millisecond)
			}
		}
	}

	t.Log("Vertical Slice Success: Event traversed Bus -> Applier -> Ledger")
}
