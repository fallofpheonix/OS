# Event Normalizer Module

Purpose: Provide a safe, userspace event normalizer for SentinelOS telemetry. This module implements schema validation, light enrichment, and a replayable normalization pipeline for simulated telemetry events.

Structure created:
- `RFC.md` — design and validation gates
- `schema.json` — normalized event schema
- `src/normalizer.py` — normalization logic (pure Python, no kernel/eBPF)
- `tests/test_normalizer.py` — unit tests
- `bench/bench_normalizer.py` — micro-benchmark
- `replay/sample_events.json` — sample recorded events for replay
- `build.sh`, `run.sh`, `test.sh`, `debug.sh`, `benchmark.sh`, `replay.sh`, `validate.sh` — lifecycle scripts

Run lifecycle scripts from repository root or this folder. These scripts are safe and use mocked inputs.
