"""
Surface Agent - Orchestrator

Purpose: coordinator-only orchestrator for project workflows.

Responsibilities implemented:
- task routing (with handler registration)
- dependency graph (simple DAG + topo ordering)
- state tracking (per-task + workflow)
- risk control (register/mitigate risks)
- workflow execution (state machine transitions)
- validation gates (pluggable checks)

Constraints: This agent is strictly orchestration-only. Any task marked
with a forbidden type will be rejected. The agent will never collect
telemetry, run kernel hooks, compute graphs, or perform forensics.
"""
from __future__ import annotations

import threading
from dataclasses import dataclass, field
from enum import Enum
from typing import Any, Callable, Dict, List, Optional, Set


class NotAllowedError(RuntimeError):
    pass


class State(str, Enum):
    IDEA = "IDEA"
    RFC = "RFC"
    RESEARCH = "RESEARCH"
    EXPERIMENT = "EXPERIMENT"
    BUILD = "BUILD"
    TEST = "TEST"
    DEBUG = "DEBUG"
    VALIDATE = "VALIDATE"
    MERGE = "MERGE"


FORBIDDEN_TASK_TYPES: Set[str] = {
    "telemetry_collection",
    "kernel_hooks",
    "compute_graphs",
    "forensics",
}


@dataclass
class Task:
    id: str
    title: str
    task_type: str
    metadata: Dict[str, Any] = field(default_factory=dict)
    state: State = State.IDEA
    depends_on: Set[str] = field(default_factory=set)


@dataclass
class Risk:
    id: str
    description: str
    severity: int  # 1-10
    mitigated: bool = False


@dataclass
class Workflow:
    id: str
    tasks: Dict[str, Task] = field(default_factory=dict)
    risks: Dict[str, Risk] = field(default_factory=dict)
    validation_gates: Dict[State, List[Callable[[Workflow], bool]]] = field(
        default_factory=dict
    )


class SurfaceOrchestrator:
    """Core orchestrator for Layer-2 Surface Agent.

    Usage:
      - create workflow, add tasks and dependencies
      - register handlers for task types (pure routing; handlers perform work)
      - start workflow execution; orchestrator advances task states through gates
    """

    def __init__(self):
        self.workflows: Dict[str, Workflow] = {}
        self.handlers: Dict[str, Callable[[Task], Any]] = {}
        self.lock = threading.RLock()

    # --- registration / creation ---
    def create_workflow(self, wf_id: str) -> Workflow:
        with self.lock:
            if wf_id in self.workflows:
                return self.workflows[wf_id]
            wf = Workflow(id=wf_id)
            self.workflows[wf_id] = wf
            return wf

    def add_task(self, wf_id: str, task: Task) -> None:
        if task.task_type in FORBIDDEN_TASK_TYPES:
            raise NotAllowedError(f"task type '{task.task_type}' is forbidden")
        with self.lock:
            wf = self.workflows.setdefault(wf_id, Workflow(id=wf_id))
            if task.id in wf.tasks:
                raise KeyError("task already exists")
            wf.tasks[task.id] = task

    def add_dependency(self, wf_id: str, task_id: str, depends_on_id: str) -> None:
        with self.lock:
            wf = self._get_wf(wf_id)
            if task_id not in wf.tasks or depends_on_id not in wf.tasks:
                raise KeyError("one of the tasks is unknown")
            wf.tasks[task_id].depends_on.add(depends_on_id)

    def register_handler(self, task_type: str, handler: Callable[[Task], Any]) -> None:
        if task_type in FORBIDDEN_TASK_TYPES:
            raise NotAllowedError(f"cannot register handler for forbidden type {task_type}")
        self.handlers[task_type] = handler

    # --- risk control ---
    def add_risk(self, wf_id: str, risk: Risk) -> None:
        with self.lock:
            wf = self._get_wf(wf_id)
            wf.risks[risk.id] = risk

    def mitigate_risk(self, wf_id: str, risk_id: str) -> None:
        with self.lock:
            wf = self._get_wf(wf_id)
            if risk_id in wf.risks:
                wf.risks[risk_id].mitigated = True

    # --- validation gates ---
    def add_validation_gate(self, wf_id: str, state: State, check: Callable[[Workflow], bool]) -> None:
        with self.lock:
            wf = self._get_wf(wf_id)
            wf.validation_gates.setdefault(state, []).append(check)

    def _run_gates(self, wf: Workflow, state: State) -> bool:
        checks = wf.validation_gates.get(state, [])
        for chk in checks:
            if not chk(wf):
                return False
        return True

    # --- execution / routing ---
    def execute_workflow(self, wf_id: str) -> None:
        """Drive tasks through the defined state machine in dependency-respecting order."""
        with self.lock:
            wf = self._get_wf(wf_id)

        order = self._topological_order(list(wf.tasks.values()))
        for task in order:
            # advance through state machine for each task
            for state in [
                State.IDEA,
                State.RFC,
                State.RESEARCH,
                State.EXPERIMENT,
                State.BUILD,
                State.TEST,
                State.DEBUG,
                State.VALIDATE,
                State.MERGE,
            ]:
                # run gates for this state
                with self.lock:
                    wf = self._get_wf(wf_id)
                if not self._run_gates(wf, state):
                    # gate blocked; stop advancing this task
                    break
                # route task to handler if one is registered for the task_type
                handler = self.handlers.get(task.task_type)
                if handler:
                    # handler must not perform forbidden ops; orchestrator only routes
                    handler(task)
                task.state = state

    def route_task(self, wf_id: str, task_id: str) -> None:
        with self.lock:
            wf = self._get_wf(wf_id)
            task = wf.tasks[task_id]
            if task.task_type in FORBIDDEN_TASK_TYPES:
                raise NotAllowedError("forbidden task type")
            handler = self.handlers.get(task.task_type)
            if not handler:
                raise KeyError("no handler registered for task type")
        handler(task)

    # --- utilities ---
    def _get_wf(self, wf_id: str) -> Workflow:
        if wf_id not in self.workflows:
            raise KeyError("unknown workflow")
        return self.workflows[wf_id]

    def _topological_order(self, tasks: List[Task]) -> List[Task]:
        # Kahn's algorithm
        id_map = {t.id: t for t in tasks}
        indeg: Dict[str, int] = {t.id: 0 for t in tasks}
        out: Dict[str, Set[str]] = {t.id: set() for t in tasks}
        for t in tasks:
            for d in t.depends_on:
                if d in id_map:
                    indeg[t.id] += 1
                    out[d].add(t.id)
        queue = [tid for tid, deg in indeg.items() if deg == 0]
        ordered: List[Task] = []
        while queue:
            n = queue.pop(0)
            ordered.append(id_map[n])
            for m in list(out[n]):
                indeg[m] -= 1
                if indeg[m] == 0:
                    queue.append(m)
        if len(ordered) != len(tasks):
            raise RuntimeError("dependency cycle detected")
        return ordered


__all__ = ["SurfaceOrchestrator", "Task", "Workflow", "State", "Risk", "NotAllowedError"]
