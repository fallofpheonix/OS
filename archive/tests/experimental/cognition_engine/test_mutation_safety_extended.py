import pytest
import json
import hashlib
from dataclasses import asdict
from pathlib import Path
from transactions.journal import FilesystemJournal, JournalEntry
from transactions.rollback import RollbackEngine
from runtime.mutation_sandbox import MutationSandbox
from transactions.diff_plan import DiffPlan, FilePatch
from shared_context import RunState

@pytest.fixture
def setup_dirs(tmp_path):
    project_root = tmp_path / "project"
    project_root.mkdir()
    data_dir = tmp_path / "data"
    data_dir.mkdir()
    return project_root, data_dir

def test_journal_integrity_verification(setup_dirs):
    project_root, data_dir = setup_dirs
    journal_path = data_dir / "journal.jsonl"
    journal = FilesystemJournal(journal_path)
    
    # Add some valid entries
    journal.record(JournalEntry(run_id="run1", operation="CREATE", path="f1.txt"))
    journal.record(JournalEntry(run_id="run1", operation="CREATE", path="f2.txt"))
    
    assert journal.verify_integrity() is True
    
    # Tamper with the journal: change a path in the middle
    lines = journal_path.read_text().splitlines()
    data = json.loads(lines[0])
    data["path"] = "tampered.txt"
    lines[0] = json.dumps(data)
    journal_path.write_text("\n".join(lines) + "\n")
    
    # The hash chain should now be broken because the first line's hash changed, 
    # but the second line's previous_hash still points to the old hash of the first line.
    assert journal.verify_integrity() is False

def test_journal_corruption_handling(setup_dirs):
    project_root, data_dir = setup_dirs
    journal_path = data_dir / "journal.jsonl"
    journal = FilesystemJournal(journal_path)
    
    journal.record(JournalEntry(run_id="run1", operation="CREATE", path="f1.txt"))
    
    # Append a corrupt line (invalid JSON)
    with open(journal_path, "a") as f:
        f.write("this is not json\n")
    
    journal.record(JournalEntry(run_id="run1", operation="CREATE", path="f2.txt"))
    
    # get_entries should skip the corrupt line
    entries = journal.get_entries("run1")
    assert len(entries) == 2
    assert entries[0].path == "f1.txt"
    assert entries[1].path == "f2.txt"

def test_rollback_partial_drift(setup_dirs):
    project_root, data_dir = setup_dirs
    journal = FilesystemJournal(data_dir / "journal.jsonl", data_dir / "backups")
    rollback_engine = RollbackEngine(project_root, journal)
    
    # Run 1 creates two files
    (project_root / "f1.txt").write_text("orig1")
    (project_root / "f2.txt").write_text("orig2")
    
    journal.record(JournalEntry(
        run_id="run1", 
        operation="CREATE", 
        path="f1.txt", 
        new_hash=FilesystemJournal.compute_hash(project_root / "f1.txt")
    ))
    journal.record(JournalEntry(
        run_id="run1", 
        operation="CREATE", 
        path="f2.txt", 
        new_hash=FilesystemJournal.compute_hash(project_root / "f2.txt")
    ))
    
    # Drift f1.txt (modify it outside the system)
    (project_root / "f1.txt").write_text("drifted")
    
    # Rollback run1
    undone = rollback_engine.rollback_run("run1")
    
    # f1.txt should NOT be undone because of drift
    assert "f1.txt" not in undone
    assert (project_root / "f1.txt").exists()
    assert (project_root / "f1.txt").read_text() == "drifted"
    
    # f2.txt should be undone (deleted)
    assert "f2.txt" in undone
    assert not (project_root / "f2.txt").exists()

