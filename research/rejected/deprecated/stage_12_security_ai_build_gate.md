# Phase 9 Build Gate: Security AI

## UEBA System

- [ ] Build behavioral analytics engine.
- [ ] Ingest logs, network telemetry, and endpoint data.
- [ ] Engineer behavioral features.
- [ ] Establish normal behavior baseline.
- [ ] Train anomaly model: autoencoder, LOF, or isolation forest.
- [ ] Generate risk scores and alerts.
- [ ] Support real-time or near-real-time detection.
- [ ] Explain why each anomaly was flagged.

Detection use cases:

- [ ] Insider threat: unusual data access, lateral movement, scope deviation.
- [ ] Account takeover: unusual location/time, password spray, credential stuffing.
- [ ] Privilege abuse: unauthorized escalation, admin-account movement, sensitive operations.

Minimum metrics:

- [ ] Insider threat precision above 80 percent on synthetic or labeled test data.
- [ ] Account takeover detection rate above 85 percent.
- [ ] Privilege abuse precision above 75 percent.
- [ ] False positive analysis completed.

## ML Malware Detector

- [ ] Use EMBER or equivalent dataset with at least 100K samples, or document why unavailable.
- [ ] Extract static features: PE headers, imports, sections, strings, entropy.
- [ ] Optionally extract dynamic features: syscalls, APIs, behavior.
- [ ] Train random forest, XGBoost, LightGBM, or equivalent.
- [ ] Achieve AUC-ROC above 95 percent or document dataset limitations.
- [ ] Produce SHAP or LIME feature explanations.
- [ ] Explain individual predictions.

Robustness:

- [ ] Create adversarial feature modifications.
- [ ] Test import, section, packing, obfuscation, and padding perturbations.
- [ ] Measure accuracy degradation.
- [ ] Apply adversarial training.
- [ ] Measure robustness improvement.

Optional:

- [ ] Cluster malware families.
- [ ] Identify emerging variants.
- [ ] Report silhouette or NMI.

## Threat Prediction

- [ ] Build threat scoring system.
- [ ] Include CVSS or EPSS.
- [ ] Include asset criticality.
- [ ] Include exploit availability.
- [ ] Include historical likelihood.
- [ ] Output calibrated risk score or probability.
- [ ] Achieve AUC-ROC above 0.8 on held-out data or document limitations.
- [ ] Analyze precision-recall trade-off.
- [ ] Analyze false positive cost.
- [ ] Produce calibration curve.

Optional:

- [ ] Predict exploitation timing.
- [ ] Predict incident type.
- [ ] Predict affected sector or asset class.

## SOC Automation and Copilot

- [ ] Build alert enrichment system.
- [ ] Ingest alerts from IDS, SIEM, EDR, or equivalent sources.
- [ ] Enrich with asset, user, and threat intelligence context.
- [ ] Correlate related alerts.
- [ ] Generate analyst summary.
- [ ] Keep enrichment latency below 5 minutes.

SOAR:

- [ ] Design playbook for phishing, malware, lateral movement, or data exfiltration.
- [ ] Define branching logic.
- [ ] Integrate with security tools or mocks.
- [ ] Test execution.
- [ ] Measure success rate and time savings.

LLM assistant optional:

- [ ] Implement RAG over incident history, docs, and threat intelligence.
- [ ] Answer security questions with cited context.
- [ ] Generate YARA, Snort, Sigma, SPL, KQL, or EQL drafts.
- [ ] Evaluate answer quality with human review.

## LLM Security and Red Teaming

- [ ] Test direct prompt injection.
- [ ] Test indirect prompt injection through retrieved context.
- [ ] Attempt system prompt extraction.
- [ ] Test jailbreaks.
- [ ] Document successful attacks and impact.
- [ ] Propose mitigations.

AI red team:

- [ ] Select attack vectors: evasion, poisoning, prompt injection, extraction.
- [ ] Design attacks.
- [ ] Implement test harness.
- [ ] Execute attacks.
- [ ] Document what succeeded and failed.
- [ ] Deploy defenses.
- [ ] Verify improvements.

## Adversarial Examples

- [ ] Generate adversarial examples for ML detector.
- [ ] Use genetic algorithm, gradient-based attack, or malware-specific feature perturbation.
- [ ] Preserve input functionality where applicable.
- [ ] Measure evasion success rate.
- [ ] Apply adversarial training.
- [ ] Measure robust accuracy.

## Poisoning and Model Attacks

- [ ] Implement label-flipping attack at 1, 5, and 10 percent poisoning.
- [ ] Measure model degradation.
- [ ] Implement backdoor trigger attack.
- [ ] Demonstrate activation and impact.
- [ ] Implement robust aggregation or data filtering defense.
- [ ] Compare defended versus undefended results.
- [ ] Document accuracy trade-off.

## Integrated Security AI System

- [ ] Integrate UEBA.
- [ ] Integrate malware detection.
- [ ] Integrate threat prediction.
- [ ] Integrate alert enrichment.
- [ ] Provide unified interface or dashboard.
- [ ] Test with real or synthetic data.
- [ ] Measure latency and throughput.
- [ ] Deploy in production-like environment.

System metrics:

- [ ] Accuracy or precision-recall.
- [ ] False positive rate.
- [ ] Alert fatigue index.
- [ ] Analyst time savings.
- [ ] Detection-rate improvement.
- [ ] Time-to-detect improvement.
- [ ] Cost-benefit analysis.

## Documentation

- [ ] Architecture and data flow diagram.
- [ ] Model training and evaluation process.
- [ ] Feature engineering methodology.
- [ ] Attack vectors tested.
- [ ] Mitigations implemented.
- [ ] Performance benchmarks.
- [ ] Deployment instructions.
- [ ] Known limitations.

Red team report:

- [ ] Executive summary.
- [ ] Technical findings.
- [ ] Proof-of-concept attacks.
- [ ] Impact assessment.
- [ ] Prioritized mitigations.
- [ ] Fix verification.

## Exit Criteria

- [ ] AI system threat model promoted to `02_docs/threat_models/`.
- [ ] Feature extractors promoted to `06_ai/features/`.
- [ ] Models promoted to `06_ai/anomaly_detection/` or `06_ai/classifiers/`.
- [ ] RAG or assistant components promoted to `06_ai/rag/` or `06_ai/agents/`.
- [ ] Playbooks promoted to `07_security/soar/`.
- [ ] Detection logic promoted to `07_security/detections/`.
- [ ] No autonomous destructive response exists without human approval and rollback.
