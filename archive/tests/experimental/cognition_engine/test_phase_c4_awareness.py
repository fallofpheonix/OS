"""Tests for Phase C.4: Architecture-aware validation and repair.

Covers:
- ContextResolver (relevance ranking, blast radius, temperature)
- Architecture-aware Critic (deterministic invariant supplement)
- Architecture-aware RepairPlanner (temperature budgets, new strategies)
"""

from __future__ import annotations



from models.protocol import ModelCallResult
from repo_indexer.models import (
    ArchitecturalNode,
    ArchitecturalTemperature,
    ArchitectureGraph,
    BoundaryViolation,
    Mutability,
    SideEffectType,
)
from validator.context_resolver import ContextResolver
from validator.critic import Critic
from validator.failure_types import FailureType
from validator.context_extractors import failure_from_error
from repair.repair_planner import RepairPlanner


class AcceptingModelClient:
    """Always accepts."""
    def generate(self, model: str, prompt: str, *, timeout_s: int = 120) -> ModelCallResult:
        return ModelCallResult(
            model=model,
            prompt=prompt,
            response='{"accepted": true, "issues": [], "notes": "ok"}',
        )


class RejectingModelClient:
    """Always rejects."""
    def generate(self, model: str, prompt: str, *, timeout_s: int = 120) -> ModelCallResult:
        return ModelCallResult(
            model=model,
            prompt=prompt,
            response='{"accepted": false, "issues": ["bad"], "notes": "reject"}',
        )


def _build_graph() -> ArchitectureGraph:
    """Build a test architecture graph with known topology."""
    graph = ArchitectureGraph()

    graph.add_node(ArchitecturalNode(
        module_path="contracts/models.py",
        subsystem="core",
        layer="contracts",
        mutability=Mutability.FROZEN,
        public_api=["TaskType", "Task"],
        imports=[],
    ))

    graph.add_node(ArchitecturalNode(
        module_path="orchestrator/control_plane.py",
        subsystem="engine",
        layer="control-plane",
        imports=["contracts/models.py", "validator/critic.py"],
        imported_by=["api/main.py"],
        public_api=["ControlPlane"],
        side_effects=[SideEffectType.FILESYSTEM_WRITE, SideEffectType.NETWORK_CALL],
    ))

    graph.add_node(ArchitecturalNode(
        module_path="validator/critic.py",
        subsystem="engine",
        layer="validation",
        imports=["contracts/models.py"],
        imported_by=["orchestrator/control_plane.py"],
        public_api=["Critic"],
    ))

    graph.add_node(ArchitecturalNode(
        module_path="api/main.py",
        subsystem="infrastructure",
        layer="api",
        imports=["orchestrator/control_plane.py"],
        public_api=["app"],
    ))

    graph.add_edge("orchestrator/control_plane.py", "contracts/models.py")
    graph.add_edge("orchestrator/control_plane.py", "validator/critic.py")
    graph.add_edge("validator/critic.py", "contracts/models.py")
    graph.add_edge("api/main.py", "orchestrator/control_plane.py")

    return graph


# ═══════════════════════════════════════════════════════════════════════════
# ContextResolver
# ═══════════════════════════════════════════════════════════════════════════


