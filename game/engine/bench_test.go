package engine

import (
	"fmt"
	"testing"

	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
	"github.com/fallofpheonix/phoenix/internal/contracts"
)

func BenchmarkWorldState_CalculateHash(b *testing.B) {
	ws := NewWorldState(1234)
	ws.Tick = 100
	ws.Epoch = 1

	// Add entities
	for i := 0; i < 100; i++ {
		InjectEntity(ws, &Entity{
			ID:     fmt.Sprintf("agent_%d", i),
			Pos:    phxmath.NewFixedPoint(int64(i)),
			Status: "ACTIVE",
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ws.CalculateHash()
	}
}

func BenchmarkWorldState_ApplyEvent(b *testing.B) {
	ws := NewWorldState(1234)

	events := make([]contracts.Event, b.N)
	for i := 0; i < b.N; i++ {
		ev := contracts.Event{
			Version: 1,
			Type:    contracts.EventMove,
		}
		// ev.Payload = ...
		events[i] = ev
	}

	rules := make(map[string]VerificationRule)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		applied := contracts.AppliedEvent{Height: uint64(i), Epoch: 0, Event: events[i]}
		_ = ws.ApplyEvent(applied, rules)
	}
}
