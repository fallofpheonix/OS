/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package governance

import "fmt"

// GuardTransition checks if a transition is allowed based on rules.
func GuardTransition(from string, to string) (bool, error) {
    // F0: CLOSED, F1: ACTIVE, F2: LOCKED
    // Rejected: F1 → Training, F1 → Proposal, Observation → Merge, Runtime → Patch
    
    if from == "F1" && to == "Training" {
        return false, fmt.Errorf("F1 to Training is forbidden")
    }
    if from == "F1" && to == "Proposal" {
        return false, fmt.Errorf("F1 to Proposal is forbidden")
    }
    if from == "Observation" && to == "Merge" {
        return false, fmt.Errorf("Observation to Merge is forbidden")
    }
    if from == "Runtime" && to == "Patch" {
        return false, fmt.Errorf("Runtime to Patch is forbidden")
    }
    
    return true, nil
}
