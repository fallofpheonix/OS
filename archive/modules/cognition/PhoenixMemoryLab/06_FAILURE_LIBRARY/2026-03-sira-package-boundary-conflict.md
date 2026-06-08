---
failure-id: 2026-03-sira-package-boundary-conflict
project: [[05_PROJECTS/ACTIVE/sira]]
severity: MEDIUM
status: OPEN
date-encountered: 2026-03
tags: [failure, python, packaging]
---
# Failure: Two conflicting Python package roots in sira

## What Was Tried
The project started with code at the root level and later adopted a src/ layout. Both were kept.

## What Happened
Two package roots exist: sira/ (root-level) and src/sira/ (src-layout). Python imports resolve differently depending on which root is on sys.path. Tests may import from one root while the API imports from another, leading to subtle bugs where changes to one don't affect the other.

## Root Cause
Migration from flat layout to src/ layout was incomplete. The old package root was not removed.

## What Was Learned
Python projects must have exactly one canonical package root. The src/ layout is preferred (prevents accidental imports from the working directory). When migrating layouts, the old root must be deleted, not kept alongside.

## Prevention / Resolution
- Delete the root-level sira/ package
- Ensure src/sira/ is the single canonical package root
- Update all import statements to use the src layout
- Update pyproject.toml [tool.pytest.ini_options] to point to src/

## Linked Concepts
- [[03_CORE_KNOWLEDGE/architecture]] — Python project layout, src layout
- [[03_CORE_KNOWLEDGE/devops]] — Python packaging, import resolution
