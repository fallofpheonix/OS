# CyberAI-OS Long-Term Structure

## Rule

`01_research/` is immutable research material. Do not place implementation code there.

Implementation work belongs in:

- `06_ai/`
- `07_security/`
- `08_forensics/`
- `09_telemetry/`
- `10_kernel/`
- `11_distribution/`
- `12_infrastructure/`
- `14_experiments/`

Project control belongs in `00_program_management/`.

Architecture, RFCs, decisions, schemas, diagrams, and specifications belong in `02_docs/`.

Repository tracking belongs in `03_repositories/`.

Datasets belong in `04_datasets/`.

Tool inventories and local tool wrappers belong in `05_tools/`.

Open source contribution work belongs in `13_open_source/`.

Retired material belongs in `15_archive/`.

## Execution Path

```text
Stage 00-04:
Foundations -> OS Internals -> Linux/Distros -> Networking -> Security

Stage 05-10:
Reverse Engineering -> Forensics -> Threat Intel -> Observability -> Telemetry -> SOC

Stage 11-14:
AI/ML -> Security AI -> Cloud -> Automation

Stage 15-18:
Security Distribution -> Kernel Extensions -> Hybrid OS -> Custom OS

Stage 19-20:
Production Platform -> Future Research
```

## Dependency Constraint

No implementation stage starts before its research prerequisites and gate conditions are complete.

Critical path:

```text
00_foundations
-> 01_system_internals
-> 02_linux_and_distros
-> 03_networking
-> 04_security
-> 08_observability
-> 09_telemetry
-> 10_soc
-> 11_ai_ml
-> 12_security_ai
-> 14_automation
-> 15_security_distribution
-> 16_kernel_extensions
-> 17_hybrid_os
-> 18_custom_os
-> 19_production
-> 20_future
```

## Migration Map

| Existing path | Long-term path | Action |
|---|---|---|
| `docs/` | `02_docs/` | Classify/copy stable docs into architecture, specs, decisions, schemas, diagrams |
| `research/` | `01_research/` | Classify/copy research only; keep immutable |
| `ai/` | `06_ai/` | Move implementation code only |
| `security/` | `07_security/` | Move IDS/IPS/EDR/SIEM/SOAR/detections implementation |
| `forensics/` | `08_forensics/` | Move DFIR implementation and reports |
| `telemetry/` | `09_telemetry/` | Move collectors, events, traces, eBPF telemetry |
| `kernel/` | `10_kernel/` | Move kernel implementation and prototypes |
| `drivers/` | `10_kernel/drivers/` | Move driver code |
| `network/` | `07_security/ids/` or `09_telemetry/collectors/` | Split by purpose |
| `userspace/` | `10_kernel/prototypes/` or `11_distribution/packages/` | Split by kernel/userspace ownership |
| `tools/` | `05_tools/` | Move tool wrappers and inventories |
| `repos/` | `03_repositories/` | Move repo catalogs |
| `papers/` | `01_research/` | Move papers into relevant stage |
| `datasets/` | `04_datasets/` | Move datasets by source type |
| `roadmap/` | `00_program_management/roadmap/` | Move planning docs |
| `notes/` | `01_research/*/notes/` or `00_program_management/reports/` | Split by content |
| `project_stages/` | `00_program_management/reports/` and `00_program_management/dependency_graphs/` | Keep generated classification as program control artifact |

## Stage Ownership

