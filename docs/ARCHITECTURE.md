### 🧠 PhoenixOS System Design: Closed‑Loop Cybernetic Pipeline

PhoenixOS is a **Deterministic Cybernetic Security Runtime** that runs **on Linux**, not *as* Linux. It is not a general‑purpose OS but a **replay‑first, telemetry‑first, bounded‑control defense layer** that survives hostile environments.

The design is summarized as:

- **Immutable, append‑only evidence ledger.**  
- **Deterministic, logical‑time‑based replay.**  
- **Warden FSM with strict budgets, cooldowns, and reversibility.**  
- **AI‑advisory only, never direct control.**  
- **Hardware‑aligned**: 8GB M3 profile, pure‑Go + eBPF, no heavy Python/JVM/PyTorch.  

***

### 1. High‑Level Cybernetic Data Flow

All events flow through a **strict, unidirectional cybernetic pipeline**. Feedback only occurs via **explicit, cryptographically‑signed map updates**.

```text
┌─────────────────────────────────────────────────────────────────────────┐
│                           EXTERNAL ENVIRONMENT                          │
└──────┬───────────────────────────────────────────────────────────▲──────┘
       │ (1. Raw Network/Syscalls)                                 │ (8. Drop/Allow)
┌──────▼───────────────────────────────────────────────────────────┴──────┐
│ LAYER 1: KERNEL FAST-PATH (eBPF / XDP / WireGuard)                      │
│   ├─ XDP Ingress Filter (NIC driver level)                              │
│   ├─ cgroup/connect4 Egress Hooks                                       │
│   └─ eBPF Ring Buffer (Telemetry Emitter)                               │
└──────┬───────────────────────────────────────────────────────────▲──────┘
       │ (2. Raw Binary Telemetry)                                 │ (7. BPF Map Update)
┌──────▼───────────────────────────────────────────────────────────┴──────┐
│ LAYER 2: DETERMINISTIC USERLAND (Pure Go / PID 1)                       │
│   ├─ Normalizer (Canonical JSON / Bencode)                              │
│   ├─ Lamport Clock (Logical Timestamping)                               │
│   └─ TCS Engine (Telemetry Confidence Score via Sliding Window)         │
└──────┬───────────────────────────────────────────────────────────▲──────┘
       │ (3. Scored, Deterministic Event)                          │ (6. Actuation Decision)
┌──────▼───────────────────────────────────────────────────────────┴──────┐
│ LAYER 3: THE WARDEN & THE LEDGER (Control Plane)                        │
│   ├─ Arbiter (Anomaly Scoring / Math & Physics Models)                  │
│   ├─ Warden FSM (PID Control / Hysteresis / Budgets)                    │
│   └─ Phoenix Ledger (Merkle DAG / SHA-256 Causality Chain) ◄──────┐     │
└──────┬────────────────────────────────────────────────────────────┼─────┘
       │ (4. Read-Only Context)                                     │ (5. Explainability)
┌──────▼────────────────────────────────────────────────────────────┴─────┐
│ LAYER 4: ADVISORY AI LAYER (PhoenixMind)                                │
│   ├─ Local LLM (Qwen2.5 3B) + SQLite Vector Memory (BGE-Small)          │
│   ├─ Attack Graph Generator / Process Lineage                           │
│   └─ Cloud Router (Redacted Summaries to DeepSeek/Groq)                 │
└─────────────────────────────────────────────────────────────────────────┘
```

**Key invariant**:  
Same external inputs → same logical‑time events → same replay‑hashes → same Warden‑decisions.

***

### 2. Component Design Breakdown

#### **Layer 1: Substrate (Kernel Fast‑Path)**

**Pattern**: Shared‑memory BPF maps, minimum latency.  
**Requirement**: Execute in ≤100 ns; no complex decisions.

- **`xdp_ingress` / `egress_policy.c`**  
  - Read from `BPF_MAP_TYPE_HASH` maps populated by Warden.  
  - If IP / cgroup not in map → **instant drop / block**.  
- **eBPF Ring Buffer**  
  - Forward raw telemetry to userspace via `ringbuf`.  
- **WireGuard integration**  
  - Allowed egress traffic routed over `wg0`; enforced by kernel‑level routing.  

This layer is **“dumb pipe”** actuation: **no AI, no complex logic, only fast enforcement.**

***

#### **Layer 2: Deterministic Userland (Go PID1)**

Run as a **single statically‑linked Go binary** (`phoenix‑os` as PID1).

- **`Normalizer (canonical JSON / Bencode)`**  
  - Eliminate non‑deterministic map iteration, float‑ordering, field‑ordering.  
  - Same binary input → same canonical‑JSON → same hash.  
- **`Lamport Clock`**  
  - Global atomic tick counter for logical time; no `time.Now()`.  
- **`TCS Engine (Telemetry Confidence Score)`**  
  - Welford‑style σ, EWMA‑style drift‑detection, packet‑loss / jitter / entropy.  
  - If sensor integrity is suspect → **degrade or halt** instead of mis‑actuating.  

This layer is **the heart of determinism** and replay‑correctness.

***

#### **Layer 3: Warden & Ledger (Control + Evidence)**

