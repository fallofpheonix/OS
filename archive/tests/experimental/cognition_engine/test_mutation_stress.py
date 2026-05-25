import pytest
import shutil
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

def test_interleaved_rollback(setup_dirs):
    project_root, data_dir = setup_dirs
    sandbox = MutationSandbox(project_root, data_dir)
    rollback_engine = RollbackEngine(project_root, sandbox.journal)
    
    state_a = RunState(run_id="run_A", goal="task A")
    state_a.snapshots.append({"id": "dummy"})
    state_b = RunState(run_id="run_B", goal="task B")
    state_b.snapshots.append({"id": "dummy"})
    
    # 1. Run A creates file1
    plan1 = DiffPlan(patches=[FilePatch(path="file1.txt", new_text="A1")])
    sandbox.apply_mutation(plan1, state_a)
    
    # 2. Run B modifies file1
    plan2 = DiffPlan(patches=[FilePatch(path="file1.txt", old_text="A1", new_text="B2")])
    sandbox.apply_mutation(plan2, state_b)
    
    # 3. Run A modifies file1 again
    plan3 = DiffPlan(patches=[FilePatch(path="file1.txt", old_text="B2", new_text="A3")])
    sandbox.apply_mutation(plan3, state_a)
    
    assert (project_root / "file1.txt").read_text() == "A3"
    
    # 4. Rollback Run A
    # It should undo A3 -> B2.
    # But it will NOT undo A1 (CREATE) because Run B modified it (drift).
    rollback_engine.rollback_run("run_A")
    
    # Drift protection prevents deleting file1.txt because Run B's modification (B2)
    # changed the state after Run A's initial creation (A1).
    assert (project_root / "file1.txt").exists()
    assert (project_root / "file1.txt").read_text() == "B2"
    # This confirms that the current system is NOT dependency-aware. 
    # If you rollback the run that created a file, the file is GONE even if others modified it.

def test_large_transaction(setup_dirs):
    project_root, data_dir = setup_dirs
    sandbox = MutationSandbox(project_root, data_dir)
    state = RunState(run_id="large_run", goal="stress")
    state.snapshots.append({"id": "dummy"})
    
    num_files = 50
    patches = [
        FilePatch(path=f"file_{i}.txt", new_text=f"content {i}")
        for i in range(num_files)
    ]
    plan = DiffPlan(patches=patches)
    
    result = sandbox.apply_mutation(plan, state)
    assert result.ok
    
    for i in range(num_files):
        assert (project_root / f"file_{i}.txt").exists()
        
    # Rollback
    rollback_engine = RollbackEngine(project_root, sandbox.journal)
    undone = rollback_engine.rollback_run("large_run")
    assert len(undone) == num_files
    
    for i in range(num_files):
        assert not (project_root / f"file_{i}.txt").exists()

def test_validation_failure_no_commit(setup_dirs):
    project_root, data_dir = setup_dirs
    sandbox = MutationSandbox(project_root, data_dir)
    state = RunState(run_id="fail_run", goal="test failure")
    state.snapshots.append({"id": "dummy"})
    
    # Create a file that SHOULD exist
    (project_root / "exists.txt").write_text("stay")
    
    plan = DiffPlan(
        patches=[
            FilePatch(path="new.txt", new_text="should not exist"),
            FilePatch(path="exists.txt", old_text="stay", new_text="changed")
        ],
        validation_commands=["false"] # This will fail
    )
    
    result = sandbox.apply_mutation(plan, state)
    assert not result.ok
    assert result.error == "validation failed"
    
    # Verify no changes were committed
    assert not (project_root / "new.txt").exists()
    assert (project_root / "exists.txt").read_text() == "stay"
    
    # Verify journal is empty for this run
    entries = sandbox.journal.get_entries("fail_run")
    assert len(entries) == 0

def test_deep_directory_rollback(setup_dirs):
    project_root, data_dir = setup_dirs
    sandbox = MutationSandbox(project_root, data_dir)
    state = RunState(run_id="deep_run", goal="test deep")
    state.snapshots.append({"id": "dummy"})
    
    path = "a/b/c/d/e.txt"
    plan = DiffPlan(patches=[FilePatch(path=path, new_text="deep content")])
    sandbox.apply_mutation(plan, state)
    
    assert (project_root / path).exists()
    
    rollback_engine = RollbackEngine(project_root, sandbox.journal)
    rollback_engine.rollback_run("deep_run")
    
    assert not (project_root / path).exists()
    # Note: directories a/b/c/d/ remain. The system only deletes the file.
    assert (project_root / "a/b/c/d").exists()
