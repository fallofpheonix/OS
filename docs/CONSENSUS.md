# PHOENIX MATRIX CONSENSUS SPECIFICATION

**Status:** PROTOTYPE
**Implementation:** 45%
**Confidence:** High (Ledger) / Medium (Identity) / Conceptual (Voting)

## Version 1.0 (Draft)

### 1. Overview [IMPLEMENTED]
Phoenix Matrix uses a **Ledger-Authoritative Deterministic State Machine**. 
Agreement is reached through a permissioned Byzantine Fault Tolerant (BFT) protocol 
operating over a stream of signed envelopes.

### 2. Validator Identity (CONSENSUS-015) [IMPLEMENTED]
*   **Identity:** Every validator possesses a persistent Ed25519 keypair.
*   **Persistence:** Private keys are stored in `~/.phoenix/identity.json` with `0600` permissions.
*   **Public Key:** Validators are identified by their hex-encoded Ed25519 public key.

**Rule:** A validator MUST use the same public key for all signatures within an Epoch.
**Example:** Validator signs an event with key `A`. Replay nodes verify using key `A`.
**Counterexample:** Validator rotates key mid-epoch without a governance event; replay nodes reject new signatures.

### 3. Validator Membership (CONSENSUS-009) [PARTIAL]
*   **Validator Set:** A set of authorized Ed25519 public keys stored in the `WorldState`.
*   **BFT Threshold:** The validator set MUST contain at least 4 nodes (`N >= 4`) to tolerate `f=1` Byzantine failure.
*   **Evolution:** The validator set evolves via `UPDATE_VALIDATOR` events.

**Rule:** Validator updates REQUIRE a Quorum Certificate (QC) proving `2f + 1` agreement.
**Example:** `ADD_VALIDATOR` event is accompanied by signatures from 3 out of 4 current validators.
**Counterexample:** A single validator attempts to add their own secondary key to gain majority control; the event is rejected by peers.

### 4. Quorum Rules (CONSENSUS-011) [PLANNED]
*   **Quorum Requirement:** `2f + 1` unique authorized signatures.
*   **Finality:** A state is final once a QC exists for its index.

### 5. Signed Envelopes (CONSENSUS-014) [IMPLEMENTED]
*   **Structure:** Wraps all consensus messages (Events, Votes, Certificates).
*   **Replay Protection:** Includes `Epoch`, `Sequence`, and `Timestamp`.

**Rule:** Envelopes with duplicate `Sequence` numbers within an `Epoch` MUST be rejected.
**Example:** Node receives message with sequence `100`, then later receiving another sequence `100` from the same validator results in an alert.
**Counterexample:** An attacker captures a signed vote and rebroadcasts it 10 minutes later; the node rejects it due to stale `Timestamp` or duplicate `Sequence`.

### 6. Ledger Integrity (LEDGER-003, CONSENSUS-003) [IMPLEMENTED]
*   **Hash Chain:** Each event contains `PrevHash`, forming an immutable chain.
*   **Temporal Bounds:** `event.Tick` must be `>= ws.Tick` and `<=` `ws.Tick + 100`.

**Rule:** `event.Tick` MUST be monotonic and bounded.
**Example:** Current tick is `10`. Event tick `11` is accepted. Event tick `111` is rejected (jump > 100).
**Counterexample:** Event tick `9` is submitted; replay nodes reject it as a "Time Paradox".

### 7. Performance Targets (CONSENSUS-005) [VALIDATED]
*   **Replay:** ~30,000 events/sec.
*   **Hashing:** ~6,000 states/sec (1k entities).

### 11. Slashing Conditions (CONSENSUS-016) [PLANNED]
Byzantine behavior results in immediate validator removal and economic penalties.

| Violation | Proof Required | Penalty |
| :--- | :--- | :--- |
| **Double Vote** | Two `Vote` messages for same Index but different `StateHash`. | Removal + Full Slashing |
| **Double Proposal** | Two `Event` envelopes for same Index but different content. | Removal + Full Slashing |
| **Time Warping** | Event `Tick` jump > `MaxTickSkip`. | Temporary Suspension |
| **Sequence Reuse** | Duplicate `Sequence` in same `Epoch`. | Warning / Reputation Loss |

**Rule:** Any node detecting a Slashing Condition MUST broadcast a `SlashingEvidence` envelope.
**Example:** Validator signs two different blocks at the same height; any peer receiving both generates evidence.
**Counterexample:** A validator goes offline due to network partition; this is not a slashing condition (liveness failure vs. safety failure).

---
*Authorized by Phoenix Sovereign Governance*
