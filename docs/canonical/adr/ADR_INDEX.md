---\nStatus: Draft\nImplementation: 0%\nConfidence: High\n---\n# ADR Index & Decision Framework\n\n## 1. ADR Lifecycle\n\n- **Proposed**: Rationale documented, awaiting review.\n- **Accepted**: Approved by architecture lead/governance.\n- **Superseded**: Replaced by a newer ADR (must link to successor).\n- **Deprecated**: No longer relevant, but kept for historical context.\n\n## 2. ADR Inventory\n\n| ID | Title | Status | Date | Tags |\n| :--- | :--- | :--- | :--- | :--- |\n| **[ADR-001](./ADR-001-Fractal-Architecture.md)** | Contract-First Architecture Model | Accepted | 2026-06-04 | Core, Boundary |\n| **[ADR-002](./ADR-002-ADR-Framework.md)** | ADR Framework & Decision Process | Proposed | 2026-06-04 | Governance |\n\n## 3. Decision Process\n\n1. **Problem Identification**: Document the context and failure modes.\n2. **Alternative Evaluation**: Compare at least two alternatives with Pros/Cons.\n3. **Selection**: Choose based on alignment with `MASTER_INVARIANTS.md`.\n4. **Validation**: Define the automated test or proof required to "close" the ADR.\n\n## 4. Canonical Specifications
- **[SPEC-001: System State](../specifications/SPEC-001-System-State.md)**
- **[SPEC-002: Runtime State](../specifications/SPEC-002-Runtime-State.md)**
- **[SPEC-003: Ledger State](../specifications/SPEC-003-Ledger-State.md)**
- **[SPEC-004: Memory State](../specifications/SPEC-004-Memory-State.md)**
- **[SPEC-005: Governance State](../specifications/SPEC-005-Governance-State.md)**
- **[SPEC-006: Intent State](../specifications/SPEC-006-Intent-State.md)**
- **[SPEC-007: Capability Model](../specifications/SPEC-007-Capability-Model.md)**

--- \n*Refer to [docs/governance/REPOSITORY_CONSTITUTION.md](../governance/REPOSITORY_CONSTITUTION.md) for engineering standards.* \n