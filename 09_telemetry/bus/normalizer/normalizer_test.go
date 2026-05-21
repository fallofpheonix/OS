package normalizer

import (
	"testing"
	"time"
)

func TestNormalize_ValidEvent(t *testing.T) {
	norm := NewEventNormalizer()
	rawJSON := `{
		"timestamp": "2026-05-21T15:00:00Z",
		"event_id": "a6b6d8e6-1234-4567-8910-1234567890ab",
		"category": "process",
		"event_type": "execve",
		"host_id": "test-host-macos",
		"pid": 1234,
		"ppid": 100,
		"uid": 501,
		"gid": 20,
		"comm": "zsh",
		"exe_path": "/bin/zsh",
		"payload": {
			"args": ["ls", "-la"],
			"env_vars": ["PATH=/usr/bin"]
		}
	}`

	evt, err := norm.Normalize([]byte(rawJSON))
	if err != nil {
		t.Fatalf("unexpected error normalizing: %v", err)
	}

	if evt.PID != 1234 {
		t.Errorf("expected PID 1234, got %d", evt.PID)
	}
	if evt.Comm != "zsh" {
		t.Errorf("expected comm 'zsh', got %q", evt.Comm)
	}
}

func TestNormalize_InvalidEvent(t *testing.T) {
	norm := NewEventNormalizer()
	// Missing required field "host_id"
	rawJSON := `{
		"timestamp": "2026-05-21T15:00:00Z",
		"event_id": "a6b6d8e6-1234-4567-8910-1234567890ab",
		"category": "process",
		"event_type": "execve",
		"pid": 1234,
		"ppid": 100,
		"uid": 501,
		"gid": 20,
		"comm": "zsh",
		"exe_path": "/bin/zsh",
		"payload": {}
	}`

	_, err := norm.Normalize([]byte(rawJSON))
	if err == nil {
		t.Error("expected normalization error for missing host_id, got nil")
	}
}

func TestEnrich_PathNormalization(t *testing.T) {
	norm := NewEventNormalizer()
	evt := &Event{
		Timestamp: time.Now(),
		EventID:   "a6b6d8e6-1234-4567-8910-1234567890ab",
		Category:  "filesystem",
		EventType: "open",
		HostID:    "test-host",
		PID:       1234,
		UID:       501,
		Comm:      "test",
		ExePath:   "/usr/bin/test",
		Payload: map[string]interface{}{
			"file_path": "/a/b/../c/./d",
		},
	}

	err := norm.Enrich(evt)
	if err != nil {
		t.Fatalf("unexpected error enriching: %v", err)
	}

	cleanedPath := evt.Payload["file_path"].(string)
	expected := "/a/c/d"
	if cleanedPath != expected {
		t.Errorf("expected cleaned path %q, got %q", expected, cleanedPath)
	}
}

func TestCalculateEntropy(t *testing.T) {
	// Zero entropy for uniform repeating byte
	uniform := []byte("aaaa")
	entropy := CalculateEntropy(uniform)
	if entropy != 0.0 {
		t.Errorf("expected 0 entropy for uniform string, got %f", entropy)
	}

	// Non-zero entropy for high variety
	varied := []byte("abcdefghijklmnopqrstuvwxyz")
	entropy = CalculateEntropy(varied)
	if entropy < 4.0 {
		t.Errorf("expected high entropy for unique alphabet, got %f", entropy)
	}

	// 8 bits of entropy per byte if we had all 256 bytes once.
	allBytes := make([]byte, 256)
	for i := 0; i < 256; i++ {
		allBytes[i] = byte(i)
	}
	maxEntropy := CalculateEntropy(allBytes)
	if mathAbs(maxEntropy-8.0) > 0.0001 {
		t.Errorf("expected 8.0 entropy for complete byte spectrum, got %f", maxEntropy)
	}
}

func mathAbs(v float64) float64 {
	if v < 0 {
		return -v
	}
	return v
}

func TestEnrich_UserCache(t *testing.T) {
	norm := NewEventNormalizer()
	evt := &Event{
		Timestamp: time.Now(),
		EventID:   "a6b6d8e6-1234-4567-8910-1234567890ab",
		Category:  "process",
		EventType: "exec",
		HostID:    "test",
		PID:       1234,
		UID:       9999, // Unlikely to exist, resolves to unknown
		Comm:      "test",
		ExePath:   "/usr/bin/test",
		Payload:   map[string]interface{}{},
	}

	start := time.Now()
	_ = norm.Enrich(evt)
	firstLookupDuration := time.Since(start)

	if evt.Username == "" {
		t.Error("expected username to be populated, even if 'unknown'")
	}

	// Second lookup must be extremely fast due to caching
	start = time.Now()
	_ = norm.Enrich(evt)
	secondLookupDuration := time.Since(start)

	t.Logf("First lookup: %v, Second lookup (cached): %v", firstLookupDuration, secondLookupDuration)
}
