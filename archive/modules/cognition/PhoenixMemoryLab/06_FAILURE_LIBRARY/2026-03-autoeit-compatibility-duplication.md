---
failure-id: 2026-03-autoeit-compatibility-duplication
project: [[05_PROJECTS/ACTIVE/autoeit-suite]]
severity: LOW
status: OPEN
date-encountered: 2026-03
tags: [failure, architecture, duplication]
---
# Failure: Duplicated compatibility shims between audio_transcription and AutoEIT-STS

## What Was Tried
Both repos independently implemented compatibility layers for shared concerns (Excel I/O, text normalization, accent handling).

## What Happened
The same compatibility code exists in both repos with slight variations. When a bug was fixed in one, it wasn't propagated to the other. The two implementations diverged over time.

## Root Cause
Two repos that should have been one suite from the start. Without a shared package, each independently solved the same problems.

## What Was Learned
When two tools share a data format or processing step, the shared code must be extracted into a common package immediately — not "later." Duplication between repos is harder to detect and fix than duplication within a repo.

## Prevention / Resolution
- Extract shared code into packages/autoeit_common/ during the suite merge (ADR-023)
- Add tests that verify both pipelines use the shared implementation
- Prevent future duplication with import linting (ban duplicate module names across packages)

## Linked Concepts
- [[03_CORE_KNOWLEDGE/architecture]] — DRY principle, shared libraries, monorepo patterns
- [[03_CORE_KNOWLEDGE/devops]] — dependency management, package extraction
