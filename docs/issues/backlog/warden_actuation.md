# Issue: [Control] Warden Actuation Feature Backlog

## Problem
Need to establish a safe, rate-limited, and auditable actuation layer that prevents self-DOS and feedback oscillations.

## Current State
Basic FSM and MockWarden established. Requires sophisticated actuation classing and budgeting.

## Required Work
- [ ] Warden FSM state‑graph visualization.
- [ ] Warden-configurable thresholds.
- [ ] Warden-actuation‑classes (Class 0-5).
- [ ] Warden-actuation‑budgeting.
- [ ] Warden-cooldown‑per‑node.
- [ ] Warden-cooldown‑per‑action.
- [ ] Warden-rollback‑tracking.
- [ ] Warden-feedback‑loop protection.
- [ ] Warden-“safe‑mode”.
- [ ] Warden-actuation‑API.
- [ ] Warden-actuation‑confirmation.
- [ ] Warden-actuation‑logs.
- [ ] Warden-multiple‑policy‑engines.
- [ ] Warden-replay‑safety.
- [ ] Warden-cross‑session limits.

## Priority
High

## Labels
control, security, architecture
