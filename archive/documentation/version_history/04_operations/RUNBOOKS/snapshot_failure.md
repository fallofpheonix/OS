# Runbook: Snapshot Failure Diagnostic

## 1. Indicator
`[RECOVERY] Error: Failed to capture system state snapshot`

## 2. Action Steps
1. **Mutex Status:** Check if there are thread-safety locks blocking audit histories:
   ```bash
   go test -race ./phoenix_os/containment/...
   ```
2. **Buffer Limits:** Check if telemetry events overflowed memory storage.
3. **Write Lock Debug:** Confirm write locks are released promptly.
