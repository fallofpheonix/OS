package guard

import (
	"bufio"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"

	"phoenix/bus"
)

type ReplayMode int

const (
	ModeExact       ReplayMode = 0
	ModeAccelerated ReplayMode = 1
	ModeSaturation  ReplayMode = 2
	ModeFault       ReplayMode = 3
)

type GuardAdapter struct {
	Bus         *bus.Bus
	sourceFile  string
	mode        ReplayMode
	speedFactor float64
	Seed        int64
	rng         *rand.Rand
}

func NewGuardAdapter(b *bus.Bus, file string, mode ReplayMode, speed float64, seed int64) *GuardAdapter {
	return &GuardAdapter{
		Bus:         b,
		sourceFile:  file,
		mode:        mode,
		speedFactor: speed,
		Seed:        seed,
		rng:         rand.New(rand.NewSource(seed)),
	}
}

type oldEvent struct {
	Timestamp string                 `json:"timestamp"`
	EventID   string                 `json:"event_id"`
	HostID    string                 `json:"host_id"`
	PID       int                    `json:"pid"`
	PPID      int                    `json:"ppid"`
	UID       int                    `json:"uid"`
	GID       int                    `json:"gid"`
	EventType string                 `json:"event_type"`
	Payload   map[string]interface{} `json:"payload"`
}

type parsedEventWithTime struct {
	time  time.Time
	event oldEvent
	raw   []byte
}

func (g *GuardAdapter) FetchEvents() ([]bus.TelemetryEvent, error) {
	file, err := os.Open(g.sourceFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var parsedEvents []parsedEventWithTime
	var count int64

	for scanner.Scan() {
		count++
		var old oldEvent
		raw := scanner.Bytes()
		// Copy bytes to avoid slice sharing issues
		rawCopy := make([]byte, len(raw))
		copy(rawCopy, raw)

		if err := json.Unmarshal(rawCopy, &old); err != nil {
			continue
		}

		var t time.Time
		if pt, err := time.Parse(time.RFC3339Nano, old.Timestamp); err == nil {
			t = pt
		} else {
			t = time.Unix(1700000000+count, 0)
		}

		parsedEvents = append(parsedEvents, parsedEventWithTime{
			time:  t,
			event: old,
			raw:   rawCopy,
		})
	}

	// Bounded Reorder Window / Deterministic Ingestion Sorting
	// We sort the events by their source epoch (Timestamp) and tie-break on EventID.
	// This guarantees a deterministic sequence order regardless of raw log entry order.
	sort.Slice(parsedEvents, func(i, j int) bool {
		if parsedEvents[i].time.Equal(parsedEvents[j].time) {
			return parsedEvents[i].event.EventID < parsedEvents[j].event.EventID
		}
		return parsedEvents[i].time.Before(parsedEvents[j].time)
	})

	var events []bus.TelemetryEvent
	var seqID int64

	// Deterministic Sequence Allocator
	for _, pe := range parsedEvents {
		seqID++

		entropy := 3.2
		if pe.event.Payload != nil {
			if scoreVal, ok := pe.event.Payload["entropy_score"].(float64); ok {
				entropy = scoreVal
			}
		}

		event := bus.TelemetryEvent{
			SeqID:        seqID,
			MonotonicNs:  seqID * 1000000,
			WallTimeUnix: pe.time.Unix(),
			Source:       "guard.mock",
			HostID:       pe.event.HostID,
			PID:          pe.event.PID,
			TID:          pe.event.PID,
			UID:          pe.event.UID,
			GID:          pe.event.GID,
			EventType:    pe.event.EventType,
			Severity:     entropy,
			Payload:      pe.raw,
			Hash:         "mock-hash",
			PrevHash:     "mock-prev",

			// Hardening fields
			EventID:     pe.event.EventID,
			CausalID:    strconv.Itoa(pe.event.PPID),
			SequenceNo:  seqID,
			SourceEpoch: pe.time.Unix(),
		}

		if g.mode == ModeFault {
			roll := g.rng.Float64()
			if roll < 0.05 {
				continue
			}
		}

		events = append(events, event)
	}

	return events, scanner.Err()
}

func (g *GuardAdapter) Start() (int64, error) {
	events, err := g.FetchEvents()
	if err != nil {
		return 0, err
	}
	for _, event := range events {
		g.Bus.Publish("telemetry.raw", event)
	}
	return int64(len(events)), nil
}

// GetSequenceHash returns a SHA-256 hash of the entire sorted sequence of EventIDs.
// This serves as a "Replay Compiler Proof" to ensure ordering consistency.
func (g *GuardAdapter) GetSequenceHash(events []bus.TelemetryEvent) string {
	h := sha256.New()
	for _, ev := range events {
		h.Write([]byte(ev.EventID))
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
