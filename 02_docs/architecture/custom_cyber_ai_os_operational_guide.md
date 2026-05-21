# Custom Cyber AI OS: Operational Guide

## Classification

- Domain: cybersecurity platform + AI/ML security engine + telemetry + SOC + hybrid OS.
- Lifecycle stage: Stage 19 Production Platform.
- Document type: operational architecture.
- Status: draft.
- Implementation status: planned.
- Critical dependencies:
  - Stage 09 eBPF and Telemetry.
  - Stage 10 SOC Stack.
  - Stage 11 AI/ML Core.
  - Stage 12 Security AI.
  - Stage 14 Automation and SOAR.
  - Stage 16 Kernel Extensions.
  - Stage 17 Hybrid OS.

## Corrections To Treat As Requirements

- “Never misses” is invalid as a system guarantee. Replace with measured recall target and false-negative budget.
- “Millions of events/sec” is invalid until proven by load tests. Current planning target is `>=10k events/sec/node`, then scale horizontally.
- Automated containment must be gated by blast-radius controls, rollback, audit logging, and human-approval policy.
- eBPF logic must stay verifier-safe and bounded. Complex policy and ML inference belong in userspace.
- Forensic collection can be privacy-sensitive and destructive if misused. Preserve chain of custody and access audit.
- Kernel configuration options are version-dependent. Every hardening option must be verified with `scripts/config`, `olddefconfig`, or `merge_config.sh` against the selected kernel tree.
- Custom LSM hooks are normally built into the kernel, not loaded as arbitrary runtime modules. Treat LSM work as kernel-build work unless the target kernel explicitly supports the desired hook path.
- Full physical memory acquisition through `/dev/mem` is commonly restricted. Prefer LiME, AVML, crash dumps, or hypervisor snapshots depending on deployment model.

## Pre-Build Readiness Checkpoint

Do not begin construction until each prerequisite can be explained, implemented without references, debugged under production failure, and taught to another engineer.

| Competency | Required evidence |
|---|---|
| C programming | pointers, allocation, file I/O, sockets, syscalls, low-level data structures |
| Rust programming | ownership, borrowing, unsafe FFI, async, `Result`/`Option` error paths |
| Linux internals | processes, scheduling, virtual memory, VFS, syscalls, modules |
| Networking | TCP/IP layers 2-4, sockets, packet capture, TLS basics, DNS |
| Security foundations | STRIDE/PASTA, authentication, authorization, crypto, RBAC, defense in depth |
| eBPF | program types, maps, verifier constraints, ring buffers, XDP, kprobes, tracepoints |
| Reverse engineering | disassembly, decompilation, Ghidra/radare2, dynamic analysis, malware patterns |
| Forensics and IR | memory analysis, timelines, artifact preservation, chain of custody, rootkit detection |
| ML fundamentals | feature engineering, supervised/unsupervised learning, deep learning, evaluation, deployment |
| Attack paths | ATT&CK mapping, kill chain, risk assessment, incident classification |

Failure condition:

```text
If any competency fails the four-part verification, return to the relevant research phase before implementation.
```

## Target Layer Model

```text
User-space security tools
-> AI threat engine: UEBA, malware detection, threat prediction, scoring
-> SOC stack: alert correlation, enrichment, playbooks, SIEM/SOAR interfaces
-> forensics platform: memory, artifacts, timeline, chain of custody
-> eBPF monitoring: syscalls, network, file access, process tracking
-> telemetry collector: aggregation, buffering, filtering, formatting, forwarding
-> security kernel extensions: LSM, MAC policy, integrity, audit hooks
-> hardened Linux kernel: KSPP-oriented config, eBPF, crypto, audit
-> firmware and hardware
```

Core invariant:

```text
kernel path observes or enforces bounded policy
userspace performs enrichment, correlation, ML, and response orchestration
```

## Complete Data Flow

