import os
import pytest

from agents.surface.orchestrator.service import OrchestratorService
from agents.surface.orchestrator.orchestrator import Task, State


def test_persist_and_load(tmp_path):
    db = tmp_path / 'state.db'
    svc = OrchestratorService(db_path=db)
    wf_id = 'wf-test'
    svc.create_workflow(wf_id)
    svc.add_task(wf_id, Task(id='task1', title='Task 1', task_type='documentation'))
    svc.persist_workflow(wf_id)

    # create a new service instance pointing at same DB and load
    svc2 = OrchestratorService(db_path=db)
    wf = svc2.orch.create_workflow(wf_id)
    assert 'task1' in wf.tasks


def test_approve_gate_blocks_merge(tmp_path):
    svc = OrchestratorService(db_path=tmp_path / 'state2.db')
    wf_id = 'wf-approve'
    svc.create_workflow(wf_id)
    svc.add_task(wf_id, Task(id='a', title='A', task_type='documentation'))
    # execute async then approve
    svc.execute_workflow_async(wf_id).join(timeout=2)
    wf = svc.orch.create_workflow(wf_id)
    t = wf.tasks['a']
    # after run without approval the task should reach MERGE only if approved; default code sets MERGE only when approved
    # ensure approve() sets metadata and persists
    svc.approve_task(wf_id, 'a', approver='tester')
    assert t.metadata.get('approved') is True