def test_path_traversal_protection_detailed(setup_dirs):
    project_root, data_dir = setup_dirs
    sandbox = MutationSandbox(project_root, data_dir)
    state = RunState(run_id="test", goal="test")
    state.snapshots.append({"id": "dummy"})
    
    # Test various traversal patterns
    bad_paths = [
        "../outside.txt",
        "foo/../../bar.txt",
        "/tmp/absolute.txt",
    ]
    
    for path in bad_paths:
        plan = DiffPlan(patches=[FilePatch(path=path, new_text="bad")])
        with pytest.raises(PermissionError) as excinfo:
            sandbox.apply_mutation(plan, state)
        assert any(msg in str(excinfo.value) for msg in [
            "outside project root",
            "unsafe patch path",
            "is forbidden",
            "Absolute or home-relative path"
        ])
def test_rollback_missing_backup_file(setup_dirs):
    project_root, data_dir = setup_dirs
    journal = FilesystemJournal(data_dir / "journal.jsonl", data_dir / "backups")
    rollback_engine = RollbackEngine(project_root, journal)
    
    # Create and then Modify a file
    (project_root / "f1.txt").write_text("initial")
    h_initial = FilesystemJournal.compute_hash(project_root / "f1.txt")
    
    old_content = (project_root / "f1.txt").read_bytes()
    (project_root / "f1.txt").write_text("modified")
    h_modified = FilesystemJournal.compute_hash(project_root / "f1.txt")
    
    entry = JournalEntry(
        run_id="run1",
        operation="MODIFY",
        path="f1.txt",
        old_hash=h_initial,
        new_hash=h_modified
    )
    journal.record(entry, original_content=old_content)
    
    # Manually delete the backup file
    backup_path = journal.get_backup_path(entry)
    assert backup_path.exists()
    backup_path.unlink()
    
    # Rollback should fail gracefully for this entry
    undone = rollback_engine.rollback_run("run1")
    assert "f1.txt" not in undone
    assert (project_root / "f1.txt").read_text() == "modified"

def test_integrity_chain_broken_by_insertion(setup_dirs):
    project_root, data_dir = setup_dirs
    journal_path = data_dir / "journal.jsonl"
    journal = FilesystemJournal(journal_path)
    
    journal.record(JournalEntry(run_id="run1", operation="CREATE", path="f1.txt"))
    journal.record(JournalEntry(run_id="run1", operation="CREATE", path="f3.txt"))
    
    lines = journal_path.read_text().splitlines()
    
    # Insert a fake entry in the middle
    fake_entry = JournalEntry(run_id="run1", operation="DELETE", path="f2.txt", previous_hash="bogus")
    fake_line = json.dumps(asdict(fake_entry))
    
    new_lines = [lines[0], fake_line, lines[1]]
    journal_path.write_text("\n".join(new_lines) + "\n")
    
    # Integrity check should fail because f3.txt's previous_hash won't match the hash of the fake line,
    # and the fake line's previous_hash doesn't match the hash of f1.txt.
    assert journal.verify_integrity() is False

from repo_indexer.semantic import SemanticTopologyEngine
from orchestrator.control_plane import ControlPlane

def test_topology_refresh_after_mutation(setup_dirs):
    project_root, data_dir = setup_dirs
    engine = ControlPlane(artifact_root=data_dir/"artifacts", data_dir=data_dir)
    
    # Create a dummy python file
    (project_root / "module1.py").write_text("def foo(): pass")
    
    # Manually initialize engine components like run_goal would
    engine.topology_engine = SemanticTopologyEngine(project_root)
    engine.graph = engine.topology_engine.build_graph()
    engine.mutation_sandbox = MutationSandbox(
        project_root=project_root, 
        data_dir=data_dir, 
        approval_manager=engine.approvals
    )
    
    assert "module1.py" in engine.graph.nodes
    assert "module2.py" not in engine.graph.nodes
    
    state = RunState(run_id="test_topo", goal="test")
    state.snapshots.append({"id": "dummy"})
    
    plan = DiffPlan(patches=[FilePatch(path="module2.py", new_text="def bar(): pass")])
    engine.apply_repo_mutation(plan, state)
    
    # Verify graph refreshed automatically because a .py file was added
    assert "module2.py" in engine.graph.nodes
    assert "bar" in engine.graph.nodes["module2.py"].public_api
