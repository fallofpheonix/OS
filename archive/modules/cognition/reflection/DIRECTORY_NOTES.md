# Reflection Subsystem

## Primary Responsibility
The Reflection subsystem implements the epistemic safety layer for PhoenixOS. It is responsible for proactive world modeling (Predictions) and cumulative error detection (Reality Drift).

## System Architecture
1. **Prediction Engine:** Generates expected future states for system actions. Verifies expectations against subsequent ledger facts using Hamming distance metrics.
2. **Reality Drift Auditor:** Aggregates prediction errors over time to measure the divergence between the internal cognitive model and measured reality. Enforces the system-wide Quarantine cycle.

## Tech Stack
- Go (Standard Library)
- github.com/fallofpheonix/Phoenix.Nucleus/ledger (Verification root)

## AI-Specific Context
- **System Boundaries:** Acts as the "Sanity Meter" for `PhoenixMind`. Can overrule AI directives if drift exceeds safety thresholds.
- **Data Flow:** Prediction (Pre-Action) -> Fact Verification (Post-Action) -> Error Signal -> Drift Accumulation -> Quarantine Trigger.
