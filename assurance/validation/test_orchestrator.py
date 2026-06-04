"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
import pytest

from agents.surface.orchestrator.orchestrator import SurfaceOrchestrator, Task, State, NotAllowedError


def test_add_and_execute_simple_workflow(tmp_path):
    orch = SurfaceOrchestrator()
    wf = orch.create_workflow('w1')
    t = Task(id='t1', title='T1', task_type='documentation')
    orch.add_task(wf.id, t)

    states = []

    def on_change(wf_id, task_id, new_state):
        states.append(new_state)

    orch.on_state_change = on_change
    orch.execute_workflow(wf.id)
    assert t.state == State.MERGE
    assert State.VALIDATE in states


def test_forbidden_task_type():
    orch = SurfaceOrchestrator()
    wf = orch.create_workflow('w2')
    t = Task(id='t2', title='T2', task_type='forensics')
    with pytest.raises(NotAllowedError):
        orch.add_task(wf.id, t)
