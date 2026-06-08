/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package sandbox

import "strings"

type CommandGuard struct{}

func (g *CommandGuard) Allow(command string) string {
    if strings.Contains(command, "kernel") || strings.Contains(command, "warden") {
        return "REJECT"
    }
    if strings.Contains(command, "runtime") || strings.Contains(command, "write") {
        return "BLOCK"
    }
    return "SAFE"
}