```text
Hardware events
-> kernel events: syscalls, packets, filesystem ops, signals
-> kernel extensions: LSM, seccomp, policy hooks
-> eBPF monitoring: syscall, XDP, filesystem, process probes
-> ring buffers
-> userspace telemetry collector
-> parser and validator
-> context enrichment
-> AI/ML scoring
-> rule evaluation
-> alert generation
-> correlation engine
-> incident enrichment
-> decision engine
-> response executor
-> feedback and learning loop
```

## Kernel Extensions

Responsibilities:

- Enforce mandatory access control.
- Block unauthorized operations at kernel level where policy is deterministic.
- Provide real-time policy hooks.

Candidate mechanisms:

- LSM hooks.
- Seccomp.
- Linux capabilities.
- Namespaces and cgroups.
- Kernel module or built-in extension where justified.

Constraints:

- No unbounded work in kernel path.
- No ML inference in kernel.
- No blocking network calls in kernel.
- Failure mode must default to documented policy:
  - fail-closed for critical enforcement.
  - fail-open for observability-only probes where availability is higher priority.

## eBPF Monitoring Layer

Probe classes:

- Syscall tracer: captures syscall number, arguments, return values, PID/TID, UID/GID, namespace IDs.
- XDP network monitor: captures packet metadata at driver level.
- Filesystem monitor: tracks file access and mutation patterns.
- Process tracker: captures fork, exec, exit, credential changes, namespace transitions.

Output:

- Ring-buffer events.
- Per-CPU counters.
- Optional BPF maps for bounded aggregation.

Constraints:

- Keep BPF programs small.
- Use bounded loops only.
- Avoid complex correlation in BPF.
- Benchmark every probe for overhead.
- Drop or sample low-priority events when backpressure occurs.

## Telemetry Collector

Responsibilities:

- Read eBPF ring buffers through mmap-backed zero-copy paths.
- Decode binary event structures.
- Validate event schema version.
- Enrich with process, user, asset, namespace, container, and threat-intel context.
- Compress buffered events with zstd above configured thresholds.
- Buffer locally using circular storage.
- Forward to local processing and/or remote SIEM.

Initial targets:

| Metric | Target |
|---|---|
| Per-event ring-buffer read latency | `<100 us` |
| Local buffer size | `1 GiB` default |
| Event loss policy | priority drop after backpressure threshold |
| Offline mode | local processing continues |

## Event Processing Pipeline

Stages:

1. Parsing and validation.
2. Context enrichment.
3. Parallel anomaly detection.
4. Threat scoring.
5. Rule evaluation.
6. Alert generation.

Threat score formula:

```text
threat_score = (ueba_score * 0.40)
             + (malware_score * 0.35)
             + (anomaly_score * 0.25)
```

Normalize to `0-100`.

Default severity:

| Score | Severity | Action |
|---:|---|---|
| `<30` | Low | archive/suppress unless correlated |
| `30-59` | Medium | review |
| `60-79` | High | alert and enrich |
| `>=80` | Critical | incident path and containment review |

Alert fields:

- Severity.
- Confidence.
- Recommended action.
- Supporting evidence.
- Model contributions.
- Rule hits.
- Raw event references.
- Chain-of-custody pointer if evidence collected.

## Correlation Engine

Inputs:

- Individual alerts.
- Enriched events.
- Threat-intelligence hits.
- Historical incident context.

Correlation dimensions:

- PID/TID.
- Process tree.
- User.
- Host.
- Asset.
- Source IP.
- Destination IP.
- File hash.
- Domain.
- Container/pod.
- Time window.

Default windows:

- Deduplication: `<5 sec`.
- Related-event correlation: `<60 sec`.
- Lateral movement: `<10 min`.
- Data exfiltration: `<30 sec` between sensitive read and external transfer.

Core patterns:

| Rule | Conditions | Action |
|---|---|---|
| Process injection | parent spawns higher-privilege child within window | elevate severity |
| Lateral movement | endpoint external connection followed by internal remote login | critical incident |
| Data exfiltration | sensitive file read followed by external upload | isolate or block pending policy |
| C2 communication | known C2 plus beaconing or encrypted high-entropy flow | block, collect, alert |
| Privilege escalation | non-admin process reaches UID 0 path | escalate and investigate |

