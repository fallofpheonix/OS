import pytest
from pathlib import Path
from repo_indexer.models import ArchitectureGraph, ArchitecturalNode, FunctionDefinition, Parameter
from contracts.invariant_engine import InvariantEngine, InvariantDefinition

def test_async_handler_missing_timeout_violation():
    graph = ArchitectureGraph()
    
    # Non-compliant handler (missing timeout)
    graph.add_node(ArchitecturalNode(
        module_path="agents/task_handler.py",
        top_level_functions=[
            FunctionDefinition(
                name="execute_task_handler",
                parameters=[Parameter(name="task")],
                is_async=True
            )
        ]
    ))
    
    engine = InvariantEngine()
    engine.invariants.append(InvariantDefinition(
        id="async_handlers_must_timeout",
        category="behavioral",
        raw={
            "id": "async_handlers_must_timeout",
            "pattern": "async def.*handler",
            "required_params": ["timeout"]
        }
    ))
    
    report = engine.evaluate(graph)
    assert not report.passed
    assert report.violation_count == 1
    assert "missing required parameter: timeout" in report.results[0].violations[0].message

def test_async_handler_with_timeout_passes():
    graph = ArchitectureGraph()
    
    # Compliant handler
    graph.add_node(ArchitecturalNode(
        module_path="agents/good_handler.py",
        top_level_functions=[
            FunctionDefinition(
                name="process_handler",
                parameters=[Parameter(name="task"), Parameter(name="timeout")],
                is_async=True
            )
        ]
    ))
    
    engine = InvariantEngine()
    engine.invariants.append(InvariantDefinition(
        id="async_handlers_must_timeout",
        category="behavioral",
        raw={
            "id": "async_handlers_must_timeout",
            "pattern": "async def.*handler",
            "required_params": ["timeout"]
        }
    ))
    
    report = engine.evaluate(graph)
    assert report.passed
