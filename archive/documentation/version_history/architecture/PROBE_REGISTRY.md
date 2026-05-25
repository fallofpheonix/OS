# PhoenixOS Kernel Probe Registry

This document tracks the eBPF sensors deployed in the PhoenixOS kernel (L2).

## Active Probes (Phase 2A)

| Probe Type | Syscall | Purpose | Status |
| :--- | :--- | :--- | :--- |
| **Process** | `execve` | Capture process execution and command lines. | PLANNED |
| **Process** | `fork` | Capture child process creation. | PLANNED |
| **Process** | `clone` | Capture thread and process cloning. | PLANNED |
| **Process** | `exit` | Capture process termination. | PLANNED |
| **I/O** | `open` | Monitor file access and descriptor creation. | PLANNED |
| **I/O** | `read` | Monitor data ingestion from files/sockets. | PLANNED |
| **I/O** | `write` | Monitor data exfiltration/modification. | PLANNED |
| **Network** | `connect` | Capture outgoing network connection attempts. | PLANNED |
| **Network** | `accept` | Capture incoming network connection successes. | PLANNED |
| **Network** | `bind` | Monitor local port listening. | PLANNED |

## Configuration
- **Hook Method:** Kprobes / Tracepoints
- **Data Transfer:** eBPF Perf Buffer / Ring Buffer
- **Filtering:** In-kernel PID/UID filtering to prevent noise.
