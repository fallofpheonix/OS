# Structured Documentation Architecture (SDA)

This directory houses the complete **Structured Documentation Architecture** for **PhoenixOS**. It divides the system documentation into 13 logical layers to ensure observability, auditability, security alignment, and agent workflow sanity.

## Layer Structure & Indexes

### [0. Governance](governance/)
Focus: Core strategy, project tracking, master vision, and health metrics.
* **[P0] [PROJECT_VISION.md](governance/PROJECT_VISION.md)** - System philosophy, 6 axioms, target users, and non-goals.
* **[P0] [ROADMAP.md](governance/ROADMAP.md)** - Phase I to Phase V evolution path.
* **[P0] [MASTER_STATUS.md](governance/MASTER_STATUS.md)** - Active features, completed stages, and immediate risk registry.
* **[P0] [DECISION_LOG.md](governance/DECISION_LOG.md)** - Tradeoff analyses, alternative architecture selection records.
* *[CHANGELOG.md](governance/CHANGELOG.md)*, *[PROJECT_HEALTH.md](governance/PROJECT_HEALTH.md)*.

### [1. Architecture](architecture/)
Focus: System interaction maps, data paths, control flows, and service topologies.
* **[P0] [SYSTEM_ARCHITECTURE.md](architecture/SYSTEM_ARCHITECTURE.md)** - The 7-layer stack, IPC paths, and boot orchestration.
* **[P0] [COMPONENT_MAP.md](architecture/COMPONENT_MAP.md)** - Subsystem ownership, module directories, and dependency chains.
* *[DATAFLOW_MAP.md](architecture/DATAFLOW_MAP.md)*, *[CONTROL_FLOW.md](architecture/CONTROL_FLOW.md)*, *[MODULE_BOUNDARIES.md](architecture/MODULE_BOUNDARIES.md)*, *[DEPENDENCY_GRAPH.md](architecture/DEPENDENCY_GRAPH.md)*, *[CAPABILITY_MATRIX.md](architecture/CAPABILITY_MATRIX.md)*, *[SERVICE_REGISTRY.md](architecture/SERVICE_REGISTRY.md)*.

### [2. Repository Integration](integration/)
Focus: Pinned dependency provenance, merge checks, and upstream syncing.
* **[P0] [EXTERNAL_REPOS.md](integration/EXTERNAL_REPOS.md)** - Registry of external repos, licenses, and modification plans.
* *[MERGE_POLICY.md](integration/MERGE_POLICY.md)*, *[FORK_STRATEGY.md](integration/FORK_STRATEGY.md)*, *[PATCH_HISTORY.md](integration/PATCH_HISTORY.md)*, *[UPSTREAM_SYNC.md](integration/UPSTREAM_SYNC.md)*, *[INTEGRATION_STATUS.md](integration/INTEGRATION_STATUS.md)*, *[API_COMPATIBILITY.md](integration/API_COMPATIBILITY.md)*, *[CONFLICT_LOG.md](integration/CONFLICT_LOG.md)*, *[DEPRECATION_TRACKER.md](integration/DEPRECATION_TRACKER.md)*.

### [3. AI Agent](agents/)
Focus: Multi-agent coordination protocols, LLM advisory bounds, and memory limits.
* **[P0] [AGENT_REGISTRY.md](agents/AGENT_REGISTRY.md)** - Roles, capabilities, and permissions of AI components (e.g. PhoenixMind).
* *[AGENT_INTERACTION_MAP.md](agents/AGENT_INTERACTION_MAP.md)*, *[AGENT_CHAINING.md](agents/AGENT_CHAINING.md)*, *[AGENT_MEMORY_POLICY.md](agents/AGENT_MEMORY_POLICY.md)*, *[AGENT_TRAINING_PIPELINE.md](agents/AGENT_TRAINING_PIPELINE.md)*, *[AGENT_FAILURE_CASES.md](agents/AGENT_FAILURE_CASES.md)*, *[AGENT_REWARD_SYSTEM.md](agents/AGENT_REWARD_SYSTEM.md)*, *[AGENT_VALIDATION.md](agents/AGENT_VALIDATION.md)*, *[AGENT_ALIGNMENT.md](agents/AGENT_ALIGNMENT.md)*, *[AGENT_EVOLUTION_LOG.md](agents/AGENT_EVOLUTION_LOG.md)*, *[MULTI_AGENT_PROTOCOL.md](agents/MULTI_AGENT_PROTOCOL.md)*, *[SELF_MODIFICATION_RULES.md](agents/SELF_MODIFICATION_RULES.md)*.

