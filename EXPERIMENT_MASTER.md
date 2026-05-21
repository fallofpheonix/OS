# Experiment Status Map

| R ID | Objective | Subsystem | Status | Validation Gate | Runtime Mapping |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **R001** | eBPF File Capture Performance | Telemetry | READY | 1% CPU overhead | `09_telemetry/ebpf` |
| **R002** | Entropy Math Calculation | Telemetry | READY | <5us per block | `09_telemetry/entropy_engine`|
| **R003** | Process DAG Latency | Graph | BLOCKED | Requires Phoenix Bus | `09_telemetry/process_graphs` |
| **R004** | SMB Protocol Emulator | Tools | READY | Emulates ransomware | `05_tools/malware` |
| **R021** | Shannon Entropy Validation | Telemetry | READY | 99% accuracy | `09_telemetry/entropy_engine`|
| **R022** | Lineage Graph Extraction | Graph | BLOCKED | Requires R003 | `09_telemetry/process_graphs`|
| **R023** | Containment Cost Opt. | Control | READY | 5% optimality | `07_security/control` |
| **R024** | SDI Workload Classifier | Physics | READY | 95% pred. rate | `07_security/physics` |
| **R026** | Wavelet Processing | Signal Proc| READY | 94% accuracy | `09_telemetry/math_filters` |
| **R027** | Minimax Matrix Allocator | Game Theory | READY | <0.5ms eq calc | `07_security/game` |
| **R031** | PID Actuation Damping | Control | BLOCKED | Requires R023 | `07_security/control` |
| **R032** | Game-Aware Scheduler | Kernel | BLOCKED | Requires R031 | `10_kernel/scheduler` |

## Missing Experiments Identified & Added
*   **R035**: Validation of Event Normalizer Latency.
*   **R036**: End-to-End Game Theory -> Actuator Pipeline Benchmark.
*   **R037**: False-Positive Replay Engine for Telemetry Datasets.