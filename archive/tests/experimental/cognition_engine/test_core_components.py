"""Core component tests."""

from __future__ import annotations

import asyncio
import json
import time
from pathlib import Path
from concurrent.futures import ThreadPoolExecutor

import pytest
from pydantic import ValidationError

from contracts.models import ModelType, Task, TaskType
from events import EventAction, EventBus
from help import HelpRequestBuilder
from memory.store import SQLiteMemoryStore
from models import ModelRegistry
from models.ollama import OllamaClient
from models.protocol import ModelCallResult
from orchestrator.dag import compute_affected_subgraph, topological_batches
from orchestrator.control_plane import ControlPlane
from orchestrator.router import TaskRouter
from planner import PlannedTask, PlannerOutput, TaskDecomposer
from repair import RepairPlanner
from repo_indexer import RepoIndexer
from runtime import ReplayEngine
from runtime.snapshots import SnapshotEngine
from sandbox.runtime import DockerSandbox
from shared_context import ArtifactStore, RunState
from tools.permissions import ApprovalManager, ApprovalRequired, PermissionLevel
from transactions import DiffPlan, TransactionRunner
from validator import Critic, FailureType, failure_from_error, validate_artifact, validate_python_syntax


class PlanningModelClient:
    """Explicit test double: returns valid planner JSON, valid code, and accepting critic."""

    _PLAN_JSON = json.dumps({
        "version": "1",
        "tasks": [
            {
                "id": "task_001",
                "type": "planning",
                "goal": "Clarify architecture and implementation constraints",
                "depends_on": [],
                "validation": ["artifact"],
                "assigned_model": "phi3:mini",
            },
            {
                "id": "task_002",
                "type": "code_generation",
                "goal": "Produce concrete Python implementation artifact",
                "depends_on": ["task_001"],
                "validation": ["artifact", "python_syntax"],
                "assigned_model": "qwen2.5-coder:7b",
            },
            {
                "id": "task_003",
                "type": "validation",
                "goal": "Review generated artifacts and summarize residual risk",
                "depends_on": ["task_002"],
                "validation": ["artifact"],
                "assigned_model": "phi3:mini",
            },
        ],
    }, indent=2)

    def generate(self, model: str, prompt: str, *, timeout_s: int = 120) -> ModelCallResult:
        if "Return JSON only" in prompt and '"tasks"' in prompt:
            response = self._PLAN_JSON
        elif "You are a critic" in prompt or "accepted:" in prompt:
            response = '{"accepted": true, "issues": [], "notes": "test double accepted"}'
        elif '"tasks"' in prompt or "corrected JSON" in prompt:
            response = self._PLAN_JSON
        else:
            response = (
                "# Artifact\n\n"
                "```python\ndef generated():\n    return 'ok'\n```\n"
            )
        return ModelCallResult(model=model, prompt=prompt, response=response)


class RepairingModelClient:
    """Explicit test double: returns broken code first, then repaired code."""

    def generate(self, model: str, prompt: str, *, timeout_s: int = 120) -> ModelCallResult:
        if "bounded repair task" in prompt.lower():
            response = "```python\ndef repaired():\n    return 'ok'\n```"
        elif "critic" in prompt.lower():
            response = '{"accepted": true, "issues": [], "notes": "ok"}'
        else:
            response = "```python\ndef broken(:\n    return 'bad'\n```"
        return ModelCallResult(model=model, prompt=prompt, response=response)


class RejectingCriticModelClient:
    """Explicit test double: always rejects with issues."""

    def generate(self, model: str, prompt: str, *, timeout_s: int = 120) -> ModelCallResult:
        return ModelCallResult(
            model=model,
            prompt=prompt,
            response='```json\n{"accepted": false, "issues": ["bad"], "notes": "reject"}\n```',
        )


def test_planner_schema_rejects_bad_dependency() -> None:
    with pytest.raises(ValueError):
        PlannerOutput.model_validate(
            {
                "version": "1",
                "tasks": [
                    {
                        "id": "task_001",
                        "type": "planning",
                        "goal": "x",
                        "depends_on": ["missing"],
                        "validation": ["artifact"],
                    }
                ],
            }
        )


def test_mock_decomposer_returns_valid_dag() -> None:
    plan = TaskDecomposer(PlanningModelClient()).decompose("build a REST API")  # type: ignore[arg-type]
    assert len(plan.tasks) >= 2
    assert topological_batches(plan)


