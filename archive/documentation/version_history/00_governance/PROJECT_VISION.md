# PhoenixOS Strategic Route & Roadmap

## 1. The Mission: "Security as a Thermodynamic State"
PhoenixOS is not just an Operating System; it is a **Deterministic Cybernetic Security Runtime**. 
Traditional security is reactive. PhoenixOS treats security as a state of **minimal entropy**. When the system detects "disorder" (System Disorder Index - SDI), it autonomously "quenches" that disorder using a deterministic actuation loop (The Warden).

### Core Philosophy:
*   **Determinism is Sacred:** Every event must be replayable. If we cannot replay a hack exactly as it happened, we haven't solved it.
*   **AI is Advisory:** PhoenixMind (AI) provides strategic advice, but the Warden (Deterministic FSM) holds the "Circuit Breaker" veto.
*   **Zero-Trust Kernel:** Security doesn't sit "on top" of the OS; it is baked into the eBPF probes and fast-path enforcement.

---

## 2. The Architecture: The 7-Layer Phoenix Matrix
We build from the hardware up to the swarm:

*   **L1 Platform Integrity (Guard):** Microsecond-latency fast-path. If a packet is 99% likely to be a kill-switch, Guard drops it before the kernel even sees it.
*   **L2 Kernel Runtime (Kernel):** Real-time telemetry via eBPF. This is our "Nervous System."
*   **L3 Telemetry Math (Monitor):** Signal processing. We convert raw syscalls into Entropy and Kalman-filtered trends.
*   **L4 Graph Intelligence (Trace):** Causal Lineage. We don't just see a process; we see the "Parent-Child-Socket" DAG that created it.
*   **L5 Actuation & Control (Warden):** The Circuit Breaker. Moves the system between SAFE, WATCH, and COMPROMISED states.
*   **L5.5 Strategic Policy (Arbiter):** The Game Theorist. Decides if the "cost" of shutting down a service is worth the "benefit" of stopping an attack.
*   **L6 System Physics (Sentinel):** The "Brain." Monitors the thermodynamics of the entire OS.
*   **L7 Swarm Coordination (Nexus):** Distributed Consensus. Multiple PhoenixOS nodes agreeing that an IP is malicious across the network.

---

## 3. The Roadmap: From Foundation to Swarm

### Phase 1: Replay-Safe Foundation (COMPLETED)
*   **Goal:** Build the "Black Box" recorder and the deterministic bus.
*   **Key Artifacts:** Phoenix Ledger (Evidence Chain), Logical Clock, Event Bus with Priority Lanes.
*   **Status:** Verified. System can replay 10,000+ events with 0% divergence.

### Phase 2: The Real-World Nervous System (ACTIVE)
*   **Goal:** Move from simulated events to real Linux kernel telemetry.
*   **Tasks:**
    *   Deploy **eBPF Probes** for syscall, network, and file I/O monitoring.
    *   Implement **Phoenix Trace (L4)**: Building real-time causal graphs of process behavior.
    *   **RAG Memory Integration:** Giving PhoenixMind a long-term memory of previous attack patterns using a local vector store.
*   **Outcome:** A system that can "see" a real exploit in real-time.

### Phase 3: Strategic Autonomy (UPCOMING)
*   **Goal:** Enable the AI to make complex "Strategic Containment" decisions.
*   **Tasks:**
    *   Activate the **Phoenix Arbiter (L5.5)**: Strategic denial based on Stackelberg Game Theory.
    *   Refine **Confidence Scoring (TCS)**: Ensuring the AI only actuates when it is >95% certain.
    *   **Automated Forensics:** System automatically snapshots the RAM and Disk of a "Compromised" process for later replay.

### Phase 4: The Phoenix Nexus (FINAL)
*   **Goal:** Distributed defense.
*   **Tasks:**
    *   Implement **Swarm Gossip**: Nodes share threat intelligence instantly.
    *   **PoA (Proof of Actuation)**: Verifying that a node actually performed the security action it claimed to.
    *   **Self-Healing Clusters:** Swarm re-allocates workloads away from compromised nodes.

---

## 4. How We Build
1.  **Research First:** Every feature starts with a Mathematical or Physical model in `02_docs/06_research/`.
2.  **Deterministic Implementation:** We use Go for the runtime (speed + safety) and eBPF/C for the kernel fast-path.
3.  **Validation Loop:** No code is merged without a "Replay Test." We simulate the failure, verify the fix, and ensure the replay matches exactly.
4.  **Audit Trail:** Every change by an AI agent or human is logged in the `DECISION_LOG.md` and `CHANGELOG.md`.

**Current Focus:** Transitioning the "Replay-Safe" core into the real Linux environment via eBPF.
