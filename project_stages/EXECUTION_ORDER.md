# Execution Order

## Critical Path

Stage_00_Foundations -> Stage_01_System_Internals -> Stage_02_Linux_and_Distro -> Stage_03_Networking -> Stage_04_Security_Fundamentals -> Stage_08_Observability -> Stage_10_SOC_Stack -> Stage_11_AI_ML_Core -> Stage_12_Security_AI -> Stage_14_Automation_and_SOAR -> Stage_15_Security_Distribution -> Stage_16_Kernel_Extensions -> Stage_17_Hybrid_OS -> Stage_18_Custom_OS -> Stage_19_Production_Platform -> Stage_20_Research_and_Future

## Parallel Work Opportunities

- After Stage_01: Stage_02, Stage_03, and Stage_04 can proceed in parallel.
- After Stage_04: Stage_05, Stage_07, and Stage_08 can proceed in parallel.
- After Stage_08: Stage_09 and Stage_10 can proceed in parallel.
- Stage_11 can begin after Stage_00, but Security AI integration waits for Stage_10 and malware/forensics evidence.
- Stage_13 can proceed after Linux, networking, security, and observability gates.

## Gate Conditions

- No advanced topic starts before its prerequisite stages pass.
- Every stage has source files copied under `docs/source/`.
- Every lab records inputs, outputs, tool versions, and failure modes.
- Every security/malware/forensics task records safety constraints.
- Every AI/ML task records dataset provenance and metrics.

## Next Executable Stage

Stage_00_Foundations. Existing Stage 0 documents already define dependency graph, repository map, topic resources, checkpoints, and 12-week schedule.

## Task Separation

- Learning: Stage_00, Stage_01, Stage_03, Stage_04, Stage_11.
- Research: Stage_20 and research-only documents.
- Coding: prototype/lab/build outputs across implementation stages.
- Contribution: upstream-aligned work from `REPOSITORY_ALIGNMENT.md`.
- Integration: SOC, telemetry, AI, distribution, hybrid/custom OS convergence.