| Stage | Research owner path | Implementation owner path |
|---|---|---|
| Stage 00 Foundations | `01_research/stage_00_foundations/` | `05_tools/development/`, `14_experiments/poc/` |
| Stage 01 System Internals | `01_research/stage_01_system_internals/` | `10_kernel/prototypes/` |
| Stage 02 Linux and Distros | `01_research/stage_02_linux_and_distros/` | `11_distribution/` |
| Stage 03 Networking | `01_research/stage_03_networking/` | `09_telemetry/collectors/`, `07_security/ids/` |
| Stage 04 Security | `01_research/stage_04_security/` | `07_security/`, `10_kernel/sandbox/` |
| Stage 05 Reverse Engineering | `01_research/stage_05_reverse_engineering/` | `05_tools/reverse/`, `07_security/yara/` |
| Stage 06 Forensics | `01_research/stage_06_forensics/` | `08_forensics/` |
| Stage 07 Threat Intelligence | `01_research/stage_07_threat_intelligence/` | `07_security/hunting/`, `07_security/detections/` |
| Stage 08 Observability | `01_research/stage_08_observability/` | `12_infrastructure/observability/` |
| Stage 09 Telemetry | `01_research/stage_09_telemetry/` | `09_telemetry/` |
| Stage 10 SOC | `01_research/stage_10_soc/` | `07_security/siem/`, `07_security/soar/`, `07_security/response/` |
| Stage 11 AI/ML | `01_research/stage_11_ai_ml/` | `06_ai/` |
| Stage 12 Security AI | `01_research/stage_12_security_ai/` | `06_ai/agents/`, `06_ai/rag/`, `07_security/detections/` |
| Stage 13 Cloud | `01_research/stage_13_cloud/` | `12_infrastructure/containers/`, `12_infrastructure/kubernetes/` |
| Stage 14 Automation | `01_research/stage_14_automation/` | `05_tools/automation/`, `07_security/soar/` |
| Stage 15 Security Distribution | `01_research/stage_15_security_distribution/` | `11_distribution/` |
| Stage 16 Kernel Extensions | `01_research/stage_16_kernel_extensions/` | `10_kernel/hooks/`, `10_kernel/drivers/` |
| Stage 17 Hybrid OS | `01_research/stage_17_hybrid_os/` | `10_kernel/`, `11_distribution/` |
| Stage 18 Custom OS | `01_research/stage_18_custom_os/` | `10_kernel/` |
| Stage 19 Production | `01_research/stage_19_production/` | `12_infrastructure/deployment/`, `00_program_management/execution/` |
| Stage 20 Future | `01_research/stage_20_future/` | none until promoted |

## Active Foundation Modules

| Module | Research path | Gate | Implementation evidence path |
|---|---|---|---|
| Phase 0: Computer Science Foundations | `01_research/stage_00_foundations/notes/phase_00_computer_science_foundations.md` | `01_research/stage_00_foundations/checkpoints/phase_00_build_gate.md` | `14_experiments/poc/stage_00_foundations/` |
| Phase 1: Computer Architecture | `01_research/stage_00_foundations/architecture/phase_01_computer_architecture.md` | `01_research/stage_00_foundations/checkpoints/phase_01_build_gate.md` | `14_experiments/poc/stage_00_foundations/phase_01_architecture/` |
| Phase 2: Low-Level Programming | `01_research/stage_00_foundations/development/phase_02_low_level_programming.md` | `01_research/stage_00_foundations/checkpoints/phase_02_build_gate.md` | `14_experiments/poc/stage_00_foundations/phase_02_low_level_programming/` |
| Phase 6: Security Foundations | `01_research/stage_04_security/phase_06_security_foundations.md` | `01_research/stage_04_security/build_gate.md` | `07_security/`, `02_docs/threat_models/` |
| Phase 7: Reverse Engineering and Malware Analysis | `01_research/stage_05_reverse_engineering/phase_07_reverse_engineering_malware.md` | `01_research/stage_05_reverse_engineering/build_gate.md` | `05_tools/reverse/`, `07_security/yara/`, `08_forensics/reports/` |
| Phase 8: AI/ML Fundamentals | `01_research/stage_11_ai_ml/phase_08_ai_ml_fundamentals.md` | `01_research/stage_11_ai_ml/build_gate.md` | `06_ai/`, `04_datasets/` |
| Phase 9: Security AI | `01_research/stage_12_security_ai/phase_09_security_ai.md` | `01_research/stage_12_security_ai/build_gate.md` | `06_ai/agents/`, `06_ai/rag/`, `07_security/detections/`, `07_security/soar/` |