## Incident Investigation And Enrichment

Automated collection for high severity:

- Process memory dump.
- Open files.
- Network connections.
- Syscall history.
- Process environment.
- Command line.
- Loaded modules/libraries.
- Parent/child process tree.

Required controls:

- SHA256 hash all artifacts.
- Record acquisition tool and version.
- Record collector identity.
- Record timestamp source.
- Store access audit.
- Preserve chain of custody.

Enrichment:

- Historical process/user/IP/domain/file-hash search.
- Threat-feed lookup.
- MITRE ATT&CK mapping.
- Asset criticality.
- Data exposure estimate.
- Attack-chain reconstruction.

## Decision Engine

Decision matrix:

| Severity | Asset type | Default action | Target urgency |
|---|---|---|---|
| Critical | Critical | isolate after policy approval or pre-approved condition | `<5 min` |
| Critical | Non-critical | investigate, contain if confirmed | `<30 min` |
| High | Critical | investigate and prepare containment | `<1 hour` |
| High | Non-critical | monitor and enrich | `<4 hours` |
| Medium | Any | analyst review | `<24 hours` |
| Low | Any | archive or suppress | none |

Playbooks:

- Malware response.
- Privilege escalation response.
- Data breach response.
- Lateral movement containment.
- C2 communication response.
- False-positive storm response.

Automation constraints:

- Destructive actions require explicit policy.
- Every action must be idempotent or have rollback.
- Every action must produce audit records.
- High-risk actions need approval unless pre-authorized.

## Response Execution

Actions:

- Endpoint isolation.
- Process termination.
- Credential revocation.
- Network blocking.
- DNS sinkholing.
- Hash blocking.
- Forensic preservation.
- Human notification.
- Ticket creation.

Execution order for critical malware:

```text
1. Preserve evidence.
2. Isolate endpoint or block C2 path.
3. Stop malicious process if policy permits.
4. Revoke credentials when compromise is likely.
5. Propagate IOCs.
6. Notify humans.
7. Open incident record.
```

Failure handling:

- Retry bounded transient failures.
- Escalate permanent failures.
- Record partial completion.
- Do not hide failed response actions behind final success states.

## Feedback And Learning

Inputs:

- Incident resolution.
- Analyst labels.
- False positive/false negative outcomes.
- Detection latency.
- Response latency.
- Playbook success/failure.

Outputs:

- Model retraining candidates.
- Threshold changes.
- Rule changes.
- Probe additions/removals.
- Threat-intelligence updates.
- Playbook changes.

Model update rule:

- False positive: add benign examples, check feature leakage, raise threshold only with precision/recall review.
- False negative: add malicious examples, add detection rule, lower threshold only with alert-volume review.
- Model deployment requires offline validation and canary rollout.

## Quality Metrics

Detection:

| Metric | Target |
|---|---|
| True positive rate | `>95%` |
| False positive rate | `<5%` |
| Precision | `>90%` |
| Recall | `>95%` |

Operations:

| Metric | Target |
|---|---|
| MTTD | `<5 min` |
| MTTR | `<10 min` for containment-ready incidents |
| Critical alerts/day | `<100` |
| Total alerts/day | `<1000` |
| Analyst time/alert | `<15 min` |

System:

| Metric | Target |
|---|---|
| Uptime | `>99.9%` |
| p95 processing latency | `<100 ms` |
| Throughput | `>10k events/sec/node` |
| CPU usage | `<40%` baseline |
| RAM usage | `<50%` baseline |

## Latency Budget

Initial planning budget:

| Component | Target latency |
|---|---:|
| Kernel event | `0 us` |
| eBPF capture and ring-buffer write | `2-5 us` |
| Userspace read | `10-20 us` |
| Parsing and validation | `50-100 us` |
| Context enrichment | `0.5-2 ms` |
| ML inference | `10-50 ms` |
| Anomaly scoring | `2-5 ms` |
| Rule evaluation | `1-2 ms` |
| Alert generation | `1 ms` |

