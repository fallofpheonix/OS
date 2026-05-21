# RFC: Phoenix Sentinel

## 1. Description
Phoenix Sentinel provides L6 physical telemetry. It processes the aggregate state of the system to determine overall "health" or "disorder".

## 2. Specification
- **Model:** Ising Spin Lattice.
- **Metric:** Security Disorder Index (SDI).

## 3. Interface
```go
func CalculateSDI(states []int8) float64
func PredictTransition(history []float64) bool
```
