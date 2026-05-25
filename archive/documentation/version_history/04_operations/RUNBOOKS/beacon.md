# Runbook: Network Beacon Mitigation

## 1. Indicator
`[ARBITER] High frequency traffic pattern matching beacon signatures`

## 2. Action Steps
1. **Apply Bandwidth Throttle:** Limit bandwidth on the target interface:
   ```bash
   tc qdisc add dev eth0 root tbf rate 10kbit burst 32kbit latency 400ms
   ```
2. **DNS Quarantine:** Divert target traffic to sinkhole servers.
3. **Verify:** Check that the outbound beacon packets drop below warning thresholds.
