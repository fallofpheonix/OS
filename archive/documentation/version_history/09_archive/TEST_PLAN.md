# Test Plan & Verification Matrix

PhoenixOS requires rigorous verification to ensure determinism and zero state drift under adversarial stress.

## 1. Test Tiers

### Tier 1: Deterministic Replay Verification
- **Command:** `go test ./guard`
- **Method:** Replay identical sequence events through the mock adapter and check that the resulting Ledger genesis hashes and sequenceProof hashes are bit-for-bit identical.

### Tier 2: Concurrency & Race Testing
- **Command:** `go test -race ./warden ./ledger/src`
- **Method:** Stress the state controllers using parallel goroutines executing concurrent reads/writes (FSM actuations, ledger entry additions, budget resets) to ensure no data races or panics occur.

### Tier 3: TCS Math Stability Tests
- **Command:** `go test ./tcs`
- **Method:** Validate out-of-order sequence insertion, missing sequence IDs, and negative SeqID filters to check that TCS confidence metrics remain within $[0.0, 1.0]$.

## 2. Invariant Gates
- **Zero Race Conditions:** Build and verify with `-race` flag must return 0 issues.
- **FSM Cooldown:** Confirm that Warden rejects back-to-back state escalations inside the 10-tick cooldown window unless severity class >= ClassLocalIsolate.
