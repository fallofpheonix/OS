---\nStatus: Planned\nImplementation: 5%\nConfidence: Conceptual\n---\n# PhoenixFormal — Formal Verification Layer

> TLA+ specifications, architecture rules, dependency policies, and invariant proofs for the PhoenixOS ecosystem.

## Overview

PhoenixFormal provides the formal verification backbone for PhoenixOS. It contains TLA+ specifications for critical system properties, architecture rules for code organization, and dependency policies for module management.

**All critical system properties must have corresponding TLA+ specifications.**

## Repository Structure

```
PhoenixFormal/
├── tla/                    # TLA+ formal specifications
│   ├── GuardFSM.tla        # Warden FSM state machine
│   ├── ConsensusSafety.tla # PoA consensus properties
│   └── LedgerSafety.tla    # Ledger append-only properties
├── contracts/              # Event and schema contracts
│   └── events/             # Event emission points and sequencing
├── dependency-map/         # Module dependency tracking
│   ├── DEPENDENCY_GRAPH.md
│   ├── MODULE_USAGE_MAP.md
│   └── PORT_ALLOCATION_MAP.md
├── build-validation/       # Build validation scripts
│   ├── validate-all.sh
│   ├── validate-go.sh
│   ├── validate-python.sh
│   ├── validate-ts.sh
│   └── validate-rust.sh
├── agent-governance/       # AI agent governance
│   ├── invariants/
│   └── verification/
├── governance/             # Dependency and purity scanning
├── health-engine/          # Ecosystem health scoring
├── repo-registry/          # Repository registration
├── sync-engine/            # Synchronization configuration
├── runtime-state/          # Runtime state tracking
├── schemas/                # JSON schemas
├── events/                 # Event emitter
├── extraction-engine/      # Fork extraction
├── extraction-tracker/     # Extraction tracking
├── git-governance/         # Git governance
├── runtime/                # Root resolution
├── bootstrap/              # Service templates
└── benchmarks/             # Runtime stress tests
```

## Core Principles

1. **TLA+ First**: Critical properties must be formally specified
2. **Architecture Rules**: Consistent code organization across repositories
3. **Dependency Policy**: No circular dependencies, no hidden coupling
4. **Build Validation**: All code must pass language-specific validators

## Build & Test

```bash
# Run all validations
bash build-validation/validate-all.sh

# Validate Go code
bash build-validation/validate-go.sh

# Run benchmarks
python benchmarks/runtime_stress_test.py
```

## Dependencies

- **Depends on**: PhoenixCore (contracts)
- **Depended by**: All repositories (architecture rules)

## Invariants

- TLA+ specifications must match Go implementations
- No circular dependencies between modules
- All builds must be reproducible
- Agent code must pass purity scanner

## License

PhoenixFormal is part of the PhoenixOS ecosystem.
