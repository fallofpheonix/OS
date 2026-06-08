/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: INTEGRATED - Formal Cognition Authoritative
 */
package knowledge

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"

	"github.com/fallofpheonix/phoenix/foundation/ledger"
)

func TestBeliefEngine_Commit(t *testing.T) {
	be := NewBeliefEngine()

	b1 := &Belief{
		ID:         "B1",
		Statement:  "Initial Belief",
		Confidence: 0.5,
		State:      BeliefActive,
	}

	be.Commit(b1)

	latest, ok := be.GetLatest("B1")
	if !ok || latest.Version != 1 {
		t.Errorf("expected version 1, got %d", latest.Version)
	}

	// Update belief
	b2 := &Belief{
		ID:         "B1",
		Version:    2,
		Statement:  "Updated Belief",
		Confidence: 0.9,
	}
	be.Commit(b2)

	latest, ok = be.GetLatest("B1")
	if !ok || latest.Version != 2 || latest.Confidence != 0.9 {
		t.Errorf("expected version 2 with high confidence, got version %d, conf %f", latest.Version, latest.Confidence)
	}

	// Verify version 1 is superseded
	versions := be.Beliefs["B1"]
	if versions[0].State != BeliefSuperseded {
		t.Errorf("expected version 1 to be superseded, got %s", versions[0].State)
	}
}

func TestBeliefEngine_Concurrency(t *testing.T) {
	be := NewBeliefEngine()
	const iterations = 100
	const workers = 10

	var wg sync.WaitGroup
	wg.Add(workers * 2)

	for i := 0; i < workers; i++ {
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				be.Commit(&Belief{
					ID:        fmt.Sprintf("B-%d", workerID),
					Statement: "Concurrent update",
				})
			}
		}(i)

		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				_, _ = be.GetLatest(fmt.Sprintf("B-%d", workerID))
			}
		}(i)
	}
	wg.Wait()
}

func TestBeliefEngine_Reconstruction(t *testing.T) {
	be := NewBeliefEngine()

	// Simulate belief event.
	beliefPayload := ledger.BeliefPayload{
		BeliefID:   "BELIEF_001",
		FactIDs:    []string{"FACT_A", "FACT_B"},
		Confidence: 0.92,
		Statement:  "Fact A caused Fact B with high confidence",
	}
	payload, _ := json.Marshal(beliefPayload)

	events := []*ledger.Event{
		{
			Type:    ledger.EventBelief,
			Payload: payload,
		},
	}

	// Simulate an ignored event type
	events = append([]*ledger.Event{{Type: ledger.EventFact}}, events...)

	if err := be.ReconstructFromLedger(events); err != nil {
		t.Fatalf("reconstruction failed: %v", err)
	}

	// Simulate bad payload
	badEvents := []*ledger.Event{{Type: ledger.EventBelief, Payload: []byte("{bad-json")}}
	if err := be.ReconstructFromLedger(badEvents); err == nil {
		t.Errorf("expected error on bad json payload")
	}

	versions, ok := be.Beliefs["BELIEF_001"]
	if !ok || len(versions) == 0 {
		t.Errorf("belief BELIEF_001 not found after reconstruction")
	}

	belief := versions[len(versions)-1]
	if belief.Confidence != 0.92 {
		t.Errorf("confidence mismatch: expected 0.92, got %v", belief.Confidence)
	}

	if belief.State != BeliefActive {
		t.Errorf("state mismatch: expected ACTIVE, got %s", belief.State)
	}
}
