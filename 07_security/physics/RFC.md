# RFC: Incident Physics (L6)

## 1. Specification
The Physics engine processes the state vector of $N$ containers or processes. Each node has a state $\sigma \in \{+1, -1\}$.

## 2. Interface
```go
type StateVector []int8
func CalculateSDI(states StateVector) float64
func CalculateEnergy(states StateVector, coupling float64, field float64) float64
```

## 3. Implementation
Uses a fixed-size partition function summation for small clusters, or Monte Carlo approximations for large swarms.
