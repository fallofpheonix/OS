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

from .orchestrator import SurfaceOrchestrator, Task, Risk, State

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
        depends_on TEXT
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


class OrchestratorService:
    def __init__(self, db_path: Optional[Path] = None):
        self.db_path = db_path or DB_PATH
        self.orch = SurfaceOrchestrator()
        self.lock = threading.RLock()
        self._init_db()
        self._load_from_db()

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
        cur.execute("SELECT id, wf_id, title, task_type, state, depends_on FROM tasks")
        for row in cur.fetchall():
            tid, wf_id, title, task_type, state, depends_on = row
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
            cur.execute(
                "INSERT OR REPLACE INTO tasks(id, wf_id, title, task_type, state, depends_on) VALUES(?,?,?,?,?,?)",
                (t.id, wf_id, t.title, t.task_type, t.state.value, deps),
            )
        for r in wf.risks.values():
            cur.execute(
                "INSERT OR REPLACE INTO risks(id, wf_id, description, severity, mitigated) VALUES(?,?,?,?,?)",
                (r.id, wf_id, r.description, r.severity, int(r.mitigated)),
            )
        conn.commit()
        conn.close()

    def register_handler(self, task_type: str, handler):
        self.orch.register_handler(task_type, handler)

    def create_workflow(self, wf_id: str):
        wf = self.orch.create_workflow(wf_id)
        self.persist_workflow(wf_id)
        return wf

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
