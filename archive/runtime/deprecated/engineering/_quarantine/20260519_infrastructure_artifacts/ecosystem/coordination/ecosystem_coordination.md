# Ecosystem Coordination Layer

## Coordination Flows

### Cleanup Flow
- **Trigger:** Status change to `ARCHIVE_CANDIDATE` in `repo_manifest.yaml`.
- **Action:** Verify GitHub sync, strip local runtime, move to `archive/`.
- **Constraint:** Must not break dependent active repos.

### Runtime Flow
- **Trigger:** New runtime requirement or upgrade.
- **Action:** Impact analysis via `ecosystem_risk_graph.json`.
- **Validation:** Build validation in `control-plane/`.

### Research Flow
- **Trigger:** Domain expansion or new research repo.
- **Action:** Update `research_domain_map.yaml` and `ecosystem_topology.json`.
- **Integration:** Map to `07_RESEARCH` in Brain.

### Archive Flow
- **Trigger:** Repo inactivity > 90 days.
- **Action:** Transition to `ARCHIVED` status.
- **Verification:** Metadata preservation in `archive_inventory.md`.

### Manifest Flow
- **Trigger:** Any repo mutation.
- **Action:** Single point of truth update in `repo_manifest.yaml`.
- **Sync:** Re-generate `ecosystem_topology.json` and `ecosystem_metrics.yaml`.

### Governance Flow
- **Trigger:** Policy violation.
- **Action:** Log in `governance/` and update `risk_level` in manifest.
- **Escalation:** Critical risks require immediate `control-plane` intervention.

## State Transitions
- **EXPERIMENTAL** → **ACTIVE** (Requires PRD + TDD)
- **ACTIVE** → **LEGACY** (Feature freeze)
- **LEGACY** → **ARCHIVED** (Removal from active workspace)
