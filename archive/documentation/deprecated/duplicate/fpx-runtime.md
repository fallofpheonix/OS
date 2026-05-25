# Module: fpx-runtime

## Source
Extracted from [[05_PROJECTS/COMPLETED/AutoMation-Engine/Project]]

## Purpose
Cross-platform executor abstraction with intelligent retry logic, exponential backoff, error classification (temporary vs permanent), and timeout handling.

## Interface
```python
from fpx_runtime import Executor, RetryEngine, ErrorClassifier

executor = Executor(platform="macos")  # or "windows", "simulation"
result = executor.run(action, params, retry_config)

classifier = ErrorClassifier()
error_type = classifier.classify(exception)  # → TEMPORARY | PERMANENT

retry = RetryEngine(max_retries=3, backoff="exponential")
retry.execute(callable, on_failure=callback)
```

## Depends On
None — this is a foundational module.

## Used By
- Banking App (transaction executor)
- Network Security Scanner (scan executor)
- All future services requiring reliable task execution

## Extraction Status
NOT_STARTED

## Location
`~/engineering/infrastructure/shared-libraries/fpx-runtime/`

## Key Files
| File | Role |
|------|------|
| `executor.py` | Platform-abstracted action executor (Mac/Windows/Simulation) |
| `retry.py` | Exponential backoff engine with configurable max retries and timeout |
| `classifier.py` | Error classification: temporary (retry-worthy) vs permanent (fail-fast) |
| `timeout.py` | Task timeout handler to prevent infinite execution |
| `metrics.py` | Success rate, avg execution time, P99 latency tracking |

## Proven Performance
- 96% success rate with exponential backoff (vs 40% with immediate retry)
- Average action: 450-500ms
- P99: 2-3 seconds
- 18+ tests passing in source project

## Quality Gates
- [ ] Tests passing
- [ ] No AutoMation-Engine-specific coupling removed
- [ ] README with usage examples
- [ ] Version pinned
- [ ] Platform executor interface documented

## Architecture Pattern
```
Caller → RetryEngine → ErrorClassifier → Executor → Result
              ↓                              ↓
         Backoff/Fail               Platform Adapter
```

#module #extracted-from/AutoMation-Engine #priority/P0
