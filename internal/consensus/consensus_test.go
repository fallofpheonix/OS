package consensus

import (
	"crypto/ed25519"
	"testing"

	"github.com/fallofpheonix/phoenix/internal/contracts"
)

func TestConsensus_CheckQuorum(t *testing.T) {
	// CONSENSUS-014 & 015: Quorum verification and uniqueness.

	// Setup 4 validators (MinValidatorsForBFT)
	privs := make([]ed25519.PrivateKey, 4)
	pubs := make([][]byte, 4)
	for i := 0; i < 4; i++ {
		pub, priv, _ := ed25519.GenerateKey(nil)
		privs[i] = priv
		pubs[i] = pub
	}

	digest := Hash{1, 2, 3}

	// 1. Successful Quorum (3 out of 4 is 2f+1 where f=1)
	var sigs []contracts.SignatureEntry
	for i := 0; i < 3; i++ {
		sig := ed25519.Sign(privs[i], digest[:])
		sigs = append(sigs, contracts.SignatureEntry{
			ValidatorID: contracts.NodeID(pubs[i]),
			Signature:   sig,
		})
	}

	ok, err := CheckQuorum(pubs, sigs, digest)
	if err != nil {
		t.Fatalf("CheckQuorum failed: %v", err)
	}
	if !ok {
		t.Error("Expected successful quorum for 3/4 validators")
	}

	// 2. Insufficient Quorum (2 out of 4)
	sigs = sigs[:2]
	ok, err = CheckQuorum(pubs, sigs, digest)
	if ok {
		t.Error("Expected failed quorum for 2/4 validators")
	}

	// 3. Unauthorized validator signature
	_, unauthPriv, _ := ed25519.GenerateKey(nil)
	unauthPub := unauthPriv.Public().(ed25519.PublicKey)
	sigs = append(sigs, contracts.SignatureEntry{
		ValidatorID: contracts.NodeID(unauthPub),
		Signature:   ed25519.Sign(unauthPriv, digest[:]),
	})

	ok, err = CheckQuorum(pubs, sigs, digest)
	if ok {
		t.Error("Expected failed quorum with unauthorized validator")
	}
}
