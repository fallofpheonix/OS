# Module: agentskill

## Source
Extracted from [[05_PROJECTS/COMPLETED/agentskill/Project]]

## Purpose
Deterministic CI/CD security gates: Agentfile parsing/validation, pip-audit dependency scanning, regex-based secret detection, and enforced stage schemas for build pipelines.

## Interface
```python
from agentskill import AgentfileParser, SecurityGate, StageValidator

# Parse and validate Agentfile
parser = AgentfileParser()
config = parser.parse("Agentfile.yaml")
parser.validate(config)  # Raises on schema violation

# Run security gates
gate = SecurityGate()
gate.pip_audit()           # Dependency vulnerability scan
gate.secret_scan("./src")  # Regex-based secret detection

# Validate CI/CD stages
validator = StageValidator(schema="tri-engine")
validator.check(stages_config)  # Ensures stage ordering + contracts
```

## Depends On
- pip-audit (Python dependency)
- GitHub Actions (for CI integration)

## Used By
- All projects (CI/CD security gates)
- Banking App (enforced build/deploy pipeline)
- Network Security Scanner (self-audit)

## Extraction Status
NOT_STARTED

## Location
`~/engineering/infrastructure/shared-libraries/agentskill/`

## Key Files
| File | Role |
|------|------|
| `parser.py` | Deterministic Agentfile parser |
| `validator.py` | Schema validation for stage contracts |
| `security/pip_audit.py` | Dependency vulnerability scanning wrapper |
| `security/secret_scan.py` | Regex-based secret detection in source files |
| `tri_engine/` | Tri-engine contract system with enforced stage schema |
| `ci_cd/` | GitHub Actions entrypoints for local validation |

## Execution Guarantees (from source project)
- **Build:** Deterministic bundle generation with `control-plane.json` + `stages.yaml`
- **Test:** ≥70% coverage for core package
- **Security:** pip-audit + regex-based secret scanning
- **CI/CD:** GitHub Actions on pull_request and push to main

## Quality Gates
- [ ] Tests passing (≥70% coverage)
- [ ] Works as standalone CLI tool
- [ ] GitHub Actions workflow template included
- [ ] README with integration guide for new projects
- [ ] Version pinned

#module #extracted-from/agentskill #priority/P2
