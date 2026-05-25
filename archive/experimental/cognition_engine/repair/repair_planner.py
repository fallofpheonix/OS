"""Architecture-aware repair planner.

FailureRecord -> minimal repair task DAG, now with:
- Blast radius awareness (scopes repair to affected boundary)
- Temperature-aware budget enforcement (hot modules get fewer retries)
- Dependency-locality repair (targets the correct subsystem model)
"""

from __future__ import annotations

from typing import Any
from pydantic import BaseModel, ConfigDict, Field

from planner.schemas import PlannedTask, PlannedTaskType
from repo_indexer.models import ArchitectureGraph
from validator.context_resolver import ContextResolver, ResolvedContext
from validator.failure_record import FailureRecord
from validator.failure_types import FailureType, ResolutionStatus


class RepairPlan(BaseModel):
    model_config = ConfigDict(extra="forbid")

    failure_id: str
    original_task_id: str | None
    repair_tasks: list[PlannedTask] = Field(default_factory=list)
    blocked_reason: str | None = None
    architectural_context: dict[str, Any] | None = None


RepairPlan.model_rebuild()


class RepairPlanner:
    """Bounded repair planner with optional architecture awareness.

    When an ArchitectureGraph is provided:
    - Hot modules get reduced repair budgets
    - Blast radius is reported in the repair plan
    - Repair prompts include dependency and boundary context
    """

    def __init__(
        self,
        normal_retry_budget: int = 2,
        repair_retry_budget: int = 1,
        architecture_graph: ArchitectureGraph | None = None,
    ):
        self.normal_retry_budget = normal_retry_budget
        self.repair_retry_budget = repair_retry_budget
        self._graph = architecture_graph
        self._resolver = ContextResolver(architecture_graph) if architecture_graph else None

    def plan(
        self,
        failure: FailureRecord,
        current_repair_count: int = 0,
        affected_file: str | None = None,
    ) -> RepairPlan:
        # Resolve architectural context if available
        arch_context: ResolvedContext | None = None
        if self._resolver:
            arch_context = self._resolver.resolve_for_repair(
                failed_task_id=failure.task_id or "unknown",
                failure_type=failure.failure_type.value,
                affected_file=affected_file,
            )

        # Temperature-adjusted budget: hot modules get fewer retries
        effective_budget = self.repair_retry_budget
        if arch_context and arch_context.temperature_warnings:
            effective_budget = max(0, effective_budget - 1)

        if current_repair_count >= effective_budget:
            failure.resolution_status = ResolutionStatus.BLOCKED
            reason = "repair budget exhausted"
            if effective_budget < self.repair_retry_budget:
                reason += " (reduced: hot module)"
            return RepairPlan(
                failure_id=failure.failure_id,
                original_task_id=failure.task_id,
                blocked_reason=reason,
                architectural_context=arch_context.as_dict() if arch_context else None,
            )

        if failure.failure_type == FailureType.IMPORT_ERROR:
            return self._import_error_plan(failure, arch_context)
        if failure.failure_type == FailureType.SYNTAX_ERROR:
            return self._syntax_error_plan(failure, arch_context)
        if failure.failure_type == FailureType.HALLUCINATED_API:
            return self._hallucinated_api_plan(failure, arch_context)
        if failure.failure_type == FailureType.HALLUCINATED_SYMBOL:
            return self._hallucinated_symbol_plan(failure, arch_context)
        if failure.failure_type == FailureType.CIRCULAR_IMPORT:
            return self._circular_import_plan(failure, arch_context)

        return RepairPlan(
            failure_id=failure.failure_id,
            original_task_id=failure.task_id,
            blocked_reason=f"no repair strategy for {failure.failure_type.value}",
            architectural_context=arch_context.as_dict() if arch_context else None,
        )

    def _import_error_plan(
        self, failure: FailureRecord, ctx: ResolvedContext | None
    ) -> RepairPlan:
        module = failure.structured_context.get("module") or "missing dependency"
        goal = f"Repair import failure for module '{module}' in task {failure.task_id}"
        if ctx and ctx.imports:
            goal += f". Available dependencies: {', '.join(ctx.imports[:8])}"

        task = PlannedTask(
            id=f"repair_{failure.failure_id}_import",
            type=PlannedTaskType.DEBUGGING,
            goal=goal,
            depends_on=[],
            validation=["artifact"],
            assigned_model="deepseek-coder:6.7b",
        )
        failure.resolution_status = ResolutionStatus.REPAIR_PLANNED
        return RepairPlan(
            failure_id=failure.failure_id,
            original_task_id=failure.task_id,
            repair_tasks=[task],
            architectural_context=ctx.as_dict() if ctx else None,
        )

    def _syntax_error_plan(
        self, failure: FailureRecord, ctx: ResolvedContext | None
    ) -> RepairPlan:
        goal = (
            f"Repair syntax error for task {failure.task_id}. "
            f"Context: {failure.structured_context}"
        )
        task = PlannedTask(
            id=f"repair_{failure.failure_id}_syntax",
            type=PlannedTaskType.CODE_GENERATION,
            goal=goal,
            depends_on=[],
            validation=["artifact", "python_syntax"],
            assigned_model="deepseek-coder:6.7b",
        )
        failure.resolution_status = ResolutionStatus.REPAIR_PLANNED
        return RepairPlan(
            failure_id=failure.failure_id,
            original_task_id=failure.task_id,
            repair_tasks=[task],
            architectural_context=ctx.as_dict() if ctx else None,
        )

    def _hallucinated_api_plan(
        self, failure: FailureRecord, ctx: ResolvedContext | None
    ) -> RepairPlan:
        goal = (
            f"Fix hallucinated API usage in task {failure.task_id}. "
            f"The generated code references APIs that do not exist."
        )
        if ctx and ctx.affected_symbols:
            goal += f" Available symbols: {', '.join(ctx.affected_symbols[:10])}"

        task = PlannedTask(
            id=f"repair_{failure.failure_id}_api",
            type=PlannedTaskType.CODE_GENERATION,
            goal=goal,
            depends_on=[],
            validation=["artifact", "python_syntax"],
            assigned_model="deepseek-coder:6.7b",
        )
        failure.resolution_status = ResolutionStatus.REPAIR_PLANNED
        return RepairPlan(
            failure_id=failure.failure_id,
            original_task_id=failure.task_id,
            repair_tasks=[task],
            architectural_context=ctx.as_dict() if ctx else None,
        )

    def _hallucinated_symbol_plan(
        self, failure: FailureRecord, ctx: ResolvedContext | None
    ) -> RepairPlan:
        symbol = failure.structured_context.get("symbol", "unknown")
        goal = (
            f"Fix reference to non-existent symbol '{symbol}' in task {failure.task_id}."
        )
        if ctx and ctx.affected_symbols:
            goal += f" Available symbols in scope: {', '.join(ctx.affected_symbols[:10])}"

        task = PlannedTask(
            id=f"repair_{failure.failure_id}_symbol",
            type=PlannedTaskType.CODE_GENERATION,
            goal=goal,
            depends_on=[],
            validation=["artifact", "python_syntax"],
            assigned_model="deepseek-coder:6.7b",
        )
        failure.resolution_status = ResolutionStatus.REPAIR_PLANNED
        return RepairPlan(
            failure_id=failure.failure_id,
            original_task_id=failure.task_id,
            repair_tasks=[task],
            architectural_context=ctx.as_dict() if ctx else None,
        )

    def _circular_import_plan(
        self, failure: FailureRecord, ctx: ResolvedContext | None
    ) -> RepairPlan:
        module = failure.structured_context.get("module", "unknown")
        goal = (
            f"Resolve circular import involving '{module}' in task {failure.task_id}. "
            "Restructure imports to break the cycle."
        )
        if ctx and ctx.imports:
            goal += f" Current import chain: {' -> '.join(ctx.imports[:6])}"

        task = PlannedTask(
            id=f"repair_{failure.failure_id}_circular",
            type=PlannedTaskType.DEBUGGING,
            goal=goal,
            depends_on=[],
            validation=["artifact"],
            assigned_model="deepseek-coder:6.7b",
        )
        failure.resolution_status = ResolutionStatus.REPAIR_PLANNED
        return RepairPlan(
            failure_id=failure.failure_id,
            original_task_id=failure.task_id,
            repair_tasks=[task],
            architectural_context=ctx.as_dict() if ctx else None,
        )
