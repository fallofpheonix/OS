# AI-Empowered IDS/IPS Design

## Goal

Detect malicious or anomalous network behavior from flow, packet, and protocol telemetry.

## Inputs

- NetFlow.
- PCAP.
- Zeek logs.
- Suricata alerts.
- DNS logs.
- Cloud VPC flow logs.

## Feature Set

- Bytes per second.
- Packets per flow.
- Flow duration.
- Source/destination fan-out.
- Destination-port histogram.
- DNS query entropy.
- Failed connection rate.
- Protocol anomalies.

## Model Candidates

| Model | Use |
|---|---|
| Random Forest | Baseline supervised classifier |
| XGBoost / LightGBM | Higher-performance tabular classifier |
| Autoencoder | Unsupervised anomaly detection |
| LSTM | Sequence behavior over time |

## Output

```text
event_id
source
destination
timestamp
score
classification
evidence
recommended_action
```

## Evaluation

- F1.
- AUC-ROC.
- AUC-PR.
- False positive rate per hour.
- Detection latency.
- Analyst review accuracy.

## Constraints

- Do not block production traffic without policy approval.
- Preserve raw evidence for replay.
- Explain every alert with source features.
- Keep deterministic IDS rules active.

