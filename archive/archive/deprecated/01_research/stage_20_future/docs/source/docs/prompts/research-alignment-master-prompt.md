# Research Alignment Master Prompt

## Purpose

Reusable prompt for aligning research data, topics, repositories, tools, dependencies, milestones, and execution planning for the Cyber AI OS project.

Use after each research cycle to:

- prevent topic drift
- reorder dependencies
- update implementation priorities
- refresh repository/tool mappings
- generate milestones and issue backlogs
- identify risks and missing prerequisites

## Recommended Inputs

Provide the AI system with:

- current research notes
- existing docs tree
- selected project direction
- known repositories
- known tools
- current milestones
- current blockers
- time horizon
- available team/resources
- preferred output format

## Master Prompt

```text
You are acting as a Technical Research Architect, Systems Engineer, Cybersecurity Planner, AI/ML Research Coordinator, and Open Source Program Manager.

Objective:

Create a complete research alignment and project execution plan for building a Cybersecurity Operating System with AI/ML support, using open source tools, repositories, frameworks, and Linux foundations before moving toward custom OS components.

Project Goal:

Build a modular security platform evolving through:

Linux Base
→ Security Distribution
→ Telemetry Layer
→ AI Security Engine
→ Threat Intelligence
→ Forensics Stack
→ SOC Platform
→ Kernel Extensions
→ Hybrid Security OS
→ Custom Cyber AI OS

Tasks:

1. Collect and organize all research domains.

Domains include:

Operating Systems
Kernel Development
Computer Architecture
Memory Systems
Filesystems
Drivers
Networking
Cybersecurity
Threat Hunting
IDS
IPS
EDR
XDR
SIEM
SOAR
Digital Forensics
Malware Analysis
Reverse Engineering
Cloud Security
Container Security
eBPF
Observability
AI
Machine Learning
LLMs
Security AI
Threat Intelligence
Sandboxing
Incident Response
DevSecOps
Distributed Systems
Performance Engineering
Privacy Engineering
Cryptography
Zero Trust
Red Team
Blue Team
Purple Team

2. Build dependency graph.

For every topic determine:

Prerequisites
Required knowledge
Tools needed
Repositories
Datasets
Complexity level
Estimated effort
Research depth
Implementation order
Parallel learning possibilities
Blocking dependencies

Output format:

Topic
├── prerequisite topics
├── required tools
├── repositories
├── papers
├── implementation phase
├── difficulty
└── next topics

3. Produce learning order.

Generate:

Stage 0:
Foundations

Stage 1:
System internals

Stage 2:
Security foundations

Stage 3:
Threat detection

Stage 4:
AI and ML

Stage 5:
Kernel telemetry

Stage 6:
SOC stack

Stage 7:
Forensics

Stage 8:
Threat intelligence

Stage 9:
Security automation

Stage 10:
Custom Linux security distribution

Stage 11:
Kernel modification

Stage 12:
Hybrid cyber OS

Stage 13:
Research OS

Stage 14:
Production architecture

4. Repository mapping.

For each area identify:

Open source projects
GitHub repos
Maintainers
Contribution difficulty
Languages used
Architecture
Documentation quality
Beginner tasks
Advanced tasks

Categorize:

Learn from
Contribute to
Fork and extend
Integrate directly
Reference only
Experimental

5. Tool alignment.

Separate into:

Development tools
Security tools
ML tools
Observability
Forensics
Cloud
Containers
Telemetry
Reverse engineering
Malware research
Threat intelligence
Automation
Testing
Simulation
Virtualization

For each include:

Purpose
Input
Output
Dependencies
Integration points

6. Build implementation roadmap.

Roadmap:

Research
↓
Prototype
↓
Integration
↓
Security platform
↓
AI support
↓
Telemetry
↓
Kernel hooks
↓
Threat engine
↓
Forensics
↓
Distribution build
↓
OS extensions
↓
Custom kernel
↓
Cyber AI OS

7. Produce artifacts.

Generate:

Research database structure

Folder hierarchy

Knowledge graph

Roadmaps

Gantt style timeline

Milestones

Risk matrix

Architecture diagrams

Dependency matrices

Repository registry

Learning backlog

Issue tracker

Documentation templates

ADR files

RFC templates

8. Risk analysis.

Identify:

Knowledge gaps
Missing prerequisites
Overengineering risks
Scope creep
Research bottlenecks
Security risks
Model risks
Performance risks
Kernel complexity risks
Maintenance burden

9. Produce final outputs.

Return:

Master research tree

Dependency graph

Ordered roadmap

Topic priority table

Repository list

Contribution plan

Phase checkpoints

Milestone schedule

Expected outputs

Deliver in strict dependency order.

Never place advanced topics before prerequisites.

Flag circular dependencies.

Mark optional paths.

Mark research-only paths.

Mark implementation-critical paths.

Prioritize open source ecosystems.

Assume long term development of a Cybersecurity + AI + Hybrid OS ecosystem.
```

## Optional Add-On Instructions

Use these when needed.

### For Weekly Planning

```text
Also convert the roadmap into a 12-week, 16-week, 6-month, and 12-month plan.
For each period include:
- goals
- deliverables
- blockers
- evaluation criteria
- required repositories
- tools to install
- documentation to produce
```

### For GitHub Issue Generation

```text
Convert the implementation roadmap into GitHub-style issues.
Each issue must include:
- title
- objective
- scope
- dependencies
- acceptance criteria
- implementation notes
- test plan
- labels
- estimated effort
```

### For Repository Contribution Planning

```text
For each open source repository, identify:
- good first issues
- documentation contribution paths
- test contribution paths
- integration opportunities
- architecture areas to study
- maintainers or governance model
- contribution risk
```

### For Architecture Review

```text
Review the generated architecture for:
- circular dependencies
- premature custom-kernel work
- overengineering
- missing telemetry contracts
- missing security boundaries
- missing rollback paths
- weak evaluation metrics
- unclear ownership
```

## Expected Output Sections

Require this structure when consistency matters:

```text
1. Executive Summary
2. Master Research Tree
3. Dependency Graph
4. Stage Order
5. Topic Priority Table
6. Repository Registry
7. Tool Alignment Matrix
8. Implementation Roadmap
9. Milestone Schedule
10. Artifact Plan
11. Risk Matrix
12. Contribution Plan
13. Backlog
14. Open Questions
15. Next Actions
```

## Usage Rule

Run this prompt after every major research addition, topic expansion, or architecture change.

If the output introduces advanced work before prerequisites, reject that section and regenerate with stricter dependency ordering.

