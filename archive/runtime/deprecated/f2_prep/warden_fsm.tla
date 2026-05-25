--------------------------- MODULE warden_fsm ---------------------------
EXTENDS Integers, Sequences

CONSTANTS States, Triggers, PID, DwellTicks

VARIABLES current_state, last_tick, deescalation_count

TypeInvariant ==
    /\ current_state \in States
    /\ last_tick \in Nat
    /\ deescalation_count \in Nat

Init ==
    /\ current_state = "SAFE"
    /\ last_tick = 0
    /\ deescalation_count = 0

Transition(newState, trigger, tick) ==
    /\ tick > last_tick
    /\ current_state' = newState
    /\ last_tick' = tick
    /\ deescalation_count' = IF newState = "SAFE" THEN 0 ELSE deescalation_count + 1

\* Property: No oscillation within DwellTicks
NoOscillation ==
    \A t1, t2 \in Nat : (t1 < t2 /\ t2 < t1 + DwellTicks) => (current_state[t1] = current_state[t2])

=============================================================================
