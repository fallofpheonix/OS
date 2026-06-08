---
failure-id: 2026-03-healingstone-module-boundary-drift
project: [[05_PROJECTS/ACTIVE/healingstone]]
severity: MEDIUM
status: OPEN
date-encountered: 2026-03
tags: [failure, architecture, duplication]
---
# Failure: Duplicated schema/core modules between 2D and 3D pipelines

## What Was Tried
Building 2D and 3D fragment reconstruction as separate pipelines, each with their own schema definitions and core utility modules.

## What Happened
Schema definitions and core utilities (descriptor matching, similarity scoring) were duplicated between pipeline_2d/ and pipeline_3d/. Changes to one must be manually propagated to the other, leading to drift.

## Root Cause
The 2D/3D pipeline split (ADR-009) was implemented without extracting shared code into a common module first. Each pipeline was built independently, and shared patterns were copy-pasted rather than abstracted.

## What Was Learned
When splitting a system into parallel pipelines, extract the shared interface FIRST, then build the pipeline variants on top. Copy-paste duplication between parallel modules always leads to drift.

## Prevention / Resolution
- Extract shared code into a healingstone/core/ module used by both pipelines
- Define a common descriptor/matching interface that both 2D and 3D implement
- Use abstract base classes or protocols to enforce the shared contract

## Linked Concepts
- [[03_CORE_KNOWLEDGE/architecture]] — module boundaries, code duplication, interface segregation
- [[03_CORE_KNOWLEDGE/algorithms]] — descriptor matching, similarity scoring
