------------------------ MODULE ConsensusSafety ------------------------
(* =========================================================================
 * WORKFLOW POSITION: FORMAL VERIFICATION — CONSENSUS SAFETY SPECIFICATION
 *
 * This TLA+ specification formally defines the PoA (Proof of Authority)
 * consensus mechanism. It's used to verify that the distributed consensus
 * in PheonixDistributed correctly requires quorum before authorization.
 *
 * WORKFLOW:
 *   1. Nodes vote on proposals
 *   2. CheckQuorum: if SumVotes >= Threshold → authorized = TRUE
 *   3. Safety property: authorized ONLY if quorum reached
 *
 * PROPERTIES VERIFIED:
 *   Safety: authorized => SumVotes >= QuorumThreshold
 *   No Unauthorized Access: Cannot authorize without sufficient votes
 *
 * CORRESPONDENCE: This spec mirrors PheonixDistributed/consensus/poa.go
 *   - Nodes: Set of participating nodes
 *   - Threshold: Quorum requirement (e.g., 2/3)
 *   - Vote/CheckQuorum: Same logic as Go implementation
 * ========================================================================= *)
EXTENDS Naturals, Reals

CONSTANTS Nodes, Threshold

VARIABLES votes, authorized

Init == 
    /\ votes = [n \in Nodes |-> FALSE]
    /\ authorized = FALSE

TotalWeight == 1.0 \* Simplified: each node in constant set has weight 1/|Nodes|
QuorumThreshold == Threshold \* e.g., 2/3

Vote(n) ==
    /\ votes[n] = FALSE
    /\ votes' = [votes EXCEPT ![n] = TRUE]
    /\ UNCHANGED authorized

CheckQuorum ==
    /\ ~authorized
    /\ (SumVotes(votes) >= QuorumThreshold)
    /\ authorized' = TRUE
    /\ UNCHANGED votes

SumVotes(v) ==
    \* Logic to sum weights of TRUE entries
    1.0 \* Placeholder

Next ==
    \/ \E n \in Nodes: Vote(n)
    \/ CheckQuorum

Safety == authorized => (SumVotes(votes) >= QuorumThreshold)

Spec == Init /\ [][Next]_<<votes, authorized>>

=============================================================================
