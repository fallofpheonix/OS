---\nStatus: Planned\nImplementation: 5%\nConfidence: Conceptual\n---\n# PhoenixFormal — Formal Verification Layer

## Agent Skills
### Issue Tracker
GitHub issue tracker. See `docs/agents/issue-tracker.md`.

### Triage Labels
Default triage label vocabulary. See `docs/agents/triage-labels.md`.

### Domain Docs
Multi-context layout. See `docs/agents/domain.md`.

## Build & Test
```bash
# Run all validations
bash build-validation/validate-all.sh

# Validate Go code
bash build-validation/validate-go.sh

# Run TLA+ model checker
java -cp tla2tools.jar tlc2.TLC tla/GuardFSM.tla
```

## Architecture
PhoenixFormal provides TLA+ specifications, architecture rules, and dependency policies. All critical system properties must have corresponding TLA+ specifications.

## Key Components
- **tla/** — TLA+ specifications (GuardFSM, ConsensusSafety, LedgerSafety)
- **ARCHITECTURE_RULES.md** — Code organization rules
- **DEPENDENCY_POLICY.md** — Module dependency rules
- **VISIBILITY_POLICY.md** — Visibility rules
- **build-validation/** — Language-specific validators
- **agent-governance/** — AI agent governance

## Invariants
- TLA+ specifications must match Go implementations
- No circular dependencies between modules
- All builds must be reproducible
- Agent code must pass purity scanner
