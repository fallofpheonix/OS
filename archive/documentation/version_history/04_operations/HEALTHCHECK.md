# Operations: Health Check & Verification Matrix

Verify the runtime status of all services using the checklist below.

## 1. Process & Ports Verification
Confirm all six core processes are active:
```bash
ps aux | grep phoenix
```

Verify that TCP listeners for the IPC channels are bound correctly:
```bash
lsof -i -P -n | grep LISTEN
```

## 2. Telemetry & Log Monitoring
Audit the active log streams for error indications:
- **eBPF Events:** `tail -f logs/kernel.log`
- **Replay parities:** `tail -f logs/replay.log`
- **Warden FSM transitions:** `tail -f logs/containment.log`
- **Recovery & snapshots:** `tail -f logs/recovery.log`

## 3. Metrics Export Audit
Verify the metrics output JSON includes latency, count, and error registers:
```bash
cat metrics/export.json
```
Ensure fields like `TransitionCount`, `Rollbacks`, and `ReplayMismatch` are incrementing correctly.