Budget target: `<100 ms p95`, `<500 ms p99` before response execution.

## Capacity Model

Initial single-node assumptions:

| Component | Planning target |
|---|---:|
| eBPF syscall capture | `100k events/sec` |
| Ring-buffer throughput | `200 MB/sec` |
| Telemetry parsing | `50k events/sec` |
| ML inference CPU | `10k events/sec` |
| ML inference GPU/quantized | `25k-50k events/sec` |
| Alert generation | `5k alerts/sec` |

Expected bottleneck:

- ML inference.

Scaling options:

- GPU inference.
- INT8 quantization.
- Per-node local inference.
- Horizontal event partitioning.
- Sampling or tiered scoring.
- Rules-first prefiltering before model inference.

## Integration Interfaces

| Interface | Method | Format | Notes |
|---|---|---|---|
| eBPF -> telemetry collector | ring buffer | binary event struct | zero-copy path |
| telemetry collector -> AI engine | queue or stream | protobuf preferred, JSON acceptable | disk fallback |
| AI engine -> correlation engine | event stream | enriched event with scores | shared correlation window |
| correlation -> decision engine | alert queue | alert object | playbook metadata |
| decision -> response executor | task queue | playbook action list | approval and status required |
| response -> EDR | REST/mTLS | action request | isolate, kill, collect |
| response -> firewall | REST/SSH/netconf | blocklist delta | reversible changes |
| response -> SIEM | HEC/REST | normalized event | rate-limited |
| threat intel -> enrichment | TAXII/REST/file | IOC/rules | local cache required |
| incident management | REST | ticket | incident timeline |
| dashboard | WebSocket/REST | incidents, alerts, metrics | human operation |

## Build Phases

### Phase 1: Security Distribution Foundation

Objectives:

- Build custom LTS Linux kernel.
- Apply and verify hardening configuration.
- Produce minimal hardened userspace.
- Enable secure boot or measured boot.
- Provide container-ready runtime for collectors and AI/SOC components.

Kernel hardening baseline:

```text
stack protector strong
fortify source
strict kernel/module RWX
KASLR
init-on-free
seccomp and seccomp-bpf
audit syscall support
kprobes and tracepoints
BPF syscall and JIT
AppArmor or SELinux
crypto primitives needed by platform
```

Build constraints:

- Pin exact kernel source version and checksum.
- Store generated `.config`.
- Diff final config against required options.
- Keep debug sanitizers such as KASAN/UBSAN out of production unless explicitly required; they are validation builds, not default production builds.
- Sign kernel artifacts when secure boot is enabled.

Userspace baseline:

- Minimal rootfs.
- C/Rust/LLVM toolchain for development image only.
- Runtime image stripped to required services.
- Audit, AIDE or equivalent integrity tooling.
- AppArmor/SELinux tools.
- eBPF loader dependencies.
- No telnet, rsh, unused GUI stack, avahi, or unnecessary daemons.

### Phase 2: Telemetry Layer

Objectives:

- Define stable event schema.
- Build collector daemon.
- Buffer locally under SIEM outage.
- Validate, enrich, compress, and forward events.

Event classes:

- Process execution.
- File access.
- Network connection.
- Privilege elevation.
- Authentication attempt.
- Syscall trace.
- Memory allocation.
- Module load.

Implementation constraints:

- Event schema version is mandatory.
- Every dropped event increments a typed counter.
- Backpressure policy is explicit: drop, sample, or block.
- Collector must expose health, lag, loss, and throughput metrics.

### Phase 3: eBPF Monitoring Layer

Objectives:

- Capture syscall, packet, filesystem, and process telemetry with bounded overhead.
- Keep policy and ML outside BPF programs.
- Use BPF maps only for bounded state and aggregation.

Probe set:

