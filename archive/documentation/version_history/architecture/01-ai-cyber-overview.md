# AI + Cybersecurity Overview

## Purpose

Define a project and study plan for applying AI/ML to cybersecurity while keeping deterministic controls, human review, and auditability intact.

## Defensive Roles

| Area | AI/ML Use |
|---|---|
| Anomaly detection | Cluster or model normal network, user, process, and cloud behavior; flag deviations |
| Malware classification | Classify binaries, scripts, documents, or behavioral traces using extracted features |
| Threat intelligence | Extract, normalize, and correlate IOCs from feeds, logs, reports, and case notes |
| SIEM/SOC | Reduce false positives, cluster related alerts, generate incident timelines |
| Network security | Detect intrusions, protocol anomalies, DDoS-like traffic, and lateral movement |
| Endpoint security | Detect suspicious process trees, file writes, API calls, and persistence behavior |
| IAM | Detect risky logins, impossible travel, MFA bypass patterns, and anomalous role usage |
| Phishing/fraud | Classify emails, URLs, domains, attachments, and transaction behavior |

## Offensive And Abuse Roles

AI can also support attackers:

- Phishing text generation.
- Password-guessing optimization.
- Adaptive signature evasion.
- Polymorphic malware mutation.
- Reconnaissance summarization.
- Social engineering at scale.

Project controls must assume attackers also use AI.

## Core Non-AI Security Controls

AI/ML is not the foundation. Baseline controls remain:

- YARA, Suricata, Sigma, and SIEM rules.
- SAST, DAST, dependency scanning, secret scanning.
- Container image scanning.
- CSPM and cloud audit logging.
- Kubernetes network policies and admission control.
- IAM least privilege.
- WAF and API gateway controls.
- EDR/runtime telemetry.

## Project Tracks

| Track | Output |
|---|---|
| AI-empowered IDS/IPS | Flow classifier and alerting design |
| ML-assisted malware detector | Static and behavioral malware classification design |
| LLM SOC assistant | Retrieval-backed alert triage and investigation assistant |

## Study Sequence

1. Logging and telemetry formats.
2. Rule-based detection.
3. Feature engineering.
4. Supervised and unsupervised ML.
5. Model evaluation.
6. Adversarial ML.
7. LLM/RAG safety.
8. Policy-gated remediation.

## References

- KPMG AI cybersecurity overview: https://kpmg.com/ch/en/insights/cybersecurity-risk/artificial-intelligence-influences.html
- CrowdStrike ML in cybersecurity: https://www.crowdstrike.com/en-us/cybersecurity-101/artificial-intelligence/machine-learning/
- Kaspersky AI cybersecurity definition: https://www.kaspersky.com/resource-center/definitions/ai-cybersecurity
- Fortinet AI cybersecurity glossary: https://www.fortinet.com/resources/cyberglossary/artificial-intelligence-in-cybersecurity

