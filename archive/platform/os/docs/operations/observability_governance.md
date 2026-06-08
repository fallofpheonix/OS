---\nStatus: Partial\nImplementation: 60%\nConfidence: Tested\n---\n# Observability Governance

This document outlines the Observation Consolidation and Governance mechanisms implemented in Task 3.

## Overview
The goal is to provide stable drift history, evidence accumulation, and governance enforcement for PhoenixOS.

## Architecture
- `observations/`: Accumulates observability data from runtime.
- `drift/`: Tracks and compares drift against baselines.
- `governance/`: Enforces phase-based transitions and readiness scoring.

## Governance Rules
- F0: CLOSED
- F1: ACTIVE
- F2: LOCKED (Review Only)
- Training: DISABLED
- Merge: HUMAN_ONLY

## Exit Criteria
- 10 stable cycles
- Baseline exists
- Drift trends computed
- Governance active
- Readiness score generated
- F2 still locked
