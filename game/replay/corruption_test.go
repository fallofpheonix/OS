package replay

import (
	"crypto/ed25519"
	"encoding/json"
	"testing"

	"github.com/fallofpheonix/phoenix/game/engine"
	"github.com/fallofpheonix/phoenix/internal/consensus"
	"github.com/fallofpheonix/phoenix/internal/contracts"
)

func TestReplay_CorruptionDetection(t *testing.T) {
	pub, priv, _ := ed25519.GenerateKey(nil)

	setup := func() (*ReplayEngine, []contracts.Event) {
		ws := engine.NewWorldState(0)
		ws.Validators = [][]byte{pub}
		r := NewReplayEngine(ws)
		r.AddAuthorizedValidator(pub)
		events := generateTestEvents(priv, 5)
		return r, events
	}

	t.Run("UnauthorizedValidator", func(t *testing.T) {
		r, _ := setup()
		_, otherPriv, _ := ed25519.GenerateKey(nil)

		e := contracts.Event{Version: 1, Type: 0, Payload: []byte("DATA")}
		payload, _ := json.Marshal(e)
		env := &contracts.SignedEnvelope{
			Type:      0,
			Payload:   payload,
			Sequence:  1,
			Validator: contracts.NodeID(otherPriv.Public().(ed25519.PublicKey)),
		}
		consensus.SignEnvelope(env, otherPriv)

		err := r.ProcessEnvelope(env)
		if err == nil {
			t.Error("Expected error for unauthorized validator, got nil")
		}
	})
}
