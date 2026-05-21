Experiment ID: R-001

Objective:
Measure eBPF syscall/event capture throughput and operational overhead to validate kernel telemetry baseline.

Threat model:
Attackers may hide activity by overwhelming telemetry or exploiting sampling gaps; we must ensure low-loss capture with acceptable overhead.

Assets:
- Host kernel telemetry
- eBPF capture shim
- Event bus ingestion pipeline

Attack path:
- Stealthy malware generates many syscalls or I/O to try to evade sampling/loss.

Telemetry required:
- Raw syscall events (open, read, write, execve, fork, clone, connect)
- Timestamps, PIDs, TTY, mount/ns context
- Event batching and drop counters

Inputs:
- Synthetic syscall generator (benign workload)
- High-frequency syscall blast generator (adversarial)
- Representative application workload (web server, DB)

Expected outputs:
- CPU cost profile (user/kernel)
- Memory usage of capture agent
- Event loss rate under load
- Replayable event logs (pcap/JSON)

Metrics:
- CPU overhead (target <5% for baseline profile)
- Memory delta
- Event loss <2% at target throughput
- 3 reproducible runs with variance <5%

Validation gates:
- CPU <5% on representative host
- Loss <2% across 3 runs
- Replays produce identical DAGs for same input

Evidence:
- Benchmark logs, perf traces, event-count histograms
- Hashes of replay logs

Failure conditions:
- Overhead exceeds threshold
- Event loss >2% or non-deterministic replays

Pilot mapping:
Hospital ransomware (telemetry reliability for pre-encryption detection)

Next integration target:
Normalized event bus schema and sampling policy module
