package serialization

import (
	"encoding/binary"
	"testing"
)

func TestEndiannessDeterminism(t *testing.T) {
	val := uint64(0x1234567890ABCDEF)
	buf := make([]byte, 8)
	
	// Axiom: We use BigEndian for all binary serialization in the Ledger/Trace path.
	binary.BigEndian.PutUint64(buf, val)
	
	expected := []byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}
	for i := range buf {
		if buf[i] != expected[i] {
			t.Errorf("Endianness mismatch at byte %d: got %02x, want %02x", i, buf[i], expected[i])
		}
	}
}