### [4. Security](security/)
Focus: Sandboxing, capability boundaries, Zero-Trust, cryptography, and playbooks.
* **[P0] [THREAT_MODEL.md](security/THREAT_MODEL.md)** - Attacker vectors, assets, entry points, and mitigation defenses.
* **[P0] [SECURITY_BOUNDARIES.md](security/SECURITY_BOUNDARIES.md)** - fast-path bypass limits, cgroups restrictions, and ledger validation.
* *[ATTACK_SURFACE.md](security/ATTACK_SURFACE.md)*, *[ZERO_TRUST_MODEL.md](security/ZERO_TRUST_MODEL.md)*, *[SANDBOX_POLICY.md](security/SANDBOX_POLICY.md)*, *[PERMISSION_MODEL.md](security/PERMISSION_MODEL.md)*, *[PRIVILEGE_ESCALATION.md](security/PRIVILEGE_ESCALATION.md)*, *[CRYPTO_DESIGN.md](security/CRYPTO_DESIGN.md)*, *[INCIDENT_RESPONSE.md](security/INCIDENT_RESPONSE.md)*, *[RED_TEAM_CASES.md](security/RED_TEAM_CASES.md)*, *[DEFENSE_PLAYBOOK.md](security/DEFENSE_PLAYBOOK.md)*, *[FORENSICS_GUIDE.md](security/FORENSICS_GUIDE.md)*, *[KILL_SWITCH.md](security/KILL_SWITCH.md)*, *[FAIL_SAFE_MODES.md](security/FAIL_SAFE_MODES.md)*, *[RECOVERY_PROTOCOL.md](security/RECOVERY_PROTOCOL.md)*.

### [5. Kernel & Core OS](kernel/)
Focus: Bootstrapping, processes, memory models, scheduling, and kernel limits.
* *[BOOT_SEQUENCE.md](kernel/BOOT_SEQUENCE.md)*, *[PROCESS_MODEL.md](kernel/PROCESS_MODEL.md)*, *[MEMORY_MODEL.md](kernel/MEMORY_MODEL.md)*, *[SCHEDULER_DESIGN.md](kernel/SCHEDULER_DESIGN.md)*, *[FILESYSTEM_DESIGN.md](kernel/FILESYSTEM_DESIGN.md)*, *[DEVICE_MODEL.md](kernel/DEVICE_MODEL.md)*, *[DRIVER_REGISTRY.md](kernel/DRIVER_REGISTRY.md)*, *[IPC_MODEL.md](kernel/IPC_MODEL.md)*, *[RESOURCE_MANAGER.md](kernel/RESOURCE_MANAGER.md)*, *[POWER_MANAGEMENT.md](kernel/POWER_MANAGEMENT.md)*, *[REALTIME_SUPPORT.md](kernel/REALTIME_SUPPORT.md)*, *[KERNEL_LIMITS.md](kernel/KERNEL_LIMITS.md)*, *[PERFORMANCE_TARGETS.md](kernel/PERFORMANCE_TARGETS.md)*.

