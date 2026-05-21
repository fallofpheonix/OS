# Stage_09_eBPF_and_Telemetry: eBPF and Telemetry

- Classification: IMPLEMENTATION
- Progression: Observe -> Extend
- Scope: eBPF, probes, audit, runtime event schema, process graph
- Prerequisites: Stage_01_System_Internals, Stage_02_Linux_and_Distro, Stage_04_Security_Fundamentals, Stage_08_Observability
- Next stage: Stage_10_SOC_Stack

## Source Files

| Item | Current location | Status | Type |
| --- | --- | --- | --- |
| Kernel Telemetry Layer | docs/research/11-kernel-telemetry-layer.md | RESEARCH_ONLY | Learning/Research |
| telemetry | telemetry | MISSING_INPUT_ROOT | Directory |

## Topics

| Topic | Difficulty | Hours | Weeks | Kind | Risk | Priority | Blocking items |
| --- | --- | --- | --- | --- | --- | --- | --- |
| eBPF tracing | High | 60 | 6 | Coding/Integration | Medium | P1 | Stage gate incomplete |
| kprobes/uprobes/tracepoints | High | 60 | 6 | Coding/Integration | Medium | P1 | Stage gate incomplete |
| Syscall telemetry | High | 60 | 6 | Coding/Integration | Medium | P1 | Stage gate incomplete |
| Auditd comparison | High | 60 | 6 | Coding/Integration | Medium | P1 | Stage gate incomplete |
| Process graph | High | 60 | 6 | Coding/Integration | Medium | P1 | Stage gate incomplete |
| Falco/Tracee model | High | 60 | 6 | Coding/Integration | Medium | P1 | Stage gate incomplete |
| Perf overhead | High | 60 | 6 | Coding/Integration | Medium | P1 | Stage gate incomplete |
| BPF verifier safety | High | 60 | 6 | Coding/Integration | Medium | P1 | Stage gate incomplete |

## Gate Conditions

- All prerequisite stages complete: Stage_01_System_Internals, Stage_02_Linux_and_Distro, Stage_04_Security_Fundamentals, Stage_08_Observability

- Source documents reviewed and contradictions logged

- Labs produce reproducible artifacts

- Risk items have owner and mitigation

- Outputs indexed in root reports
