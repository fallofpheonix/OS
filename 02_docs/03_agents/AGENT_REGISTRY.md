# Agent Registry

## 1. Autonomous Agents

| Agent ID | Role | Input | Output | Permissions |
| :--- | :--- | :--- | :--- | :--- |
| **Arbiter** | Strategic Decider | L4 Graph, L6 SDI | L5.5 Policy | Read(L1-L4), Write(L5.5) |
| **Warden** | State Actuator | L5.5 Policy | L5 State | Read(L5.5), Write(L5) |
| **Sentinel** | Physics Monitor | L3 Signals | L6 SDI | Read(L3), Write(L6) |
| **Trace-Indexer** | Lineage Manager | L2 Events | L4 DAGs | Read(L2), Write(L4) |

## 2. Agent Configuration
- **Memory:** Each agent maintains a local context and accesses the 3-tier Trace storage.
- **Actions:** Defined by the Finite-State Controller transitions.
- **Escalation:** Agents must escalate to the Arbiter if confidence scores drop below 0.6.
