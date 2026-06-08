# Phoenix Daemon (phoenixd) - Component Explanations

## Beginner (What it does)
`phoenixd` is the "brain" of PhoenixOS. It starts up all the different security tools, connects them together, and then runs through a list of security events to make sure everything is working correctly. It acts like a flight recorder and a security guard at the same time, keeping track of everything that happens and checking it against rules to keep the system safe.

## Intermediate (How it interacts)
The daemon initializes a central communication line called the "Bus". Every other part of the system (like the Monitor, the AI Orchestrator, and the Warden) talks to each other through this Bus. It also sets up a "Ledger" which is an unchangeable record of every event. When it starts, it checks its own "DNA" (boot integrity) and can even load special programs (eBPF) into the computer's core to watch for suspicious activity.

## Expert (Architectural Role)
`phoenixd` serves as the primary orchestration layer for the Phoenix Cybernetic Security Runtime. It implements a deterministic execution model where system state is derived from a replayed sequence of events, allowing for formal verification of security invariants at runtime. It integrates distributed consensus (ledger), reflexive actuation (eBPF/Warden), and AI-driven analysis into a unified security substrate. It handles the transition from boot-time integrity to runtime sovereignty.

## Code Review
- **Risk Score:** 9/10
- **Complexity Score:** 8/10
