# Module: fraud-detector-ml

## Source
*New module — built directly inside Banking App*

## Purpose
ML-based transaction fraud detection: feature engineering on transaction patterns, model training (XGBoost/LightGBM), real-time inference via internal API, and alert generation for suspicious activity.

## Interface
```python
from fraud_detector_ml import FraudDetector, FeatureEngine, AlertManager

engine = FeatureEngine()
features = engine.extract(transaction)  # → velocity, amount deviation, geo anomaly, etc.

detector = FraudDetector(model="xgboost", threshold=0.85)
prediction = detector.predict(features)  # → FraudPrediction(score, label, explanation)

if prediction.is_fraud:
    alert_mgr = AlertManager()
    alert_mgr.create(transaction, prediction)
```

## Depends On
- [[05_PROJECTS/REUSABLE_MODULES/fpx-pipeline]] (feature extraction pipeline)
- [[05_PROJECTS/REUSABLE_MODULES/fpx-observability]] (transaction event stream)
- scikit-learn, XGBoost or LightGBM

## Used By
- Banking App (primary consumer — transaction monitoring)

## Extraction Status
NOT_STARTED

## Location
`~/engineering/workspace/active/banking-app/modules/fraud-detector/`

> [!IMPORTANT]
> This module lives **inside** the Banking App, not in shared-libraries. It's too tightly coupled to banking transaction schemas to be generic. If a second consumer emerges, extract to shared-libraries at that point.

## Key Files
| File | Role |
|------|------|
| `features.py` | Transaction feature engineering (velocity, deviation, geo) |
| `detector.py` | Model loading + inference with confidence scoring |
| `trainer.py` | Model training pipeline (offline, batch) |
| `alerts.py` | Fraud alert generation and management |
| `models/` | Serialized trained models |
| `data/` | Training/evaluation datasets |

## Why Not Standalone?

| As Standalone Project | As Banking App Module |
|----------------------|----------------------|
| A notebook with scikit-learn | A production ML service with real transaction data |
| No deployment story | Deployed alongside the app it protects |
| "I trained a model" | "I designed ML-driven fraud detection integrated into transaction processing" |
| No real data contract | Consumes from fpx-observability event stream |

## Quality Gates
- [ ] Tests passing
- [ ] Model eval metrics documented (precision, recall, F1)
- [ ] Feature engineering is deterministic
- [ ] Inference latency < 50ms P99
- [ ] README with model architecture and training process
- [ ] Version pinned

#module #new-build #priority/P2 #replaces/fraud-detection-standalone
