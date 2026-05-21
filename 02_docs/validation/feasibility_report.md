# Feasibility Report

Overall Flag: PHASED

Technical feasibility
- Telemetry, entropy/Kalman math, and DAG lineage: READY — implemented in userspace and tested.
- Strategic game and control loops: EXPERIMENTAL — prototype code exists but needs hardened performance testing.
- Kernel scheduler modifications and custom kernel: HIGH_RISK — requires deep kernel expertise and incremental verification.

Team & maintenance
- Maintenance cost: High for kernel-level changes; moderate for userspace modules.
- Need for specialists: kernel engineers, control theorists, ML engineers, distributed systems.

Timeline & costs (rough)
- MVP (userspace + eBPF hooks): 3–6 months (small team).
- Kernel-weighted scheduler: 6–18 months with safeguards and formal verification.
- Cluster/swarm: additive after stable telemetry & control (12–24 months).

MVP viability
- Yes: userspace-only MVP with eBPF probes and centralized cloud control plane simulation.
- Validate via simulation and replay before kernel push.
