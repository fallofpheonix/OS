# RFC-001B: Anomaly Logic Specification

## Status
Proposed

## Context
Defining the mathematical and algorithmic foundations for detecting ransomware-like behavior (encryption bursts and rename storms) within the Pheonix telemetry pipeline.

## 1. Rename Velocity Tracker
Detects "Rename Storms" typical of ransomware moving or renaming files after encryption.

### Algorithm: Sliding Window Rate Estimation
- **Window Size ($W$):** 60 seconds (configurable).
- **Bucket Size ($B$):** 1 second.
- **Metric:** $R = \frac{\sum_{i=1}^{W} \text{rename\_events}_i}{W}$
- **Logic:** Each PID maintains a circular buffer of buckets. An alert is triggered if $R > \text{Threshold}_{rename}$.

## 2. Write Entropy Estimation
Detects high-entropy writes (indicative of encryption or compression) vs. low-entropy (standard text/logs).

### Algorithm: Shannon Entropy Approximation
To maintain performance (<5% CPU), we use a sampling-based Shannon entropy calculation on the first $N$ bytes of `vfs_write` buffers.

$$H(X) = -\sum_{i=1}^{n} P(x_i) \log_2 P(x_i)$$

- **Sampling:** Sample 256 bytes per 4KB write.
- **Frequency Table:** 256-bin array for byte values.
- **Optimization:** Use a pre-calculated log table or Taylor series approximation if floating point is too slow in the telemetry path.

## 3. Alert Triggering Thresholds
Thresholds are determined by the `anomaly_baseline` experiment but proposed as:

| Metric | Warning Threshold | Critical Threshold |
|--------|-------------------|-------------------|
| Rename Rate | 10 files/sec | 50 files/sec |
| Write Entropy | > 6.5 bits | > 7.5 bits |
| Write Velocity | 5 MB/sec | 20 MB/sec |

## 4. Correlation Strategy
A "High Confidence Ransomware Alert" is triggered if:
`Critical(Rename Rate) AND (Critical(Write Entropy) OR Critical(Write Velocity))`
