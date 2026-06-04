---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Security — Data Flow

> Last verified: 2026-06-04

This document defines violation and containment control flow data paths.

## Containment Decision Flow

```mermaid
graph TD
    Audit[Syscall Audit Event] -->|Read| TrustMatrix[trust_matrix.go]
    TrustMatrix -->|Calculate Trust Score| SDI[State Disorder Index Calculator]
    SDI -->|Evaluate Thresholds| Warden[warden.go Engine]
    Warden -->|Select Strategy| Solver[Stackelberg Solver]
    Solver -->|Determine Escalation Level| Actuator[ebpf.go & process.go Actuators]
    Actuator -->|Terminate / Sandbox| Target[Target Namespace]
```

## Violation Forensic Log Flow
- Security violations are passed to `jsonl_writer.go`.
- The writer creates JSONL-structured logs describing event IDs, violation signatures, and timestamps.
