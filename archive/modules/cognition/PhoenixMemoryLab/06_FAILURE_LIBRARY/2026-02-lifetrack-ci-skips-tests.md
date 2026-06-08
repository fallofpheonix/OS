---
failure-id: 2026-02-lifetrack-ci-skips-tests
project: [[05_PROJECTS/ACTIVE/lifetrack]]
severity: HIGH
status: OPEN
date-encountered: 2026-02
tags: [failure, ci, testing]
---
# Failure: CI configured to skip tests instead of failing

## What Was Tried
Setting up GitHub Actions CI for the LifeTrack Flutter project.

## What Happened
The CI pipeline was configured to skip failing tests rather than fail the build. Tests that would normally block a merge are silently ignored. This means broken code can be merged to main without any test gate.

## Root Cause
Likely a quick fix during initial CI setup when tests were failing — the pipeline was changed to `continue-on-error: true` or equivalent instead of fixing the tests. The temporary fix was never reverted.

## What Was Learned
CI must never skip tests. If tests are failing, the correct response is to fix the tests or mark specific tests as expected failures (skip with reason), not to configure the entire pipeline to ignore test results. A CI pipeline that cannot fail is not a CI pipeline.

## Prevention / Resolution
- Change CI to fail on test failure (remove continue-on-error or equivalent)
- Fix all currently failing tests
- Add test failure as a required check for PR merges
- Add CI health check to the project's promotion criteria

## Linked Concepts
- [[03_CORE_KNOWLEDGE/devops]] — CI/CD best practices, test gates
- [[03_CORE_KNOWLEDGE/testing]] — test reliability, CI configuration
