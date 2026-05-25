import pytest
from repo_indexer.models import ArchitectureGraph, ArchitecturalNode, FunctionDefinition, Parameter, ClassDefinition
from validator.semantic_verifier import SemanticVerifier

@pytest.fixture
def verifier():
    return SemanticVerifier()

def test_verify_signature_change(verifier):
    old_graph = ArchitectureGraph()
    old_node = ArchitecturalNode(
        module_path="math.py",
        public_api=["add"],
        top_level_functions=[
            FunctionDefinition(name="add", parameters=[Parameter("a"), Parameter("b")])
        ]
    )
    old_graph.add_node(old_node)
    
    new_graph = ArchitectureGraph()
    new_node = ArchitecturalNode(
        module_path="math.py",
        public_api=["add"],
        top_level_functions=[
            # added parameter c
            FunctionDefinition(name="add", parameters=[Parameter("a"), Parameter("b"), Parameter("c")])
        ]
    )
    new_graph.add_node(new_node)
    
    report = verifier.verify(old_graph, new_graph)
    # Signature changes are warnings, so report.passed should be True (unless strict mode enabled)
    assert report.passed
    assert any(r.regression_type == "api_modification" for r in report.regressions)
    assert any("params added" in r.message for r in report.regressions)

def test_verify_class_modification(verifier):
    old_graph = ArchitectureGraph()
    old_node = ArchitecturalNode(
        module_path="db.py",
        public_api=["Session"],
        class_definitions=[
            ClassDefinition(name="Session", methods=[FunctionDefinition(name="connect")])
        ]
    )
    old_graph.add_node(old_node)
    
    new_graph = ArchitectureGraph()
    new_node = ArchitecturalNode(
        module_path="db.py",
        public_api=["Session"],
        class_definitions=[
            ClassDefinition(name="Session", methods=[FunctionDefinition(name="close")]) # connect removed, close added
        ]
    )
    new_graph.add_node(new_node)
    
    report = verifier.verify(old_graph, new_graph)
    assert report.passed
    assert any("methods removed: {'connect'}" in r.message for r in report.regressions)
    assert any("methods added: {'close'}" in r.message for r in report.regressions)

def test_verify_api_removal(verifier):
    old_graph = ArchitectureGraph()
    old_node = ArchitecturalNode(module_path="a.py", public_api=["foo"])
    old_graph.add_node(old_node)
    
    new_graph = ArchitectureGraph()
    new_node = ArchitecturalNode(module_path="a.py", public_api=[]) # foo removed
    old_node.top_level_functions = [FunctionDefinition(name="foo")]
    # Actually need to provide the function in new_node to test removal properly if using diff logic
    
    # Simpler removal test
    old_graph = ArchitectureGraph()
    old_graph.add_node(ArchitecturalNode(module_path="a.py", public_api=["foo"], top_level_functions=[FunctionDefinition(name="foo")]))
    new_graph = ArchitectureGraph()
    new_graph.add_node(ArchitecturalNode(module_path="a.py", public_api=[], top_level_functions=[]))
    
    report = verifier.verify(old_graph, new_graph)
    assert not report.passed
    assert any(r.regression_type == "api_removal" for r in report.regressions)
