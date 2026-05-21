# RFC: Phoenix Arbiter

## 1. Description
Phoenix Arbiter provides L5.5 strategic intelligence. It uses the global system state (SDI, Trace) as input to a payoff matrix.

## 2. Specification
- **Algorithm:** Strong Stackelberg Equilibrium (SSE) solver.
- **Output:** Mixed-strategy probability distribution for monitoring targets.

## 3. Interface
```go
func SolveSSE(payoffs PayoffMatrix) Strategy
```
