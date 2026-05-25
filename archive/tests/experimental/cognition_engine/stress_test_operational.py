import asyncio
import os
import pytest
from pathlib import Path
from orchestrator.control_plane import ControlPlane
from shared_context.execution_context import ExecutionContext

# This test requires a mock or a running model server. 
# For the purpose of the "stress test", we will use the mock-enabled ControlPlane.

@pytest.mark.anyio
async def test_tiered_stress_run(tmp_path):
    """
    Tiered stress test for the Astraeus Substrate.
    - Tier 1: Small (Direct Goal)
    - Tier 2: Medium (Context-heavy Goal)
    - Tier 3: Large (Multi-step Goal with isolation)
    """
    
    artifact_root = tmp_path / "artifacts"
    data_dir = tmp_path / "data"
    
    # Use the PlanningModelClient mock internally by providing it to ControlPlane
    # In a real scenario, this would talk to Ollama or OpenAI.
    from tests.test_core_components import PlanningModelClient
    
    control_plane = ControlPlane(
        artifact_root=artifact_root,
        data_dir=data_dir,
        model_client=PlanningModelClient()
    )
    
    goals = [
        "Small: Hello World in Python",
        "Medium: Create a complex data structure in Rust with unit tests",
        "Large: Build a full-stack web application with React and Node.js",
        "Extra-Small: Quick Fibonacci",
        "Enterprise: Architect a distributed microservices system"
    ]
    
    for i, goal in enumerate(goals):
        print(f"\n--- Stress Test {i+1}/{len(goals)}: {goal} ---")
        
        # Create unique context for each run
        context = ExecutionContext(
            workspace_id=f"stress_ws_{i}",
            event_namespace=f"stress_ns_{i}"
        )
        
        result = await control_plane.run_goal(goal, context=context)
        
        if not result["ok"]:
            print(f"\nFAILURE for goal '{goal}':")
            print(f"Metrics: {result.get('metrics')}")
            # The failure might be in tasks
            for tid, tstate in result.get("state", {}).get("tasks", {}).items():
                if tstate.get("status") == "failed":
                    print(f"Task {tid} failed: {tstate.get('failure')}")
        
        # Assertions to ensure system integrity during stress
        assert result["ok"], f"Goal '{goal}' failed execution."
        assert result["run_id"] == context.workspace_id
        
        # Verify isolation
        event_log = data_dir / "events" / context.event_namespace / context.workspace_id / "main.jsonl"
        assert event_log.exists(), f"Event log missing for {goal}"
        
        memory_db = data_dir / "memory" / "default" / f"{context.workspace_id}.sqlite3"
        assert memory_db.exists(), f"Memory DB missing for {goal}"
        
        print(f"PASS: {goal} [Artifacts: {result['artifact_dir']}]")

if __name__ == "__main__":
    # If run directly
    asyncio.run(test_tiered_stress_run(Path("./stress_test_results")))
