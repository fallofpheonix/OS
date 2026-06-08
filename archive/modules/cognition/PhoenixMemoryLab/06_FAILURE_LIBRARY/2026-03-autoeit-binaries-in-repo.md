---
failure-id: 2026-03-autoeit-binaries-in-repo
project: [[05_PROJECTS/ACTIVE/autoeit-suite]]
severity: HIGH
status: OPEN
date-encountered: 2026-03
tags: [failure, git, artifacts, binaries]
---
# Failure: Model checkpoints, audio files, and workbooks committed to git

## What Was Tried
Developing audio_transcription with Whisper model checkpoints, sample audio files, and output Excel workbooks tracked in git.

## What Happened
Binary artifacts (.pt model files, .wav audio samples, .xlsx workbooks) were committed to the repository. This inflates the repo size significantly, makes cloning slow, and pollutes git history with undiffable binary blobs.

## Root Cause
No .gitignore rules for ML artifacts and data files. No pre-commit hook to catch large files. Development workflow treated git as a backup tool rather than a source code versioning tool.

## What Was Learned
Binary artifacts must NEVER be committed to source repositories. They belong in: GitHub Releases (for published versions), DVC (for versioned datasets), cloud storage (for large models), or local-only directories excluded via .gitignore.

## Prevention / Resolution
- Remove all binary artifacts from git history using BFG Repo Cleaner
- Add comprehensive .gitignore: *.pt, *.bin, *.h5, *.wav, *.mp3, *.xlsx, *.xls, models/, data/, output/
- Publish artifacts to GitHub Releases with download documentation in README
- Add pre-commit hook: check-added-large-files (max 500KB)

## Linked Concepts
- [[03_CORE_KNOWLEDGE/devops]] — git hygiene, artifact management, DVC
- [[03_CORE_KNOWLEDGE/ai-ml]] — ML model versioning, artifact lifecycle
