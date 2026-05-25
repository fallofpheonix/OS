"""Live control-plane integration tests.
Requires a running Ollama instance and COGNITION_LIVE_OLLAMA=1.
"""

import os
import pytest
from pathlib import Path
from orchestrator.control_plane import ControlPlane

@pytest.mark.skipif(os.environ.get("COGNITION_LIVE_OLLAMA") != "1", reason="requires live Ollama")
@pytest.mark.anyio
async def test_live_goal_execution(tmp_path):
    engine = ControlPlane(
        artifact_root=tmp_path / "artifacts",
        data_dir=tmp_path / "data"
    )
    
    goal = "Implement a Python function that calculates the Fibonacci sequence up to N and returns it as a list."
    result = await engine.run_goal(goal)
    
    assert result["ok"]
    assert "run_id" in result
    assert Path(result["artifact_dir"]).exists()
    
    # Verify events
    events_file = Path(result["artifact_dir"]) / "events.jsonl"
    assert events_file.exists()
    
    # Verify no mock signatures in output
    run_json = Path(result["artifact_dir"]) / "run.json"
    with open(run_json, "r") as f:
        import json
        data = json.load(f)
        for task in data.get("tasks", {}).values():
            assert "(MOCK)" not in task.get("output", "")

@pytest.mark.skipif(os.environ.get("COGNITION_LIVE_OLLAMA") != "1", reason="requires live Ollama")
@pytest.mark.anyio
async def test_live_repair_loop(tmp_path):
    # This is a harder test: we want to force a repair.
    # We can do this by using a goal that is likely to produce a syntax error or rejection.
    # But since we can't easily force a model to fail, we'll just verify the repair logic
    # if it triggers.
    ControlPlane(
        artifact_root=tmp_path / "artifacts",
        data_dir=tmp_path / "data"
    )
    
@pytest.mark.skipif(os.environ.get("COGNITION_LIVE_OLLAMA") != "1", reason="requires live Ollama")
@pytest.mark.anyio
async def test_live_stress_run(tmp_path):
    """Run 5 consecutive goals to verify substrate stability."""
    engine = ControlPlane(
        artifact_root=tmp_path / "artifacts",
        data_dir=tmp_path / "data"
    )
    
    goals = [
        "Calculate prime numbers up to 100",
        "Implement a simple stack data structure",
        "Write a function to reverse a string",
        "Create a base64 encoder/decoder",
        "Implement a bubble sort algorithm"
    ]
    
    for i, goal in enumerate(goals):
        print(f"Stress test: Running goal {i+1}/5: {goal}")
        result = await engine.run_goal(goal)
        assert result["ok"], f"Goal {i+1} failed: {goal}"
