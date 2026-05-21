# Staged Restructuring Master Prompt

## Purpose

Reusable prompt for taking the full stored Cyber AI OS project plan and separating it into:

- dependency-driven stages
- filesystem layout
- milestones
- checkpoints
- dependency reports
- execution phases
- risks
- critical path

Run this before Stage 1 implementation work.

## Master Prompt

```text
You are acting as a Research Program Architect, Systems Planner, Cybersecurity Lead, AI/ML Coordinator, and Repository Organizer.

Objective:

Take the entire existing project plan already stored in the project folder and reorganize everything into dependency-driven stages.

Input Source:

Project root contains:

docs/
research/
ai/
kernel/
security/
network/
drivers/
userspace/
tools/
forensics/
telemetry/
notes/
roadmap/
repos/
papers/
datasets/

Task:

Scan all plans, notes, markdown files, research outputs, roadmaps, ADRs, diagrams, repository lists, datasets, and documentation.

Do NOT modify content initially.

Only classify, stage, and reorganize.

Create stages based on dependency order.

Stage Structure:

Stage_00_Foundations
Stage_01_System_Internals
Stage_02_Linux_and_Distro
Stage_03_Networking
Stage_04_Security_Fundamentals
Stage_05_Malware_and_RE
Stage_06_Forensics
Stage_07_Threat_Intelligence
Stage_08_Observability
Stage_09_eBPF_and_Telemetry
Stage_10_SOC_Stack
Stage_11_AI_ML_Core
Stage_12_Security_AI
Stage_13_Containers_and_Cloud
Stage_14_Automation_and_SOAR
Stage_15_Security_Distribution
Stage_16_Kernel_Extensions
Stage_17_Hybrid_OS
Stage_18_Custom_OS
Stage_19_Production_Platform
Stage_20_Research_and_Future

For every discovered item determine:

Topic name
Current location
Dependencies
Prerequisites
Difficulty
Estimated effort
Research depth
Implementation status
Stage assignment
Cross references
Blocking items

Generate:

project_stages/
│
├── Stage_00_Foundations/
│   ├── topics/
│   ├── repos/
│   ├── docs/
│   ├── checkpoints/
│   ├── labs/
│   └── milestones/
│
├── Stage_01_System_Internals/
│
├── Stage_02_Linux_and_Distro/
│
...
│
└── Stage_20_Research_and_Future/

Create these outputs:

MASTER_STAGE_MAP.md
DEPENDENCY_GRAPH.md
TOPIC_MATRIX.md
REPOSITORY_ALIGNMENT.md
DATASET_ALIGNMENT.md
TOOL_ALIGNMENT.md
CHECKPOINTS.md
MILESTONES.md
RISK_REGISTER.md
EXECUTION_ORDER.md

Rules:

1. Never place advanced topics before prerequisites.

2. Detect circular dependencies.

3. Mark:

FOUNDATIONAL
IMPLEMENTATION
RESEARCH_ONLY
OPTIONAL
FUTURE

4. Separate:

Learning tasks
Research tasks
Coding tasks
Contribution tasks
Integration tasks

5. Build progression:

Learn
→ Prototype
→ Integrate
→ Extend
→ Secure
→ Observe
→ Automate
→ Scale
→ OS Extension
→ Custom Kernel
→ Cyber AI OS

6. Estimate:

Hours
Weeks
Difficulty
Risk
Priority

7. Preserve all files.

No deletion.

No overwrite.

Only classify and copy.

8. Produce final report:

Stage count
Topic count
Repo count
Dependency bottlenecks
Critical path
Parallel work opportunities
Gate conditions
Next executable stage

Goal:

Transform one large research ecosystem into staged execution units suitable for long term development of:

Cybersecurity Platform
+
AI/ML Security Engine
+
Telemetry Stack
+
SOC
+
Forensics
+
Kernel Extensions
+
Hybrid Security OS
+
Custom Cyber AI Operating System
```

## Required Output Shape

Ask the AI system to return this structure:

```text
1. Scan Summary
2. Stage Assignment Rules
3. Master Stage Map
4. Dependency Graph
5. Topic Matrix
6. Repository Alignment
7. Dataset Alignment
8. Tool Alignment
9. Checkpoints
10. Milestones
11. Risk Register
12. Execution Order
13. Critical Path
14. Parallel Work Plan
15. Gate Conditions
16. Next Executable Stage
17. Filesystem Operations Plan
```

## Filesystem Operation Constraints

Use these constraints when the AI system has file access:

```text
Before copying files:
- produce a dry-run file operation plan
- list every source file
- list every destination file
- list duplicate classifications
- list ambiguous files
- request approval if overwrites would occur

Allowed:
- create new directories
- copy files
- create index files
- create stage reports

Forbidden:
- delete files
- overwrite original files
- move files without explicit approval
- rewrite source content during initial classification
```

## Stage Index Template

Each stage should include an index:

```md
# Stage XX: Name

## Purpose

TBD

## Classification

FOUNDATIONAL / IMPLEMENTATION / RESEARCH_ONLY / OPTIONAL / FUTURE

## Prerequisites

- TBD

## Topics

- TBD

## Source Files

- TBD

## Tools

- TBD

## Repositories

- TBD

## Labs

- TBD

## Milestones

- TBD

## Gate Conditions

- TBD

## Next Stage

TBD
```

## Topic Matrix Template

```md
| Topic | Stage | Prerequisites | Tools | Repos | Datasets | Difficulty | Effort | Risk | Priority | Status |
|---|---|---|---|---|---|---|---|---|---|---|
| TBD | TBD | TBD | TBD | TBD | TBD | TBD | TBD | TBD | TBD | TBD |
```

## Risk Register Template

```md
| Risk | Stage | Cause | Impact | Likelihood | Severity | Mitigation | Owner | Status |
|---|---|---|---|---|---|---|---|---|
| TBD | TBD | TBD | TBD | TBD | TBD | TBD | TBD | TBD |
```

## Usage Rule

Run this prompt after the research alignment prompt when the project has enough material to organize into stages.

Order:

```text
1. Run research-alignment-master-prompt.md
2. Review generated dependency ordering
3. Run staged-restructuring-master-prompt.md
4. Review dry-run filesystem plan
5. Approve copy/classification work
6. Begin Stage_00 or next executable stage
```

