package guard

import (
	"bufio"
	"encoding/json"
	"math/rand"
	"os"
	"time"

	"github.com/fallofpheonix/phoenix_os/bus"
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
}

func NewGuardAdapter(b *bus.Bus, file string, mode ReplayMode, speed float64) *GuardAdapter {
	return &GuardAdapter{
		Bus:         b,
		sourceFile:  file,
		mode:        mode,
		speedFactor: speed,
	}
}

// Legacy test_events.jsonl structures
type oldPayload struct {
	EntropyScore float64 `json:"entropy_score"`
}
type oldEvent struct {
	Timestamp string     `json:"timestamp"`
	EventID   string     `json:"event_id"`
	EventType string     `json:"event_type"`
	HostID    string     `json:"host_id"`
	PID       int        `json:"pid"`
	UID       int        `json:"uid"`
	GID       int        `json:"gid"`
	Payload   oldPayload `json:"payload"`
}

func (g *GuardAdapter) Start() (int64, error) {
	file, err := os.Open(g.sourceFile)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var prevMono int64
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

		// Replay timing
		if prevMono != 0 {
			diff := event.MonotonicNs - prevMono
			if diff > 0 {
				switch g.mode {
				case ModeExact:
					time.Sleep(time.Duration(diff))
				case ModeAccelerated:
					time.Sleep(time.Duration(float64(diff) / g.speedFactor))
				case ModeSaturation:
					// no sleep
				case ModeFault:
					roll := rand.Float64()
					if roll < 0.05 {
						continue
					} else if roll < 0.10 {
						time.Sleep(time.Duration(diff) * 10)
					}
				}
			}
		}
		prevMono = event.MonotonicNs

		g.Bus.Publish("telemetry.raw", event)
	}

	return seqID, scanner.Err()
}
