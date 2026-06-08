/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package observations

import (
	"encoding/json"
	"os"
	"sync"
)

// Index manages the list of stored observations.
type Index struct {
	Cycles []string `json:"cycles"`
}

var (
	indexMu sync.Mutex
	indexPath = "phoenix_os/logs/DRIFT_INDEX.json"
)

// AddToIndex adds a new cycle to the index and persists it.
func AddToIndex(cycle string) error {
	indexMu.Lock()
	defer indexMu.Unlock()

	var idx Index
	data, err := os.ReadFile(indexPath)
	if err == nil {
		_ = json.Unmarshal(data, &idx)
	}

	// Check if cycle already exists
	for _, c := range idx.Cycles {
		if c == cycle {
			return nil
		}
	}

	idx.Cycles = append(idx.Cycles, cycle)

	newData, err := json.MarshalIndent(idx, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(indexPath, newData, 0644)
}
