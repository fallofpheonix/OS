# PhoenixOS Documentation Index

Welcome to the PhoenixOS canonical documentation. The system is divided into functional governance tiers.

## 00 Identity & Strategy
- [System Identity](identity/system_identity.md) - Vision, Axioms, and Core Definition.
- [Execution Roadmap](identity/roadmap.md) - Strategic stages (A-D) and Phased milestones.
- [Phase Lock](identity/phase_lock.md) - Current active blockers and status board.

## 01 Architecture & Models
- [State Model](architecture/state_model.md) - Warden FSM and Isolation transitions.
- [Replay Specification](architecture/replay_spec.md) - Determinism and Replay semantics.
- [Truth Model](architecture/truth_model.md) - Immutable Ledger and Evidence chains.
- [Decision Model](architecture/decision_model.md) - Strategy and Game Theory (Stackelberg).

## 02 Validation & Reports
- [Validation Rules](validation/validation_rules.md) - Testing standards and CI/CD gates.
- [Runtime Reality Audit](validation/runtime_reality_audit.md) - Empirical readiness of all modules.
- [Kernel Determinism](validation/kernel_runtime_report.md) - Real-world kernel validation results.
- [Replay Identity](validation/replay_identity_report.md) - 1000-run consistency proof.
- [Truth Immutability](validation/truth_immutability_report.md) - Evidence non-repudiation proof.

## 03 Security & Trust
- [Threat Model](security/threat_model.md) - Attack surface and vector analysis.
- [Boundaries](security/boundaries.md) - Strict isolation zones and illegal dependency paths.
- [RedTeam Report](validation/redteam_runtime_report.md) - Adversarial runtime validation.

## 04 External Governance
- [External Policy](external_governance/external_policy.md) - Intake standards for third-party code.
- [Repository Matrix](external_governance/external_repo_matrix.md) - Classification of all integrated repos.
- [Archive Import](external_governance/archive_import.md) - Forensic rules for restoring archived code.

## 05 Research & Theory
- [Research Policy](research_governance/research_policy.md) - Experimental integration standards.
- [Quantum Policy](research_governance/quantum_policy.md) - Specific rules for quantum OS modules.

## 06 Operations
- [Risk Matrix](operations/risk_matrix.md) - Strategic risk register and mitigation.

---
*Note: All historical and duplicate documents are stored in `archive/documentation/version_history/`.*
