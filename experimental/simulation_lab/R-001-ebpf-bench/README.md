R-001 — eBPF Capture Benchmark

Goal:
Measure agent overhead, event loss, and throughput for syscall/file/network probes.

Artifacts:
- `run_benchmark.sh` — harness to generate synthetic load and collect metrics
- `results/` — structured results (JSON) with hashes for reproducibility
- `perf/` — perf traces and BPF logs

Acceptance gates:
- CPU overhead < 5% baseline
- Event loss < 2%
- 3 reproducible runs with same result hash

Next steps:
- Implement `run_benchmark.sh` stub
- Add CI job to run in a controlled VM/container
