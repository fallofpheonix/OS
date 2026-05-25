# Runbook: Rollback Execution Failure

## 1. Indicator
`[RECOVERY] Error: Rollback state verification failed`

## 2. Action Steps
1. **Transition Path:** Confirm the target rollback state transition path is authorized.
2. **Hash Check:** Verify that the restored global snapshot matches the current hash signature.
3. **Audit Log:** Review logs for trace events blocking state transitions.
