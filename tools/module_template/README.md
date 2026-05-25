# __MODULE_NAME__

## Purpose
Brief description of the module's responsibility, primary use-cases, and the problem it solves. Replace `__MODULE_NAME__` with the real module name. Example: "Provides a deterministic telemetry replay engine used to validate kernel-actuator interactions under stress." 

## Structure
- `src/`: Core implementation
- `tests/`: Unit and integration tests
- `bench/`: Performance benchmarks
- `replay/`: Telemetry replay data and logic
- `debug/`: Debugging traces
- `artifacts/`: Build outputs

## Performance Budget
- Latency p95: 50ms (example)
- Latency p99: 200ms (example)
- Throughput: 1000 events/s (example)

Adjust values per-module. Use benchmarks in `bench/` to validate and record results.

## Validation Gates
- [ ] Build success
- [ ] Unit test pass (100% coverage)
- [ ] Performance within budget
- [ ] Replay deterministic output
