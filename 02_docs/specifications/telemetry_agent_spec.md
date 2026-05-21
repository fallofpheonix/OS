Telemetry Agent Module Spec

Purpose:
Collect kernel-level telemetry and forward normalized events onto the Phoenix Bus.

Inputs:
- eBPF probes: syscall, file ops, network connect
- Agent config

Outputs:
- Normalized JSON events
- Local artifacts (pcap, snapshots)

Dependencies:
- Kernel headers
- libbpf / BCC
- Phoenix Bus (Kafka/Redis/HTTP)

Metrics:
- Events/sec
- CPU and memory usage
- Event loss rate

CPU budget: 5% at baseline workload
Memory budget: 200MB resident for agent

Tests:
- R-001 benchmarking harness
- Event schema unit tests

Artifacts:
- Event schema docs
- Perf traces
