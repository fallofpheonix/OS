# Runbook: Replay Corruption Detection

## 1. Indicator
`[REPLAY] Error: Hash mismatch detected on replayed transaction`

## 2. Action Steps
1. **Locate Mismatch:** Identify the packet index causing the mismatch.
2. **Normalizer Check:** Verify that timestamp variables have been zeroed during normalizer execution.
3. **Ledger Audit:** Check if the transaction logs were modified or tampered with in transit.