- **`Warden FSM`**  
  - States: `NORMAL → SUSPICIOUS → CONTAINED → SAFE_MODE → RECOVER`.  
  - Uses **PID‑style control** (proportional‑integral) for isolation budgets, throttling ceilings, and hysteresis to prevent “yoyo” oscillation.  
  - Each action must be:
    - Budgeted (`per‑minute / per‑cluster`),  
    - Cooldown‑aware,  
    - Reversible.  
- **`Arbiter`**  
  - Applies **math‑ and physics‑based models** (variance, entropy, Kalman‑like filters, signal‑processing) to score anomalies.  
- **`Phoenix Ledger (Merkle‑style DAG)`**  
  - Every Warden‑actuation writes to a hash‑chain:
    - `hash = SHA‑256(ActionID + CauseID + PrevHash + LogicalTime)`.  
  - Stored as append‑only WAL‑like blocks on‑disk.  
  - Replay‑+LEDGER becomes the **authoritative source of truth**.

This layer is **bounded cybernetic control**: observable, explainable, never wild.

***

#### **Layer 4: Advisory AI (PhoenixMind)**

- **Read‑only consumption** of:
  - Ledger + Telemetry + Graphs.  
- **Air‑gapped from actuation**:  
  - No direct control over kernel, Warden, or replay.  
- **AI stack**:
  - Local 3B‑model (`Qwen2.5‑3B‑Instruct‑Q4_K_M` on `llama.cpp`) + `SQLite` + `BGE‑Small` embeddings.  
  - Local‑only “fast triage” (anomaly‑clustering, priority‑ranking).  
  - Cloud‑router for redacted‑summaries (structural‑refactoring, large‑context reports).  
- **Functions**:
  - Attack‑graph / process‑lineage generation.  
  - Explainability for Warden‑decisions.  
  - Workflow‑suggestion (e.g., “next‑step in IR playbook”).  

This keeps **AI strictly advisory** and bounded.

***

### 3. Data Integrity & Storage Architecture

- **Hot** (eBPF ringbuf)  
  - Bounded, ring‑buffered; if overflow, `TCS` drops and degrades gracefully.  
- **Warm** (SQLite Mem‑mapped)  
  - Canonical events, evidence‑blocks, Ledger‑roots.  
  - Used by local AI for fast querying and graph‑building.  
- **Cold** (Cryptographic Snapshots)  
  - Every `N` logical ticks, Ledger root‑hash signed and exported.  
  - Survives kernel‑rootkits because **external verifier** can check hashes.  

Model: **Append‑only, immutable storage** with checksummed blocks.

***

### 4. Deployment Architecture (LinuxKit Immutable Appliance)

- **Build**: `make os-image`  
  - Mainline Linux 6.x kernel, `security=lockdown`, `cgroup_no_v1=all`.  
  - Santa‑Barbara‑style **no‑SSH, no standard shells, no‑GNU‑coreutils**.  
  - Go‑PID1 + eBPF binaries + PhoenixGuard packed into `squashfs`.  
- **Output**: Single `.iso` or raw QEMU image.  
- **Boot**:  
  - Kernel + eBPF maps + PID1 Phoenix‑OS runtime initialize in milliseconds.  
  - Ready for live‑boot SOC‑style usage.  

***

### 5. Why This Is the “Best” Design for PhoenixOS

- **Replay‑as‑truth enforced**:  
  - Deterministic normalization + Lamport time + Merkle‑Ledger guarantee **exact reconstruction**.  
- **AI‑separation preserved**:  
  - PhoenixMind never touches kernel or Warden; only **read‑only, RAG‑style** consumption.  
- **Worst‑case robustness**:  
  - Kernel rootkits? Ledger + snapshots still verifiable.  
  - AI‑hallucinations? Warden‑budgets clamp damage.  
- **Hardware‑aligned performance**:  
  - Pure‑Go + eBPF + minimal userspace stay under 1–2 GB RSS, leaving 5–6 GB for M3‑hosted local‑LLM context.  
- **No scope‑singularity**:  
  - Not a “Linux‑killer”, not a Kubernetes‑competitor; a **secure, replay‑aware defense layer**.

***

### 6. Concrete Package / Folder Layout (for GitHub)

```bash
phoenix_os/
├── Makefile                        # build, assemble, package appliance image
├── phoenix_os.yml                  # LinuxKit manifest
├── Dockerfile                      # multi‑stage cross‑compile bridge
├── go.work                         # root workspace
└── phoenix_os/
    ├── main.go                     # PID1; ties Clock, TCS, Replayer, Warden, Ledger, AI
    ├── clock/
    │   └── lamport.go              # logical‑time counters
    ├── normalizer/
    │   └── normalizer.go           # canonical JSON/Bencode
    ├── tcs/
    │   ├── tcs.go                  # sliding‑window scoring
    │   └── monitor.go              # Welford‑variance, Kalman‑style drift
    ├── replay/
    │   └── replayer.go             # deterministic replay engine
    ├── ledger/
    │   └── ledger.go               # Merkle‑DAG / hash‑chain evidence blocks
    ├── warden/
    │   ├── warden.go               # FSM, budgets, cooldowns
    │   └── ebpf/
    │       ├── xdp_ingress.c       # XDP‑packet‑drop / allow
    │       └── egress_policy.c     # cgroup/connect4 egress block
    └── ai/
        ├── local_llm.go            # llama.cpp / Qwen2.5‑3B
        └── cloud_router.go         # redacted‑summary router
```
