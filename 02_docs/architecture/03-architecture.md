# AI + Cybersecurity Architecture

## Data Pipeline

```text
Telemetry Sources
  -> Collection Agents
  -> Normalization
  -> Feature Engineering
  -> Model Inference
  -> Correlation
  -> Alert / Block / Quarantine Recommendation
  -> Analyst Review
  -> Remediation
  -> Audit Log
```

## Telemetry Sources

| Source | Examples |
|---|---|
| Network | NetFlow, PCAP, Zeek, Suricata |
| Endpoint | auditd, EDR events, process trees, file access |
| Cloud | IAM, storage, VPC flow logs, control-plane audit logs |
| Containers | Kubernetes audit logs, image scans, runtime syscalls |
| AppSec | SAST, DAST, dependency scans, secrets, IaC scans |
| Identity | Login events, MFA events, role usage |

## Feature Engineering

Network:

- Bytes per second.
- Packet count.
- Destination-port distribution.
- Flow duration.
- DNS entropy.
- Connection fan-out.

Endpoint:

- Process ancestry.
- File write paths.
- API/syscall sequence.
- Persistence locations.
- Binary metadata.

Identity:

- Login location.
- Device posture.
- Role usage.
- Session duration.
- Failed-login cadence.

## Model Layer

| Use Case | Candidate Model |
|---|---|
| Flow classification | Random Forest, XGBoost, LightGBM |
| Sequence behavior | LSTM, temporal CNN, transformer encoder |
| Anomaly detection | Isolation Forest, autoencoder, clustering |
| Malware static features | LightGBM, deep neural network, autoencoder |
| Alert summarization | LLM with retrieval and source citation |
| Attack graph ranking | Graph algorithms, graph ML |

## Action Layer

Actions:

- Alert.
- Enrich.
- Open ticket.
- Recommend block.
- Recommend quarantine.
- Execute policy-gated containment.

Default:

- Alert and summarize.
- Do not block production traffic without approval.

## Storage

| Store | Purpose |
|---|---|
| Raw log store | Forensics and replay |
| Feature store | Training and inference consistency |
| Model registry | Versioned model artifacts |
| Vector index | SOC assistant retrieval |
| Audit log | Decisions, prompts, outputs, remediations |

## Safety Constraints

- Model outputs must cite source telemetry.
- Secrets must be redacted before LLM context.
- Remediation must use typed APIs, not free-form shell.
- Production actions require explicit approval unless pre-authorized.
- Every recommendation must be replayable from stored evidence.