- Raw syscall enter/exit tracer.
- Targeted socket/connect tracer.
- XDP packet metadata monitor.
- Filesystem open/write/unlink/chmod monitor.
- Process fork/exec/exit and credential-change monitor.

Failure constraints:

- Probe load failure falls back to reduced monitoring.
- Verifier rejection is a deployment failure.
- Probe overhead must be benchmarked before production enablement.

### Phase 4: SOC Stack Integration

Objectives:

- Forward normalized events to SIEM.
- Correlate related alerts.
- Enrich with asset, user, and threat-intel context.
- Execute bounded SOAR playbooks.
- Provide investigation timeline and impact estimate.

Required components:

- Event forwarder with retry and disk-backed buffering.
- Correlation engine with explicit time windows.
- Playbook engine with dry-run mode.
- Investigation context builder.
- Ticketing and notification integration.

### Phase 5: Forensics Platform

Objectives:

- Preserve evidence before destructive response.
- Collect process memory, open files, connections, syscall history, command line, environment, and loaded modules.
- Reconstruct timeline.
- Preserve chain of custody.

Required controls:

- Hash every artifact with SHA256 or stronger.
- Record acquisition timestamp and source clock.
- Record collector identity and tool version.
- Store access audit.
- Separate evidence storage from operational logs.

### Phase 6: AI Threat Engine

Objectives:

- Score events with UEBA, malware, and anomaly models.
- Produce explainable risk score.
- Feed correlation and decision engines.

Threat score:

```text
score = 0.40 * ueba + 0.35 * malware + 0.25 * anomaly
```

Controls:

- Every prediction records model version and feature schema version.
- Unknown or invalid features trigger fallback scoring.
- Model confidence is separate from severity.
- Model update requires offline validation and canary rollout.

### Phase 7: Kernel Extensions And Security Modules

Objectives:

- Enforce deterministic policy in LSM/seccomp/capability layer.
- Log security-sensitive hooks.
- Apply process-level syscall policy.

Constraints:

- LSM changes require kernel integration planning.
- Seccomp filters must be per-workload and tested against normal behavior.
- Enforcement policy must define fail-open versus fail-closed behavior.
- Runtime protections must not depend on network availability.

### Phase 8: Integration Testing

Required tests:

- End-to-end alert generation from syscall to dashboard.
- eBPF probe reload and failure recovery.
- SIEM outage buffering and replay.
- Model inference timeout and fallback to rules.
- Playbook dry-run and approval path.
- Forensic collection and chain-of-custody validation.
- False-positive storm handling.

### Phase 9: Deployment And Hardening

Required controls:

- Multi-node deployment test.
- High availability for collector, inference, and correlation services.
- Backup and restore procedures.
- Secure boot or measured boot validation.
- Key rotation procedure.
- Patch and kernel upgrade process.
- Production monitoring and on-call runbooks.

## Failure Modes

| Failure | Detection | Recovery | Impact |
|---|---|---|---|
| eBPF program failure | stale ring buffer, missing heartbeat | reload probes; fallback to audit/strace where feasible | monitoring gap |
| ML inference hang | watchdog, latency SLO breach | kill worker, fallback to rules, restart service | reduced detection quality |
| telemetry disk full | write failure, disk watermark | compress, offload, drop low-priority events | possible low-priority data loss |
| SIEM offline | network timeout | local processing, buffer, backoff retry | no central visibility |
| false-positive storm | alert rate threshold | raise temporary threshold, require analyst review | alert fatigue |
| missed ransomware | encrypted files or ransom note | post-mortem, new rules, retrain, backup recovery | incident response/data loss |

## Ransomware Scenario Flow

Attack:

```text
Phishing email
-> user downloads executable
-> user executes payload
-> C2 connection
-> child process spawn
-> payload drop
-> attempted encryption
```

Detection:

```text
execve event
-> context enrichment
-> UEBA score
-> malware score
-> anomaly score
-> threat score
-> correlated incident
```

