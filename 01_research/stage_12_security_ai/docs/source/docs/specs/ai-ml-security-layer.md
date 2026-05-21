# AI/ML Security Support Layer

## Purpose

Define how AI and machine learning support security operations for Linux-derived, cloud-connected, containerized, or security-focused OS images.

AI/ML augments security controls. It does not replace:

- Kernel hardening.
- Least-privilege IAM.
- Package signing.
- Network segmentation.
- Runtime policy enforcement.
- Human incident ownership.

## Scope

Applies to:

- Kali-derived security images.
- Arch-derived developer/security workstations.
- Cloud or container runtime deployments.
- CI/CD security scanning.
- Telemetry processing and alert triage.

Out of scope for early scratch-kernel milestones:

- In-kernel ML inference.
- Autonomous exploit generation.
- Uncontrolled automated remediation.
- Training on sensitive telemetry without data governance.

## Capability Matrix

| Capability | Technology | Function |
|---|---|---|
| Anomaly detection | Machine learning | Builds baselines from network traffic, access logs, process behavior, and container events; flags deviations such as lateral movement or data exfiltration |
| Malware classification | Machine learning | Extracts binary, metadata, syscall, or behavioral features to classify suspicious files beyond static hashes |
| Threat intelligence | NLP / graph ML | Extracts and correlates indicators of compromise from logs, reports, feeds, and internal case notes |
| Alert triage | Generative AI / LLM | Converts high-volume low-fidelity alerts into incident summaries, timelines, and analyst action queues |
| Automated remediation | Generative AI with policy guardrails | Suggests or executes constrained actions such as IAM reduction, pod isolation, host quarantine, or firewall updates |
| Code vulnerability review | ML / LLM / SAST assist | Reviews source, dependency metadata, and CI output for vulnerability patterns |
| Telemetry parsing | Specialized models | Normalizes logs from kernel, auditd, Falco, cloud APIs, Kubernetes, and application traces |

## Data Flow

```text
Sensors
  -> Log collection
  -> Normalization
  -> Feature extraction
  -> Detection models
  -> Correlation
  -> Triage summary
  -> Human approval or policy-gated remediation
  -> Audit log
```

## Required Inputs

Linux host telemetry:

- `auditd` logs.
- Kernel logs.
- Systemd journal.
- Process start/exit events.
- Package inventory.
- Authentication logs.

Container telemetry:

- Image scan results.
- Runtime syscall events.
- Kubernetes audit logs.
- Pod identity and namespace metadata.
- Network policy events.

Cloud telemetry:

- IAM changes.
- Object storage access.
- VPC flow logs.
- Security group changes.
- Control-plane audit logs.

CI/CD telemetry:

- SAST output.
- DAST output.
- Dependency scans.
- Container image scans.
- IaC scan results.

## Controls Around AI

Required:

- Human approval for destructive remediation.
- Allowlist of permitted automated actions.
- Audit log for every model-generated recommendation.
- Prompt and model version tracking for generated incident summaries.
- Redaction of secrets before LLM processing.
- No direct model access to production credentials.

Prohibited:

- Blind execution of generated shell commands.
- Automatic IAM broadening.
- Training on customer or target data without approval.
- Sending sensitive telemetry to external APIs without data classification.

## Core Security Tooling

AI/ML sits above these controls.

### Cloud And Container Security

| Tool Class | Role |
|---|---|
| CSPM | Audits AWS, Azure, or GCP posture for exposed storage, permissive IAM, public services, and drift |
| Container scanning | Detects vulnerable packages and unsafe image contents before registry push |
| Runtime detection | Uses syscall, process, and file events to detect suspicious runtime behavior |
| Kubernetes policy | Restricts pod-to-pod and pod-to-service communication |

Typical components:

- Image scanner.
- Kubernetes admission controller.
- NetworkPolicy enforcement.
- Falco or equivalent runtime detector.
- Cloud audit log collector.

### Zero Trust Architecture

Requirements:

- Authenticate every workload.
- Authorize every request.
- Enforce least privilege.
- Segment networks by service boundary.
- Rotate credentials.
- Log identity decisions.

Controls:

- IAM policy review.
- Service identity.
- Short-lived credentials.
- Micro-segmentation.
- Device or workload posture checks.

### Application Security

Required CI/CD gates:

- SAST.
- Dependency scanning.
- Container image scanning.
- IaC scanning.
- Secret scanning.

Optional runtime controls:

- DAST.
- WAF.
- RASP.
- API schema validation.

### Rule-Based And Graph Systems

Baseline controls:

- YARA.
- Suricata.
- SIEM rules.
- Sigma rules.
- Graph analytics for privilege paths, lateral movement, and attack chains.

AI/ML may prioritize, tune, or summarize these systems. It must not silently remove deterministic detection logic.

### Fuzzing And Vulnerability Discovery

AI/ML-supported techniques:

- Coverage-guided input generation.
- Bug-prone function ranking from code metrics and history.
- Crash deduplication.
- Triage summary generation.

Model output is advisory. Reproduction and root cause analysis remain required.

## Automated Remediation Policy

Allowed low-risk actions:

- Open ticket.
- Attach incident summary.
- Add temporary alert suppression with expiry.
- Quarantine non-production pod.
- Disable known-compromised test credential.

Requires approval:

- Production workload isolation.
- IAM policy changes.
- Firewall changes.
- Host quarantine.
- Package removal.
- User lockout.

Forbidden by default:

- Data deletion.
- External notification without review.
- Credential rotation without owner confirmation.
- Production network-wide deny rules.

## Model Evaluation

Track:

- False positive rate.
- False negative rate.
- Mean time to triage.
- Mean time to containment.
- Analyst override rate.
- Remediation rollback rate.

Minimum validation:

- Replay known benign telemetry.
- Replay known incident telemetry.
- Validate model does not expose secrets in summaries.
- Validate generated recommendations match policy.

## AI-Originated Threats

Threats:

- AI-generated phishing.
- Password-guessing optimization.
- Signature evasion.
- Polymorphic malware generation.
- Model evasion.
- Training-data poisoning.
- Prompt injection against SOC assistants.

Controls:

- Adversarial test suites.
- Input sanitization.
- Model monitoring.
- Retrieval source allowlists.
- Human approval for incident conclusions.
- Immutable forensic logs.

## Failure Considerations

- Model outage must not block core security controls.
- High-confidence anomaly detection still requires evidence.
- LLM summaries are not authoritative forensic records.
- All generated recommendations must cite source telemetry.
- Automated actions must be reversible where possible.

## Project Documents

Detailed project-plan documents:

- [../ai-cyber/01-ai-cyber-overview.md](../ai-cyber/01-ai-cyber-overview.md)
- [../ai-cyber/02-threat-model.md](../ai-cyber/02-threat-model.md)
- [../ai-cyber/03-architecture.md](../ai-cyber/03-architecture.md)
- [../ai-cyber/04-evaluation.md](../ai-cyber/04-evaluation.md)
- [../ai-cyber/projects/01-ai-ids-design.md](../ai-cyber/projects/01-ai-ids-design.md)
- [../ai-cyber/projects/02-ml-malware-detector.md](../ai-cyber/projects/02-ml-malware-detector.md)
- [../ai-cyber/projects/03-llm-soc-assistant.md](../ai-cyber/projects/03-llm-soc-assistant.md)
