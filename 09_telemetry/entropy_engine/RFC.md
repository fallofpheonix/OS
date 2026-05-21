# RFC: Entropy Engine (L3)

## 1. Problem Statement
Ransomware encryption is often invisible to signature-based detectors. We need a mathematical primitive to detect the phase transition from structured data to high-entropy cipher text.

## 2. Specification
### 2.1. Shannon Entropy
Calculates the average information per byte:
$$H(X) = -\sum_{i=0}^{255} P(x_i) \log_2 P(x_i)$$

### 2.2. KL Divergence
Measures the distance between the observed distribution $P$ and a normal file-type baseline $Q$:
$$D_{KL}(P \parallel Q) = \sum P(x) \log \frac{P(x)}{Q(x)}$$

## 3. Interface
```go
type Result struct {
    Entropy      float64 `json:"entropy"`
    KLDivergence float64 `json:"kl_divergence"`
    IsAnomaly    bool    `json:"is_anomaly"`
}

func Calculate(data []byte, baseline []float64) Result
```

## 4. Failure Modes
- **Buffer underflow:** Input < 256 bytes may produce unreliable entropy.
- **Floating point precision:** Stability during log sum.
- **Zero Probability:** KL Divergence handling for $Q(x) = 0$.
