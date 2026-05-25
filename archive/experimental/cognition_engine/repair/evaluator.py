"""Evaluates the outcome of repair attempts."""

from __future__ import annotations

from enum import Enum
from typing import Any
from pydantic import BaseModel, ConfigDict, Field
from shared_context.state import RunState, TaskStatus


class RepairStatus(str, Enum):
    SUCCESS = "success"
    PARTIAL = "partial"
    REGRESSED = "regressed"
    FAILED = "failed"


class RepairOutcome(BaseModel):
    model_config = ConfigDict(extra="forbid")

    status: RepairStatus
    fixed_task: bool
    downstream_failures: list[str] = Field(default_factory=list)
    notes: str = ""


class RepairEvaluator:
    def __init__(self, failure_memory: Any = None) -> None:
        self.failure_memory = failure_memory

    def evaluate(self, state: RunState, failure_task_id: str, original_failure_type: str) -> RepairOutcome:
        task_record = state.tasks.get(failure_task_id)
        if not task_record:
            return RepairOutcome(status=RepairStatus.FAILED, fixed_task=False, notes="Task not found")

        fixed_task = task_record.status == TaskStatus.SUCCEEDED
        
        downstream_failures = []
        for t_id, record in state.tasks.items():
            if t_id != failure_task_id and record.status == TaskStatus.FAILED:
                # Naive downstream check - assume any failed task after repair is a regression
                # Better would be to use DAG to check true downstream dependencies
                if failure_task_id in _get_all_dependencies(state, t_id):
                    downstream_failures.append(t_id)

        if fixed_task and not downstream_failures:
            status = RepairStatus.SUCCESS
        elif fixed_task and downstream_failures:
            status = RepairStatus.REGRESSED
        elif not fixed_task and task_record.failure and original_failure_type not in task_record.failure:
            status = RepairStatus.PARTIAL
        else:
            status = RepairStatus.FAILED

        outcome = RepairOutcome(status=status, fixed_task=fixed_task, downstream_failures=downstream_failures)

        if self.failure_memory:
            self.failure_memory.record_repair_outcome(
                failure_id=failure_task_id,
                outcome_status=status.value,
                notes=outcome.notes
            )

        return outcome

def _get_all_dependencies(state: RunState, task_id: str) -> set[str]:
    deps = set()
    stack = list(state.tasks[task_id].task.depends_on)
    while stack:
        curr = stack.pop()
        if curr not in deps:
            deps.add(curr)
            if curr in state.tasks:
                stack.extend(state.tasks[curr].task.depends_on)
    return deps
