# RFC: Phoenix Monitor

## 1. Description
The Phoenix Monitor provides L3 telemetry analysis. It consumes raw events from the Phoenix Bus and produces anomaly signals.

## 2. Specification
- **Entropy Engine:** Shannon Entropy + KL Divergence.
- **Signal Filters:** Discrete-time Kalman filters for noise reduction.

## 3. Interface
```go
func AnalyzeStream(data []byte) AnomalyScore
func FilterSignal(val float64) float64
```
