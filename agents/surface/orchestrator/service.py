"""Persistence and service helpers for SurfaceOrchestrator.

Provides lightweight SQLite-backed persistence and simple programmatic
API functions to create workflows, tasks, register handlers, and execute.
This keeps the core orchestrator free of storage concerns.
"""
from __future__ import annotations

import sqlite3
import threading
from pathlib import Path
from typing import Optional

from .orchestrator import SurfaceOrchestrator, Task, Risk, State, Workflow
from . import github_cli

DB_PATH = Path(__file__).parent / "state.db"


def _ensure_db(conn: sqlite3.Connection):
    cur = conn.cursor()
    cur.execute("""
    CREATE TABLE IF NOT EXISTS workflows (
        id TEXT PRIMARY KEY
    )
    """)
    cur.execute("""
    CREATE TABLE IF NOT EXISTS tasks (
        id TEXT PRIMARY KEY,
        wf_id TEXT,
        title TEXT,
        task_type TEXT,
        state TEXT,
        depends_on TEXT,
        issue_number INTEGER,
        issue_closed INTEGER DEFAULT 0
    )
    """)
    cur.execute("""
    CREATE TABLE IF NOT EXISTS risks (
        id TEXT PRIMARY KEY,
        wf_id TEXT,
        description TEXT,
        severity INTEGER,
        mitigated INTEGER
    )
    """)
    conn.commit()

    # Ensure older DBs get new columns if missing
    try:
        cur.execute("ALTER TABLE tasks ADD COLUMN issue_number INTEGER")
    except Exception:
        pass
    try:
        cur.execute("ALTER TABLE tasks ADD COLUMN issue_closed INTEGER DEFAULT 0")
    except Exception:
        pass
    conn.commit()


