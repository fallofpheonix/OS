# Module: fraud-detector-ml

## Identity
- **Slug**: fraud-detector-ml
- **Owner**: you
- **Location**: ~/engineering/infrastructure/shared-libraries/fraud-detector-ml/
- **Status**: ACTIVE
- **Version**: 0.1.0
- **Language**: Python
- **Created**: 2026-05-12
- **Last updated**: 2026-05-12

## One-liner
Applies a Random Forest classifier to transaction features (amount, distance, velocity) to detect and flag fraudulent activity.

## API interface
```python
FraudDetectorML()
Initializes the ML model (RandomForestClassifier).

FraudDetectorML.train(features: list, labels: list)
Trains the model on a dataset of transactions.

FraudDetectorML.predict(transaction_features: list) -> dict
Runs inference on a new transaction vector and returns fraud probability.
```

## Installation / import
```bash
pip install scikit-learn numpy
```

```python
# Import from shared library
import sys
sys.path.append('../../infrastructure/shared-libraries/fraud-detector-ml')
from detector import FraudDetectorML
```

## Usage example
```python
ml = FraudDetectorML()
ml.train(historical_data, labels)
result = ml.predict([txn.amount, txn.distance, txn.time_delta, txn.velocity])
if result['is_fraud']:
    flag_transaction()
```

## Configuration
Requires a trained model state or live training data. In production, the model weights should be pickled/serialized and loaded at runtime rather than training on the fly.

## Used by
| Project | Since | Notes |
|---------|-------|-------|
| [[05_PROJECTS/ACTIVE/machine-learning-fraud-detection]] | 2026-05-12 | uses v0.1.0 |
| [[05_PROJECTS/ACTIVE/ledger-core]] | 2026-05-12 | placeholder |

## Version history
| Version | Date | Change |
|---------|------|--------|
| 0.1.0 | 2026-05-12 | initial extraction |

## Test coverage
- Unit tests: yes
- Test file: ~/engineering/infrastructure/shared-libraries/fraud-detector-ml/tests/
- Coverage: partial
- Last test run: 2026-05-12

## Known failure modes
| Failure | Trigger condition | Workaround | Fixed in version |
|---------|-------------------|------------|-----------------|
| Concept Drift | Spending habits change over holidays | Retrain model weekly | planned |

## Dependencies
- External: `scikit-learn`, `numpy`
- Internal modules: pairs well with `impossible-travel` to generate the `distance_km` feature.

## Performance characteristics
- Time complexity: O(trees * depth) for prediction
- Space complexity: Size of the model object (~MBs)

## Status transition log
- EXPERIMENTAL → ACTIVE: verified in machine-learning-fraud-detection project.

## Related
- Concept notes: [[03_CORE_KNOWLEDGE/ml-quant/algorithms/Random Forest]], [[03_CORE_KNOWLEDGE/ml-quant/concepts/Anomaly Detection]]
- Failure library: [[06_FAILURE_LIBRARY/module-failures/fraud-detector-ml]]
