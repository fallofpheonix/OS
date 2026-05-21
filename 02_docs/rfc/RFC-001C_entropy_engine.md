# RFC-001C: SentinelOS Entropy Engine

## Status
Proposed

## Context
Implementation of Information Theory primitives for real-time telemetry analysis.

## Mathematical Primitives
### 1. Real-time Entropy (H)
Calculated on $N$-byte samples of `vfs_write` buffers.
$$H(X) = -\sum P(x) \log_2 P(x)$$

### 2. Relative Entropy (KL Divergence)
Measures the "surprise" of a sequence of syscalls relative to a known-good profile.
$$D_{KL}(P || Q) = \sum P(x) \log \frac{P(x)}{Q(x)}$$

## Operational Targets
- **H > 7.5 bits:** High suspicion of encryption.
- **D_KL Spike:** Indicates departure from normal process behavior.
