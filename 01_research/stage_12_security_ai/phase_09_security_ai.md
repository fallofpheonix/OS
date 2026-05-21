# Phase 9: Security AI

## Objective

Apply AI/ML to advanced security problems: UEBA, malware detection, threat prediction, SOC automation, LLM security, adversarial ML, and AI red teaming.

## UEBA

Required capabilities:

- Baseline establishment for normal user and entity behavior.
- Anomaly detection across user actions, endpoint activity, network flows, and application access.
- Peer group analysis.
- Risk scoring and adaptive thresholds.
- Context-aware analytics using time, location, device, network, and asset context.
- Concept drift detection.
- Explainability for anomaly decisions.

Data sources:

- Network flows, DNS, proxy logs.
- Authentication logs.
- Application access logs.
- Endpoint process and file telemetry.
- Email metadata and content signals.
- Web destination and browsing metadata.

Techniques:

- z-score, IQR, Gaussian mixture models.
- Isolation Forest, LOF, DBSCAN.
- Autoencoders.
- LSTM sequence models.
- Time-series trend and seasonality analysis.
- Graph analysis over user-resource and host-host interactions.

Primary use cases:

- Insider threat detection.
- Account takeover detection.
- Lateral movement detection.
- Privilege abuse detection.
- Exploited application behavior detection.

Failure considerations:

- High-dimensional sparse features.
- Rare labels and class imbalance.
- Concept drift.
- Legitimate behavior changes.
- Privacy and policy constraints.
- Analyst trust.

## ML Malware Detection

Feature classes:

- Static: file size, entropy, headers, sections, imports, strings, timestamps, packer signatures, CFG complexity.
- Dynamic: syscall sequences, API frequency, process behavior, file and registry writes, network behavior, memory allocation.
- Hybrid: execution timeline and behavior graphs.

Models:

- Random forest.
- XGBoost or LightGBM.
- Neural networks.
- SVM with RBF kernel.
- Stacking, voting, and cascade ensembles.

Required evaluation:

- AUC-ROC.
- Precision, recall, F1.
- False positive rate under benign enterprise samples.
- Evasion testing under packing, string changes, import manipulation, padding, and dead-code insertion.
- Explainability using SHAP or LIME.

Robustness controls:

- Adversarial training.
- Ensemble diversity.
- Stable feature selection.
- Feature validation.
- Model monitoring for drift and evasion.

## Threat Prediction

Prediction targets:

- Endpoint or account compromise probability over defined horizon.
- Vulnerability exploitation probability.
- Attack type.
- Threat actor or campaign class where evidence supports it.
- Time to detection and recovery.
- Breach cost or impact range.

Input features:

- CVSS and EPSS.
- Asset criticality.
- Exploit availability.
- Internet exposure.
- Historical incident data.
- Threat intelligence feeds.
- Campaign and actor TTPs.
- Organizational geography, sector, and control maturity.

Models:

- Time-series forecasting: ARIMA, Prophet, LSTM, Transformer.
- Binary and multi-class classifiers.
- Regression for cost, recovery time, and detection latency.
- Calibrated risk scoring.

Constraints:

- Calibrate predicted probabilities.
- Quantify false positive cost.
- Separate correlation from causation.
- Treat attribution as low-confidence unless independently corroborated.

## SOC Automation and Copilots

Alert enrichment:

- Asset context.
- User context.
- IOC reputation.
- Historical incidents.
- Malware lineage.
- ATT&CK mapping.
- Timeline reconstruction.

Correlation:

- Deduplicate related alerts.
- Merge events into attack chains.
- Prioritize using likelihood, impact, and confidence.
- Escalate based on policy.

SOAR playbooks:

- Phishing: block sender, quarantine message, notify user.
- Suspicious login: trigger MFA, revoke sessions, check lateral movement.
- Malware: isolate endpoint, collect artifacts, kill process.
- Data exfiltration: block destination, revoke credentials, notify stakeholders.

LLM SOC assistant capabilities:

