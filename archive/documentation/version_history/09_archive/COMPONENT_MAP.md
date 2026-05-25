# Subsystem Component Map

Every module within PhoenixOS is assigned ownership, directory structure, and dependencies.

| Subsystem | Folder | Owner | Core Dependencies | Purpose |
| :--- | :--- | :--- | :--- | :--- |
| **Bus** | [bus/](file:///Users/fallofpheonix/os/phoenix_os/bus) | Platform Team | Go channels | Ingestion and queue-pressure scheduling |
| **Guard** | [guard/](file:///Users/fallofpheonix/os/phoenix_os/guard) | Kernel Team | eBPF (Stage 2) | Telemetry sorting and sequence proof hashing |
| **Ledger** | [ledger/](file:///Users/fallofpheonix/os/phoenix_os/ledger) | Security Team | SHA-256 / Allocator | Verifiable state evidence Merkle-DAG logging |
| **Monitor** | [monitor/](file:///Users/fallofpheonix/os/phoenix_os/monitor) | Research Team | Kalman / EWMA | Entropy drift modeling and frequency metrics |
| **TCS** | [tcs/](file:///Users/fallofpheonix/os/phoenix_os/tcs) | Research Team | Event Bus | Jitter and sequence-loss confidence scores |
| **Warden** | [warden/](file:///Users/fallofpheonix/os/phoenix_os/warden) | Platform Team | Event Bus | Finite state containment actuation |
| **Arbiter** | [arbiter/](file:///Users/fallofpheonix/os/phoenix_os/arbiter) | Research Team | Policy maps | Strategic game-theoretic decision solver |
| **AI Orchestrator** | [ai/](file:///Users/fallofpheonix/os/phoenix_os/ai) | AI Team | Subsystem Features | Asynchronous advisory loop and LLM driver |

## Subsystem Boundaries
Modules are strictly decoupled. No strategic layer (e.g. Arbiter) should directly manipulate Kernel states; all strategic modifications must route through the Warden FSM.
