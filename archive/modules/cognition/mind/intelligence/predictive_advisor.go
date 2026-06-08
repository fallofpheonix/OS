/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 *
 * [PRAS HEADER]
 * Purpose: Monitors drift scores and calculates throttling weights to manage system load and anomalies.
 * Subsystem: Cognition Intelligence
 * Dependencies: os, sync, encoding/json
 * Dependents: AIOrchestrator
 * Security Considerations: Medium. Throttling decisions impact system responsiveness. Log files use restricted permissions.
 * Performance Considerations: Low overhead weight calculation. Thread-safe logging with mutex.
 * Labels: throttling, monitoring, thread-safe
 */
package intelligence

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

// PredictiveAdvisor monitors drift scores and calculates throttling weights.
// Thread-safe for concurrent access from multiple event handlers.
type PredictiveAdvisor struct {
	Threshold float64
	VMax      float64
	logFile   *os.File
	mu        sync.Mutex
}

/*
 * [FUNCTION HEADER]
 * Purpose: Initializes a new PredictiveAdvisor and opens its log file.
 * Responsibilities: Validate log path, open file with restricted permissions (0600).
 * Inputs: threshold (float64), vMax (float64), logPath (string)
 * Outputs: *PredictiveAdvisor, error
 * Complexity: O(1)
 */
// NewPredictiveAdvisor creates a new advisor with the given threshold and max weight value.
// Opens the log file at the specified path with restricted permissions.
func NewPredictiveAdvisor(threshold, vMax float64, logPath string) (*PredictiveAdvisor, error) {
	if logPath == "" {
		logPath = "/tmp/predictive_advisor.log"
	}
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o600)
	if err != nil {
		return nil, err
	}
	return &PredictiveAdvisor{
		Threshold: threshold,
		VMax:      vMax,
		logFile:   f,
	}, nil
}

/*
 * [FUNCTION HEADER]
 * Purpose: Maps a drift score to a throttling weight.
 * Responsibilities: Perform linear interpolation between 0.1 and 1.0 based on drift intensity.
 * Inputs: v (float64) - The current drift score.
 * Outputs: float64 - The calculated weight (0.1 to 1.0).
 * Complexity: O(1)
 */
// CalculateWeight maps a drift score to a throttling weight between 0.1 and 1.0.
// Returns 1.0 if v is at or below threshold (no throttling).
// Returns a linearly interpolated weight between 0.1 and 1.0 for v between threshold and VMax.
// Complexity: O(1)
func (pa *PredictiveAdvisor) CalculateWeight(v float64) float64 {
	if v <= pa.Threshold {
		return 1.0
	}
	if pa.VMax <= pa.Threshold {
		return 0.1
	}
	weight := 1.0 - (v-pa.Threshold)/(pa.VMax-pa.Threshold)
	if weight < 0.1 {
		return 0.1
	}
	return weight
}

/*
 * [FUNCTION HEADER]
 * Purpose: Records a prediction entry to the log file in JSON format.
 * Responsibilities: Marshal data, acquire mutex, and perform thread-safe file write.
 * Inputs: eventID (string), drift (float64), weight (float64)
 * Outputs: None
 * Complexity: O(1)
 */
// LogPrediction records a prediction entry to the log file.
// Thread-safe: acquires mutex before file write.
func (pa *PredictiveAdvisor) LogPrediction(eventID string, drift, weight float64) {
	entry := map[string]interface{}{
		"timestamp": time.Now().Format(time.RFC3339),
		"event_id":  eventID,
		"drift":     drift,
		"weight":    weight,
		"action":    "PREDICTED_THROTTLE",
	}
	data, err := json.Marshal(entry)
	if err != nil {
		fmt.Printf("[DEBUG] Failed to marshal prediction entry: %v\n", err)
		return
	}
	pa.mu.Lock()
	defer pa.mu.Unlock()
	if _, err := pa.logFile.Write(append(data, '\n')); err != nil {
		fmt.Printf("[DEBUG] Failed to write to predictive log: %v\n", err)
	}
}

/*
 * [FUNCTION HEADER]
 * Purpose: Releases the underlying log file handle.
 * Responsibilities: Close the file if open.
 * Inputs: None
 * Outputs: error
 * Complexity: O(1)
 */
// Close releases the underlying log file handle.
// Must be called when the advisor is no longer needed.
func (pa *PredictiveAdvisor) Close() error {
	if pa.logFile != nil {
		return pa.logFile.Close()
	}
	return nil
}
