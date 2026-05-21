# Stage 00: Foundations

## Purpose

Establish core computer science fundamentals required for system design, systems programming, security engineering, AI/ML engineering, telemetry pipelines, and OS development.

## Scope

- Low-level data representation.
- Memory and pointer semantics.
- Fundamental data structures.
- Core algorithms.
- Complexity analysis.
- Basic concurrency and state models.

## Classification

- Type: `FOUNDATIONAL`
- Status: `RESEARCH_ONLY`
- Difficulty: foundation level
- Estimated duration: 4-6 weeks
- Downstream blockers:
  - Stage 01 System Internals
  - Stage 03 Networking
  - Stage 04 Security
  - Stage 11 AI/ML Core

## Research Modules

| Module | Path |
|---|---|
| Computer Science Foundations | `notes/phase_00_computer_science_foundations.md` |
| Computer Architecture | `architecture/phase_01_computer_architecture.md` |
| Low-Level Programming | `development/phase_02_low_level_programming.md` |
| Phase 0 Build Gate | `checkpoints/phase_00_build_gate.md` |
| Phase 1 Build Gate | `checkpoints/phase_01_build_gate.md` |
| Phase 2 Build Gate | `checkpoints/phase_02_build_gate.md` |
| Phase 0 Implementation Labs | `labs/phase_00_implementation_labs.md` |
| Phase 1 Architecture Labs | `labs/phase_01_architecture_labs.md` |
| Phase 2 Low-Level Programming Labs | `labs/phase_02_low_level_programming_labs.md` |

## Internal Dependency Order

```text
Phase 0: Computer Science Foundations
-> Phase 1: Computer Architecture
-> Phase 2: Low-Level Programming
```

## Gate

Do not start Stage 01 implementation work until Phase 0, Phase 1, and Phase 2 build gates are complete.
