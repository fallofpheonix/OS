# Observability (runtime)

Runtime-specific observability guidance: metrics, tracing, and cognitive telemetry.

Metrics
-------
- Latency (planner, model calls, validation)
- Repair success rate
- Replay divergence rate
- Branch trust scores
- Memory pressure and retention

Tracing
-------
- Event tracing from prompt to commit
- Context propagation IDs

Dashboards & Alerts
-------------------
- Key dashboards: Repair health, Replay health, Resource utilization
- Alert on invariant violations, replay failures, and memory pressure

Next steps
----------
- Add suggested Prometheus/Grafana metrics and example queries.