### [6. Research](research/)
Focus: Control loop models, game-theoretic solvers, signal processing math, and ML notes.
* *[RESEARCH_INDEX.md](research/RESEARCH_INDEX.md)*, *[PAPERS_TRACKER.md](research/PAPERS_TRACKER.md)*, *[EXPERIMENT_LOG.md](research/EXPERIMENT_LOG.md)*, *[FAILED_IDEAS.md](research/FAILED_IDEAS.md)*, *[THEORY_NOTES.md](research/THEORY_NOTES.md)*, *[MATH_MODELS.md](research/MATH_MODELS.md)*, *[GAME_THEORY.md](research/GAME_THEORY.md)*, *[CONTROL_SYSTEMS.md](research/CONTROL_SYSTEMS.md)*, *[PHYSICS_REFERENCES.md](research/PHYSICS_REFERENCES.md)*, *[ML_NOTES.md](research/ML_NOTES.md)*, *[OPTIMIZATION_LOG.md](research/OPTIMIZATION_LOG.md)*.

### [7. Validation](validation/)
Focus: Test matrices, latency reporting, fault injections, and verification.
* **[P0] [TEST_PLAN.md](validation/TEST_PLAN.md)** - Replay validation test plan, race testing, and invariant gates.
* *[UNIT_TEST_MATRIX.md](validation/UNIT_TEST_MATRIX.md)*, *[INTEGRATION_TESTS.md](validation/INTEGRATION_TESTS.md)*, *[SYSTEM_TESTS.md](validation/SYSTEM_TESTS.md)*, *[CHAOS_TESTING.md](validation/CHAOS_TESTING.md)*, *[STRESS_TESTS.md](validation/STRESS_TESTS.md)*, *[FAULT_INJECTION.md](validation/FAULT_INJECTION.md)*, *[LATENCY_REPORT.md](validation/LATENCY_REPORT.md)*, *[BENCHMARKS.md](validation/BENCHMARKS.md)*, *[BUG_DATABASE.md](validation/BUG_DATABASE.md)*, *[KNOWN_LIMITATIONS.md](validation/KNOWN_LIMITATIONS.md)*, *[VERIFICATION_STATUS.md](validation/VERIFICATION_STATUS.md)*, *[MODEL_ACCURACY.md](validation/MODEL_ACCURACY.md)*.

### [8. Deployment](deployment/)
Focus: CI/CD integration, installer scripts, versioning policy, and update schemas.
* **[P0] [BUILD_PIPELINE.md](deployment/BUILD_PIPELINE.md)** - Automated builds, static linting, and image compilation.
* *[CI_CD.md](deployment/CI_CD.md)*, *[RELEASE_PROCESS.md](deployment/RELEASE_PROCESS.md)*, *[VERSION_POLICY.md](deployment/VERSION_POLICY.md)*, *[ROLLBACK.md](deployment/ROLLBACK.md)*, *[INSTALLATION.md](deployment/INSTALLATION.md)*, *[UPDATE_SYSTEM.md](deployment/UPDATE_SYSTEM.md)*, *[MIGRATION_GUIDE.md](deployment/MIGRATION_GUIDE.md)*, *[BACKUP_POLICY.md](deployment/BACKUP_POLICY.md)*, *[RESTORE_GUIDE.md](deployment/RESTORE_GUIDE.md)*.

### [9. Operations](operations/)
Focus: Event schemas, telemetry streams, threshold alerts, and anomaly logs.
* *[TELEMETRY.md](operations/TELEMETRY.md)*, *[LOGGING_SCHEMA.md](operations/LOGGING_SCHEMA.md)*, *[METRICS.md](operations/METRICS.md)*, *[EVENT_STREAMS.md](operations/EVENT_STREAMS.md)*, *[OBSERVABILITY.md](operations/OBSERVABILITY.md)*, *[ALERTING_RULES.md](operations/ALERTING_RULES.md)*, *[HEALTH_MONITORING.md](operations/HEALTH_MONITORING.md)*, *[ANOMALY_DETECTION.md](operations/ANOMALY_DETECTION.md)*.

