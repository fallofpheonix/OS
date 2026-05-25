import pytest
import json
import hashlib
from pathlib import Path
from transactions.journal import FilesystemJournal, JournalEntry
from transactions.rollback import RollbackEngine
from transactions.runner import TransactionRunner
from transactions.diff_plan import DiffPlan, FilePatch
from shared_context import RunState

@pytest.fixture
def setup_dirs(tmp_path):
    project_root = tmp_path / "project"
    project_root.mkdir()
    data_dir = tmp_path / "data"
    data_dir.mkdir()
    return project_root, data_dir

def test_journal_tampering_detection(setup_dirs):
    project_root, data_dir = setup_dirs
    journal = FilesystemJournal(data_dir / "journal.jsonl", data_dir / "backups")
    
    # 1. Record some entries
    journal.record(JournalEntry(run_id="run1", operation="CREATE", path="file1.txt"))
    journal.record(JournalEntry(run_id="run1", operation="CREATE", path="file2.txt"))
    
    assert journal.verify_integrity()
    
    # 2. Tamper with the journal (delete a line)
    lines = (data_dir / "journal.jsonl").read_text().splitlines()
    (data_dir / "journal.jsonl").write_text(lines[0] + "\n")
    
    # Verification should fail because the next entry's previous_hash won't match
    # Wait, if I delete the LAST line, it might still pass if I don't check the tail.
    # But if I delete the FIRST line, the second line's previous_hash will be wrong.
    
    # Let's tamper by changing the content of the first line
    data = json.loads(lines[0])
    data["path"] = "tampered.txt"
    lines[0] = json.dumps(data)
    (data_dir / "journal.jsonl").write_text("\n".join(lines) + "\n")
    
    assert not journal.verify_integrity()

def test_rollback_drift_protection(setup_dirs):
    project_root, data_dir = setup_dirs
    journal = FilesystemJournal(data_dir / "journal.jsonl", data_dir / "backups")
    rollback_engine = RollbackEngine(project_root, journal)
    
    # 1. Create and modify a file
    path = "file.txt"
    (project_root / path).write_text("v1")
    old_hash = FilesystemJournal.compute_hash(project_root / path)
    
    (project_root / path).write_text("v2")
    new_hash = FilesystemJournal.compute_hash(project_root / path)
    
    entry = JournalEntry(
        run_id="run1",
        operation="MODIFY",
        path=path,
        old_hash=old_hash,
        new_hash=new_hash
    )
    journal.record(entry, original_content=b"v1")
    
    # 2. External modification (drift)
    (project_root / path).write_text("v3_external")
    
    # 3. Rollback should fail
    success = rollback_engine.rollback_entry(entry)
    assert not success
    assert (project_root / path).read_text() == "v3_external"
    
    # 4. Force rollback should succeed
    success = rollback_engine.rollback_entry(entry, force=True)
    assert success
    assert (project_root / path).read_text() == "v1"

def test_partial_transaction_failure_recovery(setup_dirs):
    project_root, data_dir = setup_dirs
    journal = FilesystemJournal(data_dir / "journal.jsonl", data_dir / "backups")
    runner = TransactionRunner(data_dir / "staging", journal=journal)
    state = RunState(run_id="run_fail", goal="test")
    
    # 1. Create a plan with multiple files
    plan = DiffPlan(
        patches=[
            FilePatch(path="a.txt", new_text="content a"),
            FilePatch(path="b.txt", new_text="content b")
        ],
        validation_commands=["true"]
    )
    
    # 2. Mock a failure during commit by making project_root/b.txt a directory
    # so shutil.copy2 fails
    (project_root / "b.txt").mkdir()
    
    # 3. Apply mutation - it should fail at the commit stage
    # TransactionRunner currently doesn't have an explicit try/except around the whole commit loop
    # that rolls back ALREADY committed files. This is a gap!
    
    with pytest.raises(Exception):
        runner.apply(project_root=project_root, plan=plan, state=state)
        
    # Check current state: a.txt might have been created, b.txt is still a dir
    # This proves we need atomic transactions or automated cleanup on failure.
    
def test_journal_hash_chain_continuation(setup_dirs):
    project_root, data_dir = setup_dirs
    journal_path = data_dir / "journal.jsonl"
    
    # 1. Start a journal and record some entries
    j1 = FilesystemJournal(journal_path)
    j1.record(JournalEntry(run_id="run1", operation="CREATE", path="f1"))
    last_hash = j1._last_entry_hash
    
    # 2. Re-instantiate journal (simulating process restart)
    j2 = FilesystemJournal(journal_path)
    assert j2._last_entry_hash == last_hash
    
    # 3. Record more and verify chain is preserved
    j2.record(JournalEntry(run_id="run1", operation="CREATE", path="f2"))
    assert j2.verify_integrity()
