---
project-id: lifetrack
repo: https://github.com/fallofpheonix/LifeTrack
status: REFACTOR
audit-verdict: "REFACTOR / POSSIBLE PRIMARY APP"
audit-scores: ED7/AQ7/SC5/RU6/MT6/OR6/PV8/CP8/RU5/EC7/LV8/RS7/OS5
language: Dart
started: 2026-02
last-commit: 2026-02-22
tags: [project, existing, mobile, health-tech, flutter]
---

# LifeTrack

## What It Is
A privacy-first, offline-first health management app built with Flutter. Tracks steps, vitals (BP, heart rate, glucose, weight), hydration, sleep, and calories with a domain/data/presentation architecture. All data stays local — zero-cloud architecture. Features an intelligence engine with consistency scoring, plateau detection, and context-aware health suggestions. Built for personal health tracking without surrendering data to third parties.

## Current State
- **What is working:** Flutter app structure with domain/data/presentation separation, large and organized codebase, Riverpod state management, Drift (SQLite) local persistence, modular feature architecture (Vitals, Activity, Hydration, Nutrition, Medical, Profile), dark-mode UI.
- **What is broken or incomplete:** No backend/sync/security model, TODO platform integrations (HealthKit/Google Fit) not implemented, no encrypted storage, CI skips tests instead of failing, generated platform scaffold files tracked in git, no threat model, no auth system.
- **Audit-identified defects:**
  - No backend/sync/security model — health data stored without encryption
  - CI skips tests instead of failing — CI is configured to hide failures, not catch them
- **Production readiness:** NOT READY — security model is a hard prerequisite before any feature work.

**HARD REQUIREMENT:** Build auth-bcrypt and totp-2fa modules first. LifeTrack must import both before any new features are added.

## Architecture

**Current:**
```
LifeTrack/
├── lib/
│   ├── core/
│   │   ├── database/          ← Drift schema, migrations, repositories
│   │   └── theme/
│   ├── domain/models/         ← POJO domain entities
│   ├── features/
│   │   ├── activity/
│   │   ├── hydration/
│   │   ├── medical/
│   │   ├── nutrition/
│   │   ├── profile/
│   │   └── vitals/
│   └── state/providers/       ← Riverpod providers for analytics
├── android/                   ← generated scaffold (should not be tracked)
├── ios/                       ← generated scaffold
├── test/
└── pubspec.yaml
```

**Target (from audit):**
```
LifeTrack/
├── lib/
│   ├── core/
│   │   ├── database/          ← Drift + encrypted storage
│   │   ├── auth/              ← import auth-bcrypt + totp-2fa
│   │   ├── security/          ← encryption at rest, threat model
│   │   └── platform/          ← HealthKit / Google Fit adapters
│   ├── domain/
│   ├── features/
│   └── state/
├── test/                      ← CI must FAIL on test failure
├── docs/
│   ├── threat-model.md
│   └── adr/
└── pubspec.yaml
```

## Tech Stack
- **Language:** Dart
- **Framework:** Flutter 3.10.8+
- **State Management:** flutter_riverpod
- **Persistence:** Drift (SQLite), code generation
- **Visualization:** fl_chart
- **Design:** Atomic UI with dark-mode focus

## Modules Used
| Module | Status | Notes |
|--------|--------|-------|
| auth-bcrypt | NOT YET BUILT | Hard prerequisite — must be built before any LifeTrack features |
| totp-2fa | NOT YET BUILT | Hard prerequisite — must be built before any LifeTrack features |

## Retroactive ADRs
- [[04_ENGINEERING/decision-logs/ADR-018-lifetrack-local-first]] — why local-first architecture with no backend sync
- [[04_ENGINEERING/decision-logs/ADR-019-lifetrack-domain-data-presentation]] — why domain/data/presentation Flutter architecture pattern
- [[04_ENGINEERING/decision-logs/ADR-020-lifetrack-flutter-over-native]] — why Flutter over React Native or native iOS/Android

## Known Failure Modes
- [[06_FAILURE_LIBRARY/2026-02-lifetrack-ci-skips-tests]] — CI configured to skip tests instead of failing, hiding real failures
- [[06_FAILURE_LIBRARY/2026-02-lifetrack-no-threat-model]] — health data stored without encryption or threat model

## Committed Artifact Violations
- Generated platform scaffold files (android/, ios/, linux/, windows/, macos/) tracked in git
- .dart_tool/ or build/ artifacts — verify with git status

## Refactor Actions
- [ ] Write threat model (what data is stored, who can access, attack surfaces)
- [ ] Add encrypted local storage (flutter_secure_storage or equivalent)
- [ ] Add FHIR export (standard health data interop format)
- [ ] Implement real HealthKit (iOS) and Google Fit (Android) integration — remove TODOs
- [ ] Fix CI: make it FAIL on test failure, not skip
- [ ] Remove generated platform scaffold files from git tracking (.gitignore)
- [ ] Build and import auth-bcrypt module
- [ ] Build and import totp-2fa module
- [ ] Write architecture diagram

## Spec Kit Status
- specify init: NOT YET RUN
- Next feature to spec: Threat model + encrypted storage (security prerequisite)
- When ready: `specify init . --integration claude` inside `~/engineering/workspace/active/lifetrack/`
- **BLOCKED:** Do not start new features until threat model is written

## Linked Concepts
- [[03_CORE_KNOWLEDGE/security]] — encryption at rest, threat modeling, authentication
- [[03_CORE_KNOWLEDGE/databases]] — SQLite, Drift ORM, local-first data architecture

## Promotion Criteria
- [ ] Threat model written and reviewed
- [ ] Encrypted storage implemented
- [ ] Auth modules (bcrypt + 2FA) imported and working
- [ ] CI fails on test failure (not skips)
- [ ] Platform scaffolds removed from git tracking
- [ ] At least one real platform integration (HealthKit or Google Fit)
- [ ] All 3 ADRs written
- [ ] README follows docs contract


## Independence Checklist
- [ ] Unique port assigned?
- [ ] Dedicated DB or schema?
- [ ] Unique `.env` injected?
- [ ] Isolated Docker network?
