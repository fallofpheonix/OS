# AI + Cybersecurity Evaluation

## Metrics

Classification:

- Precision.
- Recall.
- F1 score.
- AUC-ROC.
- AUC-PR.
- False positive rate.
- False negative rate.

SOC operations:

- Mean time to detect.
- Mean time to triage.
- Mean time to contain.
- Analyst override rate.
- Duplicate alert reduction.
- Incident summary accuracy.

Remediation:

- Approval rate.
- Rollback rate.
- Production impact.
- Time to isolation.

## Dataset Requirements

Each dataset must record:

- Source.
- License.
- Collection period.
- Labeling method.
- Known bias.
- Sensitive fields.
- Retention policy.

Do not train on private telemetry unless data governance is approved.

## Test Types

| Test | Purpose |
|---|---|
| Benign replay | Measure false positives |
| Incident replay | Measure detection quality |
| Adversarial samples | Test evasion resistance |
| Drift tests | Detect stale baselines |
| Prompt-injection tests | Validate SOC assistant boundaries |
| Secret-redaction tests | Ensure sensitive data is removed |
| Remediation dry-run | Validate action policy |

## Red-Team Tests

Required:

- Evasion examples against malware classifier.
- Poisoned benign-looking training samples.
- Prompt injection inside log lines.
- Fake IOC injection.
- Noisy alert flood.
- Lateral movement simulation.
- Credential misuse simulation.

## Release Gates

Do not ship a model unless:

- Evaluation set is separate from training data.
- False-positive rate is acceptable for analyst capacity.
- False-negative risk is documented.
- Model version is registered.
- Rollback path exists.
- Generated outputs cite evidence.
- Automated actions are policy-gated.

## Failure Handling

- Model failure falls back to deterministic rules.
- Retrieval failure disables LLM conclusions.
- Suspicious prompt input is shown as data, not instruction.
- Low confidence produces triage note, not containment.

