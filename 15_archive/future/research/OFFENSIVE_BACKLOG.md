# PhoenixOS: Offensive Research Backlog

This backlog tracks development within the **PhoenixRange**, **PhoenixRedLab**, and **PhoenixSim** isolated environments.

## 1. Primary Research Tasks

| Issue ID | Title | Subsystem | Description |
| :--- | :--- | :--- | :--- |
| **ATTACK-001** | Recon Simulation Engine | PhoenixSim | Passive/Active scanning emulation to test counter-recon misdirection. |
| **ATTACK-002** | Persistence Scenario Library | PhoenixRedLab | Curated set of persistence methods (Cron, systemd) for detection validation. |
| **ATTACK-003** | Credential Misuse Replay | PhoenixRedLab | Auth telemetry validation using replayed credential-access traces. |
| **ATTACK-004** | Beacon Emulation Runtime | PhoenixSim | C2 traffic pattern generator for long-horizon exfiltration testing. |
| **ATTACK-005** | Ransomware Impact Simulator | PhoenixRedLab | High-entropy file write bursts to test Fast-Path encryption detection. |
| **ATTACK-006** | Lateral Movement Graph Model | PhoenixRange | Validation of causal process lineage during simulated spread. |
| **ATTACK-007** | Exfiltration Path Simulator | PhoenixSim | Network flow analysis across different exfiltration protocols. |
| **ATTACK-008** | Insider Scenario Generator | PhoenixSim | Simulates behavioral drift from normal process baselines. |
| **ATTACK-009** | Campaign Replay Framework | PhoenixRange | Multi-stage trace execution for training and evaluation. |
| **ATTACK-010** | Recovery Benchmark Suite | PhoenixRedLab | Automated verification of "Time-to-Safe" metrics after containment. |

## 2. Environment Readiness

- [ ] **Lab Mode:** Basic script execution for emulation.
- [ ] **Training Mode:** Interactive scenario orchestration in PhoenixRange.
- [ ] **Replay Mode:** High-fidelity trace loading for forensic analysis.
- [ ] **Simulation Mode:** Noise injection and saturation testing in PhoenixSim.
- [ ] **Validation Mode:** Automated AXIO-compliance checking.

## 3. Strict Boundary Checks
- [ ] PR Audit: Ensure no research code leaks into `phoenix_os/`.
- [ ] Binary Scan: Verify production artifacts contain zero emulation logic.
