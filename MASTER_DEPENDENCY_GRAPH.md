# Phoenix Master Dependency Graph

## Execution Path Constraint
No implementation stage starts before its research prerequisites and gate conditions are complete.

### Core Dependency DAG
```mermaid
graph TD
    %% Telemetry Foundation
    Math[01: Telemetry Math / Entropy] --> T_Kernel[02: eBPF Kernel Probes]
    T_Kernel --> E_Bus[03: Phoenix Bus]

    %% Graphs
    E_Bus --> Graph[04: Process Lineage DAGs]
    
    %% State and Physics
    Graph --> Physics[05: Phoenix Sentinel / SDI]
    
    %% Decision and Games
    Physics --> Game[06: Phoenix Arbiter / Stackelberg]
    
    %% Actuation
    Game --> Control[07: PID Control / Throttling]
    Control --> Sched[08: Game-Aware Kernel Scheduler]
    
    %% Final
    Sched --> Swarm[09: Autonomous MARL Swarm]
```

## Violations Found and Resolved
1. **Premature AI (Resolved):** AI correlator (RFC-004) was listed without explicit dependency on Event Normalizer (RFC-007) and Phoenix Trace (RFC-006). Enforcing graph and normalized telemetry dependency.
2. **Missing State Estimation (Resolved):** Game theory (Stackelberg) must receive SDI (Physics) and Bayesian posterior probabilities before making mixed-strategy decisions.
3. **Kernel Work Before Validation (Resolved):** Phase E (In-Kernel Schedulers) strictly deferred until Phase D (Closed-loop PID) is experimentally validated via R031.
