package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type SystemState uint8

const (
	StateSafe        SystemState = 0
	StateWatch       SystemState = 1
	StateSuspicious  SystemState = 2
	StateCritical    SystemState = 3
	StateCompromised SystemState = 4
)

type Config struct {
	Thresholds struct {
		Safe       float64 `json:"safe"`
		Watch      float64 `json:"watch"`
		Suspicious float64 `json:"suspicious"`
		Critical   float64 `json:"critical"`
	} `json:"thresholds"`
}

type Warden struct {
	CurrentState SystemState
	Throttling   float64 // 0.0 (None) to 1.0 (Full Block)
	Config       Config
}

func NewWarden(configPath string) (*Warden, error) {
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := json.Unmarshal(configFile, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &Warden{
		CurrentState: StateSafe,
		Throttling:   0.0,
		Config:       config,
	}, nil
}

// EvaluateSDI maps the System Disorder Index to a discrete state
func (w *Warden) EvaluateSDI(sdi float64) {
	oldState := w.CurrentState
	
	switch {
	case sdi < w.Config.Thresholds.Safe:
		w.CurrentState = StateSafe
	case sdi < w.Config.Thresholds.Watch:
		w.CurrentState = StateWatch
	case sdi < w.Config.Thresholds.Suspicious:
		w.CurrentState = StateSuspicious
	case sdi < w.Config.Thresholds.Critical:
		w.CurrentState = StateCritical
	default:
		w.CurrentState = StateCompromised
	}

	if oldState != w.CurrentState {
		fmt.Printf("[WARDEN] State Transition: %v -> %v (SDI: %.2f)\n", oldState, w.CurrentState, sdi)
		w.ApplyAction()
	}
}

func (w *Warden) ApplyAction() {
	switch w.CurrentState {
	case StateSafe:
		w.Throttling = 0.0
		fmt.Println("[WARDEN] Action: Observation only. No throttling.")
	case StateWatch:
		w.Throttling = 0.1
		fmt.Println("[WARDEN] Action: Increased sampling. Minimal throttling (10%).")
	case StateSuspicious:
		w.Throttling = 0.5
		fmt.Println("[WARDEN] Action: Process Throttling (50%).")
	case StateCritical:
		w.Throttling = 0.9
		fmt.Println("[WARDEN] Action: Process Isolation (90%).")
	case StateCompromised:
		w.Throttling = 1.0
		fmt.Println("[WARDEN] Action: Forensic Snapshot + Termination.")
	}
}

func main() {
	fmt.Println("Phoenix Warden starting with Finite-State Controller...")
	warden, err := NewWarden("../../config/warden.json")
	if err != nil {
		fmt.Printf("Error initializing Warden: %v\n", err)
		os.Exit(1)
	}

	// Simulate rising threat
	warden.EvaluateSDI(0.15) // Safe
	warden.EvaluateSDI(0.45) // Watch
	warden.EvaluateSDI(0.65) // Suspicious
	warden.EvaluateSDI(0.95) // Compromised
}