Containment:

```text
collect evidence
-> isolate endpoint
-> terminate process if policy permits
-> block C2
-> revoke credentials
-> notify incident responders
-> open incident ticket
```

Investigation questions:

- How did the attacker reach the user?
- Was anything encrypted before containment?
- Did lateral movement occur?
- Were other users targeted?
- Was data exfiltrated?
- Which indicators should be promoted into detection logic?

## Production Deployment

Pre-deployment:

- Components compiled and tested.
- Load tests pass at `>=10k events/sec/node`.
- Failover tested.
- Security review complete.
- No hardcoded secrets.
- Architecture, runbooks, and playbooks complete.
- Incident response team trained.

Deployment:

1. Test network.
2. Production read-only mode.
3. Non-disruptive automation.
4. Staged blocking and isolation.
5. Full production mode.

Single-node deployment order:

```text
build and install custom kernel
-> deploy telemetry collector
-> load eBPF probes
-> start AI threat engine
-> initialize SOC stack
-> run health check
-> enable alerting
-> enable response only after approval policy is loaded
```

Enterprise deployment model:

- Agent runs as privileged DaemonSet or host service only where kernel telemetry is required.
- `hostPID`, `hostNetwork`, `/sys`, and debugfs access are high-risk privileges and require node pool isolation.
- Threat engine can run separately from node agent if event latency budget allows.
- SIEM endpoint, keys, and certificates must come from secrets manager, not static manifests.
- Production deployment needs network policies limiting egress to SIEM, model, update, and management endpoints.

Benchmark commands:

```text
perf stat -e cycles,instructions,cache-misses <collector>
bpftool prog profile <prog-id>
bpftool map show
ps -o pid,ppid,pcpu,pmem,rss,args -C telemetry-collector
```

Tuning defaults:

| Parameter | Initial value | Notes |
|---|---:|---|
| Telemetry local buffer | `1 GiB` | increase for disconnected nodes |
| eBPF ring buffer | `256-512 MiB` | size by event burst rate |
| Compression | `zstd` | compress events above `1 KiB` |
| Anomaly threshold | `0.70` | tune from precision/recall review |
| Correlation window | `60 sec` | override per attack pattern |
| Playbook timeout | `30 sec` | bounded retry only |

Ongoing operations:

- Daily alert and health review.
- Weekly model/rule threshold review.
- Monthly patches and disaster recovery drill.
- Quarterly red team and architecture review.

Maintenance cadence:

| Cadence | Tasks |
|---|---|
| Daily | alert trend review, high-severity review, resource checks |
| Weekly | model drift review, rule tuning, playbook success review |
| Monthly | security patches, kernel upgrade review, model validation, recovery drill |
| Quarterly | red team exercise, penetration test, architecture review, threat landscape update |

## Production Readiness Exit Criteria

- [ ] Readiness competencies verified.
- [ ] Kernel source, config, patches, and signatures are reproducible.
- [ ] Rootfs package manifest is minimized and version-pinned.
- [ ] eBPF programs pass verifier and overhead tests.
- [ ] Telemetry loss, backpressure, and buffering policies are tested.
- [ ] SIEM/SOAR integration works under outage and replay.
- [ ] Forensic artifacts preserve chain of custody.
- [ ] AI models have model cards, metrics, feature schema versions, and rollback.
- [ ] Playbooks run in dry-run mode before enforcement.
- [ ] Destructive response actions require policy approval and rollback.
- [ ] End-to-end ransomware scenario test passes.
- [ ] Backup, restore, key rotation, and disaster recovery procedures pass.

## Open Design Questions

- What exact event schema is stable enough for model training?
- Which response actions are allowed without human approval?
- What is the legal boundary for memory dumps and user data capture?
- Which workloads require fail-closed enforcement?
- What is the minimum local-only mode when disconnected from SIEM?
- What is the retention policy for raw events, enriched events, alerts, incidents, and forensic artifacts?
- Which model types are acceptable for explainability and audit?
