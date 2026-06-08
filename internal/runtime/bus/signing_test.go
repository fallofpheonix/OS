/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package bus

import (
	"encoding/hex"
	"testing"
)

func TestEventSigningAndVerification(t *testing.T) {
	b := NewBus()
	topic := "test.signing"
	ch := b.Subscribe(topic)

	event := TelemetryEvent{
		SeqID:     1,
		Source:    "test-source",
		EventType: "test-event",
		Payload:   []byte(`{"data": "test"}`),
	}

	// Publish without hash, Bus should sign it
	b.Publish(topic, event)

	received := <-ch
	if received.Hash == "" {
		t.Fatal("Expected event to be signed by the Bus")
	}

	// Verify the signature
	if !VerifyEvent(received, nil) {
		t.Error("Event verification failed for Bus-signed event")
	}

	// Tamper with the event
	received.Payload = []byte(`{"data": "tampered"}`)
	if VerifyEvent(received, nil) {
		t.Error("Event verification should have failed for tampered event")
	}
}

func TestSetSigningKey(t *testing.T) {
	key := []byte("secret-key")
	SetSigningKey(key)

	event := TelemetryEvent{SeqID: 2}
	sig := ComputeEventSignature(event, nil)

	// Verify it was signed with the new key
	if !VerifyEvent(TelemetryEvent{SeqID: 2, Hash: hex.EncodeToString(sig)}, key) {
		t.Error("Verification failed with explicitly set key")
	}

	// Verify it fails with default/wrong key
	if VerifyEvent(TelemetryEvent{SeqID: 2, Hash: hex.EncodeToString(sig)}, []byte("wrong-key")) {
		t.Error("Verification should have failed with wrong key")
	}
}
