import pytest
import shutil
import json
import asyncio
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

def test_rollback_conflict_interleaved(setup_dirs):
    project_root, data_dir = setup_dirs
    sandbox = MutationSandbox(project_root, data_dir)
    rollback_engine = RollbackEngine(project_root, sandbox.journal)
    
    state_a = RunState(run_id="run_A", goal="task A")
    state_a.snapshots.append({"id": "dummy"})
    state_b = RunState(run_id="run_B", goal="task B")
    state_b.snapshots.append({"id": "dummy"})

    
    # 1. Run A creates file
    plan1 = DiffPlan(patches=[FilePatch(path="conflict.txt", new_text="A1")])
    sandbox.apply_mutation(plan1, state_a)
    
    # 2. Run B modifies file
    plan2 = DiffPlan(patches=[FilePatch(path="conflict.txt", old_text="A1", new_text="B2")])
    sandbox.apply_mutation(plan2, state_b)
    
    # 3. Rollback Run A (Should FAIL to rollback conflict.txt because hash changed)
    undone = rollback_engine.rollback_run("run_A")
    
    # The file should still exist and equal "B2"
    assert (project_root / "conflict.txt").read_text() == "B2"
    assert "conflict.txt" not in undone

def test_journal_corruption_detection(setup_dirs):
    project_root, data_dir = setup_dirs
    sandbox = MutationSandbox(project_root, data_dir)
    state = RunState(run_id="run_corr", goal="task")
    state.snapshots.append({"id": "dummy"})

    
    plan1 = DiffPlan(patches=[FilePatch(path="file.txt", new_text="data")])
    sandbox.apply_mutation(plan1, state)
    
    assert sandbox.journal.verify_integrity() is True
    
    # Corrupt the journal by appending a fake line with a bad previous_hash
    with open(sandbox.journal.journal_path, "a") as f:
        fake_entry = {"run_id": "fake", "operation": "CREATE", "path": "x.txt", "previous_hash": "bad"}
        f.write(json.dumps(fake_entry) + "\n")
        
    assert sandbox.journal.verify_integrity() is False

def test_force_rollback(setup_dirs):
    project_root, data_dir = setup_dirs
    sandbox = MutationSandbox(project_root, data_dir)
    rollback_engine = RollbackEngine(project_root, sandbox.journal)
    
    state_a = RunState(run_id="run_A", goal="task A")
    state_a.snapshots.append({"id": "dummy"})
    plan1 = DiffPlan(patches=[FilePatch(path="force.txt", new_text="A1")])
    sandbox.apply_mutation(plan1, state_a)
    
    # Manual tamper
    (project_root / "force.txt").write_text("tampered")
    
    # Normal rollback fails
    undone = rollback_engine.rollback_run("run_A")
    assert "force.txt" not in undone
    assert (project_root / "force.txt").read_text() == "tampered"
    
    # Force rollback succeeds
    undone_force = rollback_engine.rollback_run("run_A", force=True)
    assert "force.txt" in undone_force
    assert not (project_root / "force.txt").exists()

@pytest.mark.anyio
async def test_concurrent_journal_appends(setup_dirs):
    project_root, data_dir = setup_dirs
    journal = FilesystemJournal(data_dir / "journal.jsonl", data_dir / "backups")
    
    async def write_entries(run_id: str, count: int):
        for i in range(count):
            journal.record(JournalEntry(run_id=run_id, operation="CREATE", path=f"{run_id}_{i}.txt"))
            await asyncio.sleep(0.001)

    # Run two coroutines that write to the journal simultaneously
    await asyncio.gather(
        write_entries("run_1", 20),
        write_entries("run_2", 20)
    )
    
    # Verify the journal has 40 entries and integrity is maintained
    entries = journal.get_entries("run_1") + journal.get_entries("run_2")
    assert len(entries) == 40
    assert journal.verify_integrity() is True
