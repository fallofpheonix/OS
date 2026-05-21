# AI + Cybersecurity Threat Model

## Assets

- Telemetry streams.
- Training datasets.
- Feature extraction code.
- Model artifacts.
- Prompt templates.
- Retrieval index.
- Incident summaries.
- Remediation actions.
- Analyst decisions.
- Audit logs.

## Defenders

- SOC analyst.
- Security engineer.
- Platform engineer.
- Incident commander.
- Model owner.

## Adversaries

| Adversary | Capability |
|---|---|
| External attacker | Network attacks, phishing, credential attacks |
| Malware operator | Evasion, persistence, lateral movement |
| Insider | Legitimate access abuse |
| Cloud attacker | IAM abuse, storage exposure, control-plane changes |
| Model attacker | Prompt injection, data poisoning, evasion examples |

## Threats Defended Against

- Phishing.
- Credential stuffing.
- MFA bypass attempts.
- Insider misuse.
- Lateral movement.
- Data exfiltration.
- Zero-day exploitation signals.
- Malicious container images.
- Runtime container escape indicators.
- Cloud misconfiguration abuse.

## AI-Specific Threats

| Threat | Impact | Control |
|---|---|---|
| Model evasion | Malicious input classified benign | Adversarial tests, ensemble checks, rule fallback |
| Data poisoning | Corrupted training baseline | Dataset provenance, holdout validation, drift monitoring |
| Prompt injection | SOC assistant returns attacker-controlled guidance | Retrieval filtering, instruction hierarchy, output validation |
| Secret leakage | Sensitive logs exposed in model context | Redaction, classification, access controls |
| Hallucinated IOCs | Analyst wastes time or blocks wrong asset | Require source citations from telemetry |
| Unsafe remediation | Production outage | Approval gates and action allowlists |

## Trust Boundaries

```text
Production systems
  -> Telemetry collector
  -> Security data lake
  -> Feature pipeline
  -> Model inference
  -> Triage assistant
  -> Analyst
  -> Remediation system
```

Boundary requirements:

- Authenticate collectors.
- Validate telemetry schema.
- Redact secrets before model access.
- Separate training and production data.
- Log every generated recommendation.
- Require approval for destructive actions.

## Assumptions

- Logs are incomplete during active compromise.
- Models can be wrong.
- Attackers can adapt after observing detections.
- Human analysts remain accountable for incident closure.

## Open Decisions

- Cloud provider.
- SIEM platform.
- Model hosting location.
- Data retention period.
- Allowed automated remediation actions.

