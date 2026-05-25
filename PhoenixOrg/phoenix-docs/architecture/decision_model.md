# PhoenixOS: Decision Model

The Decision Model governs the transition between strategic policy (Arbiter) and physical action (Warden). It is the implementation of the OS's "Mind."

## 1. Stochastic Decision Components
To handle the "Fog of War" in cyber-security, the Decision Model uses probabilistic evaluation for importance scores between 0.4 and 0.8 (The Uncertainty Zone).

- **Algorithm**: Monte Carlo Simulation for risk impact.
- **Goal**: Introduce nuance to prevent binary oscillation between SAFE and CRITICAL states.
- **Rule**: Kernel Fast-Path (L1) remains strictly deterministic. Stochasticity is reserved for the Advisory/Arbiter layers.

## 2. Meta-Learning Framework
PhoenixOS observes its own execution to define its operational boundaries.
- **Self-Auditing**: The system tracks its "False Positive" rate by comparing Arbiter decisions against Replay Truth.
- **Limitation Awareness**: If detection confidence drops below 60% for a specific category, the system marks that module as "UNTREATED" and escalates to Human-in-the-Loop.

## 3. Reflective Feedback Loops
Every actuation triggers a three-stage reflection:
1. **Detection Accuracy**: Was the threat real (Replay verified)?
2. **Containment Success**: Did the action stop the disorder (SDI reduction)?
3. **Rollback Cost**: What was the system impact of the recovery?

These metrics are fed back into the Stackelberg weights in the Arbiter to optimize future defense strategies.
