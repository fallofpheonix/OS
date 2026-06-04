--------------------------- MODULE warden ---------------------------
EXTENDS Naturals, Sequences

VARIABLES state, lamportClock, evidenceWeight

States == {"SAFE", "WATCH", "SUSPICIOUS", "CRITICAL", "COMPROMISED"}
Weights == 0..100

Init == 
    /\ state = "SAFE"
    /\ lamportClock = 0
    /\ evidenceWeight = 100

Next == 
    \E w \in Weights:
        /\ evidenceWeight' = w
        /\ lamportClock' = lamportClock + 1
        /\ \/ (state = "SAFE" /\ state' \in {"SAFE", "WATCH"})
           \/ (state = "WATCH" /\ (state' \in {"SAFE", "WATCH"} \/ (state' = "SUSPICIOUS" /\ w >= 50)))
           \/ (state = "SUSPICIOUS" /\ (state' \in {"WATCH", "SUSPICIOUS"} \/ (state' = "CRITICAL" /\ w >= 80)))
           \/ (state = "CRITICAL" /\ (state' \in {"SUSPICIOUS", "CRITICAL"} \/ (state' = "COMPROMISED" /\ w >= 95)))
           \/ (state = "COMPROMISED" /\ state' \in {"CRITICAL", "COMPROMISED"})

TypeOK == 
    /\ state \in States
    /\ lamportClock \in Nat
    /\ evidenceWeight \in Weights

NoIllegalJumps == 
    [][state = "SAFE" => state' \in {"SAFE", "WATCH"}]_state /\
    [][state = "WATCH" => state' \in {"SAFE", "WATCH", "SUSPICIOUS"}]_state /\
    [][state = "SUSPICIOUS" => state' \in {"WATCH", "SUSPICIOUS", "CRITICAL"}]_state /\
    [][state = "CRITICAL" => state' \in {"SUSPICIOUS", "CRITICAL", "COMPROMISED"}]_state /\
    [][state = "COMPROMISED" => state' \in {"CRITICAL", "COMPROMISED"}]_state

MonotonicClock == 
    [][lamportClock' > lamportClock]_lamportClock

EvidenceGates ==
    [][state' = "SUSPICIOUS" /\ state \in {"WATCH"} => evidenceWeight' >= 50]_state /\
    [][state' = "CRITICAL" /\ state \in {"SUSPICIOUS"} => evidenceWeight' >= 80]_state /\
    [][state' = "COMPROMISED" /\ state \in {"CRITICAL"} => evidenceWeight' >= 95]_state

=============================================================================