def test_planner_schema_coerces_common_task_type_aliases() -> None:
    plan = PlannerOutput.model_validate(
        {
            "version": "1",
            "tasks": [
                {
                    "id": "task_001",
                    "type": "pytest testing",
                    "goal": "write pytest tests",
                    "depends_on": [],
                    "validation": ["artifact", "python_syntax"],
                }
            ],
        }
    )
    assert plan.tasks[0].type.value == "testing"


def test_planner_schema_rejects_unknown_model_and_enforces_routing() -> None:
    with pytest.raises(ValidationError, match="hallucinated or invalid model name"):
        PlannerOutput.model_validate(
            {
                "version": "1",
                "tasks": [
                    {
                        "id": "task_001",
                        "type": "code_generation",
                        "goal": "write python",
                        "depends_on": [],
                        "validation": ["artifact"],
                        "assigned_model": "qwen2anb:7b",
                    }
                ],
            }
        )

    # Test routing enforcement
    plan = PlannerOutput.model_validate(
        {
            "version": "1",
            "tasks": [
                {
                    "id": "task_001",
                    "type": "code_generation",
                    "goal": "write python",
                    "depends_on": [],
                    "validation": ["artifact"],
                }
            ],
        }
    )
    assert plan.tasks[0].assigned_model == "qwen2.5-coder:7b"
    assert "python_syntax" in plan.tasks[0].validation


def test_event_bus_appends_jsonl(tmp_path) -> None:
    bus = EventBus(tmp_path / "events.jsonl")
    bus.emit(run_id="run_1", action=EventAction.RUN_STARTED)
    bus.emit(run_id="run_1", action=EventAction.RUN_FINISHED)
    events = bus.read_all()
    assert [event.action for event in events] == [
        EventAction.RUN_STARTED,
        EventAction.RUN_FINISHED,
    ]


def test_artifact_store_never_overwrites_run(tmp_path) -> None:
    store = ArtifactStore(tmp_path)
    store.create_run("run_1")
    with pytest.raises(FileExistsError):
        store.create_run("run_1")


def test_sqlite_memory_search(tmp_path) -> None:
    store = SQLiteMemoryStore(tmp_path / "memory.sqlite3")
    store.record_failure_lesson("run_1", "task_1", "syntax error from missing colon")
    rows = store.search("syntax error", kinds=["failure"])
    assert rows
    assert rows[0]["kind"] == "failure"


def test_task_router_legacy_and_planned_routing() -> None:
    router = TaskRouter()
    legacy = Task(id="t", task_type=TaskType.DEBUGGING, description="fix bug")
    assert router.route(legacy).assigned_model == ModelType.DEEPSEEK

    plan = PlannerOutput.model_validate(
        {
            "version": "1",
            "tasks": [
                {
                    "id": "task_001",
                    "type": "code_generation",
                    "goal": "write code",
                    "depends_on": [],
                    "validation": ["artifact"],
                }
            ],
        }
    )
    assert router.route_planned(plan.tasks[0]) == ModelType.QWEN.value


def test_validators() -> None:
    assert validate_artifact("content")["ok"]
    assert validate_python_syntax("```python\ndef x():\n    return 1\n```")["ok"]
    assert not validate_python_syntax("```python\ndef x(:\n```")["ok"]


def test_critic_rejects_when_model_json_says_rejected() -> None:
    result = Critic(RejectingCriticModelClient()).review("goal", "non-empty output")  # type: ignore[arg-type]
    assert not result["ok"]
    assert result["issues"] == ["bad"]


def test_critic_rejection_does_not_become_syntax_failure(tmp_path) -> None:
    client = OllamaClient(allow_mock=True)
    engine = ControlPlane(artifact_root=tmp_path / "artifacts", data_dir=tmp_path / "data", model_client=client)
    failure_type = engine._failure_type_from_validation(
        {
            "ok": False,
            "checks": [
                {"name": "artifact", "ok": True},
                {"name": "python_syntax", "ok": True},
                {"name": "critic", "ok": False, "issues": ["lacks specific implementation detail"]},
            ],
        }
    )
    assert failure_type == FailureType.UNCLASSIFIED


