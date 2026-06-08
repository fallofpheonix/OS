/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: CYCLE 2b — EVENT NORMALIZATION AND ENRICHMENT
//
// The EventNormalizer transforms raw JSON telemetry into validated,
// enriched Event structs. It's an ALTERNATIVE to the telemetry/normalizer
// package, operating at the Bus level rather than the kernel level.
//
// WORKFLOW:
//   Raw JSON → EventNormalizer.Normalize(raw)
//     → json.Unmarshal → Event struct
//     → Validate required fields (timestamp, event_id, category, etc.)
//     → Return validated Event
//   → EventNormalizer.Enrich(event)
//     → Resolve UID to username (cached)
//     → Normalize filesystem paths
//     → Return enriched Event
//
// VALIDATION: Required fields: timestamp, event_id, category, event_type,
// host_id, comm, exe_path, payload. Missing fields cause normalization failure.
//
// ENRICHMENT: UID→username resolution uses in-memory cache for performance.
// Filesystem paths are normalized using filepath.Clean().
//
// ENTROPY: CalculateEntropy() provides Shannon entropy for byte buffers.
// Used by the Detector for threat analysis.
// =========================================================================
package normalizer

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os/user"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

// Event represents the normalized telemetry event routed through the event bus
type Event struct {
	Timestamp   time.Time              `json:"timestamp"`
	EventID     string                 `json:"event_id"`
	Category    string                 `json:"category"`
	EventType   string                 `json:"event_type"`
	HostID      string                 `json:"host_id"`
	PID         uint32                 `json:"pid"`
	PPID        uint32                 `json:"ppid"`
	UID         uint32                 `json:"uid"`
	GID         uint32                 `json:"gid"`
	Username    string                 `json:"username,omitempty"`
	Comm        string                 `json:"comm"`
	ExePath     string                 `json:"exe_path"`
	ContainerID string                 `json:"container_id,omitempty"`
	Payload     map[string]interface{} `json:"payload"`
}

// Normalizer interface defines the parser and enricher contract
type Normalizer interface {
	Normalize(raw []byte) (*Event, error)
	Enrich(event *Event) error
}

// EventNormalizer implements the Normalizer interface
type EventNormalizer struct {
	userCache sync.Map // Cache UID (uint32) -> Username (string)
}

// NewEventNormalizer initializes a new EventNormalizer
func NewEventNormalizer() *EventNormalizer {
	return &EventNormalizer{}
}

// Normalize parses and validates a raw telemetry event
func (n *EventNormalizer) Normalize(raw []byte) (*Event, error) {
	var event Event
	if err := json.Unmarshal(raw, &event); err != nil {
		return nil, fmt.Errorf("malformed JSON payload: %w", err)
	}

	// Validate required header fields
	if event.Timestamp.IsZero() {
		return nil, errors.New("missing required field: timestamp")
	}
	if event.EventID == "" {
		return nil, errors.New("missing required field: event_id")
	}
	switch event.Category {
	case "process", "syscall", "filesystem", "network", "container", "memory":
		// valid
	default:
		return nil, fmt.Errorf("invalid event category: %q", event.Category)
	}
	if event.EventType == "" {
		return nil, errors.New("missing required field: event_type")
	}
	if event.HostID == "" {
		return nil, errors.New("missing required field: host_id")
	}
	if event.Comm == "" {
		return nil, errors.New("missing required field: comm")
	}
	if event.ExePath == "" {
		return nil, errors.New("missing required field: exe_path")
	}
	if event.Payload == nil {
		return nil, errors.New("missing required field: payload")
	}

	return &event, nil
}

// Enrich resolves UIDs to usernames and sanitizes paths
func (n *EventNormalizer) Enrich(event *Event) error {
	// 1. UID -> Username resolution with in-memory caching
	if cachedVal, ok := n.userCache.Load(event.UID); ok {
		event.Username = cachedVal.(string)
	} else {
		resolvedUsername := "unknown"
		u, err := user.LookupId(strconv.Itoa(int(event.UID)))
		if err == nil {
			resolvedUsername = u.Username
		}
		n.userCache.Store(event.UID, resolvedUsername)
		event.Username = resolvedUsername
	}

	// 2. Filesystem path normalization
	if event.Category == "filesystem" {
		if rawPath, ok := event.Payload["file_path"]; ok {
			if pathStr, ok := rawPath.(string); ok {
				event.Payload["file_path"] = filepath.Clean(pathStr)
			}
		}
	}

	return nil
}

// CalculateEntropy calculates the Shannon entropy of a given byte buffer
func CalculateEntropy(data []byte) float64 {
	if len(data) == 0 {
		return 0.0
	}
	frequencies := make(map[byte]int)
	for _, b := range data {
		frequencies[b]++
	}
	var entropy float64
	length := float64(len(data))
	for _, count := range frequencies {
		p := float64(count) / length
		entropy -= p * math.Log2(p)
	}
	return entropy
}
