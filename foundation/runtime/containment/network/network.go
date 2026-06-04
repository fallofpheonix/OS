/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: CYCLE 9h — NETWORK CONTAINMENT PRIMITIVES (Layer 4)
//
// This file defines the primitive actions available for network containment.
// These are the BUILDING BLOCKS that the SandboxWarden uses to implement
// network isolation.
//
// PRIMITIVES:
//   MONITOR_CONNECTION → Observe network connections without interference
//   THROTTLE_BANDWIDTH → Rate-limit network throughput
//   PAUSE_EGRESS      → Block all outgoing connections
//   QUARANTINE_NAMESPACE → Migrate to isolated network namespace
//   RESTORE_NETWORK   → Remove network restrictions
//
// Each primitive maps to a specific kernel operation:
//   MONITOR_CONNECTION → netfilter/conntrack
//   THROTTLE_BANDWIDTH → tc (traffic control)
//   PAUSE_EGRESS      → iptables DROP
//   QUARANTINE_NAMESPACE → setns(CLONE_NEWNET)
//   RESTORE_NETWORK   → restore original namespace
// =========================================================================
package network

import (
	"time"
)

type NetworkActionType string

const (
	ActionMonitor    NetworkActionType = "MONITOR_CONNECTION"
	ActionThrottle   NetworkActionType = "THROTTLE_BANDWIDTH"
	ActionPause      NetworkActionType = "PAUSE_EGRESS"
	ActionQuarantine NetworkActionType = "QUARANTINE_NAMESPACE"
	ActionRestore    NetworkActionType = "RESTORE_NETWORK"
)

// NetworkAction defines the safe primitive set for network quarantine.
type NetworkAction struct {
	Src        string
	Dst        string
	Port       int
	Action     NetworkActionType
	Reason     string
	EvidenceID string
	DecisionID string
	Timestamp  time.Time
	Sequence   int
	Hash       string
}