def test_python_syntax_validation_requires_code_block_for_python_tasks(tmp_path) -> None:
    client = OllamaClient(allow_mock=True)
    engine = ControlPlane(artifact_root=tmp_path / "artifacts", data_dir=tmp_path / "data", model_client=client)
    task = PlannedTask(
        id="task_001",
        type="code_generation",
        goal="write python",
        depends_on=[],
        validation=["artifact", "python_syntax"],
    )
    validation = engine._validate(task, "plain prose")
    syntax_check = next(check for check in validation["checks"] if check["name"] == "python_syntax")
    assert not validation["ok"]
    assert syntax_check["errors"] == ["missing fenced Python code block"]


def test_approval_manager_gates_dangerous_actions() -> None:
    manager = ApprovalManager()
    manager.require("read status", PermissionLevel.READ_ONLY)
    with pytest.raises(ApprovalRequired) as exc:
        manager.require("delete files", PermissionLevel.DANGEROUS)
    decision = manager.resolve(exc.value.approval_id, False, "unsafe")
    assert not decision.approved


def test_repo_indexer_scans_python_symbols(tmp_path) -> None:
    source = tmp_path / "module.py"
    source.write_text("class A:\n    pass\n\ndef f():\n    return 1\n", encoding="utf-8")
    index = RepoIndexer(tmp_path).scan()
    names = {symbol["name"] for symbol in index["symbols"]}
    assert {"A", "f"} <= names


def test_docker_sandbox_reports_missing_or_result(tmp_path) -> None:
    result = DockerSandbox().run_python("print('ok')", tmp_path)
    assert "ok" in result
    assert "stderr" in result


def test_cognition_engine_full_mock_run(tmp_path) -> None:
    engine = ControlPlane(
        artifact_root=tmp_path / "artifacts",
        data_dir=tmp_path / "data",
        model_client=PlanningModelClient(),  # type: ignore[arg-type]
    )
    result = asyncio.run(engine.run_goal("build a REST API skeleton with tests", str(tmp_path)))
    assert result["run_id"].startswith("run_")
    assert result["artifact_dir"]
    assert result["plan"]["tasks"]
    assert result["metrics"]["counters"]["planner_valid"] == 1
    # Check for events in the new isolated namespace
    event_log = tmp_path / "data" / "events" / "default" / result["run_id"] / "main.jsonl"
    assert event_log.exists()
    assert (Path(result["artifact_dir"]) / "plan.json").exists()
    
    from shared_context.execution_context import ExecutionContext
    replay = ReplayEngine(tmp_path / "data", context=ExecutionContext(workspace_id=result["run_id"])).reconstruct_state()
    assert replay.run_id == result["run_id"]
    assert len(replay.tasks) == 3


def test_model_registry_legacy_api() -> None:
    models = ModelRegistry.list_models()
    assert ModelType.MISTRAL in models
    assert ModelRegistry.get_adapter(ModelType.PHI).__class__.__name__ == "PhiAdapter"


def test_failure_record_classification_and_persistence(tmp_path) -> None:
    failure = failure_from_error(
        run_id="run_1",
        task_id="task_1",
        raw_error='ModuleNotFoundError: No module named "jwt"',
        attempt=1,
    )
    assert failure.failure_type == FailureType.IMPORT_ERROR
    assert failure.structured_context["module"] == "jwt"

    store = SQLiteMemoryStore(tmp_path / "memory.sqlite3")
    store.record_failure_record(failure.storage_payload())
    row = store.conn.execute(
        "SELECT failure_type, structured_context FROM failure_records WHERE failure_id = ?",
        (failure.failure_id,),
    ).fetchone()
    assert row["failure_type"] == FailureType.IMPORT_ERROR.value
    assert "jwt" in row["structured_context"]


