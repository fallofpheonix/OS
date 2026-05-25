# RFC: Phoenix Warden

## 1. Description
Phoenix Warden provides L5 actuation. It receives target metrics (CPU/Entropy) and computes control outputs to damp anomalous behavior.

## 2. Specification
- **Controller:** PID (Proportional-Integral-Derivative).
- **Actuator:** `cgroups.cpu.max` or `kill -STOP`.

## 3. Interface
```go
func ComputeThrottle(setpoint, measured float64) float64
```
