# PhoenixOS Recovery Model

## 1. Objective
Define the protocols for restoring the system to a `SAFE` state after a compromise or failure.

## 2. Recovery Protocol
1. **Detection**: Warden enters `RECOVERY` state.
2. **Isolation**: Current session/container is locked or discarded.
3. **Evidence Retrieval**: Fetch the last stable snapshot from the Truth Ledger.
4. **Integrity Verification**: Verify the hash-chain integrity of the snapshot.
5. **Rollback**: Revert the system state to the last verified stable point.
6. **Authorization**: Re-authorize system contracts and reset actuation budgets.
7. **Resume**: Transition Warden back to `SAFE`.

## 3. Stabilization & Limits
- **Recovery Budget**: Maximum of 3 automatic de-escalations per session.
- **Manual Override**: If the budget is exhausted, a manual operator reset is required.
- **Evidence-First Recovery**: No recovery is permitted without a validated evidence chain.