def test_failure_record_extended_classification() -> None:
    circ = failure_from_error(run_id="r", task_id="t", raw_error="circular import 'foo'", attempt=1)
    assert circ.failure_type == FailureType.CIRCULAR_IMPORT
    assert circ.structured_context["module"] == "foo"

    dep = failure_from_error(run_id="r", task_id="t", raw_error="pydantic version conflict", attempt=1)
    assert dep.failure_type == FailureType.DEPENDENCY_CONFLICT
    assert dep.structured_context["package"] == "pydantic"

    sym = failure_from_error(run_id="r", task_id="t", raw_error="module has no attribute 'missing_func'", attempt=1)
    assert sym.failure_type == FailureType.HALLUCINATED_SYMBOL
    assert sym.structured_context["symbol"] == "missing_func"

    patch = failure_from_error(run_id="r", task_id="t", raw_error="patching file 'app.py' failed", attempt=1)
    assert patch.failure_type == FailureType.INVALID_PATCH
    assert patch.structured_context["file"] == "app.py"

    broken = failure_from_error(run_id="r", task_id="t", raw_error="assert expected 'a', actual 'b' in test_foo", attempt=1)
    assert broken.failure_type == FailureType.BROKEN_TEST_ASSUMPTION
    assert broken.structured_context["test_name"] == "test_foo"

    stale = failure_from_error(run_id="r", task_id="t", raw_error="stale context for 'models.py'", attempt=1)
    assert stale.failure_type == FailureType.STALE_CONTEXT
    assert stale.structured_context["stale_ref"] == "models.py"


def test_repair_evaluator_classifies_outcomes() -> None:
    from repair.evaluator import RepairEvaluator, RepairStatus
    from shared_context.state import RunState, TaskStatus
    from planner.schemas import PlannedTask, PlannedTaskType

    state = RunState(run_id="r", goal="g")
    
    t1 = PlannedTask(id="t1", type=PlannedTaskType.CODE_GENERATION, goal="g1", depends_on=[], validation=[])
    t2 = PlannedTask(id="t2", type=PlannedTaskType.TESTING, goal="g2", depends_on=["t1"], validation=[])
    
    state.add_task(t1)
    state.add_task(t2)
    
    evaluator = RepairEvaluator()
    
    # Test Success
    state.set_status("t1", TaskStatus.SUCCEEDED)
    state.set_status("t2", TaskStatus.SUCCEEDED)
    outcome = evaluator.evaluate(state, "t1", "syntax_error")
    assert outcome.status == RepairStatus.SUCCESS
    assert outcome.fixed_task is True
    assert outcome.downstream_failures == []

    # Test Regressed
    state.set_status("t1", TaskStatus.SUCCEEDED)
    state.set_status("t2", TaskStatus.FAILED)
    outcome = evaluator.evaluate(state, "t1", "syntax_error")
    assert outcome.status == RepairStatus.REGRESSED
    assert outcome.fixed_task is True
    assert outcome.downstream_failures == ["t2"]

    # Test Partial
    state.set_status("t1", TaskStatus.FAILED)
    state.tasks["t1"].failure = "import_error"
    outcome = evaluator.evaluate(state, "t1", "syntax_error")
    assert outcome.status == RepairStatus.PARTIAL
    assert outcome.fixed_task is False

    # Test Failed
    state.set_status("t1", TaskStatus.FAILED)
    state.tasks["t1"].failure = "syntax_error"
    outcome = evaluator.evaluate(state, "t1", "syntax_error")
    assert outcome.status == RepairStatus.FAILED
    assert outcome.fixed_task is False


def test_dag_differential_invalidation(tmp_path) -> None:
    from orchestrator.dag import invalidate_downstream_if_changed
    from planner.schemas import PlannedTask, PlannedTaskType, PlannerOutput
    from shared_context.state import RunState, TaskStatus
    from events.immutable_log import EventBus
    from events.schema import EventAction

    t1 = PlannedTask(id="t1", type=PlannedTaskType.CODE_GENERATION, goal="g1", depends_on=[], validation=[])
    t2 = PlannedTask(id="t2", type=PlannedTaskType.TESTING, goal="g2", depends_on=["t1"], validation=[])
    plan = PlannerOutput(tasks=[t1, t2])
    
    state = RunState(run_id="r", goal="g")
    state.add_task(t1)
    state.add_task(t2)
    state.set_status("t2", TaskStatus.SUCCEEDED)
    
    bus = EventBus(tmp_path / "events.jsonl")
    events = []
    bus.subscribe(lambda e: events.append(e))

    # Identical output
    invalidate_downstream_if_changed(plan, "t1", "old_code", "old_code", state, bus)
    assert any(e.action == EventAction.SUBTREE_SKIP for e in events)
    assert state.tasks["t2"].status == TaskStatus.SUCCEEDED
    events.clear()

    # Changed output
    invalidate_downstream_if_changed(plan, "t1", "old_code", "new_code", state, bus)
    assert any(e.action == EventAction.SUBTREE_RERUN for e in events)
    assert state.tasks["t2"].status == TaskStatus.PENDING

