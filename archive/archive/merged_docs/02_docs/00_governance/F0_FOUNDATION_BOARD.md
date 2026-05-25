# Phase F0: Foundation Sprints (Hardening & Verification)

This document is the master execution board for the final push of PhoenixOS Foundation (Phase F0).

## Wave-Based Execution Strategy
- **Recovery Sprint (F0-R):** Teams A0 (Contracts), B0 (Truth), C0 (State), D0 (Imports), E0 (Audit). **[COMPLETE]**
- **Wave 1:** Contracts (A), Truth Layer (B), State Runtime (C).
- **Wave 2:** Replay Hardening (E), Decision Layer (G).
- **Wave 3:** Warden Refactor (F), Determinism Lab (D).
- **Wave 4:** Security (I), Validation (J), Documentation (H).

---

## 0. RECOVERY SPRINT (F0-R) - COMPLETE
- [X] **Team A0: Contract Recovery**
  - [X] A0.1 Restore semantic Version struct
  - [X] A0.2 Keep Version/APILevel
  - [X] A0.3 Add CompatibilityMatrix
  - [X] A0.4 Add ContractHash
  - [X] A0.5 Add Deprecation registry
  - [X] A0.6 Restore ActuationClass
  - [X] A0.7 Restore PolicyContext
  - [X] A0.8 Add contract validator
  - [X] A0.9 Interface snapshot tests
  - [X] A0.10 Contract freeze CI (via `validate_imports.py`)
- [X] **Team B0: Truth Recovery**
  - [X] B0.1 Restore TruthLedger
  - [X] B0.2 Restore hash chain
  - [X] B0.3 Add payload storage
  - [X] B0.4 Add verification pass
  - [X] B0.5 Add replay snapshots
  - [X] B0.6 Add evidence wrapper
  - [X] B0.7 Add corruption detector
  - [X] B0.8 Add replay export
  - [X] B0.9 Add ledger benchmarks
  - [X] B0.10 Add integrity tests
- [X] **Team C0: State Recovery**
  - [X] C0.1 Restore Registry
  - [X] C0.2 Restore history
  - [X] C0.3 Restore current state
  - [X] C0.4 Add transition audit
  - [X] C0.5 Add rollback
  - [X] C0.6 Add metrics
  - [X] C0.7 Add illegal detector
  - [X] C0.8 Add replay state tests
  - [X] C0.9 Add serializer
  - [X] C0.10 Add snapshot export
- [X] **Team D0: Import Freeze**
  - [X] D0.1 Import validator (Block AI/Memory/Nexus)
  - [X] D0.2 Cross-module checker
  - [X] D0.3 Dependency graph export
  - [X] D0.4 CI enforcement (Scripted in `05_tools/validate_imports.py`)
- [X] **Team E0: Foundation Audit**
  - [X] E0.1 Contracts Audit: **YELLOW** (Restored, needs full suite)
  - [X] E0.2 Truth Audit: **RED** (Ledger restored, hash chain verified, but coverage low)
  - [X] E0.3 State Audit: **RED** (Registry restored, history active, needs transition rules)
  - [X] E0.4 Replay Audit: **GREEN** (Deterministic replay verified in Arbiter)
  - [X] E0.5 Decision Audit: **GREEN** (Deterministic scoring verified in Arbiter)
  - [X] E0.6 Containment Audit: **YELLOW** (Warden stable, needs cooldown/budget audit)

---

## 1. TEAM A: CONTRACTS (Priority: CRITICAL)
- [ ] A1: Finalize semantic Versioning across all modules
- [ ] A2: Implement universal `Policy` validation in Arbiter
- [ ] A3: Add interface compatibility tests (Snapshot tests)
- [X] A4: Remove direct imports (Replay->Warden, Arbiter->Internal Replay, Decision->Telemetry) - **VERIFIED**
- [ ] A5: Build dependency graph checker
- [ ] A6: Add import freeze validator (ACTIVE: `validate_imports.py`)
- [ ] A7: Create `CONTRACT_MATRIX.md`
- [ ] A8: Create interface evolution rules
- [ ] A9: Add semantic versioning
- [ ] A10: Contract regression tests

## 2. TEAM B: TRUTH LAYER (Priority: CRITICAL)
- [ ] B1: Ledger integrity verifier (Unit tests for `TruthLedger.Verify`)
- [ ] B2: Replay snapshot engine (Integrate `TruthLedger` with Replay)
- [ ] B3: Evidence garbage collector
- [ ] B4: Evidence compression
- [ ] B5: Hash chain repair detector
- [ ] B6: Replay consistency checker
- [ ] B7: Timeline merger
- [ ] B8: PID evidence index
- [ ] B9: Event provenance tracker
- [ ] B10: Immutable archive export

