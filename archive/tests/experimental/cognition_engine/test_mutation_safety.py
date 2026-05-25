import pytest
from transactions.journal import FilesystemJournal, JournalEntry
from transactions.rollback import RollbackEngine
from runtime.risk_engine import CommandRiskEngine, RiskLevel
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

def test_journal_and_rollback_basic(setup_dirs):
    project_root, data_dir = setup_dirs
    journal = FilesystemJournal(data_dir / "journal.jsonl", data_dir / "backups")
    rollback_engine = RollbackEngine(project_root, journal)
    
    # 1. Create file1
    path1 = "file1.txt"
    (project_root / path1).write_text("initial")
    journal.record(JournalEntry(
        run_id="run1",
        operation="CREATE",
        path=path1,
        new_hash=FilesystemJournal.compute_hash(project_root / path1)
    ))
    
    # 2. Modify file1
    old_content = (project_root / path1).read_bytes()
    old_hash = FilesystemJournal.compute_hash(project_root / path1)
    (project_root / path1).write_text("modified")
    journal.record(JournalEntry(
        run_id="run1",
        operation="MODIFY",
        path=path1,
        old_hash=old_hash,
        new_hash=FilesystemJournal.compute_hash(project_root / path1)
    ), original_content=old_content)
    
    # 3. Delete file1
    old_content = (project_root / path1).read_bytes()
    old_hash = FilesystemJournal.compute_hash(project_root / path1)
    (project_root / path1).unlink()
    journal.record(JournalEntry(
        run_id="run1",
        operation="DELETE",
        path=path1,
        old_hash=old_hash
    ), original_content=old_content)
    
    assert not (project_root / path1).exists()
    
    # Rollback run1
    undone = rollback_engine.rollback_run("run1")
    assert path1 in undone
    
    # After rolling back DELETE, MODIFY, and CREATE, the file should be gone.
    # Step 1: Undo DELETE -> Restore "modified"
    # Step 2: Undo MODIFY -> Restore "initial"
    # Step 3: Undo CREATE -> Delete file
    assert not (project_root / path1).exists()

def test_rollback_partial_run(setup_dirs):
    project_root, data_dir = setup_dirs
    journal = FilesystemJournal(data_dir / "journal.jsonl", data_dir / "backups")
    rollback_engine = RollbackEngine(project_root, journal)
    
    # Run 1 creates file1
    (project_root / "file1.txt").write_text("run1")
    journal.record(JournalEntry(run_id="run1", operation="CREATE", path="file1.txt"))
    
    # Run 2 modifies file1
    old_content = (project_root / "file1.txt").read_bytes()
    (project_root / "file1.txt").write_text("run2")
    journal.record(JournalEntry(run_id="run2", operation="MODIFY", path="file1.txt", old_hash="h1"), original_content=old_content)
    
    # Rollback Run 2 only
    undone = rollback_engine.rollback_run("run2")
    assert "file1.txt" in undone
    assert (project_root / "file1.txt").read_text() == "run1"
    
    # file1.txt still exists because Run 1 CREATE was not undone
    assert (project_root / "file1.txt").exists()

def test_risk_engine_classification():
    engine = CommandRiskEngine()
    
    assert engine.evaluate("ls").level == RiskLevel.SAFE
    assert engine.evaluate("ls -la").level == RiskLevel.SAFE
    assert engine.evaluate("rm -rf /tmp/foo").level == RiskLevel.CRITICAL
    assert engine.evaluate("rm -rf /").level == RiskLevel.CRITICAL
    assert engine.evaluate("git commit -m 'fix'").level == RiskLevel.MODERATE
    assert engine.evaluate("curl http://api.com").level == RiskLevel.HIGH
    assert engine.evaluate("pip install requests").level == RiskLevel.MODERATE
    assert engine.evaluate("python -c 'import os; os.system(\"rm -rf /\")'").level == RiskLevel.CRITICAL # caught by destructive pattern

from tools.permissions import ApprovalRequired

def test_mutation_sandbox_safety(setup_dirs):
    project_root, data_dir = setup_dirs
    sandbox = MutationSandbox(project_root, data_dir)
    state = RunState(run_id="test_run", goal="test")
    state.snapshots.append({"id": "dummy"})
    
    # 1. Unsafe path (outside project root)
    plan = DiffPlan(
        patches=[FilePatch(path="../outside.txt", new_text="bad")],
        validation_commands=["ls"]
    )
    with pytest.raises(PermissionError, match="outside project root"):
        sandbox.apply_mutation(plan, state)
        
    # 2. Unsafe validation command (HIGH risk)
    plan = DiffPlan(
        patches=[FilePatch(path="safe.txt", new_text="good")],
        validation_commands=["curl http://attacker.com"]
    )
    with pytest.raises(ApprovalRequired, match="approval required"):
        sandbox.apply_mutation(plan, state)

def test_mutation_sandbox_apply(setup_dirs):
    project_root, data_dir = setup_dirs
    sandbox = MutationSandbox(project_root, data_dir)
    state = RunState(run_id="run_apply", goal="test")
    state.snapshots.append({"id": "dummy"})
    
    # Create a file via sandbox
    plan = DiffPlan(
        patches=[FilePatch(path="new.txt", new_text="hello world")],
        validation_commands=["grep 'hello' new.txt"]
    )
    result = sandbox.apply_mutation(plan, state)
    assert result.ok
    assert (project_root / "new.txt").read_text() == "hello world"
    
    # Verify journaled
    entries = sandbox.journal.get_entries("run_apply")
    assert len(entries) == 1
    assert entries[0].operation == "CREATE"
    assert entries[0].path == "new.txt"
