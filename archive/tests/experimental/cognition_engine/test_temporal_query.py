import pytest
from datetime import UTC, datetime
from events import EventBus, EventAction
from planner.schemas import PlannedTask, PlannedTaskType
from runtime.replay_engine import ReplayEngine
from runtime.temporal_query import TemporalQueryLayer
from transactions.journal import FilesystemJournal, JournalEntry

@pytest.fixture
def setup_query(tmp_path):
    run_id = "test_run_query"
    run_dir = tmp_path / "artifacts" / run_id
    run_dir.mkdir(parents=True)
    event_bus = EventBus(run_dir / "events.jsonl")
    replay_engine = ReplayEngine(tmp_path)
    
    journal_path = tmp_path / "journal.jsonl"
    journal = FilesystemJournal(journal_path)
    
    query_layer = TemporalQueryLayer(replay_engine, journal=journal)
    return tmp_path, run_id, run_dir, event_bus, query_layer, journal

def test_mutation_history(setup_query):
    tmp_path, run_id, run_dir, event_bus, query_layer, journal = setup_query
    
    # Record some mutations
    entry1 = JournalEntry(run_id=run_id, task_id="task1", operation="CREATE", path="main.py", new_hash="h1")
    journal.record(entry1)
    
    entry2 = JournalEntry(run_id=run_id, task_id="task2", operation="MODIFY", path="main.py", old_hash="h1", new_hash="h2")
    journal.record(entry2)
    
    history = query_layer.get_mutation_history("main.py", run_id=run_id)
    assert len(history) == 2
    assert history[0].task_id == "task1"
    assert history[1].task_id == "task2"

def test_task_timeline(setup_query):
    tmp_path, run_id, run_dir, event_bus, query_layer, journal = setup_query
    
    event_bus.emit(run_id=run_id, action=EventAction.RUN_STARTED)
    task1 = PlannedTask(id="task1", type=PlannedTaskType.CODE_GENERATION, goal="g1", validation=[])
    event_bus.emit(run_id=run_id, task_id="task1", action=EventAction.TASK_CREATED, metadata=task1.model_dump(mode="json"))
    
    # Record a mutation for task1
    entry = JournalEntry(run_id=run_id, task_id="task1", operation="CREATE", path="main.py", new_hash="h1")
    journal.record(entry)
    
    event_bus.emit(run_id=run_id, task_id="task1", action=EventAction.TASK_SUCCEEDED, metadata={"output": "out1"})
    
    timeline = query_layer.get_task_timeline(run_id, "task1")
    
    # 2 events (CREATED, SUCCEEDED) + 1 mutation
    assert len(timeline) == 3
    assert timeline[0]["type"] == "event"
    assert timeline[0]["action"] == "task_created"
    assert timeline[1]["type"] == "mutation"
    assert timeline[1]["operation"] == "CREATE"
    assert timeline[2]["type"] == "event"
    assert timeline[2]["action"] == "task_succeeded"

def test_branch_history(setup_query):
    tmp_path, run_id, run_dir, event_bus, query_layer, journal = setup_query
    
    event_bus.emit(run_id=run_id, action=EventAction.RUN_STARTED)
    
    # task1 output changes, triggering rerun of task2
    event_bus.emit(run_id=run_id, task_id="task1", action=EventAction.SUBTREE_RERUN, metadata={"reason": "output_hash_changed"})
    event_bus.emit(run_id=run_id, task_id="task2", action=EventAction.ARTIFACT_INVALIDATED, metadata={"failed_task_id": "task1"})
    
    branches = query_layer.get_branch_history(run_id)
    assert len(branches) == 1
    assert branches[0].origin_task_id == "task1"
    assert branches[0].affected_tasks == ["task2"]

def test_diff_states(setup_query):
    tmp_path, run_id, run_dir, event_bus, query_layer, journal = setup_query
    
    # T1: Task 1 created
    event_bus.emit(run_id=run_id, action=EventAction.RUN_STARTED)
    task1 = PlannedTask(id="task1", type=PlannedTaskType.CODE_GENERATION, goal="g1", validation=[])
    event_bus.emit(run_id=run_id, task_id="task1", action=EventAction.TASK_CREATED, metadata=task1.model_dump(mode="json"))
    t1 = datetime.now(UTC).isoformat()
    
    # T2: Task 1 succeeded, Task 2 created
    event_bus.emit(run_id=run_id, task_id="task1", action=EventAction.TASK_SUCCEEDED, metadata={"output": "out1"})
    task2 = PlannedTask(id="task2", type=PlannedTaskType.CODE_GENERATION, goal="g2", depends_on=["task1"], validation=[])
    event_bus.emit(run_id=run_id, task_id="task2", action=EventAction.TASK_CREATED, metadata=task2.model_dump(mode="json"))
    t2 = datetime.now(UTC).isoformat()
    
    diff = query_layer.diff_states(run_id, t1, t2)
    
    assert "task2" in diff.new_tasks
    assert "task1" in diff.status_changes
    assert diff.status_changes["task1"] == ("pending", "succeeded")
    assert "task1" in diff.changed_tasks # output changed from None to "out1"

def test_trace_lineage(setup_query):
    tmp_path, run_id, run_dir, event_bus, query_layer, journal = setup_query
    
    event_bus.emit(run_id=run_id, action=EventAction.RUN_STARTED)
    
    # task1 -> task2 -> task3
    task1 = PlannedTask(id="task1", type=PlannedTaskType.CODE_GENERATION, goal="g1", validation=[])
    event_bus.emit(run_id=run_id, task_id="task1", action=EventAction.TASK_CREATED, metadata=task1.model_dump(mode="json"))
    
    task2 = PlannedTask(id="task2", type=PlannedTaskType.CODE_GENERATION, goal="g2", depends_on=["task1"], validation=[])
    event_bus.emit(run_id=run_id, task_id="task2", action=EventAction.TASK_CREATED, metadata=task2.model_dump(mode="json"))
    
    task3 = PlannedTask(id="task3", type=PlannedTaskType.CODE_GENERATION, goal="g3", depends_on=["task2"], validation=[])
    event_bus.emit(run_id=run_id, task_id="task3", action=EventAction.TASK_CREATED, metadata=task3.model_dump(mode="json"))
    
    lineage = query_layer.trace_lineage(run_id, "task3")
    assert sorted(lineage) == ["task1", "task2"]
    
    lineage_t1 = query_layer.trace_lineage(run_id, "task1")
    assert lineage_t1 == []
