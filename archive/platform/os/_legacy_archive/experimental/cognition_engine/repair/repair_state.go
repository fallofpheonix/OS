/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package repair

type RepairState string

const (
    Success   RepairState = "SUCCESS"
    Failed    RepairState = "FAILED"
    Regressed RepairState = "REGRESSED"
    Blocked   RepairState = "BLOCKED"
    Unknown   RepairState = "UNKNOWN"
)

type RiskLevel string

const (
    Low       RiskLevel = "LOW"
    Medium    RiskLevel = "MEDIUM"
    High      RiskLevel = "HIGH"
    Dangerous RiskLevel = "DANGEROUS"
)
