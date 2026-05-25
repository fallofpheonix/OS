# Runbook: Reverse Shell Containment

## 1. Indicator
`[WARDEN] Alert: Unauthorized interactive bash execution on network port`

## 2. Action Steps
1. **Network Quarantine:** Call network containment to block the socket and quarantine the namespace:
   ```bash
   ip netns exec phoenix_ns iptables -A OUTPUT -j DROP
   ```
2. **Process Pause:** Pause the process executing the shell:
   ```bash
   kill -STOP <target-pid>
   ```
3. **Evidence Extraction:** Extract state snapshot and dump it into `TruthLedger` for forensics analysis.
