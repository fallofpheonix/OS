# Idea Validation

Verdict: PARTIAL — Novel and defensible, but scope, kernel work, and operational costs require staged validation.

Summary
- Problem: Real-time autonomous containment of fast-moving host threats (ransomware, supply-chain attacks). Problem exists; target users include enterprise SOCs, cloud providers, and high-assurance environments.
- Frequency & Impact: High-impact but moderate frequency; cascade failures are rare but costly.
- Need for custom OS / kernel: Partial — many goals can be achieved in userspace + eBPF; full kernel replacement is high-risk and should be postponed until userspace models and replay/forensics mature.
- Need for AI: Useful for correlation and policy recommendation; not required for core telemetry, SDI, or control loops.

Alternatives
- Existing IDS/EDR + SIEM stacks (Velociraptor, OSQuery, Zeek, sysdig) for telemetry and detection.
- eBPF-based in-kernel filtering + userspace graph/evidence chain is a pragmatic compromise.

Recommendation
- Continue: telemetry, trace, evidence ledger, warden controls in userspace with eBPF hooks.
- Defer: invasive kernel scheduler changes until replay and end-to-end validation exist (Phase E scheduled but gated).