class TestContextResolver:
    def test_resolve_known_module(self) -> None:
        graph = _build_graph()
        resolver = ContextResolver(graph)
        ctx = resolver.resolve_for_task(
            task_goal="fix the engine",
            task_type="code_generation",
            target_modules=["orchestrator/control_plane.py"],
        )
        assert ctx.layer == "control-plane"
        assert ctx.subsystem == "engine"
        assert "contracts/models.py" in ctx.imports
        assert "api/main.py" in ctx.imported_by
        assert ctx.blast_radius_size > 0

    def test_resolve_unknown_module_returns_empty(self) -> None:
        graph = _build_graph()
        resolver = ContextResolver(graph)
        ctx = resolver.resolve_for_task(
            task_goal="do something with nonexistent",
            task_type="code_generation",
            target_modules=["nonexistent.py"],
        )
        assert ctx.layer == ""
        assert ctx.blast_radius_size == 0

    def test_resolve_infers_modules_from_goal(self) -> None:
        graph = _build_graph()
        resolver = ContextResolver(graph)
        ctx = resolver.resolve_for_task(
            task_goal="fix the critic validation logic",
            task_type="debugging",
        )
        assert "validator/critic.py" in ctx.relevant_modules

    def test_temperature_warnings(self) -> None:
        graph = _build_graph()
        graph.temperatures["orchestrator/control_plane.py"] = ArchitecturalTemperature(
            module_path="orchestrator/control_plane.py",
            recent_mutations=8,
            invariant_failures=2,
        )
        resolver = ContextResolver(graph)
        ctx = resolver.resolve_for_task(
            task_goal="modify engine",
            task_type="code_generation",
            target_modules=["orchestrator/control_plane.py"],
        )
        assert len(ctx.temperature_warnings) > 0
        assert "HOT" in ctx.temperature_warnings[0]

    def test_as_prompt_section_is_bounded(self) -> None:
        graph = _build_graph()
        resolver = ContextResolver(graph)
        ctx = resolver.resolve_for_task(
            task_goal="fix engine",
            task_type="code_generation",
            target_modules=["orchestrator/control_plane.py"],
        )
        section = ctx.as_prompt_section(max_chars=200)
        assert len(section) <= 200

    def test_resolve_for_repair(self) -> None:
        graph = _build_graph()
        resolver = ContextResolver(graph)
        ctx = resolver.resolve_for_repair(
            failed_task_id="task_001",
            failure_type="syntax_error",
            affected_file="orchestrator/control_plane.py",
        )
        assert ctx.layer == "control-plane"
        assert ctx.blast_radius_size > 0


# ═══════════════════════════════════════════════════════════════════════════
# Architecture-aware Critic
# ═══════════════════════════════════════════════════════════════════════════


class TestArchitectureAwareCritic:
    def test_critic_without_graph_behaves_normally(self) -> None:
        critic = Critic(model_client=AcceptingModelClient())  # type: ignore[arg-type]
        result = critic.review("write code", "some output")
        assert result["ok"] is True
        assert result["architectural_context"] is None

    def test_critic_with_graph_includes_context(self) -> None:
        graph = _build_graph()
        critic = Critic(
            model_client=AcceptingModelClient(),  # type: ignore[arg-type]
            architecture_graph=graph,
        )
        result = critic.review(
            "fix the engine",
            "some output",
            target_modules=["orchestrator/control_plane.py"],
        )
        assert result["ok"] is True
        assert result["architectural_context"] is not None
        assert result["architectural_context"]["layer"] == "control-plane"

    def test_critic_rejects_hot_module_mutations(self) -> None:
        graph = _build_graph()
        graph.temperatures["orchestrator/control_plane.py"] = ArchitecturalTemperature(
            module_path="orchestrator/control_plane.py",
            recent_mutations=10,
            invariant_failures=3,
        )
        critic = Critic(
            model_client=AcceptingModelClient(),  # type: ignore[arg-type]
            architecture_graph=graph,
        )
        result = critic.review(
            "modify engine",
            "some code",
            target_modules=["orchestrator/control_plane.py"],
        )
        # Model says accepted, but deterministic invariant check adds temperature issue
        assert result["ok"] is False
        assert any("TEMPERATURE" in issue for issue in result["issues"])

    def test_critic_flags_existing_violations(self) -> None:
        graph = _build_graph()
        graph.boundary_violations.append(BoundaryViolation(
            invariant_id="test",
            source_module="orchestrator/control_plane.py",
            target_module="validator/critic.py",
            violation_type="forbidden_import",
            message="orchestrator imports validator directly",
        ))
        critic = Critic(
            model_client=AcceptingModelClient(),  # type: ignore[arg-type]
            architecture_graph=graph,
        )
        result = critic.review(
            "fix engine",
            "code",
            target_modules=["orchestrator/control_plane.py"],
        )
        assert result["ok"] is False
        assert any("EXISTING VIOLATION" in issue for issue in result["issues"])


