Experiment ID: R-004

Objective:
Build an SMB lateral movement emulator to simulate propagation, session spikes, credential reuse, and remote execution for validation of detection and containment.

Threat model:
Adversary uses SMB to move laterally via credential theft and remote command execution.

Assets:
- SMB service endpoints
- Host authentication logs
- Network packet capture

Attack path:
- Credential harvest → SMB auth attempts → remote execution (psexec-like) → payload deployment

Telemetry required:
- SMB session events, auth failures/successes
- Process creation on target hosts
- Network connection metadata and pcaps

Inputs:
- Synthetic credential lists
- Remote execution payloads (benign test agents)

Expected outputs:
- Host spread map, session spike metrics, MTTD/recall metrics
- Forensic traces per host

Metrics:
- MTTD for lateral hop
- Recall of lateral movement detection
- Host spread count over time

Validation gates:
- Emulator reproduces propagation patterns deterministically
- Detection achieves configured recall on lab runs

Evidence:
- PCAPs, host logs, event bus captures, containment actions

Failure conditions:
- Non-deterministic propagation, noisy benign-based triggers

Pilot mapping:
Hospital pilot (internal spread), enterprise SOC testing

Next integration target:
Automated cross-host containment playbook and credential alerting
