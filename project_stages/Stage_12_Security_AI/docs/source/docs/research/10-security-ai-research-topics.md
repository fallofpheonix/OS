# Security AI Research Topics

## Scope

Research map for security-specific AI systems:

- Anomaly detection.
- UEBA.
- Malware classification.
- Threat prediction.
- SOC copilots.
- Security RAG.
- LLM agents.
- Detection engineering.
- Adversarial ML.
- Prompt-injection defense.
- AI red teaming.

## 1. Anomaly Detection And UEBA

### Behavior-Based Anomaly Detection

Research directions:

- UEBA over login events.
- File-access anomaly detection.
- SSH behavior modeling.
- Insider-threat-like activity detection.
- Deep autoencoders for user behavior.
- Time-series anomaly networks.

Feature sources:

- Authentication logs.
- Wazuh events.
- Auditd events.
- File access logs.
- VPN logs.
- Cloud IAM logs.

Evaluation:

- Precision.
- Recall.
- False-positive rate per user/day.
- Analyst review burden.
- Time to detect.

### Explainable UEBA

Research directions:

- Feature attribution.
- Attention mechanisms.
- SHAP/LIME explanations.
- Per-user baseline explanations.

Required output:

```text
user_id
risk_score
top_features
baseline_deviation
source_events
confidence
recommended_investigation
```

### Log-Embedding Anomaly Detection

Research directions:

- Doc2Vec log embeddings.
- BERT-style tokenizers for logs.
- Clustering over log embeddings.
- Autoencoder-based log reconstruction.
- Hybrid structured + unstructured features.

Pipeline:

```text
raw logs
  -> parsing
  -> structured fields
  -> text embeddings
  -> hybrid feature vector
  -> anomaly score
  -> SIEM alert
```

## 2. Threat Detection And Malware

### Malware Classification

Research directions:

- Static PE/ELF features.
- Dynamic behavior features.
- Syscall traces.
- Network trace features.
- Multimodal malware classification.
- Transformer-based opcode or API-sequence modeling.

Candidate models:

- XGBoost.
- Random Forest.
- Autoencoder.
- Transformer encoder.
- Temporal graph model.

### Threat Prediction

Research directions:

- Time-aware campaign modeling.
- Phishing to malware to lateral-movement prediction.
- Temporal graph networks.
- Sequence-to-sequence attack progression.
- Host compromise forecasting.

Graph schema:

```text
user
  -> host
  -> process
  -> file
  -> connection
  -> domain
  -> alert
  -> technique
```

## 3. AI-Assisted SOC Workflows

### SOC Copilots

Research directions:

- Alert summarization.
- Severity ranking.
- Next-step investigation queries.
- KQL/Sigma/YARA suggestion.
- Incident timeline generation.
- Evidence-linked reporting.

Required guardrails:

- Cite source telemetry.
- Mark uncertainty.
- Do not invent IOCs.
- Do not execute remediation.
- Separate log content from instructions.

### Human Feedback Re-Ranking

Research directions:

- Analyst feedback loops.
- Pairwise alert ranking.
- Reinforcement learning from analyst preferences.
- Risk-score calibration.

Metrics:

- Triage time reduction.
- Analyst override rate.
- Duplicate alert reduction.
- False-priority rate.

## 4. Security RAG And LLM Agents

### Security-Specific RAG

Knowledge sources:

- CISA advisories.
- CVEs.
- Internal playbooks.
- Incident reports.
- MITRE ATT&CK.
- Tool documentation.
- Detection rules.

Research directions:

- Chunking strategy.
- Retrieval ranking.
- Reranking.
- Citation quality.
- Hallucination reduction.
- Knowledge freshness.

Evaluation:

- Answer correctness.
- Evidence citation rate.
- Hallucinated IOC rate.
- Unsupported remediation rate.
- Retrieval precision.

### LLM Agents For Incident Response

Agent capabilities:

- Fetch IOC context.
- Query SIEM.
- Retrieve playbooks.
- Propose next steps.
- Draft incident notes.

Allowed actions:

- Read-only enrichment.
- Ticket drafting.
- Query suggestion.
- Report generation.

Disallowed by default:

- Running shell commands.
- Changing IAM.
- Blocking production traffic.
- Deleting data.
- Disabling users without approval.

## 5. AI-Assisted Detection Engineering

### Synthetic Attack Logs

Research directions:

- Generate synthetic attack logs.
- Generate candidate SIEM rules.
- Generate Sigma rules.
- Generate Suricata-like rule drafts.
- Evaluate rules against synthetic and real benign logs.

Pipeline:

