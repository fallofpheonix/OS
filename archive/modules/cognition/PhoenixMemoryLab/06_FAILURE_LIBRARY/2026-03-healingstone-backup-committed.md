---
failure-id: 2026-03-healingstone-backup-committed
project: [[05_PROJECTS/ACTIVE/healingstone]]
severity: HIGH
status: OPEN
date-encountered: 2026-03
tags: [failure, git, artifacts]
---
# Failure: backup_untracked directory committed to git in healingstone

## What Was Tried
During development, backup copies of untracked files were saved to a backup_untracked/ directory. This directory was not in .gitignore.

## What Happened
The entire backup_untracked/ tree was committed to git, adding binary/generated artifacts to the repository history. This inflates the repo size, pollutes git log, and makes cloning slower.

## Root Cause
Missing .gitignore entry for backup directories. No pre-commit hook to catch large binary files.

## What Was Learned
Never commit backup directories. Add common backup patterns to .gitignore from day one: backup*, *_backup, *_old, __pycache__, *.pyc, .DS_Store. Use pre-commit hooks (e.g., check-added-large-files) to catch accidentally committed binaries.

## Prevention / Resolution
- Remove backup_untracked/ from git history using BFG Repo Cleaner or git filter-repo
- Add comprehensive .gitignore with backup, generated, and binary patterns
- Add pre-commit hook: check-added-large-files (max 500KB)

## Linked Concepts
- [[03_CORE_KNOWLEDGE/devops]] — git hygiene, .gitignore, pre-commit hooks
- [[03_CORE_KNOWLEDGE/architecture]] — repository structure, artifact management
