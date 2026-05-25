# PX-012 Truth immutability hardening

**Label**: core-hardening, F0-exit
**Status**: CLOSED

## Problem
`TestEvidenceMutation` in `phoenix_os/truth/audit_test.go` confirms that the current in-memory implementation allows mutation of the `Entries` slice, violating the core axiom of evidence non-repudiation.

## Tasks
- [x] Implement copy-on-write for ledger entries
- [x] Create read-only wrappers for evidence access
- [x] Implement immutable snapshot mechanism
- [x] Add periodic hash seal verification
- [x] Generate `TRUTH_IMMUTABILITY_REPORT.md`

## Verification
- `TestEvidenceMutation` now FAILS to mutate internal state due to deep copy returns.
- `VerifySeal` detects out-of-band mutations.
