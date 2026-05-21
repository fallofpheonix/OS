# AI And ML Layer

## AI Stack

```text
Input sources
  -> Feature extraction
  -> ML engine
  -> Threat classification
  -> Response engine
  -> Firewall / sandbox / alert
```

## Input Sources

- System logs.
- Network packets.
- Processes.
- Memory dumps.
- File changes.
- User activity.
- Kernel audit events.
- Container runtime events.

## Models

| Model | Use |
|---|---|
| Isolation Forest | Unsupervised anomaly detection |
| Random Forest | Baseline classification |
| XGBoost | Tabular event classification |
| LSTM | Sequence behavior |
| Transformer | Opcode, log, or event sequence modeling |
| Autoencoder | Outlier detection and reconstruction error scoring |

## IDS Use Case

```text
Packet stream
  -> Feature extraction
  -> Anomaly detector
  -> Risk score
  -> IDS alert
  -> Containment recommendation
```

## AI Pipeline

```text
Events
  -> Kernel collector
  -> Feature builder
  -> ML inference
  -> Threat engine
  -> Response manager
  -> Firewall / sandbox / alert
```

## Runtime Constraints

- Inference must not block kernel-critical paths.
- Kernel sends events; userspace performs feature building and model inference.
- Models must be versioned.
- Every model decision must retain source event references.
- Automated containment requires policy gates.

## Technology Stack

| Layer | Candidates |
|---|---|
| Kernel/user probes | eBPF, auditd, tracepoints |
| ML runtime | PyTorch, ONNX Runtime, TensorRT |
| Storage | SQLite, Parquet, object storage, graph DB |
| Observability | Prometheus, Grafana, OpenTelemetry |
| Security sensors | Suricata, Zeek, YARA, Sigma |

