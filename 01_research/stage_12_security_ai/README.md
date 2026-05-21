# Stage 12: Security AI

## Purpose

Synthesize security engineering, reverse engineering, telemetry, SOC workflows, and AI/ML into production-grade security AI systems.

## Scope

- UEBA and behavioral analytics.
- ML-based malware detection and clustering.
- Threat prediction and vulnerability risk scoring.
- SOC alert enrichment, correlation, and SOAR automation.
- LLM-powered SOC copilots with RAG.
- LLM security: prompt injection, jailbreaks, prompt leakage, indirect injection.
- Data poisoning, model poisoning, model extraction, and model inversion.
- Adversarial ML attacks and robustness evaluation.
- AI red teaming and security testing.

## Classification

- Type: `SECURITY_AI_RESEARCH`
- Status: `RESEARCH_ONLY`
- Difficulty: expert
- Estimated duration: 10-12 weeks
- Upstream prerequisites:
  - Stage 04 Security
  - Stage 05 Reverse Engineering
  - Stage 09 Telemetry
  - Stage 10 SOC
  - Stage 11 AI/ML
- Downstream blockers:
  - Stage 14 Automation
  - Stage 15 Security Distribution
  - Stage 19 Production

## Research Modules

| Module | Path |
|---|---|
| Phase 9 Research Plan | `phase_09_security_ai.md` |
| Phase 9 Build Gate | `build_gate.md` |

## Internal Dependency Order

```text
Behavioral analytics
-> Malware ML
-> Threat prediction
-> SOC automation
-> LLM security
-> Adversarial ML
-> AI red teaming
-> Integrated deployment
```

## Gate

Do not deploy autonomous security response without explicit containment limits, audit logging, rollback paths, and human approval boundaries.
