# Runbook: Fork Bomb Containment

## 1. Indicator
System load spikes, process count rises exponentially, and console displays:
`[WARDEN] Posture transition Alert: CPU load anomalous`

## 2. Action Steps
1. **Warden Auto-Containment:** Warden will automatically apply cgroup constraints:
   ```bash
   cgset -r pids.max=20 /sys/fs/cgroup/pids/phoenix_jail
   ```
2. **Retrieve Lineage:** Query the trace graph engine to find the ancestor process PID:
   ```bash
   go run cmd/trace/main.go --query-ancestor --pid=<target-pid>
   ```
3. **Throttle:** Apply process throttling to slow down scheduler slices.
4. **Kill Parent:** Once confirmed safe, kill the root fork bomb launcher parent process.
