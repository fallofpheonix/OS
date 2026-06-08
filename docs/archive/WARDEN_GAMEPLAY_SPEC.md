---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# WARDEN.EXE: Gameplay Specification (Vertical Slice)

## 1. Overview
WARDEN.EXE is the primary human-machine interface for the PhoenixOS substrate. The player takes on the role of a **Sovereign Auditor** (or Warden), responsible for maintaining system stability (Coherence) against the forces of decay and chaos (Entropy).

## 2. The First 5-Minute Loop

### 0:00 - 0:30: The Boot Sequence
- **Visuals:** CRT-style flickering, bios-style log scrolling.
- **Audio:** High-pitched digital hum, rhythmic clicking.
- **Player Input:** `auth admin`
- **Information:** 
  - Substrate Status: ACTIVE
  - Entropy: 12% (STABLE)
  - Coherence: 0.95

### 0:30 - 1:30: Observation & The First Anomaly
- **Event:** A rhythmic pulse in the "Drift Monitor." 
- **Notification:** `[!] ANOMALY DETECTED: PID 5011 (AuthService) | Drift: 0.42`
- **Choice:**
  - `ignore`: Drift continues; global entropy rises (+1%).
  - `investigate`: Opens the Causal Graph view.
- **Tutorial Tip:** "Small drifts are common. High drifts indicate a breach or incompetence."

### 1:30 - 3:00: Investigation & Evidence Gathering
- **Visuals:** A node graph showing PID 5011 connected to `/etc/shadow` and `NetworkStack`.
- **Telemetry Feed:**
  - `5011: Attempted read of /etc/shadow`
  - `Security Context: CAP_NONE` (Anomaly confirmed)
- **Player Input:** `audit 5011`
- **Cost:** 5 Verification Credits.
- **Information:** "Evidence suggests unauthorized credential harvesting. Certainty: 0.88."

### 3:00 - 4:30: Actuation & The Warden FSM
- **Visuals:** The Warden FSM Status Panel.
- **Current State:** `WATCH`
- **Action:** Transition 5011 to `CONTAINED`.
- **System Response:** 
  - "Process isolated."
  - "Resource limits applied (CPU: 0%)."
  - "Network interface dropped."

### 4:30 - 5:00: Feedback & Reward
- **Summary:** "Anomaly 5011 Neutralized."
- **Rewards:**
  - +10 Trust (Global Resource)
  - +20 Verification Credits (Action Resource)
  - System Coherence: 0.97 (+0.02)
- **Progression:** Unlock "Causal Trace" (Level 2).

---

## 3. Design Rationale (Audit Response)

This section addresses the 35+ critical questions (101-135) raised during the Level 27 Architecture Audit.

### 3.1 Accountability & Courts (Q101-110)
- **Q101 (Bad Telemetry):** We distinguish "Ignoring Truth" from "Ignoring Bad Telemetry" via **Sensor Reputation**. If a sensor claims drift but independent sensors (N+1) disagree, the sensor's `ReputationScore` decays.
- **Q102 (Doctrine Validity):** Measured via **Counterfactual Simulation**. The system replays the ledger with alternative actions to see if a better outcome was possible.
- **Q103 (Intent):** Proven via **Pattern of Conduct**. Malice is defined as a sequence of actions that bypasses secondary safety checks (e.g., silencing an alarm before melting a reactor).
- **Q104-110 (Appeals):** Verdicts are `PROPOSED` until the **Appeals Cycle** completes. Players or other agents can challenge a verdict by providing counter-evidence or better simulations.

### 3.2 Civilizational Decay (Q111-120)
- **Q111 (Decay Priority):** Trust → Knowledge → Institutions → Education → Technology.
- **Q112 (Knowledge vs Tech):** Yes, civilization can retain technology it no longer understands (e.g., using a reactor as a "Cargo Cult" object).
- **Q115 (Preservation Cost):** Paid by **Guilds** and **Federations**. If a Guild collapses, its specialized knowledge is lost to the substrate.

### 3.3 Game Resources (Q121-125)
- **Primary Resource (Q121):** **Verification Capacity** (represented as Credits). It is the energy required to prove a truth.
- **Rarest Resource (Q122):** **Operator Attention** (Time). 
- **HP Equivalent (Q124):** **System Coherence**. When it hits zero, the substrate collapses (Game Over).
- **XP Equivalent (Q125):** **Governance Level**. Increases as you successfully adjudicate complex cases.

### 3.4 ECS & Simulation (Q126-135)
- **Entities:** Processes, Sensors, Guilds, Cities.
- **Components:** `TrustComponent`, `EntropyComponent`, `ReputationComponent`, `KnowledgeComponent`.
- **Systems:** `AuditSystem` (runs event-driven), `DecaySystem` (runs every 100 ticks), `EntropySystem` (runs every tick).
- **Determinism (Q129):** All state-mutating actions must be deterministic for Replay Verification. Visuals (VFX, UI animations) are asynchronous.
