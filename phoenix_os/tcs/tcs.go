package tcs

import (
	"time"
)

type TelemetryEvent struct {
	Timestamp  time.Time
	SequenceID uint64
	Payload    []byte
}

type SlidingWindow struct{}

func NewSlidingWindow(d time.Duration) *SlidingWindow { return &SlidingWindow{} }
func (s *SlidingWindow) AddEvent(e TelemetryEvent) {}
func (s *SlidingWindow) Evaluate() float64 { return 1.0 }

type DegradationMonitor struct{}

func NewDegradationMonitor(w *SlidingWindow, i interface{}) *DegradationMonitor { return &DegradationMonitor{} }
func (d *DegradationMonitor) IsDegraded() bool { return false }
func (d *DegradationMonitor) Evaluate(score float64) {}
