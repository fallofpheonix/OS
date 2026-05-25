# PhoenixOS: Validation Rules

This document defines the gates for validating PhoenixOS components, including the new cognitive architecture mandates.

## 1. Fast-Path Determinism
- All L1 Guard heuristics must execute in <10ms.
- 100% replay hash matching is required for all kernel event traces.

## 2. Stochastic Decision Validation
- Probabilistic decisions must be reproducible using a fixed random seed in Replay Lab.
- Monte Carlo simulations must converge on the "Target Policy" within 1000 iterations.

## 3. Meta-Learning Performance Gates
- The system must correctly identify a 10% artificial drift in detection confidence.
- `HIGH_UNCERTAINTY_DETECTED` flags must be raised when average accuracy drops below 60% across 10 evaluation ticks.

## 4. Reflective Feedback Verification
- Actuation outcomes must be recorded in the Truth Ledger within 50ms of physical action.
- Feedback loop must demonstrably adjust Arbiter weights in the subsequent policy evaluation cycle.
