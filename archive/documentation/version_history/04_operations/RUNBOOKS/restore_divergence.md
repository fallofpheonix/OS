# Runbook: Restore Divergence Detection

## 1. Indicator
`[RECOVERY] Error: State divergence detected post-restore`

## 2. Action Steps
1. **Compare History:** Compare original process/network/file history arrays with the restored history.
2. **Epoch Checks:** Verify timestamp normalizer reset values to `time.Unix(0, 0)`.
3. **Audit Verification:** Ensure sequence counters match exactly.