### [10. GitHub / Agent Workflow](github/)
Focus: PR guidelines, code review checklists, merge automation rules, and doc versioning.
* *[ISSUE_TEMPLATES.md](github/ISSUE_TEMPLATES.md)*, *[PR_RULES.md](github/PR_RULES.md)*, *[CODE_REVIEW.md](github/CODE_REVIEW.md)*, *[MERGE_CHECKLIST.md](github/MERGE_CHECKLIST.md)*, *[BRANCH_POLICY.md](github/BRANCH_POLICY.md)*, *[LABELS.md](github/LABELS.md)*, *[AUTOMATION_RULES.md](github/AUTOMATION_RULES.md)*, *[BOT_BEHAVIOR.md](github/BOT_BEHAVIOR.md)*, *[REPO_HYGIENE.md](github/REPO_HYGIENE.md)*, *[DOCUMENTATION_POLICY.md](github/DOCUMENTATION_POLICY.md)*.

### [11. Runtime Intelligence](runtime_ai/)
Focus: Local learning policies, knowledge graph bindings, drift warnings, and model lifecycle.
* *[MEMORY_SYSTEM.md](runtime_ai/MEMORY_SYSTEM.md)*, *[KNOWLEDGE_GRAPH.md](runtime_ai/KNOWLEDGE_GRAPH.md)*, *[PERSONALIZATION.md](runtime_ai/PERSONALIZATION.md)*, *[LEARNING_POLICY.md](runtime_ai/LEARNING_POLICY.md)*, *[MODEL_SELECTION.md](runtime_ai/MODEL_SELECTION.md)*, *[LOCAL_MODEL_REGISTRY.md](runtime_ai/LOCAL_MODEL_REGISTRY.md)*, *[MODEL_LIFECYCLE.md](runtime_ai/MODEL_LIFECYCLE.md)*, *[MODEL_PRUNING.md](runtime_ai/MODEL_PRUNING.md)*, *[SELF_TRAINING.md](runtime_ai/SELF_TRAINING.md)*, *[ADAPTATION_RULES.md](runtime_ai/ADAPTATION_RULES.md)*, *[DRIFT_DETECTION.md](runtime_ai/DRIFT_DETECTION.md)*, *[MODEL_ROLLBACK.md](runtime_ai/MODEL_ROLLBACK.md)*.

### [12. Emergency Operations](emergency/)
Focus: Safe mode booting, manual override budget resets, and recovery triggers.
* *[DISASTER_RECOVERY.md](emergency/DISASTER_RECOVERY.md)*, *[CORRUPTION_RECOVERY.md](emergency/CORRUPTION_RECOVERY.md)*, *[BOOT_FAILURES.md](emergency/BOOT_FAILURES.md)*, *[AI_FAILURE_PROTOCOL.md](emergency/AI_FAILURE_PROTOCOL.md)*, *[SAFE_MODE.md](emergency/SAFE_MODE.md)*, *[MANUAL_OVERRIDE.md](emergency/MANUAL_OVERRIDE.md)*, *[EMERGENCY_OPERATIONS.md](emergency/EMERGENCY_OPERATIONS.md)*.

---

## Mandates for AI Agents Maintaining this OS

1. **No Code Edits Without Documentation:** No agent may edit functional source code without verifying and updating the corresponding architectural, dependency, threat, or test documents.
2. **Every PR Updates the Core 4:** Every PR must update or assert impact in:
   - *Architecture Impact* (`/docs/architecture/SYSTEM_ARCHITECTURE.md`)
   - *Dependency Impact* (`/docs/architecture/COMPONENT_MAP.md` / `EXTERNAL_REPOS.md`)
   - *Threat Impact* (`/docs/security/THREAT_MODEL.md`)
   - *Test Impact* (`/docs/validation/TEST_PLAN.md`)
3. **No Deletion of Failed Experiments:** All failed experiment logs or mathematical theories must be archived in `experiments/` or `archives/` rather than deleted.
4. **Self-Modifying System Rules:** Any automated runtime modifications to policies or FSM bounds require verification through an immutable audit trail written to the cryptographic Evidence Ledger.
