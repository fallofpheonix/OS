# Runtime Lab

Purpose: empirical runtime truth under hostile conditions.

## Test Domains
- floods
- soak
- restart-storms
- sink-poisoning
- memory-pressure
- contention
- lifecycle
- fuzzing
- scheduler
- degradation

## Lab Rules
- Define explicit pass and fail SLOs before running tests.
- Capture artifacts for every run (metrics, traces, incidents, logs).
- Promote only validated behavior into production-facing runtimes.
- No feature work should bypass lab validation for critical runtime paths.
