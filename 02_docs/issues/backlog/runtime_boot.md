# Issue: [Runtime] Boot & Core Runtime Feature Backlog

## Problem
Need to establish the foundational, immutable, and deterministic runtime environment for PhoenixOS.

## Current State
PhoenixOS currently operates as a Go binary on a macOS host.

## Required Work
- [ ] Deterministic boot checksum.
- [ ] Immutable initrd with Warden as PID1.
- [ ] Deterministic CPU-idle / tick behavior.
- [ ] Per-boot replay-session ID.
- [ ] Boot-time telemetry snapshot.
- [ ] Hardware-attestation-like self-hash.
- [ ] Boot-mode selector (--replay, --record, --monitor).
- [ ] Minimal LinuxKit-style appliance build.
- [ ] Deterministic entropy-source configuration.
- [ ] Replay-compatible resume from snapshot.

## Acceptance Criteria
[ ] All features implemented
[ ] Verified deterministic boot hashes
[ ] Verified immutable initrd behavior

## Priority
High

## Labels
runtime, architecture, security
