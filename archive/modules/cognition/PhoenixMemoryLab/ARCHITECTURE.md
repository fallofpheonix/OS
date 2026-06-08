# Master Engineering Ecosystem Structure

This document is the canonical reference for the long-term engineering ecosystem, institutionalizing the extraction strategy and architectural standards.

## 1. Master Filesystem Structure

```txt
~/engineering/
│
├── brain/                     # Cognition only (Knowledge, ADRs, Failure logs, etc.)
│
├── workspace/                 # Active executable repos
│   ├── active/                # Current high-priority systems
│   ├── incubating/            # Early prototypes
│   ├── maintenance/           # Stable projects with low development
│   └── migration/             # Projects being modularized/extracted
│
├── infrastructure/            # Shared infra/templates (Docker, K8s, Terraform, CI/CD)
│
├── environments/              # Runtime isolation (Rust, Python, Node, Go, Containers)
│
├── modules/                   # Extracted reusable systems
│   ├── core/                  # Logging, Config, Auth, Observability
│   ├── shared/
│   ├── infra/
│   ├── ai/
│   ├── systems/
│   ├── networking/
│   ├── automation/
│   ├── security/
│   ├── data/
│   └── experimental/
│
├── services/                  # Deployable independent services
│
├── sdk/                       # Language bindings (Python, TS, Rust, Go)
│
├── forks/                     # Temporary extraction repos
│   ├── active/
│   ├── extracting/
│   ├── mined/
│   └── archived/
│
├── research/                  # Experiments + POCs
│
├── archives/                  # Deprecated systems
│
└── control-plane/             # Meta-control layer
    ├── repo-registry/         # Single source of truth for all repos
    ├── sync-engine/
    ├── dependency-map/
    ├── health-monitor/
    └── extraction-tracker/
```

## 2. Core Architectural Principles

### A. Module Extraction Rules
- **Rule of 2**: Only extract after a pattern repeats in 2+ projects.
- **Independence**: Every module must have its own `README.md`, `API.md`, `VERSION`, `tests/`, and `src/`.
- **Interoperability**: Use stable interfaces (REST, gRPC) and contracts (OpenAPI, Protobuf).

### B. Fork Lifecycle Strategy
Forks are **temporary extraction environments**, not permanent foundations.
1. **Active**: Understand architecture and identify candidates.
2. **Extraction**: Move logic into `modules/`, `services/`, or `sdk/`.
3. **Mined**: Reference for architecture notes and lessons.
4. **Archive/Delete**: Once fully extracted, move to `archives/forks/`.

### C. Language Strategy (Best-Fit)
- **Rust/Go**: Core Infra, Networking, System Tools, High-throughput services.
- **Python**: AI/ML execution, automation scripting.
- **TypeScript**: Frontend/Web ecosystems, rapid iteration.
- **Bash**: Glue code and simple automation.

## 3. Implementation Roadmap
- **Phase 1-2**: Inventory and Classify everything.
- **Phase 3**: Extract Tier 1 Foundations (Logging, Config, Auth, Observability).
- **Phase 4-5**: Create control-plane layer and standardize CI/CD.
- **Phase 6**: Convert major projects into ecosystem consumers.

---
*"The repo is temporary. The architecture is permanent."*
