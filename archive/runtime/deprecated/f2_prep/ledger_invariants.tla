--------------------------- MODULE ledger_invariants ---------------------------
EXTENDS Integers, Sequences, hashes

VARIABLES entries

\* Property: Hash Chain Integrity
HashChainIntegrity ==
    \A i \in 2..Len(entries) : entries[i].prev_hash = hash(entries[i-1])

\* Property: Immutability
NoMutation ==
    [][\A i \in 1..Len(entries) : UNCHANGED entries[i]]_entries

=============================================================================
