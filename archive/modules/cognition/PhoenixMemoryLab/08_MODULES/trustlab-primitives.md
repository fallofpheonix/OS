# Module: trustlab-primitives

## Source
Extracted from [[05_PROJECTS/COMPLETED/SecureForg/Project]] + TrustLab repo

## Purpose
Runtime vulnerability detection primitives: sandboxed code execution, payload-based security testing, and behavioral comparison between baseline and attack runs.

## Interface
```python
from trustlab_primitives import Sandbox, PayloadRunner, BehaviorDiff

sandbox = Sandbox(timeout=5, memory_limit="256MB")
baseline = sandbox.run(target_code, benign_inputs)

runner = PayloadRunner(vulnerability_types=["sql_injection", "cmd_injection", "code_injection"])
attack_results = runner.run(target_code)

diff = BehaviorDiff(baseline, attack_results)
report = diff.analyze()  # → VulnerabilityReport
```

## Depends On
None — self-contained security primitives.

## Used By
- Network Security Scanner (core scanning engine)
- Banking App (security audit pipeline)
- All projects via CI security gates

## Extraction Status
NOT_STARTED

## Location
`~/engineering/infrastructure/shared-libraries/trustlab-primitives/`

## Key Files
| File | Role |
|------|------|
| `sandbox.py` | Isolated code execution with resource limits |
| `payloads/` | Fixed attack payload library (SQL, cmd, code injection) |
| `runner.py` | Executes baseline + attack sequences |
| `diff.py` | Behavioral comparison engine |
| `report.py` | Structured vulnerability report generator |
| `ast_analyzer.py` | Static AST analysis from SecureForg |

## SecureForg Architecture Being Extracted
```
Target Code → Sandbox(benign inputs) → Baseline Behavior
                                            ↓
Target Code → Sandbox(attack payloads) → Attack Behavior
                                            ↓
                                    BehaviorDiff → Report
```

## Quality Gates
- [ ] Tests passing
- [ ] Sandbox is truly isolated (no file system leaks)
- [ ] Payload library is extensible
- [ ] README with security testing tutorial
- [ ] Version pinned

> [!WARNING]
> **TrustLab repo has no README.** Must audit actual repo contents before merging with SecureForg primitives. TrustLab may be empty scaffolding or may contain trust-scoring logic that complements SecureForg's runtime detection.

#module #extracted-from/SecureForg #extracted-from/TrustLab #priority/P1
