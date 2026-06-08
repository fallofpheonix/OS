---
project-id: autoeit-suite
repo: https://github.com/fallofpheonix/AutoEIT-STS
status: REFACTOR
audit-verdict: "CONVERT TO PACKAGE / MERGE INTO AUTOEIT SUITE"
language: Python
started: 2026-03
last-commit: 2026-03-17
tags: [project, existing, linguistics, nlp, audio-processing]
---

# AutoEIT Suite

## What It Is
A merged audio transcription + deterministic EIT linguistic scoring pipeline. Combines two previously separate repos — audio_transcription (Whisper-based Spanish ASR) and AutoEIT-STS (rule-based elicited imitation task scorer) — into a single end-to-end pipeline: audio → segmentation → ASR → sentence alignment → workbook output → deterministic rubric scoring → agreement/audit report. Built for Spanish language acquisition researchers who need reproducible, auditable scoring.

## Current State
- **audio_transcription:** Committed model/audio/workbook/output binaries in git, compatibility-layer duplication across modules.
- **AutoEIT-STS:** Clean deterministic domain package with tests, Streamlit UI, and auditable scoring. Conservative 2/3 boundary for inter-rater reliability.
- **Combined status:** Natural upstream/downstream pipeline that should be one repo. The merge has NOT been performed yet — this Project.md covers the target merged state.
- **Audit-identified defects:**
  - Committed model/audio/workbook binaries — .pt, .bin, .wav, .xlsx files in git history
  - Compatibility-layer duplication — shared code duplicated between the two repos
- **Production readiness:** NOT READY — merge must happen first, then binary purge.

## Architecture

**Current (two separate repos):**
```
audio_transcription/          AutoEIT-STS/
├── models/ ← VIOLATION      ├── src/autoeit/
├── audio/  ← VIOLATION      │   ├── scorer.py
├── output/ ← VIOLATION      │   ├── rubric.py
├── src/                      │   └── io/
│   ├── transcriber.py        ├── tests/
│   └── compatibility/        ├── streamlit_app.py
└── tests/                    └── configs/
```

**Target (merged suite from audit):**
```
autoeit-suite/
├── packages/
│   ├── autoeit_transcribe/   ← from audio_transcription
│   ├── autoeit_score/        ← from AutoEIT-STS
│   └── autoeit_common/       ← shared primitives (extracted)
├── apps/
│   └── streamlit_review/     ← Streamlit scoring review UI
├── datasets/sample/          ← sample audio + expected output (via DVC/releases)
├── docs/validation/          ← inter-rater reliability reports
├── tests/
│   ├── unit/
│   └── integration/          ← end-to-end: audio in → audit report out
└── pyproject.toml
```

**Pipeline flow:** audio → segmentation → ASR (Whisper) → sentence alignment → workbook output → deterministic rubric scoring → agreement/audit report

## Tech Stack
- **Language:** Python 3.10+
- **ASR:** faster-whisper (CPU/GPU inference, configurable model sizes)
- **Scoring:** Rule-based rubric engine (0-4 scale, accent-insensitive)
- **UI:** Streamlit
- **I/O:** openpyxl (Excel workbook processing)
- **Testing:** pytest

## Modules Used
| Module | Status | Notes |
|--------|--------|-------|
| [[08_MODULES/fpx-runtime]] | Not yet imported | Executor for transcription jobs |
| [[08_MODULES/fpx-pipeline]] | Not yet imported | End-to-end audio → report pipeline |

## Retroactive ADRs
- [[04_ENGINEERING/decision-logs/ADR-021-autoeit-deterministic-scoring]] — why deterministic rule-based scoring over ML (audit trail requirement)
- [[04_ENGINEERING/decision-logs/ADR-022-autoeit-excel-output-format]] — why Excel workbook as output format (stakeholder requirement)
- [[04_ENGINEERING/decision-logs/ADR-023-autoeit-suite-merge-decision]] — why merging two repos into one suite

## Known Failure Modes
- [[06_FAILURE_LIBRARY/2026-03-autoeit-binaries-in-repo]] — model checkpoints, audio files, and output workbooks committed to git
- [[06_FAILURE_LIBRARY/2026-03-autoeit-compatibility-duplication]] — duplicated compatibility shims between the two repos

## Committed Artifact Violations (CRITICAL)
- Model checkpoints (.pt, .bin files) — move to GitHub Releases or DVC
- Sample audio files (.wav) — move to Releases/DVC
- Output workbooks (.xlsx) — move to Releases
- Local database backups — delete permanently
- ALL must be removed from git history with BFG before promotion

## Refactor Actions
- [ ] Create new autoeit-suite repo (or rename audio_transcription and restructure)
- [ ] Move audio_transcription code into `packages/autoeit_transcribe/`
- [ ] Move AutoEIT-STS code into `packages/autoeit_score/`
- [ ] Extract shared code into `packages/autoeit_common/`
- [ ] Remove ALL committed binaries from git history (BFG Repo Cleaner)
- [ ] Publish model artifacts to GitHub Releases + document download in README
- [ ] Add pipeline integration test (end-to-end: audio in → audit report out)
- [ ] Write inter-rater reliability validation docs

## Spec Kit Status
- specify init: NOT YET RUN
- Next feature to spec: Repo merge + monorepo restructure
- When ready: `specify init . --integration claude` inside `~/engineering/workspace/active/autoeit-suite/`

## Linked Concepts
- [[03_CORE_KNOWLEDGE/ai-ml]] — speech recognition, Whisper ASR
- [[03_CORE_KNOWLEDGE/algorithms]] — rule-based systems, deterministic scoring, text comparison

## Promotion Criteria
- [ ] Two repos merged into one monorepo
- [ ] All binaries removed from git history
- [ ] Model artifacts published to GitHub Releases with download docs
- [ ] End-to-end integration test passing
- [ ] Inter-rater reliability validation documented
- [ ] All 3 ADRs written
- [ ] README follows docs contract


## Independence Checklist
- [ ] Unique port assigned?
- [ ] Dedicated DB or schema?
- [ ] Unique `.env` injected?
- [ ] Isolated Docker network?