def test_assess_rollback_risk(tmp_path) -> None:
    from runtime.snapshots import assess_rollback_risk, RollbackRisk
    
    repo = tmp_path / "repo"
    run_dir = tmp_path / "artifacts" / "run_1"
    repo.mkdir(parents=True)
    run_dir.mkdir(parents=True)
    
    (run_dir / "foo.txt").touch()
    assert assess_rollback_risk([str(run_dir / "foo.txt")], repo, run_dir) == RollbackRisk.LOW
    
    (repo / "src").mkdir()
    (repo / "src" / "app.py").touch()
    assert assess_rollback_risk([str(repo / "src" / "app.py")], repo, run_dir) == RollbackRisk.MEDIUM
    
    (repo / "Makefile").touch()
    assert assess_rollback_risk([str(repo / "Makefile")], repo, run_dir) == RollbackRisk.DANGEROUS
    
    (repo / "config.yaml").touch()
    assert assess_rollback_risk([str(repo / "config.yaml")], repo, run_dir) == RollbackRisk.DANGEROUS

    (repo / "Dockerfile").touch()
    assert assess_rollback_risk([str(repo / "src" / "app.py"), str(repo / "Dockerfile")], repo, run_dir) == RollbackRisk.DANGEROUS

    out = tmp_path / "out.txt"
    out.touch()
    assert assess_rollback_risk([str(out)], repo, run_dir) == RollbackRisk.DANGEROUS

def test_failure_memory_semantic_retrieval(tmp_path) -> None:
    from memory.failure_memory import FailureMemory
    from memory.store import SQLiteMemoryStore
    
    store = SQLiteMemoryStore(tmp_path / "memory.sqlite3")
    fm = FailureMemory(sqlite_store=store)
    
    fid1 = fm.record_failure("syntax_error", "bad syntax", {"task_id": "t1"}, None)
    fid2 = fm.record_failure("syntax_error", "other syntax", {"task_id": "t2"}, None)
    
    fm.record_repair_outcome(fid1, "success")
    fm.record_repair_outcome(fid2, "failed")
    
    similar = fm.find_similar_failures("syntax_error", "new error")
    assert len(similar) == 1
    assert similar[0]["id"] == fid1
    assert similar[0]["resolution"] == "success"

def test_repair_planner_is_budgeted_and_minimal() -> None:
    failure = failure_from_error(
        run_id="run_1",
        task_id="task_1",
        raw_error="SyntaxError: invalid syntax at line 3",
        attempt=1,
    )
    planner = RepairPlanner(repair_retry_budget=1)
    plan = planner.plan(failure, current_repair_count=0)
    assert len(plan.repair_tasks) == 1
    assert plan.repair_tasks[0].id.startswith("repair_")

    blocked = planner.plan(failure, current_repair_count=1)
    assert blocked.blocked_reason == "repair budget exhausted"


def test_affected_subgraph_returns_transitive_dependents() -> None:
    plan = PlannerOutput.model_validate(
        {
            "version": "1",
            "tasks": [
                {"id": "a", "type": "planning", "goal": "a", "depends_on": [], "validation": ["artifact"]},
                {"id": "b", "type": "planning", "goal": "b", "depends_on": ["a"], "validation": ["artifact"]},
                {"id": "c", "type": "planning", "goal": "c", "depends_on": ["b"], "validation": ["artifact"]},
                {"id": "d", "type": "planning", "goal": "d", "depends_on": [], "validation": ["artifact"]},
            ],
        }
    )
    assert compute_affected_subgraph(plan, "a") == {"a", "b", "c"}


def test_snapshot_create_and_rollback_excludes_runtime_dirs(tmp_path) -> None:
    project = tmp_path / "project"
    project.mkdir()
    source = project / "module.py"
    source.write_text("value = 1\n", encoding="utf-8")
    (project / "artifacts").mkdir()
    (project / "artifacts" / "ignored.txt").write_text("ignore", encoding="utf-8")

    state = RunState(run_id="run_1", goal="test")
    engine = SnapshotEngine(tmp_path / "run")
    snapshot = engine.create(run_id="run_1", root=project, state=state)
    source.write_text("value = 2\n", encoding="utf-8")

    engine.rollback_files(snapshot)
    assert source.read_text(encoding="utf-8") == "value = 1\n"
    assert all(not item["path"].startswith("artifacts/") for item in snapshot["manifest"])


