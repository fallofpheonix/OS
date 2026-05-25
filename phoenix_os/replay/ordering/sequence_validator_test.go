package ordering

import (
	"testing"
)

func TestSequenceValidator_Validate(t *testing.T) {
	sv := NewSequenceValidator()

	// Test 1: Empty sequence (should be valid)
	err := sv.Validate([]*Event{})
	if err != nil {
		t.Errorf("Expected no error for empty sequence, got: %v", err)
	}

	// Test 2: Correctly ordered sequence
	events1 := []*Event{
		{SequenceID: 1, Payload: []byte("event A")},
		{SequenceID: 2, Payload: []byte("event B")},
		{SequenceID: 3, Payload: []byte("event C")},
	}
	err = sv.Validate(events1)
	if err != nil {
		t.Errorf("Expected no error for correctly ordered sequence, got: %v", err)
	}

	// Test 3: Sequence with regression (decreasing ID)
	events2 := []*Event{
		{SequenceID: 1, Payload: []byte("event A")},
		{SequenceID: 3, Payload: []byte("event C")},
		{SequenceID: 2, Payload: []byte("event B")}, // Regression
	}
	err = sv.Validate(events2)
	if err == nil {
		t.Error("Expected error for sequence regression, got nil")
	} else {
		expectedErr := "sequence regression or duplication detected: event at index 2 (ID 2) is not strictly greater than event at index 1 (ID 3)"
		if err.Error() != expectedErr {
			t.Errorf("Expected error '%s', got: '%v'", expectedErr, err)
		}
	}

	// Test 4: Sequence with duplicate sequence IDs
	events3 := []*Event{
		{SequenceID: 1, Payload: []byte("event A")},
		{SequenceID: 2, Payload: []byte("event B")},
		{SequenceID: 2, Payload: []byte("event C")}, // Duplicate
	}
	err = sv.Validate(events3)
	if err == nil {
		t.Error("Expected error for duplicate sequence ID, got nil")
	} else {
		expectedErr := "duplicate sequence ID 2 found at index 2"
		if err.Error() != expectedErr {
			t.Errorf("Expected error '%s', got: '%v'", expectedErr, err)
		}
	}

	// Test 5: Sequence with nil event
	events4 := []*Event{
		{SequenceID: 1, Payload: []byte("event A")},
		nil, // Nil event
		{SequenceID: 3, Payload: []byte("event C")},
	}
	err = sv.Validate(events4)
	if err == nil {
		t.Error("Expected error for nil event, got nil")
	} else {
		expectedErr := "event at index 1 is nil"
		if err.Error() != expectedErr {
			t.Errorf("Expected error '%s', got: '%v'", expectedErr, err)
		}
	}

	// Test 6: Sequence with zero sequence ID
	events5 := []*Event{
		{SequenceID: 0, Payload: []byte("event A")}, // Zero ID
		{SequenceID: 1, Payload: []byte("event B")},
	}
	err = sv.Validate(events5)
	if err == nil {
		t.Error("Expected error for zero sequence ID, got nil")
	} else {
		expectedErr := "event at index 0 has a zero sequence ID"
		if err.Error() != expectedErr {
			t.Errorf("Expected error '%s', got: '%v'", expectedErr, err)
		}
	}

	// Test 7: Correct sequence with non-consecutive IDs
	events6 := []*Event{
		{SequenceID: 1, Payload: []byte("event A")},
		{SequenceID: 3, Payload: []byte("event C")}, // Valid, just non-consecutive
		{SequenceID: 5, Payload: []byte("event E")},
	}
	err = sv.Validate(events6)
	if err != nil {
		t.Errorf("Expected no error for non-consecutive but ordered sequence, got: %v", err)
	}
}
