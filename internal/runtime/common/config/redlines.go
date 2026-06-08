/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package config

import (
	"encoding/json"
	"os"
)

type RedLines struct {
	Actuation struct {
		Mode            string  `json:"mode"`
		DefaultPriority float64 `json:"default_priority"`
	} `json:"actuation"`
	Gates struct {
		MinOracleConfidence    float64 `json:"min_oracle_confidence"`
		MinTelemetryConfidence float64 `json:"min_telemetry_confidence"`
		MaxReactionTimeMs      int     `json:"max_reaction_time_ms"`
	} `json:"gates"`
	PhaseLocks struct {
		KernelWrite    string `json:"kernel_write"`
		TrainingEngine string `json:"training_engine"`
		ProposalEngine string `json:"proposal_engine"`
	} `json:"phase_locks"`
	Escalation struct {
		ExtremeEntropyThreshold float64 `json:"extreme_entropy_threshold"`
		AllowStateSkipping      bool    `json:"allow_state_skipping"`
	} `json:"escalation"`
}

func LoadRedLines(path string) (*RedLines, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var rl RedLines
	err = json.Unmarshal(data, &rl)
	if err != nil {
		return nil, err
	}
	return &rl, nil
}
