# Project Status — PhoenixOS

Short summary
- Vision: Autonomous, thermodynamics-inspired OS security stack using telemetry → graph → physics → control → game → cloud.
- Current state: Mixed — telemetry, trace, basic SDI math reported "STABLE" in README; control/warden and kernel patches are partial or experimental.
- Branch activity: active PRs (e.g., Kalman drift PR on branch `feature/issue-706-kalman-drift-detection`).
- Tests: Unit tests exist for FSM/warden area; test coverage across repo is partial.

Key modules
- Telemetry (L3): `phoenix_os/monitor`, `02_docs/specifications/telemetry_agent_spec.md`
- Trace/Graph (L4): `phoenix_os/trace`
- Evidence Ledger: `phoenix_os/ledger`
- Warden/Control (L5): `07_security/control`, `phoenix_os/warden`
- Kernel/Evangelism (L1/L2): `10_kernel/`, eBPF fast paths
- Nexus/Swarm (L7): `07_security/nexus` (gossip, MARL) — in progress

Next immediate actions
- Create a PR grouping documentation validation artifacts and the small FSM threshold refactor (if not already included).
- Run repository-wide CI and unit tests; collect coverage.
- Triage and prioritize P0/P1 risks: kernel patch safety, evidence integrity, and CI stabilization.