class OrchestratorService:
    def __init__(self, db_path: Optional[Path] = None):
        self.db_path = db_path or DB_PATH
        self.orch = SurfaceOrchestrator()
        self.lock = threading.RLock()
        self._init_db()
        self._load_from_db()
        # wire orchestrator callbacks
        self.orch.on_state_change = self._on_state_change
        self._bg_stop = threading.Event()
        self._bg_thread: Optional[threading.Thread] = None

    def _init_db(self):
        conn = sqlite3.connect(self.db_path)
        _ensure_db(conn)
        conn.close()

    def _load_from_db(self):
        conn = sqlite3.connect(self.db_path)
        cur = conn.cursor()
        cur.execute("SELECT id FROM workflows")
        for (wf_id,) in cur.fetchall():
            self.orch.create_workflow(wf_id)
        cur.execute("SELECT id, wf_id, title, task_type, state, depends_on, issue_number, issue_closed FROM tasks")
        for row in cur.fetchall():
            tid, wf_id, title, task_type, state, depends_on, issue_number, issue_closed = row
            t = Task(id=tid, title=title, task_type=task_type)
            try:
                t.state = State(state)
            except Exception:
                t.state = State.IDEA
            if depends_on:
                t.depends_on = set(depends_on.split(','))
            try:
                self.orch.add_task(wf_id, t)
            except KeyError:
                # workflow may have been created above
                pass
            # attach issue metadata into task metadata for convenience
            if issue_number:
                t.metadata.setdefault("issue_number", issue_number)
                t.metadata.setdefault("issue_closed", bool(issue_closed))

        # add approval gate automatically: tasks must be approved before MERGE
        for wf in list(self.orch.workflows.values()):
            self.add_approval_gate(wf.id)

        cur.execute("SELECT id, wf_id, description, severity, mitigated FROM risks")
        for row in cur.fetchall():
            rid, wf_id, desc, sev, mitigated = row
            r = Risk(id=rid, description=desc, severity=int(sev), mitigated=bool(mitigated))
            self.orch.add_risk(wf_id, r)
        conn.close()

    def persist_workflow(self, wf_id: str):
        conn = sqlite3.connect(self.db_path)
        cur = conn.cursor()
        cur.execute("INSERT OR IGNORE INTO workflows(id) VALUES(?)", (wf_id,))
        wf = self.orch.create_workflow(wf_id)
        for t in wf.tasks.values():
            deps = ','.join(t.depends_on) if t.depends_on else None
            issue_number = t.metadata.get("issue_number") or None
            issue_closed = int(bool(t.metadata.get("issue_closed", False)))
            cur.execute(
                "INSERT OR REPLACE INTO tasks(id, wf_id, title, task_type, state, depends_on, issue_number, issue_closed) VALUES(?,?,?,?,?,?,?,?)",
                (t.id, wf_id, t.title, t.task_type, t.state.value, deps, issue_number, issue_closed),
            )
        for r in wf.risks.values():
            cur.execute(
                "INSERT OR REPLACE INTO risks(id, wf_id, description, severity, mitigated) VALUES(?,?,?,?,?)",
                (r.id, wf_id, r.description, r.severity, int(r.mitigated)),
            )
        conn.commit()
        conn.close()

    def set_task_issue(self, wf_id: str, task_id: str, issue_number: int):
        """Associate a GitHub issue number with a task and persist it."""
        with self.lock:
            wf = self.orch.create_workflow(wf_id)
            t = wf.tasks.get(task_id)
            if not t:
                raise KeyError("unknown task")
            t.metadata["issue_number"] = int(issue_number)
            t.metadata["issue_closed"] = False
            self.persist_workflow(wf_id)

    def create_issue_for_task(self, wf_id: str, task_id: str, repo: Optional[str] = None, extra_labels: Optional[list] = None) -> int:
        """Create a GitHub issue for the given task using `gh` CLI and persist the mapping.

        `repo` should be owner/repo; if omitted we'll detect from git remote.
        """
        with self.lock:
            wf = self.orch.create_workflow(wf_id)
            t = wf.tasks.get(task_id)
            if not t:
                raise KeyError("unknown task")
        if not repo:
            repo = self._get_repo_from_git()
            if not repo:
                raise RuntimeError("could not determine repo; provide repo argument")

        # label mapping: base label + state label
        base_label = "surface:todo"
        state_label = self._label_for_state(t.state)
        labels = [base_label, state_label]
        if extra_labels:
            labels.extend(extra_labels)

        # ensure labels exist
        github_cli.ensure_label(repo, base_label, color="ff8800", description="Surface agent tasks")
        github_cli.ensure_label(repo, state_label, color="c2f0c2", description="State label for surface agent")

        title = f"[{wf_id}] {t.title}"
        body_lines = [f"Task ID: {t.id}", f"State: {t.state.value}", f"Type: {t.task_type}", "", "Metadata:"]
        for k, v in (t.metadata or {}).items():
            body_lines.append(f"- {k}: {v}")
        body = "\n".join(body_lines)

        issue_number = github_cli.create_issue_cli(repo, title, body, labels=labels)
        self.set_task_issue(wf_id, task_id, issue_number)
        return issue_number

    def _label_for_state(self, state: State) -> str:
        mapping = {
            State.IDEA: "surface:todo",
            State.RFC: "surface:triage",
            State.RESEARCH: "surface:in-progress",
            State.EXPERIMENT: "surface:in-progress",
            State.BUILD: "surface:in-progress",
            State.TEST: "surface:in-progress",
            State.DEBUG: "surface:in-progress",
            State.VALIDATE: "surface:ready",
            State.MERGE: "surface:ready-to-merge",
        }
        return mapping.get(state, "surface:todo")

    def sync_issue_labels_for_task(self, wf_id: str, task_id: str):
        """Update issue labels to reflect current task state."""
        with self.lock:
            wf = self.orch.create_workflow(wf_id)
            t = wf.tasks.get(task_id)
            if not t:
                raise KeyError("unknown task")
            issue_number = t.metadata.get("issue_number")
            if not issue_number:
                raise RuntimeError("task has no associated issue")
            repo = self._get_repo_from_git()
            if not repo:
                raise RuntimeError("could not determine repo")

        # compute desired labels and ensure they exist
        desired = ["surface:todo", self._label_for_state(t.state)]
        for lbl in set(desired):
            github_cli.ensure_label(repo, lbl)

        # remove other surface:* labels
        all_surface = ["surface:todo", "surface:triage", "surface:in-progress", "surface:ready", "surface:ready-to-merge", "surface:closed"]
        to_remove = [l for l in all_surface if l not in desired]

        # apply labels via gh CLI
        github_cli.add_labels(repo, int(issue_number), desired)
        github_cli.remove_labels(repo, int(issue_number), to_remove)


    def _get_repo_from_git(self) -> Optional[str]:
        import subprocess, re
        try:
            out = subprocess.check_output(["git", "remote", "get-url", "origin"]).decode().strip()
            m = re.search(r"github.com[:/](.+?)(?:\.git)?$", out)
            if m:
                return m.group(1)
        except Exception:
            return None

    def close_merged_issues(self):
        """Close GitHub issues for tasks that reached the MERGE state and haven't been closed yet."""
        repo = self._get_repo_from_git()
        if not repo:
            raise RuntimeError("could not determine GitHub repo from git remote")
        conn = sqlite3.connect(self.db_path)
        cur = conn.cursor()
        cur.execute("SELECT id, wf_id, issue_number FROM tasks WHERE state = ? AND issue_number IS NOT NULL AND issue_closed = 0", (State.MERGE.value,))
        rows = cur.fetchall()
        for tid, wf_id, issue_number in rows:
            try:
                # use gh CLI to close the issue
                import subprocess
                subprocess.check_call(["gh", "issue", "close", str(issue_number), "--repo", repo])
                # mark closed
                cur.execute("UPDATE tasks SET issue_closed = 1 WHERE id = ?", (tid,))
                conn.commit()
            except Exception as e:
                print(f"failed to close issue #{issue_number}: {e}")
        conn.close()

    def register_handler(self, task_type: str, handler):
        self.orch.register_handler(task_type, handler)

    def _on_state_change(self, wf_id: str, task_id: str, new_state: State):
        # persist state and sync labels when a task advances state
        try:
            # persist workflow state
            self.persist_workflow(wf_id)
            # if task has an associated issue, sync labels
            try:
                self.sync_issue_labels_for_task(wf_id, task_id)
            except Exception:
                pass
            # when a task reaches VALIDATE, create a PR if one doesn't exist
            try:
                if new_state == State.VALIDATE:
                    wf = self.orch.create_workflow(wf_id)
                    t = wf.tasks.get(task_id)
                    if t and not t.metadata.get("pr_url"):
                        try:
                            # if no GITHUB_TOKEN present, do a dry-run to avoid errors
                            import os
                            dry = not bool(os.environ.get("GITHUB_TOKEN"))
                            self.create_pr_for_task(wf_id, task_id, dry_run=dry)
                        except Exception:
                            pass
            except Exception:
                pass
        except Exception:
            pass

    def start_background_worker(self, interval_seconds: int = 60):
        """Start a background thread that periodically syncs labels and closes merged issues."""
        if self._bg_thread and self._bg_thread.is_alive():
            return

        def run():
            while not self._bg_stop.wait(interval_seconds):
                try:
                    # sync labels for all tasks with issues
                    for wf in list(self.orch.workflows.values()):
                        for tid in list(wf.tasks.keys()):
                            t = wf.tasks[tid]
                            if t.metadata.get("issue_number"):
                                try:
                                    self.sync_issue_labels_for_task(wf.id, tid)
                                except Exception:
                                    pass
                    # close any merged issues
                    try:
                        self.close_merged_issues()
                    except Exception:
                        pass
                except Exception:
                    pass

        self._bg_stop.clear()
        self._bg_thread = threading.Thread(target=run, daemon=True)
        self._bg_thread.start()

    def stop_background_worker(self):
        if self._bg_thread and self._bg_thread.is_alive():
            self._bg_stop.set()
            self._bg_thread.join(timeout=5)

    def create_workflow(self, wf_id: str):
        wf = self.orch.create_workflow(wf_id)
        self.persist_workflow(wf_id)

    def approve_task(self, wf_id: str, task_id: str, approver: Optional[str] = None) -> None:
        """Mark a task as approved (used to satisfy the MERGE gate)."""
        with self.lock:
            wf = self.orch.create_workflow(wf_id)
            t = wf.tasks.get(task_id)
            if not t:
                raise KeyError("unknown task")
            t.metadata["approved"] = True
            if approver:
                t.metadata["approved_by"] = approver
            import time
            t.metadata["approved_at"] = time.time()
            self.persist_workflow(wf_id)
            # sync labels to reflect approval
            try:
                self.sync_issue_labels_for_task(wf_id, task_id)
            except Exception:
                pass

    def add_approval_gate(self, wf_id: str) -> None:
        """Add a validation gate that requires per-task approval before MERGE."""
        def check_approved(wf: Workflow, task: Task) -> bool:
            # allow merge only if the specific task is approved
            return bool(task.metadata.get("approved", False))
        # delegate to the underlying orchestrator's gate registration
        self.orch.add_validation_gate(wf_id, State.MERGE, check_approved)

    def add_task(self, wf_id: str, task: Task):
        self.orch.add_task(wf_id, task)
        self.persist_workflow(wf_id)

    def add_dependency(self, wf_id: str, task_id: str, depends_on_id: str):
        self.orch.add_dependency(wf_id, task_id, depends_on_id)
        self.persist_workflow(wf_id)

    def execute_workflow_async(self, wf_id: str):
        t = threading.Thread(target=self._exec_and_persist, args=(wf_id,), daemon=True)
        t.start()
        return t

    def _exec_and_persist(self, wf_id: str):
        self.orch.execute_workflow(wf_id)
        self.persist_workflow(wf_id)

    def create_pr_for_task(self, wf_id: str, task_id: str, base: str = "main", repo: Optional[str] = None, dry_run: bool = False) -> str:
        """Create a branch, commit a small PR note, push, and open a PR for the task.

        Returns the PR URL and stores it in task.metadata['pr_url'].
        """
        import subprocess
        import os
        from . import github_integration

        with self.lock:
            wf = self.orch.create_workflow(wf_id)
            t = wf.tasks.get(task_id)
            if not t:
                raise KeyError("unknown task")

        if not repo:
            repo = self._get_repo_from_git()

        if not repo:
            raise RuntimeError("could not determine repo; provide repo argument")

        head = f"surface/{wf_id}/{task_id}"
        prs_dir = Path(__file__).parent / "prs"
        prs_dir.mkdir(parents=True, exist_ok=True)
        note_path = prs_dir / f"{task_id}.md"
        body_lines = [f"# PR for task {task_id}", f"Workflow: {wf_id}", f"Task title: {t.title}", f"State: {t.state.value}", "", "Metadata:"]
        for k, v in (t.metadata or {}).items():
            body_lines.append(f"- {k}: {v}")
        body = "\n".join(body_lines)

        note_path.write_text(body)

        # create branch, add file, commit, push
        try:
            if dry_run:
                pr_url = f"dry-run://{repo}/{head}"
            else:
                subprocess.check_call(["git", "fetch", "origin"])
                # create branch based off origin/<base>
                subprocess.check_call(["git", "checkout", "-b", head, f"origin/{base}"])
                subprocess.check_call(["git", "add", str(note_path)])
                subprocess.check_call(["git", "commit", "-m", f"surface: PR for {wf_id}/{task_id}"])
                subprocess.check_call(["git", "push", "--set-upstream", "origin", head])
                # create PR via API helper
                owner_repo = repo
                pr = github_integration.create_pull_request(owner_repo, title=f"Surface: {t.title}", head=head, base=base, body=body)
                pr_url = pr.get("html_url")
        except Exception as e:
            # try to cleanup branch locally
            try:
                subprocess.check_call(["git", "checkout", base])
                subprocess.check_call(["git", "branch", "-D", head])
            except Exception:
                pass
            raise

        # persist PR url into task metadata
        with self.lock:
            t.metadata["pr_url"] = pr_url
            self.persist_workflow(wf_id)

        return pr_url
