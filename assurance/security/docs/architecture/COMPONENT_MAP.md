---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Security — Component Map

> Last verified: 2026-06-04

The `assurance/security` package implements system-level containment, adversarial game solvers, process isolation sandbox controllers, security audit logging, and hardware emergency killswitches.

## Component Breakdown

```
assurance/security/
├── go.mod                     # Module configuration
├── warden.go                  # Main entry point for policy routing
├── engine/                    # Main Warden monitoring loop
│   └── warden.go
├── actuators/                 # Low-level enforcement hooks
│   ├── ebpf.go                # System call denial via eBPF
│   └── process.go             # SIGKILL process actuator
├── actuation/                 # Sandbox creation
│   ├── sandbox.go             # Namespace isolation setup
│   └── executor.go            # Command executor inside sandbox
├── audit/                     # Violation trace logger
│   ├── violation_log.go
│   └── jsonl_writer.go
├── policies/                  # Dynamic trust scores calculation
│   └── trust_matrix.go
├── emergency/                 # Global containment trigger
│   └── killswitch.go
├── physics/                   # Thermodynamic disorder modeling
│   ├── thermo.go
│   └── disorder/sdi.go        # State Disorder Index formulas
└── game/stackelberg/          # Defensive policy solver
    └── solver.go
```

### Component Details

1. **Warden Engine (`engine/`)**
   - Receives events from syscall adapters.
   - Evaluates actions using the trust matrix and thermodynamic indexes.
   - Dispatches isolation commands to actuators.

2. **Process Actuators (`actuators/` & `actuation/`)**
   - Spawns target processes in isolated namespaces (VFS mount, IPC namespaces, network drop).
   - Enforces syscall restrictions using loaded eBPF programs.

3. **Physics / Disorder Index (`physics/`)**
   - Measures system anomalies in thermodynamic entropy terms, computing a State Disorder Index (SDI) to trigger adaptive escalation.

4. **Stackelberg Solver (`game/stackelberg/`)**
   - Solves the game-theoretic Nash equilibrium for system defenses against attackers.