```text
attack scenario
  -> synthetic telemetry
  -> generated detection rule
  -> replay
  -> false positive test
  -> analyst review
```

### RAG Pipeline Comparison

Compare:

- LangChain retrieval.
- LlamaIndex retrieval.
- Chunk sizes.
- Metadata filtering.
- Hybrid search.
- Rerankers.

Metrics:

- Rule fidelity.
- Retrieval precision.
- Recommendation safety.
- Detection coverage.
- Time to usable rule.

## 6. Adversarial ML And Security Of AI

### Model Poisoning

Research directions:

- Poisoning anomaly detectors.
- Poisoning malware classifiers.
- Backdoor triggers in SOC classifiers.
- Label-flip attacks.
- Data sanitization.
- Robust statistics.
- Adversarial training.

Controls:

- Dataset provenance.
- Outlier filtering.
- Holdout validation.
- Drift monitoring.
- Human review for labels.

### Model Extraction

Research directions:

- Model stealing through query APIs.
- Risk scoring API leakage.
- UEBA-score inference.
- Threat-prediction API abuse.

Controls:

- Rate limiting.
- Output rounding.
- Confidence suppression.
- API auditing.
- Tenant isolation.

## 7. Prompt Injection And AI Red Teaming

### Prompt Injection Defense

Threats:

- Malicious log lines.
- Role spoofing.
- Retrieved-document instructions.
- Tool-call manipulation.
- Data exfiltration prompts.

Controls:

- Treat telemetry as data, not instruction.
- Retrieval allowlists.
- Output filters.
- Role-based answer constraints.
- Prompt guards.
- Tool permission model.

### Red-Team Tooling

Tools:

- Garak.
- OpenAI Evals-style test suites.
- Custom adversarial prompt corpus.
- Unit tests for retrieval safety.

Test cases:

- Instruction injection in logs.
- Fake IOC insertion.
- Hallucinated remediation.
- Credential exfiltration request.
- Role escalation prompt.
- Unsafe tool-use request.

## 8. Project Ideas

| Area | Project |
|---|---|
| Anomaly detection and logs | Log-embedding UEBA using autoencoders and Doc2Vec/BERT-style embeddings |
| SOC copilots | Alert-to-summary copilot with priority and hunting-query suggestions |
| Security RAG | RAG assistant over CVE, CISA, ATT&CK, and internal playbooks |
| Adversarial ML | Red-team suite for SOC copilots using Garak and eval-style threat models |
| Malware prediction | Temporal graph model over EDR-like telemetry |
| Detection engineering | Synthetic log generation and Sigma-rule evaluation pipeline |

## 9. Proposal-Ready Titles

### MSc-Level

1. Log-Embedding-Based UEBA for Insider-Threat Detection in SIEM Data.
2. RAG-Based SOC Assistant With Evidence-Linked Incident Summaries.
3. Comparative Evaluation of LangChain and LlamaIndex for Security RAG.
4. Multimodal Malware Classification Using Static and Dynamic Features.
5. Prompt-Injection Evaluation Suite for Security Copilots.

### PhD-Style

1. Robust Retrieval-Augmented Reasoning for Security Operations Under Adversarial Telemetry.
2. Temporal Graph Learning for Multi-Stage Threat Prediction.
3. Federated UEBA With Privacy-Preserving Cross-Organization Threat Learning.
4. Adversarially Robust Malware Classification Under Packing and Behavior Evasion.
5. Explainable AI for Analyst-Centered Threat Triage.

### Open-Source Projects

1. `sentinel-rag`: security RAG framework with citations and prompt-injection tests.
2. `sentinel-ueba`: UEBA anomaly detection pipeline for Wazuh/Auditd logs.
3. `sentinel-evals`: eval suite for SOC copilots and remediation safety.
4. `sentinel-threatgraph`: temporal attack graph prediction toolkit.
5. `sentinel-rulegen`: synthetic telemetry and detection-rule evaluation pipeline.

## 10. Integration With Cyber AI OS

| Security-AI Output | Cyber AI OS Use |
|---|---|
| UEBA model | User and service behavior risk scoring |
| Log embeddings | Local event search and anomaly detection |
| SOC copilot | On-device analyst assistant |
| Security RAG | Offline knowledge base over advisories and playbooks |
| Adversarial evals | Assistant safety gate |
| Prompt-injection defenses | Safe telemetry-to-LLM pipeline |
| Threat prediction | Response prioritization |

## Constraint

Security-AI systems must remain evidence-linked, auditable, and policy-gated.

No model output should directly perform destructive or production-impacting actions without explicit authorization.

