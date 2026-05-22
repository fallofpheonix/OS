"""Demo runner for Surface Orchestrator service.

Creates a workflow, registers simple handlers, persists state, and executes.
"""
from __future__ import annotations

from agents.surface.orchestrator.service import OrchestratorService
from agents.surface.orchestrator.orchestrator import Task


def simple_logger_handler(task: Task):
    print(f"[handler] Task {task.id} ({task.title}) at state {task.state}")


def demo():
    svc = OrchestratorService()
    svc.create_workflow("demo-pheonix")

    svc.register_handler("documentation", simple_logger_handler)
    svc.register_handler("build", simple_logger_handler)
    svc.register_handler("test", simple_logger_handler)

    svc.add_task("demo-pheonix", Task(id="doc-1", title="Write RFC", task_type="documentation"))
    svc.add_task("demo-pheonix", Task(id="build-1", title="Build Component", task_type="build"))
    svc.add_task("demo-pheonix", Task(id="test-1", title="Run Unit Tests", task_type="test"))

    svc.add_dependency("demo-pheonix", "build-1", "doc-1")
    svc.add_dependency("demo-pheonix", "test-1", "build-1")

    print("Starting async execution...")
    t = svc.execute_workflow_async("demo-pheonix")
    # wait for background thread to finish (simple join)
    t.join(timeout=30)
    print("Execution finished. Persisted state in sqlite DB at:", svc.db_path)


if __name__ == "__main__":
    demo()
