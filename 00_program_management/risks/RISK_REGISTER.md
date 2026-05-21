# Risk Register

| Risk ID | Risk | Impact | Likelihood | Mitigation | Stage |
| --- | --- | --- | --- | --- | --- |
| R01 | Foundational gaps delay all downstream kernel/security work | High | High | Keep Stage_00 and Stage_01 gates strict; do not start custom kernel branch until QEMU/GDB/debugging competency exists | Stage_00_Foundations |
| R02 | Scope splits between Linux distro, hybrid OS, and scratch OS | High | High | Treat Linux/distro as implementation path, scratch kernel as later research/extension path; gate with ADR updates | Stage_17_Hybrid_OS |
| R03 | Unsafe malware handling | High | Medium | Encrypted samples, isolated lab, no live execution outside sandbox, chain-of-custody and legal policy | Stage_05_Malware_and_RE |
| R04 | Telemetry overhead or verifier failures | Medium | High | Benchmark eBPF programs; keep BPF logic minimal; push policy to userspace | Stage_09_eBPF_and_Telemetry |
| R05 | Dataset leakage, weak labels, or non-reproducible ML metrics | Medium | High | Dataset provenance, split discipline, baseline models, fixed seeds, documented metrics | Stage_11_AI_ML_Core |
| R06 | SOC stack becomes tool collection without normalized schema | High | Medium | Define SOC event schema before integrations; every alert keeps evidence references | Stage_10_SOC_Stack |
| R07 | Cloud/container expansion outruns local platform maturity | Medium | Medium | Gate Stage_13 on observability, networking, distro reproducibility | Stage_13_Containers_and_Cloud |
| R08 | Automation performs unsafe response actions | High | Medium | Human approval gates, dry-run mode, audit logs, rollback paths | Stage_14_Automation_and_SOAR |
| R09 | Production claims without operational controls | High | Medium | Require release gates, retention policy, RBAC, monitoring, performance budgets | Stage_19_Production_Platform |
