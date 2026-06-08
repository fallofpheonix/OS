package engine

import (
	"math/rand"
	"testing"

	"github.com/fallofpheonix/phoenix/internal/contracts"
)

func FuzzWorldState_ApplyEvent(f *testing.F) {
	f.Add([]byte("initial_state"))
	f.Fuzz(func(t *testing.T, data []byte) {
		ws := NewWorldState(0)

		r := rand.New(rand.NewSource(int64(len(data))))

		ev := contracts.Event{
			Version: 1,
			Type:    contracts.EventType(r.Intn(4)),
			Payload: data,
		}

		applied := contracts.AppliedEvent{
			Height: uint64(r.Intn(1000)),
			Epoch:  uint64(r.Intn(100)),
			Event:  ev,
		}

		_ = ws.ApplyEvent(applied, make(map[string]VerificationRule))
	})
}
