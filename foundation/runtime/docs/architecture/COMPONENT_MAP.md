---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Runtime — Component Map

> Last verified: 2026-06-04

The `foundation/runtime` package is the core execution substrate of Phoenix OS. It coordinates consensus, policy enforcement, fork detection, invariant checks, and kernel system call instrumentation.

## Component Breakdown

```
foundation/runtime/
├── constitution/              # Invariant enforcement engine
│   ├── engine.go
│   └── invariant.go
├── nexus_coordination/        # BFT consensus layer
│   ├── consensus.go
│   └── bft.go
├── arbiter/                   # Multi-node consensus bridges
│   ├── consensus_bridge.go
│   └── validator.go
├── truth/                     # Forensic truth and fork detection
│   ├── seal.go
│   ├── fork_detector.go
│   └── recovery.go
├── authority/                 # Policy management and audits
│   ├── manager.go
│   └── policy.go
├── kernel/                    # System-level call hooks & eBPF
│   ├── ebpf_probe.go
│   ├── ebpf_loader.go
│   └── enforcer.go
└── adapters/                  # Contracts boundary wrappers
    ├── replay_adapter.go
    └── warden_adapter.go
```

### Component Details

1. **Invariants Engine (`constitution/`)**
   - Implements validation loops that assert system invariants after each state modification.

2. **BFT Consensus Layer (`nexus_coordination/` & `arbiter/`)**
   - Coordinates replication state machines using a Byzantine Fault Tolerant consensus engine.
   - Enforces transaction validities and hashes consistency checks.

3. **Truth & Forensic Audit (`truth/`)**
   - Uses cryptographic seals to lock log fragments.
   - Detects linear history forks and drives rollback recovery.

4. **eBPF Kernel Agent (`kernel/`)**
   - Loads and manages eBPF probes for tracing process activity, disk I/O, and IPC.
   - Triggers process kill actuators upon policy violations.
