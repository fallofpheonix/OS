# PhoenixOS: Offensive Research Strategy

To ensure PhoenixOS remains a **Deterministic Cybernetic Security Runtime** and does not drift into becoming an offensive platform, all offensive research is strictly isolated into dedicated environments.

## 1. Structural Separation

The project is split into four distinct entities to maintain strict boundaries between production defense and offensive research:

*   **PhoenixOS-Core:** The production-grade deterministic runtime, telemetry, ledger, and warden.
*   **PhoenixRange:** A cyber range for training, exercises, and validating defense scenarios.
*   **PhoenixRedLab:** A lab environment for adversary emulation, detection testing, and validation.
*   **PhoenixSim:** A simulation environment for synthetic incidents, agent stress testing, and world modeling.

### Directory Mapping

```text
phoenix_os/           # PhoenixOS-Core
    core/             # Deterministic logic
    replay/           # Replay authority
    telemetry/        # L2/L3 ingestion
    ledger/           # Evidence chain
    warden/           # Actuation FSM

14_experiments/phoenix_range/    # PhoenixRange
    scenarios/                   # Exercise scenarios
    attack-graphs/               # Visualized attack paths
    replay-lab/                  # Lab for trace analysis

14_experiments/phoenix_red_lab/  # PhoenixRedLab
    emulation/                   # Adversary emulation scripts
    detection-tests/             # Tests for Warden/Arbiter
    validation/                  # Outcome verification

14_experiments/phoenix_sim/      # PhoenixSim
    synthetic-world/             # Environment modeling
    agents/                      # Synthetic attacker/user agents
    stress/                      # Saturation and noise testing
```

---

## 2. Research Areas

The following areas are authorized for research within the isolated labs:

1.  **Adversary Emulation:** Modeling specific TTPs.
2.  **Detection Validation:** Quantifying the effectiveness of Warden policies.
3.  **Red-Team Simulations:** Controlled offensive runs to test survival objectives.
4.  **Purple-Team Workflows:** Collaborative tuning of telemetry and actuation.
5.  **Attack Replay Datasets:** Curating traces of real/synthetic attacks for regression testing.
6.  **IOC Generation:** Extracting indicators from replayed sessions.
7.  **Campaign Modeling:** Grouping related events into logical campaigns.
8.  **Persistence Simulations:** Testing detection of long-dwell persistence mechanisms.
9.  **Lateral Movement Modeling:** Graph-based analysis of spread patterns.
10. **Exfiltration Simulations:** Testing throughput-based exfiltration detection.

---

## 3. Offensive Research Backlog (Isolated)

The following build issues are tracked for the research environments:

*   **ATTACK-001:** Recon simulation engine (Passive/Active scanning).
*   **ATTACK-002:** Persistence scenario library (Cron, systemd, hooks).
*   **ATTACK-003:** Credential misuse replay (Auth telemetry validation).
*   **ATTACK-004:** Beacon emulation runtime (C2 traffic patterns).
*   **ATTACK-005:** Ransomware impact simulator (File entropy bursts).
*   **ATTACK-006:** Lateral movement graph model (Causal link validation).
*   **ATTACK-007:** Exfiltration path simulator (Network flow analysis).
*   **ATTACK-008:** Insider scenario generator (Behavioral drift).
*   **ATTACK-009:** Campaign replay framework (Multi-stage trace execution).
*   **ATTACK-010:** Recovery benchmark suite (Time-to-safe verification).

---

## 4. Operational Boundaries

Strict adherence to the following boundaries is mandatory:

*   **Production != Research:** Production binaries must never contain emulation or attack code.
*   **Detection != Emulation:** The Warden detects; the Lab emulates. Never combine these logic paths.
*   **Defense != Real-world operations:** PhoenixOS provides counter-pressure for defense, not offensive intrusion tools.

### Authorized Runtime Modes
*   **Lab Mode:** For testing specific emulation scripts.
*   **Training Mode:** For human-in-the-loop exercises in PhoenixRange.
*   **Replay Mode:** For deterministic analysis of previous sessions.
*   **Simulation Mode:** For synthetic world-state testing in PhoenixSim.
*   **Validation Mode:** For automated regression testing of security axioms.

---

## 5. Long-term Branch Layout

*   **PhoenixOS-Core:** Deterministic runtime (Mainline).
*   **PhoenixSOC:** Detection, replay, and alert management.
*   **PhoenixDFIR:** Evidence, timelines, and forensic export.
*   **PhoenixRange:** Training, exercises, and range orchestration.
*   **PhoenixSim:** Synthetic incidents and stress modeling.
*   **PhoenixResearch:** Advanced experiments (e.g., microkernel, custom scheduling).
