package guard

import (
	"bufio"
	"encoding/json"
	"math/rand"
	"os"
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
	Timestamp string `json:"timestamp"`
	HostID    string `json:"host_id"`
	PID       int    `json:"pid"`
	UID       int    `json:"uid"`
	GID       int    `json:"gid"`
	EventType string `json:"event_type"`
	Payload   struct {
		EntropyScore float64 `json:"entropy_score"`
	} `json:"payload"`
}

func (g *GuardAdapter) FetchEvents() ([]bus.TelemetryEvent, error) {
	file, err := os.Open(g.sourceFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var events []bus.TelemetryEvent
	var seqID int64

	for scanner.Scan() {
		seqID++
		var old oldEvent
		if err := json.Unmarshal(scanner.Bytes(), &old); err != nil {
			continue
		}

		var wallUnix int64
		if t, err := time.Parse(time.RFC3339Nano, old.Timestamp); err == nil {
			wallUnix = t.Unix()
		} else {
			wallUnix = 1700000000 + seqID
		}

		event := bus.TelemetryEvent{
			SeqID:        seqID,
			MonotonicNs:  seqID * 1000000,
			WallTimeUnix: wallUnix,
			Source:       "guard.mock",
			HostID:       old.HostID,
			PID:          old.PID,
			TID:          old.PID,
			UID:          old.UID,
			GID:          old.GID,
			EventType:    old.EventType,
			Severity:     old.Payload.EntropyScore,
			Payload:      []byte("{}"),
			Hash:         "mock-hash",
			PrevHash:     "mock-prev",
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
