# PhoenixOS Layer 2: Internal Orchestrators

Layer 2 agents manage the internal health and strategic coordination of PhoenixOS. Unlike Layer 1 (External), these agents have direct read-access to the Phoenix Bus and can issue commands to the Phoenix Warden (Actuation).

## Agents

### 1. Swarm Coordinator (`agents/internal/swarm/`)
- **Purpose:** Multi-agent coordination and node consensus.
- **Owns:** MARL (Multi-Agent Reinforcement Learning) state, swarm logic.
- **Inputs:** `phoenix_os/bus` (all telemetry).
- **Outputs:** Coordinated defense tasks.

### 2. Resource Optimizer (`agents/internal/optimizer/`)
- **Purpose:** Economic resource allocation.
- **Owns:** VCG Auctions, performance vs. security trade-offs.
- **Inputs:** System load, threat temperature (SDI).
- **Outputs:** Resource priority updates.

### 3. Self-Healer (`agents/internal/healer/`)
- **Purpose:** Automated recovery and system restoration.
- **Owns:** Rollback triggers, container/service restarts.
- **Inputs:** Incident physics alerts (L6).
- **Outputs:** Recovery actions.
