Surface Agent - Orchestrator

Purpose
- Coordinator-only agent (Layer 2 Surface)

Responsibilities
- Task routing
- Dependency graph (simple DAG)
- State tracking
- Risk control
- Workflow execution
- Validation gates

State machine
IDEA -> RFC -> RESEARCH -> EXPERIMENT -> BUILD -> TEST -> DEBUG -> VALIDATE -> MERGE

Constraints
- The surface agent is orchestration-only. It must NOT collect telemetry, run kernel hooks, compute graphs, or do forensics.

Quick example

```py
from agents.surface.orchestrator import SurfaceOrchestrator, Task, State

orch = SurfaceOrchestrator()
wf = orch.create_workflow('demo')
orch.add_task('demo', Task(id='t1', title='Write RFC', task_type='documentation'))

# register a simple handler that just logs routing
orch.register_handler('documentation', lambda t: print('handled', t.id, t.state))

orch.execute_workflow('demo')
```

GitHub Integration

Set `GITHUB_TOKEN` in your environment to enable creating issues from workflow tasks. Use `github_integration.create_issues_for_workflow(service, wf_id)`.
