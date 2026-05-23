# INVARIANT MATRIX - Astraeus Ecosystem

This matrix tracks the translation of governance documentation into executable architectural invariants.

| ID | Category | Rule Title | Source Doc | Status | Enforcement Layer |
|---|---|---|---|---|---|
| INV-STR-001 | Structural | No Validator/Orchestrator Cycle | invariants.yaml | ✅ Active | InvariantEngine |
| INV-STR-002 | Structural | Contracts are Pure | invariants.yaml | ✅ Active | InvariantEngine |
| INV-STR-003 | Structural | No Global Cycles | invariants.yaml | ✅ Active | InvariantEngine |
| INV-STR-004 | Structural | Module/App Dependency Rule | DEPENDENCY_POLICY.md | ✅ Active | EcosystemValidator |
| INV-STR-005 | Structural | Infra/Experiment Dependency Rule | DEPENDENCY_POLICY.md | ✅ Active | EcosystemValidator |
| INV-BEH-001 | Behavioral | All Model Adapters Implement Protocol | invariants.yaml | ✅ Active | InvariantEngine |
| INV-BEH-002 | Behavioral | Async Handlers Must Timeout | invariants.yaml | ⏳ Partial | InvariantEngine |
| INV-BEH-003 | Behavioral | Service Contract Compliance | ARCHITECTURE_RULES.md | ⏳ Planned | ContractValidator |
| INV-OPE-001 | Operational | Mutations Require Snapshot | invariants.yaml | 🚧 Stub | runtime/sandbox |
| INV-OPE-002 | Operational | Dangerous Commands Approval | invariants.yaml | 🚧 Stub | runtime/risk_engine |
| INV-OPE-003 | Operational | Repair Budget Enforced | invariants.yaml | ✅ Active | repair/planner |
| INV-IDN-001 | Identity | Identity Docs Protection | CONSTITUTION.md | ✅ Active | runtime/sandbox |
| INV-MEM-001 | Memory | Semantic/Episodic Separation | INVARIANTS.md | ⏳ Planned | MemoryStore |

## Legend
- ✅ **Active:** Fully implemented and enforced.
- ⏳ **Planned:** Defined in docs, implementation pending.
- 🚧 **Stub:** Placeholder exists in code, but logic is incomplete.
- ❌ **Missing:** Documented but no trace in code.
