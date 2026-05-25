# PhoenixOS System Metrics

## 1. Core Metrics

### 1.1 System Disorder Index (SDI)
- **Definition**: Thermodynamic measurement of system entropy.
- **Target**: < 1.0 (Normal), > 7.9 (Critical).

### 1.2 Telemetry Confidence Score (TCS)
- **Definition**: Confidence in incoming telemetry based on loss and drift.
- **Target**: > 0.95.

### 1.3 Actuation Budget
- **Definition**: CPU/Network/Memory consumed by security actions.
- **Limit**: < 5% total host resources.

## 2. Operational Metrics
- **Settling Time**: Time for the Warden to stabilize post-actuation.
- **Recall Rate**: Percentage of attacks successfully detected and contained.
- **Replay Mismatch**: Frequency of non-deterministic events detected during replay.
