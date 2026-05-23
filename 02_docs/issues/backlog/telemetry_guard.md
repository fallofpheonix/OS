# Issue: [Telemetry] Guard & eBPF Feature Backlog

## Problem
Need to establish granular, performant, and tamper-resistant telemetry capture in the kernel.

## Current State
Core eBPF architecture mapped and drafted, but lacks advanced filtering, rate limiting, and robust health monitoring.

## Required Work
- [ ] Guard-mode selectable.
- [ ] Hot-plug telemetry adapters.
- [ ] Telemetry-filtering at kernel.
- [ ] Telemetry-rate-limiting.
- [ ] Telemetry-prioritization.
- [ ] Telemetry-level toggling.
- [ ] Telemetry-checksumming.
- [ ] Telemetry-replay-proof.
- [ ] Guard-configuration-via-configmap.
- [ ] Guard-health-monitoring.
- [ ] Guard-cold-start warm-up.
- [ ] Guard-support for XDP-style networking.
- [ ] Guard-support for kernel-security-module hooks.
- [ ] Guard-support for container-runtime hooks.
- [ ] Guard-“noisy‑process” throttling.
- [ ] Guard-“sensitive” telemetry‑only‑on‑request.
- [ ] Guard-tamper‑resistant counters.
- [ ] Guard-delay‑analysis.
- [ ] Guard-Linux‑compatibility‑matrix.
- [ ] Guard-multi‑node telemetry‑replay.

## Priority
High

## Labels
telemetry, security, kernel
