---
failure-id: 2026-03-particle-archive-in-main
project: [[05_PROJECTS/ACTIVE/particle-stimulator]]
severity: HIGH
status: OPEN
date-encountered: 2026-03
tags: [failure, git, artifacts]
---
# Failure: Large archive directory committed to main branch

## What Was Tried
Historical simulation results and previous iterations of the project were kept in an archive/ directory within the main branch.

## What Happened
The archive/ directory adds significant size to the repository. Every clone downloads the full archive. The archive contains files that are not part of the current working version of the software.

## Root Cause
Archive was used as a local backup strategy instead of git tags or releases. No .gitignore rule prevented committing archive files.

## What Was Learned
Historical versions belong in git tags or GitHub Releases, not in directories on the main branch. The main branch should contain only the current working version of the software.

## Prevention / Resolution
- Move archive/ contents to a versioned git tag (e.g., v0.1-archive)
- Remove archive/ from main branch
- Add archive/ to .gitignore
- Use GitHub Releases for published artifacts

## Linked Concepts
- [[03_CORE_KNOWLEDGE/devops]] — git branching, release management
- [[03_CORE_KNOWLEDGE/architecture]] — repository structure, artifact management
