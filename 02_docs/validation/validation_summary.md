# Validation Summary

Coverage snapshot
- Unit tests: Present for FSM/warden components. Run results: 4 tests passed (local run context).
- Integration tests: Limited; no full end-to-end CI verification of kernel patches or swarm coordination.
- Benchmarks: Throughput claim (10M+ events/sec) in README — need traceable benchmark artifacts.

Missing validation
- Replay determinism tests for evidence ledger and trace ingestion.
- Performance microbenchmarks for eBPF filters under realistic load.
- Fault-injection tests for controller stability under adversarial load.

Recommendations
- Add test harness for replay/forensic validation.
- Add benchmark artifacts to `02_docs/benchmarks/` and CI job to run them.
- Expand unit tests to cover telemetry normalization and Kalman drift detectors.