def test_help_request_builder_includes_failure_evidence() -> None:
    state = RunState(run_id="run_1", goal="test")
    failure = failure_from_error(
        run_id="run_1",
        task_id="task_1",
        raw_error="planner output was not valid json",
        attempt=1,
        forced_type=FailureType.PLANNER_INVALID,
    )
    request = HelpRequestBuilder().build(state=state, failures=[failure], attempted_repairs=[])
    assert request.status == "awaiting_human_input"
    assert request.evidence[0]["failure_type"] == "planner_invalid"


def test_engine_executes_syntax_repair_plan(tmp_path) -> None:
    engine = ControlPlane(
        artifact_root=tmp_path / "artifacts",
        data_dir=tmp_path / "data",
        model_client=RepairingModelClient(),  # type: ignore[arg-type]
    )
    task = PlannedTask(
        id="task_1",
        type="code_generation",
        goal="write valid python",
        validation=["artifact", "python_syntax"],
    )
    plan = PlannerOutput(tasks=[task])
    state = RunState(run_id="run_1", goal="repair")
    state.add_task(task)
    run_dir = engine.artifacts.create_run("run_1")
    bus = EventBus(run_dir / "events.jsonl")

    output, validation = asyncio.run(engine._execute_task(task, plan, state, None, bus))

    assert validation["ok"]
    assert "repaired" in output
    assert state.tasks["task_1"].repair_count == 1
    repair_tasks = [task_id for task_id in state.tasks if task_id.startswith("repair_")]
    assert repair_tasks
    assert state.tasks[repair_tasks[0]].status.value == "succeeded"
    assert EventAction.REPAIR_PLANNED in {event.action for event in bus.read_all()}


def test_transaction_runner_stages_validates_and_commits(tmp_path) -> None:
    project = tmp_path / "project"
    project.mkdir()
    source = project / "module.py"
    source.write_text("def old():\n    return 1\n", encoding="utf-8")

    state = RunState(run_id="run_1", goal="transaction")
    snapshot_engine = SnapshotEngine(tmp_path / "run")
    runner = TransactionRunner(tmp_path / "tx")
    plan = DiffPlan(
        patches=[
            {
                "path": "module.py",
                "old_text": "def old():\n    return 1\n",
                "new_text": "def new():\n    return 2\n",
            }
        ],
        validation_commands=["python -m py_compile module.py"],
    )

    result = runner.apply(
        project_root=project,
        plan=plan,
        state=state,
        snapshot_engine=snapshot_engine,
    )

    assert result.ok
    assert result.changed_files == ["module.py"]
    assert source.read_text(encoding="utf-8") == "def new():\n    return 2\n"


def test_transaction_runner_blocks_failed_validation(tmp_path) -> None:
    project = tmp_path / "project"
    project.mkdir()
    source = project / "module.py"
    source.write_text("def old():\n    return 1\n", encoding="utf-8")

    runner = TransactionRunner(tmp_path / "tx")
    result = runner.apply(
        project_root=project,
        plan=DiffPlan(
            patches=[{"path": "module.py", "new_text": "def broken(:\n"}],
            validation_commands=["python -m py_compile module.py"],
        ),
        state=RunState(run_id="run_1", goal="transaction"),
        snapshot_engine=SnapshotEngine(tmp_path / "run"),
    )

    assert not result.ok
    assert result.error == "validation failed"
    assert source.read_text(encoding="utf-8") == "def old():\n    return 1\n"


def test_ollama_client_enforces_single_active_generation(monkeypatch) -> None:
    client = OllamaClient(allow_mock=True)
    active = 0
    max_active = 0

    def slow_mock(model: str, prompt: str, start: float, error: str) -> ModelCallResult:
        nonlocal active, max_active
        active += 1
        max_active = max(max_active, active)
        time.sleep(0.02)
        active -= 1
        return ModelCallResult(model=model, prompt=prompt, response="ok")

    monkeypatch.setattr(client, "_mock_response", slow_mock)
    with ThreadPoolExecutor(max_workers=4) as pool:
        results = list(pool.map(lambda index: client.generate("phi3:mini", f"prompt {index}"), range(8)))

    assert all(result.ok for result in results)
    assert max_active == 1
