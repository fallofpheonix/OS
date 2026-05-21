Experiment ID: R-002

Objective:
Develop and validate a file-activity entropy detector for ransomware-like write patterns (mass rename, high-entropy writes, rapid extension changes).

Threat model:
Ransomware attempts to encrypt files quickly while avoiding detection; detection should trigger pre-encryption containment.

Assets:
- Host filesystem events
- File content sampling (entropy calculation)
- Process lineage

Attack path:
- Initial compromise → process executes encryption payload → mass file writes and renames

Telemetry required:
- File create/rename/write events
- Per-file write sizes and timing
- File sample bytes for entropy
- Process PID/cmdline and parent lineage

Inputs:
- Known ransomware samples (test lab) and benign file churn workloads
- Synthetic rename/extension burst scripts

Expected outputs:
- Entropy model & thresholds
- Detection rule with confidence score
- Forensic evidence: list of modified files, sample hashes, timeline

Metrics:
- Detection precision & recall on labeled runs
- Time-to-detect pre-encryption window
- False positive rate under benign churn

Validation gates:
- Detect encryption activity before >5% of target filesystem encrypted in 90% of runs
- Precision and recall meet pilot thresholds (configurable)
- Evidence snapshot (RAM/disk) captured on trigger

Evidence:
- File lists, entropy logs, sample hashes, snapshot artifacts

Failure conditions:
- High false positives during backups or benign churn
- Too-late detection after large portion encrypted

Pilot mapping:
Hospital ransomware pilot

Next integration target:
Automatic filesystem freeze + snapshot playbook
