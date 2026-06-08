# Module: fpx-pipeline

## Source
Extracted from [[05_PROJECTS/COMPLETED/AutoMation-Engine/Project]]

## Purpose
Staged pipeline control-plane pattern: decompose complex workflows into ordered stages that communicate through shared typed contracts.

## Interface
```python
from fpx_pipeline import Pipeline, Stage, Orchestrator

pipeline = Pipeline(name="transaction-processing")
pipeline.add_stage(Stage("validate", validator_fn))
pipeline.add_stage(Stage("execute", executor_fn))
pipeline.add_stage(Stage("log", logger_fn))

orchestrator = Orchestrator(pipeline)
result = orchestrator.run(input_data)
```

## Depends On
- [[05_PROJECTS/REUSABLE_MODULES/fpx-runtime]] (executor layer for individual stages)

## Used By
- Banking App (transaction pipeline: validate → execute → settle → log)
- Personal Knowledge Graph (ingestion pipeline: scrape → normalize → tag → store)
- Any multi-step workflow

## Extraction Status
NOT_STARTED

## Location
`~/engineering/infrastructure/shared-libraries/fpx-pipeline/`

## Key Files
| File | Role |
|------|------|
| `pipeline.py` | Pipeline definition with ordered stage registry |
| `stage.py` | Individual stage abstraction with input/output contracts |
| `orchestrator.py` | Coordinates stage execution, handles inter-stage data flow |
| `contracts.py` | Typed data contracts between stages |

## Proven Pattern
Used in AutoMation-Engine (task → orchestrator → executor → logger) and ParticleStimulator (BeamSource → BeamDynamics → CollisionEngine → DetectorSimulator → EventReconstructor → PhysicsAnalyser).

## Quality Gates
- [ ] Tests passing
- [ ] Generic stage interface (no task-automation coupling)
- [ ] README with pipeline construction examples
- [ ] Version pinned
- [ ] Supports async stages

#module #extracted-from/AutoMation-Engine #priority/P0
