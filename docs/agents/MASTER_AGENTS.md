---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# PhoenixOS Agent Master Guidelines (v1.0)

## 1. Substrate Domains
- **Nucleus**: Foundation of Execution (contracts, FSM, kernel).
- **Cognition**: AI Intelligence Layer (memory, knowledge).
- **Arbiter**: Strategic Governance (oversight, compliance).
- **Terminus**: Interface Layer (CLI, Oracle).

## 2. Agent Guidelines
- **Logic Preservation**: Determinism and Replay must NEVER be compromised.
- **Workflow Position**: Every function must have standardized metadata labels.
- **Verification**: Run `make test` before declaring a task complete.

## 3. Operations
### Issue Tracker
GitHub-based tracking. See local issue records for context.

### Triage Labels
Standardized labels: `bug`, `feature`, `invariant-violation`, `refactor`.

## 4. Maintenance
### Build Commands
```bash
make ignite  # Sovereign node launch
make test    # Proof verification
```
