# AI/ML For Threat Detection And Analysis

Practical implementation patterns for machine-learning models in security.

## 1. Anomaly Detection

### Definition

Anomaly detection identifies events or patterns that deviate from expected behavior.

Security examples:

- Unusual login times.
- Abnormal data transfer volume.
- Unexpected process execution.
- New destination geographies.
- Lateral movement patterns.

### Algorithms

- Isolation Forest: fast, scalable, high-dimensional anomaly detection.
- Local Outlier Factor: density-based local anomaly detection.
- Autoencoders: neural reconstruction-error anomaly detection.
- Gaussian Mixture Models: statistical distribution modeling.

### Implementation Example

```python
from sklearn.ensemble import IsolationForest
import pandas as pd

model = IsolationForest(contamination=0.05)
model.fit(feature_matrix)
scores = model.decision_function(feature_matrix)
labels = model.predict(feature_matrix)
```

## 2. Intrusion Detection

### Classification Models

- Random Forest: robust baseline and handles mixed features.
- XGBoost: high-accuracy gradient boosting for tabular data.
- Neural networks: complex sequence and feature learning.
- SVM: binary classification where feature space is controlled.

### Training Data Requirements

- Labeled benign and malicious examples.
- Feature engineering from raw telemetry.
- Held-out test set.
- Class imbalance handling.
- Clear label provenance.

## 3. Network Behavior Analytics

### UEBA

User and Entity Behavior Analytics models normal behavior and flags deviations.

Use cases:

- Login behavior.
- Data access patterns.
- Network traffic.
- Privilege escalation.
- Service-account misuse.

### Features To Analyze

- Login time and location.
- Bytes sent and received.
- Number of connections.
- Protocols used.
- Port diversity.
- Failed connection rate.
- DNS query entropy.

## 4. Malware Classification

### Feature Extraction

- Byte n-grams.
- API calls.
- Imports.
- File metadata.
- Entropy.
- Size.
- Sections.
- Strings.

### Model Training

Pipeline:

```text
Samples
  -> feature extraction
  -> label validation
  -> train model
  -> evaluate false positives
  -> deploy scanner
```

Do not upload private binaries or sensitive samples to third-party services without approval.

## 5. Data Sources For Training

### Public Datasets

- KDD Cup 99.
- NSL-KDD.
- UNSW-NB15.
- CTF datasets.
- Malware feature datasets where license permits.

### Internal Data

- Zeek logs.
- Suricata alerts.
- Osquery data.
- Auditd logs.
- Incident labels.

## 6. Evaluation Metrics

- Precision: how many raised alerts are true positives.
- Recall: how many real attacks are detected.
- F1 score: harmonic mean of precision and recall.
- ROC-AUC: performance across thresholds.
- False positive rate: analyst-load indicator.
- Detection latency: time from event to alert.

## 7. Deployment Patterns

### Real-Time Scoring

Deploy model as API.

Flow:

```text
Event
  -> feature extraction
  -> prediction endpoint
  -> score
  -> alert or enrichment
```

### Batch Analysis

Process large datasets daily or hourly.

Use:

- Historical anomaly discovery.
- Model validation.
- Top-risk event ranking.

### Edge Deployment

Deploy lightweight models close to sensors.

Use:

- Network ingress sensors.
- Endpoint agents.
- High-volume telemetry filtering.

