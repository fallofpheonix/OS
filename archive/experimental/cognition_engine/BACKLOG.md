# Cognition Engine Backlog

This document tracks pending issues and tasks for the Cognition Engine (Astraeus Core), as identified in the May 2026 status report.

## Actionable Tasks (Astraeus Core)
These tasks apply to the archived codebase in `experimental/cognition_engine/astraeus_core`.

- [ ] **Commit Phase A Hardening**: The working directory state from the last session needs to be staged and committed. Includes live Ollama support, strict routing, critic rejection, and Python syntax enforcement.
- [ ] **Fix Repair Loop Mock**: The repair loop currently uses a mock critic. It must be updated to use real model evaluation to ensure success metrics are accurate.
- [ ] **Stabilize CI Gates**: Add `.trufflehogignore` entries for confirmed mock secrets in test fixtures to prevent result count drift.
- [ ] **WebSocket Integration**: Connect the `frontend-console` HTML interface to the live API via WebSockets and verify the 6 message types (`dag_ready`, `task_started`, etc.).
- [ ] **MyPy Strict Mode**: Transition from partial MyPy passing to full strict mode coverage (`--strict`).

## Blocked Tasks (OpenHuman Repository)
The following tasks are blocked because the `openhuman` repository has been removed from the workspace. They are preserved here for forensic completeness.

- [ ] **Stripe Key Investigation**: Verify if `sk_live_12345678901234567890` in `src/openhuman/memory/safety/mod.rs` is a mock or a compromised live key.
- [ ] **Cleanup Diagnostic Scripts**: Remove `scan.py`, `scan_urls.py`, etc., from `tests/fixtures/` to prevent test collection pollution.
- [ ] **Dynamic Fixture Fetching**: Migrate `composio_github.json` from a git-committed file to a dynamically fetched asset with SHA-256 checksum validation.
- [ ] **macOS Shasum Compatibility**: Update security documentation to use `shasum -a 256` instead of `sha256sum`, which is not natively available on macOS.

## References
- **Status Report Source**: `15_archive/parts/engineering/_quarantine/20260519_1834/research/Pasted text.txt`
- **Granular Roadmap**: `experimental/cognition_engine/roadmap/COGNITION_ENGINE_ROADMAP.md`
