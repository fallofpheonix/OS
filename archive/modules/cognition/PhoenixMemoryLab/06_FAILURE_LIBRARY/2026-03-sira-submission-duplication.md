---
failure-id: 2026-03-sira-submission-duplication
project: [[05_PROJECTS/ACTIVE/sira]]
severity: MEDIUM
status: OPEN
date-encountered: 2026-03
tags: [failure, git, structure]
---
# Failure: Full source tree duplicated in submission/ directory

## What Was Tried
Preparing a submission (course assignment, competition, or paper) by copying the source tree into a submission/ directory within the main repo.

## What Happened
A complete copy of the source code exists at both the root level and inside submission/. This doubles the effective codebase size and creates confusion about which copy is authoritative.

## Root Cause
Submission was created by copying files rather than using git tags, branches, or release artifacts.

## What Was Learned
Submissions and releases should be created using git tags or GitHub Releases, not by copying source trees into subdirectories. Tags preserve the exact state of the code at a point in time without duplicating it.

## Prevention / Resolution
- Delete submission/ from main branch
- Create a git tag (e.g., v1.0-submission) pointing to the relevant commit
- If submission had modifications, create a branch (submission/v1) instead
- Add submission/, dist/, and build/ to .gitignore

## Linked Concepts
- [[03_CORE_KNOWLEDGE/devops]] — git branching, tagging, release management
- [[03_CORE_KNOWLEDGE/architecture]] — repository structure
