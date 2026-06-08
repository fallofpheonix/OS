# PhoenixMind Cyber Operations Layer

This directory houses the Cyber Operations agents for PhoenixOS.

## Agents

### 1. Red Team Agent (`red_team.py`)
Responsible for "Heat Injection" and threat simulation. It replays failure patterns from the `06_FAILURE_LIBRARY`.

### 2. Blue Team Agent (`blue_team.py`)
Responsible for intercepting threats, monitoring drift, performing forensic analysis, and projecting future threats using Stackelberg models.

### 3. Cyber Ops Orchestrator (`orchestrator.py`)
Coordinates the interaction between the Red and Blue teams in a controlled scenario.

## Usage

To run the default "executable-code-in-brain" scenario:

```bash
python3 orchestrator.py
```

## Workflows

### Master Cyber Ops Loop
1.  **Red Team** selects a threat pattern and injects "Heat" into the system.
2.  **Blue Team** monitors telemetry for drift and evaluates the **Threat Projector**.
3.  If a threat is detected or projected, the **Blue Team** triggers forensic analysis to derive a causal proof.
4.  The **Blue Team** requests actuation (e.g., `ISOLATE_PID`) and logs the decision to the **State Scribe** (Ledger).

## Future Enhancements
- Integration with the Go-based `AIOrchestrator` via a bridge.
- Expanded failure library mapping.
- Real-time Stackelberg game-theory modeling for multi-step attacks.
