import pytest
import json
from datetime import UTC, datetime
from pathlib import Path
from events import EventBus, EventAction, Event, EventResult
from shared_context import RunState
from shared_context.state import TaskStatus
from runtime.replay_engine import ReplayEngine
from planner.schemas import PlannedTask, PlannedTaskType

@pytest.fixture
def setup_replay(tmp_path):
    run_id = "test_run_123"
    run_dir = tmp_path / "artifacts" / run_id
    run_dir.mkdir(parents=True)
    event_bus = EventBus(run_dir / "events.jsonl")
    replay_engine = ReplayEngine(tmp_path)
    return tmp_path, run_id, run_dir, event_bus, replay_engine

def test_reconstruct_basic_run(setup_replay):
    tmp_path, run_id, run_dir, event_bus, replay_engine = setup_replay
    
    # 1. Emit events
    event_bus.emit(run_id=run_id, action=EventAction.RUN_STARTED, metadata={"goal": "test goal"})
    
    task1 = PlannedTask(id="task1", type=PlannedTaskType.CODE_GENERATION, goal="goal1", validation=[])
    event_bus.emit(run_id=run_id, task_id="task1", action=EventAction.TASK_CREATED, metadata=task1.model_dump(mode="json"))
    
    event_bus.emit(run_id=run_id, task_id="task1", action=EventAction.TASK_STARTED)
    event_bus.emit(run_id=run_id, task_id="task1", action=EventAction.TASK_SUCCEEDED, metadata={"output": "final code"})
    
    event_bus.emit(run_id=run_id, action=EventAction.RUN_FINISHED)
    
    # 2. Reconstruct
    state = replay_engine.reconstruct_state(run_id)
    
    # 3. Verify
    assert state.run_id == run_id
    assert state.goal == "test goal"
    assert "task1" in state.tasks
    assert state.tasks["task1"].status == TaskStatus.SUCCEEDED
    assert state.tasks["task1"].output == "final code"

def test_point_in_time_replay(setup_replay):
    tmp_path, run_id, run_dir, event_bus, replay_engine = setup_replay
    
    # Time T0: Start
    event_bus.emit(run_id=run_id, action=EventAction.RUN_STARTED, metadata={"goal": "goal"})
    t0 = datetime.now(UTC).isoformat()
    
    # Time T1: Task Created
    task1 = PlannedTask(id="task1", type=PlannedTaskType.CODE_GENERATION, goal="g1", validation=[])
    event_bus.emit(run_id=run_id, task_id="task1", action=EventAction.TASK_CREATED, metadata=task1.model_dump(mode="json"))
    t1 = datetime.now(UTC).isoformat()
    
    # Time T2: Task Success
    event_bus.emit(run_id=run_id, task_id="task1", action=EventAction.TASK_SUCCEEDED, metadata={"output": "out"})
    
    # Replay up to T1
    state_t1 = replay_engine.reconstruct_state(run_id, up_to_timestamp=t1)
    assert "task1" in state_t1.tasks
    assert state_t1.tasks["task1"].status == TaskStatus.PENDING # Hasn't succeeded yet at T1
    
    # Replay full
    state_full = replay_engine.reconstruct_state(run_id)
    assert state_full.tasks["task1"].status == TaskStatus.SUCCEEDED

def test_artifact_invalidation_replay(setup_replay):
    tmp_path, run_id, run_dir, event_bus, replay_engine = setup_replay
    
    event_bus.emit(run_id=run_id, action=EventAction.RUN_STARTED)
    task1 = PlannedTask(id="task1", type=PlannedTaskType.CODE_GENERATION, goal="g1", validation=[])
    event_bus.emit(run_id=run_id, task_id="task1", action=EventAction.TASK_CREATED, metadata=task1.model_dump(mode="json"))
    event_bus.emit(run_id=run_id, task_id="task1", action=EventAction.TASK_SUCCEEDED)
    
    # Invalidate
    event_bus.emit(run_id=run_id, task_id="task1", action=EventAction.ARTIFACT_INVALIDATED, metadata={"failed_task_id": "other"})
    
    state = replay_engine.reconstruct_state(run_id)
    assert state.tasks["task1"].status == TaskStatus.PENDING
    assert state.tasks["task1"].subtree_resets == 1

def test_consistency_verification(setup_replay):
    tmp_path, run_id, run_dir, event_bus, replay_engine = setup_replay
    
    # 1. State from events
    event_bus.emit(run_id=run_id, action=EventAction.RUN_STARTED, metadata={"goal": "g"})
    event_bus.emit(run_id=run_id, action=EventAction.RUN_FINISHED)
    
    # 2. Persist a DIFFERENT state to run.json (simulate hidden mutation)
    run_json = run_dir / "run.json"
    bad_state = {
        "run_id": run_id,
        "goal": "CORRUPTED GOAL",
        "tasks": {}
    }
    run_json.write_text(json.dumps(bad_state))
    
    # 3. Verify consistency fails
    assert replay_engine.verify_consistency(run_id) is False