- Natural-language security queries.
- Incident timeline generation.
- Alert summarization.
- Runbook retrieval.
- Threat intelligence retrieval.
- Detection rule generation for YARA, Snort, Sigma, SPL, KQL, or EQL.

LLM constraints:

- No direct execution of generated response actions.
- Tool calls require authorization boundaries.
- Retrieved context is untrusted.
- Outputs need citation or evidence links.
- Sensitive logs require redaction.

## LLM Security

Attack classes:

- Direct prompt injection.
- Indirect prompt injection through retrieved documents.
- Prompt leakage.
- Jailbreaks.
- Tool-use manipulation.
- Data exfiltration through generated queries or summaries.

Defenses:

- Input validation and suspicious-pattern detection.
- Context isolation between instructions and retrieved data.
- Output filtering.
- Secret redaction.
- Tool allowlists and parameter validation.
- Rate limiting.
- Adversarial test suites.
- Human approval for sensitive actions.

## Poisoning and Model Attacks

Data poisoning:

- Label flipping.
- Feature poisoning.
- Backdoor triggers.
- Source-specific poisoning.

Model attacks:

- Model extraction through API queries.
- Model inversion.
- Gradient leakage.
- Membership inference.

Defenses:

- Data provenance and source authentication.
- Distribution validation.
- Robust statistics.
- Multi-source corroboration.
- Poison-sample detection.
- Differential privacy where acceptable.
- API rate limits.
- Watermarking and fingerprinting.

## Adversarial ML

Attack classes:

- Evasion at test time.
- Training-time poisoning.
- Black-box transfer attacks.
- Model extraction.

Techniques:

- FGSM.
- PGD.
- Carlini-Wagner.
- Genetic algorithms.
- Evolutionary search.
- Malware-specific feature perturbation.

Security examples:

- Malware feature modification.
- Spam token obfuscation.
- Fraud timing and pattern manipulation.
- Network-flow mimicry.

Defenses:

- Adversarial training.
- Input validation.
- Out-of-distribution detection.
- Feature squeezing.
- Defensive distillation.
- Diverse ensembles.
- Randomized smoothing where applicable.

Evaluation:

- Attack success rate.
- Perturbation budget.
- Query count for black-box attacks.
- Robust accuracy.
- Gradient obfuscation checks.

## AI Red Teaming

Framework:

```text
Threat model
-> Attack design
-> Harness implementation
-> Execution
-> Impact analysis
-> Remediation
-> Verification
```

Test classes:

- Behavioral probing.
- Prompt injection.
- Evasion attacks.
- Poisoning attempts.
- Model extraction.
- Distribution shift.
- Fuzzing.
- Genetic adversarial search.
- Saliency and feature-importance inspection.

Metrics:

- Vulnerability discovery rate.
- Time to exploit.
- Impact severity.
- Fix deployment rate.
- Residual risk.

Disclosure:

- Keep findings internal unless coordinated disclosure is required.
- Preserve evidence.
- Track remediation and verification.

## Integration Architecture

Target system components:

```text
telemetry ingestion
-> feature extraction
-> model inference
-> risk scoring
-> alert enrichment
-> correlation
-> analyst review
-> constrained response automation
-> feedback loop
```

Invariants:

- Raw data provenance is preserved.
- Model version is recorded for every prediction.
- Features are reproducible.
- Alerts include explanation and evidence.
- Automated response is bounded and reversible.
- Human approval exists for destructive or business-impacting actions.

## Promotion Targets

Research outcomes promote to:

- `06_ai/anomaly_detection/` for UEBA models.
- `06_ai/classifiers/` for malware and risk classifiers.
- `06_ai/features/` for feature extractors.
- `06_ai/rag/` for SOC copilot retrieval systems.
- `06_ai/agents/` for constrained assistant workflows.
- `07_security/detections/` for detection logic.
- `07_security/soar/` for playbooks.
- `07_security/response/` for response workflows.
- `02_docs/threat_models/` for AI system threat models.
