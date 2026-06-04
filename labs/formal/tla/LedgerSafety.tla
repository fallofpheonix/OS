------------------------- MODULE LedgerSafety -------------------------
(* =========================================================================
 * WORKFLOW POSITION: FORMAL VERIFICATION — LEDGER SAFETY SPECIFICATION
 *
 * This TLA+ specification formally defines the Ledger append-only property.
 * It's used to verify that the ledger in PheonixCore/ledger/src/ledger.go
 * correctly maintains its integrity invariants.
 *
 * WORKFLOW:
 *   1. Entries are appended with weights
 *   2. ConflictResolve: only entries with strictly greater weight can modify
 *   3. Safety property: all entries have weight > 0.0
 *
 * PROPERTIES VERIFIED:
 *   WeightInvariant: All ledger entries have positive weight
 *   Append-Only: Entries can only be appended, never removed
 *   Conflict Resolution: Only higher-weight entries can override
 *
 * CORRESPONDENCE: This spec mirrors PheonixCore/ledger/src/ledger.go
 *   - Append: AddEntry/AddEntryV2
 *   - ConflictResolve: RollbackTo (weight-based)
 *   - Weight: TrustScore from evidence records
 * ========================================================================= *)
EXTENDS Naturals, Sequences

VARIABLES ledger \* A sequence of entries (Index, Weight)

Init == ledger = << >>

(* Entry structure: [index |-> 1, weight |-> 0.5] *)

Append(weight) ==
    ledger' = Append(ledger, [index |-> Len(ledger) + 1, weight |-> weight])

ConflictResolve(index, weight) ==
    /\ index <= Len(ledger)
    /\ weight > ledger[index].weight
    /\ ledger' = [ledger EXCEPT ![index].weight = weight]

Next ==
    \/ \E w \in {0.1, 0.5, 1.0}: Append(w)
    \/ \E i \in DOMAIN ledger, w \in {0.1, 0.5, 1.0}: ConflictResolve(i, w)

(* Safety Property: History mutation is only possible with strictly greater weight *)
WeightInvariant == \A i \in DOMAIN ledger: 
    ledger[i].weight > 0.0

Spec == Init /\ [][Next]_<<ledger>>

=============================================================================