# ═══════════════════════════════════════════════════════════════════════════
# Architecture-aware RepairPlanner
# ═══════════════════════════════════════════════════════════════════════════


class TestArchitectureAwareRepairPlanner:
    def test_planner_without_graph_behaves_normally(self) -> None:
        planner = RepairPlanner(repair_retry_budget=1)
        failure = failure_from_error(
            run_id="r", task_id="t", raw_error="SyntaxError: invalid syntax", attempt=1
        )
        plan = planner.plan(failure, current_repair_count=0)
        assert len(plan.repair_tasks) == 1
        assert plan.architectural_context is None

    def test_planner_with_graph_includes_context(self) -> None:
        graph = _build_graph()
        planner = RepairPlanner(repair_retry_budget=1, architecture_graph=graph)
        failure = failure_from_error(
            run_id="r", task_id="t", raw_error="SyntaxError: invalid syntax", attempt=1
        )
        plan = planner.plan(
            failure, current_repair_count=0, affected_file="orchestrator/control_plane.py"
        )
        assert len(plan.repair_tasks) == 1
        assert plan.architectural_context is not None
        assert plan.architectural_context["layer"] == "control-plane"

    def test_hot_module_reduces_budget(self) -> None:
        graph = _build_graph()
        graph.temperatures["orchestrator/control_plane.py"] = ArchitecturalTemperature(
            module_path="orchestrator/control_plane.py",
            recent_mutations=10,
            invariant_failures=3,
        )
        planner = RepairPlanner(repair_retry_budget=1, architecture_graph=graph)
        failure = failure_from_error(
            run_id="r", task_id="t", raw_error="SyntaxError: invalid syntax", attempt=1
        )
        # Budget is 1, but hot module reduces it to 0 → blocked immediately
        plan = planner.plan(
            failure, current_repair_count=0, affected_file="orchestrator/control_plane.py"
        )
        assert plan.blocked_reason is not None
        assert "hot module" in plan.blocked_reason

    def test_hallucinated_api_strategy(self) -> None:
        graph = _build_graph()
        planner = RepairPlanner(architecture_graph=graph)
        failure = failure_from_error(
            run_id="r", task_id="t",
            raw_error="module has no attribute 'nonexistent_api'",
            attempt=1,
            forced_type=FailureType.HALLUCINATED_API,
        )
        plan = planner.plan(failure, current_repair_count=0)
        assert len(plan.repair_tasks) == 1
        assert "api" in plan.repair_tasks[0].id

    def test_circular_import_strategy(self) -> None:
        graph = _build_graph()
        planner = RepairPlanner(architecture_graph=graph)
        failure = failure_from_error(
            run_id="r", task_id="t",
            raw_error="circular import 'models'",
            attempt=1,
        )
        plan = planner.plan(failure, current_repair_count=0)
        assert len(plan.repair_tasks) == 1
        assert "circular" in plan.repair_tasks[0].id

    def test_import_error_includes_available_deps(self) -> None:
        graph = _build_graph()
        planner = RepairPlanner(architecture_graph=graph)
        failure = failure_from_error(
            run_id="r", task_id="t",
            raw_error="ModuleNotFoundError: No module named 'missing'",
            attempt=1,
        )
        plan = planner.plan(
            failure, current_repair_count=0, affected_file="orchestrator/control_plane.py"
        )
        assert len(plan.repair_tasks) == 1
        # The repair goal should mention available dependencies
        assert "Available dependencies" in plan.repair_tasks[0].goal
