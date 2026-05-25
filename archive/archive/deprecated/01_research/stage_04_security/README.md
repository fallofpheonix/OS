# Stage 04: Security Foundations

## Purpose

Establish defensive security foundations required before SOC, telemetry, reverse engineering, security AI, and security distribution work.

## Scope

- Security principles: confidentiality, integrity, authenticity, non-repudiation, availability.
- Threat modeling and risk analysis.
- Authentication, identity, OAuth 2.0, OIDC, JWT, MFA.
- Authorization and access control: DAC, MAC, RBAC, ABAC, ACLs, capabilities.
- Cryptography basics for system design.
- Security logging, monitoring, alerting, and incident response.
- Detection and response systems: IDS, IPS, EDR, XDR, SIEM, SOAR.
- Threat hunting, MITRE ATT&CK, cyber kill chain, and CTI fundamentals.

## Classification

- Type: `FOUNDATIONAL_SECURITY`
- Status: `RESEARCH_ONLY`
- Difficulty: advanced
- Estimated duration: 7-9 weeks
- Upstream prerequisites:
  - Stage 00 Foundations
  - Stage 03 Networking
- Downstream blockers:
  - Stage 05 Reverse Engineering
  - Stage 07 Threat Intelligence
  - Stage 10 SOC
  - Stage 12 Security AI
  - Stage 15 Security Distribution

## Research Modules

| Module | Path |
|---|---|
| Phase 6 Research Plan | `phase_06_security_foundations.md` |
| Phase 6 Build Gate | `build_gate.md` |

## Internal Dependency Order

```text
Security principles
-> Threat modeling and risk analysis
-> Identity and access control
-> Cryptography and secure communication
-> Logging and monitoring
-> Detection and response frameworks
-> Incident response and forensics process
```

## Gate

Do not start implementation work in `07_security/`, SOC workflows, or security AI detection logic until the Phase 6 build gate is complete.
