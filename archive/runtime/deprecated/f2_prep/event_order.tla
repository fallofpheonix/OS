--------------------------- MODULE event_order ---------------------------
EXTENDS Integers, Sequences

VARIABLES bus_queue

\* Property: Monotonic Ticks
MonotonicTicks ==
    \A i \in 2..Len(bus_queue) : bus_queue[i].tick >= bus_queue[i-1].tick

=============================================================================
