---
failure-id: 2026-02-lifetrack-no-threat-model
project: [[05_PROJECTS/ACTIVE/lifetrack]]
severity: HIGH
status: OPEN
date-encountered: 2026-02
tags: [failure, security, health-data]
---
# Failure: Health data stored without encryption or threat model

## What Was Tried
Storing health data (vitals, medical records, activity logs) in a local SQLite database via Drift without any encryption or security analysis.

## What Happened
The app claims to be "privacy-first" but stores sensitive health data in plaintext on the device. No threat model exists. No encryption at rest. No assessment of what happens if the device is lost, stolen, or forensically examined.

## Root Cause
"Privacy-first" was interpreted as "no cloud" rather than "data is protected." Local-only does not mean secure — physical device access or a malicious app with storage permissions can read the SQLite database.

## What Was Learned
Privacy requires active security measures, not just the absence of a cloud backend. A threat model must be written before storing sensitive data. Minimum requirements for health data: encryption at rest, secure key storage, and a documented security model.

## Prevention / Resolution
- Write a threat model documenting: what data is stored, sensitivity levels, access vectors, and mitigations
- Implement encrypted local storage (flutter_secure_storage for keys, SQLCipher for database encryption)
- Add biometric/PIN authentication for app access
- Document security model in docs/threat-model.md

## Linked Concepts
- [[03_CORE_KNOWLEDGE/security]] — threat modeling, encryption at rest, data classification
- [[03_CORE_KNOWLEDGE/architecture]] — security architecture, privacy by design