## 3. TEAM C: STATE RUNTIME
- [ ] C1: State migration map
- [ ] C2: Compat registry
- [ ] C3: Illegal transition detector (Implement rules in `Registry.Transition`)
- [ ] C4: State audit logs (Verified active in `Registry`)
- [X] C5: Recovery validator
- [ ] C6: Transition metrics
- [ ] C7: FSM replay tests
- [ ] C8: Runtime snapshots
- [ ] C9: State serializer
- [X] C10: State rollback support (Verified active in Registry)

## 4. TEAM D: DETERMINISM LAB (Priority: MAX)
- [ ] D1: Fork bomb scenario
- [ ] D2: Reverse shell scenario
- [ ] D3: Beacon scenario
- [ ] D4: Port scan scenario
- [ ] D5: File exfiltration scenario
- [ ] D6: Ransomware simulation
- [ ] D7: SSH abuse scenario
- [ ] D8: Process injection scenario
- [ ] D9: Persistence attempt scenario
- [ ] D10: CPU exhaustion scenario
- [ ] D11: Memory exhaustion scenario
- [ ] D12: Log flooding scenario
- [ ] D13: Replay corruption scenario
- [ ] D14: Timeline divergence scenario
- [ ] D15: Hash tampering scenario

## 5. TEAM E: REPLAY HARDENING
- [ ] E1: Event ordering verifier
- [ ] E2: Duplicate removal
- [ ] E3: Missing event detector
- [ ] E4: Divergence engine v2
- [ ] E5: Timeline density scoring
- [ ] E6: Replay checkpointing
- [ ] E7: Replay cache
- [ ] E8: Multi-run comparator
- [ ] E9: Event normalization validator
- [ ] E10: Replay benchmark suite

## 6. TEAM F: WARDEN REFACTOR
- [ ] F1: Remove legacy state logic
- [ ] F2: Move aliases to `compat.go`
- [ ] F3: Add transition guards
- [ ] F4: Add recovery budgets
- [ ] F5: Add containment cooldown
- [ ] F6: Add state lock protection
- [ ] F7: Add audit events
- [ ] F8: Add replay gating
- [ ] F9: Add authorization hooks
- [ ] F10: Add rollback path

## 7. TEAM G: DECISION LAYER
- [ ] G1: Decision bus persistence
- [ ] G2: Confidence normalization
- [ ] G3: Evidence weighting
- [ ] G4: Policy weighting
- [ ] G5: Vote replay
- [ ] G6: Consensus history
- [ ] G7: Authorization audit
- [ ] G8: Confidence decay
- [ ] G9: Merge validator
- [ ] G10: Decision export

## 8. TEAM H: DOCUMENTATION TEAM
- [X] H1: `SYSTEM_IDENTITY.md`
- [X] H2: `NON_GOALS.md`
- [X] H3: `IMPORT_RULES.md`
- [X] H4: `CONTRACT_RULES.md`
- [X] H5: `STATE_MODEL.md`
- [X] H6: `TRUTH_MODEL.md`
- [X] H7: `DECISION_MODEL.md`
- [X] H8: `LAB_GUIDE.md`
- [X] H9: `VALIDATION_RULES.md`
- [X] H10: `REPLAY_SPEC.md`
- [X] H11: `EVIDENCE_SPEC.md`
- [X] H12: `PHASE_LOCK.md`

## 9. TEAM I: SECURITY TEAM
- [ ] I1: Threat taxonomy
- [ ] I2: Attack catalog
- [ ] I3: Reverse shell signatures
- [ ] I4: Fork bomb indicators
- [ ] I5: Exfiltration indicators
- [ ] I6: Beacon patterns
- [ ] I7: Privilege escalation map
- [ ] I8: Persistence techniques
- [ ] I9: Recovery playbooks
- [ ] I10: Risk matrix

## 10. TEAM J: VALIDATION TEAM
- [ ] J1: Replay determinism test
- [ ] J2: Evidence determinism test
- [ ] J3: Decision determinism test
- [ ] J4: State determinism test
- [ ] J5: Containment determinism test
- [ ] J6: Cross-run validation
- [ ] J7: Multi-seed validation
- [ ] J8: Failure injection
- [ ] J9: Load validation
- [ ] J10: Regression suite

---

## Final Unlock Condition (ALL MUST BE GREEN)
- [ ] Contracts stable
- [ ] Truth immutable
- [ ] Replay deterministic
- [ ] State reproducible
- [ ] Decision repeatable
- [ ] Containment gated
