Experiment ID: R-003

Objective:
Accurately reconstruct parent-child process lineage and produce an incident DAG for timeline replay and causal analysis.

Threat model:
Adversary uses process injection, forks, and service creation to hide execution paths; process lineage must be robust to PID reuse and namespaces.

Assets:
- Process lifecycle telemetry (fork, execve, clone, exit)
- cgroups / namespaces context
- Binary hashes and command-lines

Attack path:
- Initial execution → fork/exec sequences → background service creation → lateral tooling execution

Telemetry required:
- Syscalls: fork/clone/execve, setns, ptrace, prctl
- PID/PPID mapping with boot-time monotonic IDs
- Namespace and cgroup identifiers

Inputs:
- Malware samples exhibiting complex process trees
- Legitimate process churn workloads

Expected outputs:
- Incident DAG export (nodes: process events, edges: parent-child)
- Timeline replayable format

Metrics:
- Reconstruction completeness (percent of true edges found)
- False edges introduced
- Resilience to PID reuse and container namespaces

Validation gates:
- ≥95% edge recall on controlled traces
- Reproducible DAGs across 3 runs

Evidence:
- Exported DAG files, raw syscall traces, mapping tables

Failure conditions:
- Missing critical edges, corrupted timelines, non-reproducible outputs

Pilot mapping:
SOC incident graph generation; ransomware and lateral movement scenarios

Next integration target:
Incident graph database (Neo4j) ingestion pipeline
