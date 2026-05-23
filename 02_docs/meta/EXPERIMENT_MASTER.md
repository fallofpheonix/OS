# Experiment Status Map

| R ID | Objective | Subsystem | Status | Validation Gate | Runtime Mapping |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **R001** | eBPF File Capture Performance | Telemetry | READY | 1% CPU overhead | `09_telemetry/ebpf` |
| **R002** | Entropy Math Calculation | Telemetry | READY | <5us per block | `09_telemetry/entropy_engine`|
| **R003** | Process DAG Latency | Graph | PASSED | <1ms extraction | `09_telemetry/process_lineage` |
| **R004** | SMB Protocol Emulator | Tools | READY | Emulates ransomware | `05_tools/malware` |
| **R021** | Shannon Entropy Validation | Telemetry | READY | 99% accuracy | `09_telemetry/entropy_engine`|
| **R022** | Lineage Graph Extraction | Graph | PASSED | 100k nodes / 16MB | `09_telemetry/process_lineage`|
| **R023** | Containment Cost Opt. | Control | READY | 5% optimality | `07_security/control` |
| **R024** | SDI Workload Classifier | Physics | PASSED | 0.6931 Max Ent | `07_security/physics` |
| **R026** | Wavelet Processing | Signal Proc| READY | 94% accuracy | `09_telemetry/math_filters` |
| **R027** | Minimax Matrix Allocator | Game Theory | PASSED | 1.21 Avg Utility | `07_security/game` |
| **R031** | PID Actuation Damping | Control | PASSED | 99% accuracy | `agents/internal/control` |
| **R032** | Game-Aware Scheduler | Kernel | PASSED | Deterministic | `agents/internal/kernel` |
| **R036** | End-to-End Game Theory -> Actuator Pipeline Benchmark | Integrated | PASSED | <50ms loop | `agents/tests/integration_test.go` |