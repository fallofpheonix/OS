"""SQLite-first structured memory store."""

from __future__ import annotations

import json
import sqlite3
from pathlib import Path
from typing import Any, Iterable


class SQLiteMemoryStore:
    def __init__(self, path: Path | str = "data/memory.sqlite3"):
        self.path = Path(path)
        # Ensure the filename is used if it ends in .sqlite3, 
        # but the directory structure is preserved.
        if self.path.suffix != ".sqlite3":
             self.path = self.path / "memory.sqlite3"
             
        self.path.parent.mkdir(parents=True, exist_ok=True)
        self.conn = sqlite3.connect(self.path, check_same_thread=False)
        self.conn.row_factory = sqlite3.Row
        self._init_schema()

    def _init_schema(self) -> None:
        self.conn.executescript(
            """
            CREATE TABLE IF NOT EXISTS goals (
                id TEXT PRIMARY KEY,
                goal TEXT NOT NULL,
                created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
            );

            CREATE TABLE IF NOT EXISTS decompositions (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                goal_id TEXT NOT NULL,
                planner_json TEXT NOT NULL,
                valid INTEGER NOT NULL,
                created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
            );

            CREATE TABLE IF NOT EXISTS artifacts (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                run_id TEXT NOT NULL,
                task_id TEXT NOT NULL,
                content TEXT NOT NULL,
                path TEXT,
                created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
            );

            CREATE TABLE IF NOT EXISTS failure_lessons (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                run_id TEXT,
                task_id TEXT,
                lesson TEXT NOT NULL,
                created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
            );

            CREATE TABLE IF NOT EXISTS failure_records (
                failure_id TEXT PRIMARY KEY,
                run_id TEXT NOT NULL,
                task_id TEXT,
                failure_type TEXT NOT NULL,
                raw_error TEXT NOT NULL,
                structured_context TEXT NOT NULL,
                attempt INTEGER NOT NULL,
                timestamp TEXT NOT NULL,
                resolution_status TEXT NOT NULL
            );

            CREATE TABLE IF NOT EXISTS help_requests (
                help_id TEXT PRIMARY KEY,
                run_id TEXT NOT NULL,
                payload_json TEXT NOT NULL,
                status TEXT NOT NULL,
                created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
            );

            CREATE TABLE IF NOT EXISTS snapshots (
                snapshot_id TEXT PRIMARY KEY,
                run_id TEXT NOT NULL,
                root_path TEXT NOT NULL,
                manifest_json TEXT NOT NULL,
                state_json TEXT NOT NULL,
                created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
            );

            CREATE TABLE IF NOT EXISTS architecture_summaries (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                source TEXT NOT NULL,
                summary TEXT NOT NULL,
                created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
            );

            CREATE TABLE IF NOT EXISTS task_leases (
                run_id TEXT NOT NULL,
                task_id TEXT NOT NULL,
                worker_id TEXT NOT NULL,
                expires_at TEXT NOT NULL,
                PRIMARY KEY (run_id, task_id)
            );

            CREATE VIRTUAL TABLE IF NOT EXISTS memory_fts USING fts5(
                kind,
                ref,
                content
            );
            """
        )
        self.conn.commit()

    def record_goal(self, goal_id: str, goal: str) -> None:
        self.conn.execute(
            "INSERT OR REPLACE INTO goals (id, goal) VALUES (?, ?)",
            (goal_id, goal),
        )
        self.conn.execute(
            "INSERT INTO memory_fts (kind, ref, content) VALUES ('goal', ?, ?)",
            (f"{goal_id}", goal),
        )
        self.conn.commit()

    def record_decomposition(self, goal_id: str, planner_json: str, valid: bool) -> None:
        self.conn.execute(
            "INSERT INTO decompositions (goal_id, planner_json, valid) VALUES (?, ?, ?)",
            (goal_id, planner_json, int(valid)),
        )
        self.conn.commit()

    def record_artifact(self, run_id: str, task_id: str, content: str, path: str | None = None) -> None:
        self.conn.execute(
            "INSERT INTO artifacts (run_id, task_id, content, path) VALUES (?, ?, ?, ?)",
            (run_id, task_id, content, path),
        )
        self.conn.execute(
            "INSERT INTO memory_fts (kind, ref, content) VALUES ('artifact', ?, ?)",
            (f"{run_id}:{task_id}", content),
        )
        self.conn.commit()

    def record_failure_lesson(self, run_id: str, task_id: str | None, lesson: str) -> None:
        self.conn.execute(
            "INSERT INTO failure_lessons (run_id, task_id, lesson) VALUES (?, ?, ?)",
            (run_id, task_id, lesson),
        )
        self.conn.execute(
            "INSERT INTO memory_fts (kind, ref, content) VALUES ('failure', ?, ?)",
            (f"{run_id}:{task_id or 'run'}", lesson),
        )
        self.conn.commit()

    def record_failure_record(self, payload: dict) -> None:
        self.conn.execute(
            """
            INSERT OR REPLACE INTO failure_records (
                failure_id, run_id, task_id, failure_type, raw_error,
                structured_context, attempt, timestamp, resolution_status
            ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
            """,
            (
                payload["failure_id"],
                payload["run_id"],
                payload.get("task_id"),
                payload["failure_type"],
                payload["raw_error"],
                payload["structured_context"],
                int(payload["attempt"]),
                payload["timestamp"],
                payload["resolution_status"],
            ),
        )
        self.conn.execute(
            "INSERT INTO memory_fts (kind, ref, content) VALUES ('failure_record', ?, ?)",
            (payload["failure_id"], f"{payload['failure_type']}\n{payload['raw_error']}"),
        )
        self.conn.commit()

    def record_help_request(self, help_id: str, run_id: str, payload_json: str, status: str) -> None:
        self.conn.execute(
            """
            INSERT OR REPLACE INTO help_requests (help_id, run_id, payload_json, status)
            VALUES (?, ?, ?, ?)
            """,
            (help_id, run_id, payload_json, status),
        )
        self.conn.commit()

    def record_snapshot(
        self,
        snapshot_id: str,
        run_id: str,
        root_path: str,
        manifest_json: str,
        state_json: str,
    ) -> None:
        self.conn.execute(
            """
            INSERT OR REPLACE INTO snapshots (
                snapshot_id, run_id, root_path, manifest_json, state_json
            ) VALUES (?, ?, ?, ?, ?)
            """,
            (snapshot_id, run_id, root_path, manifest_json, state_json),
        )
        self.conn.commit()

    def record_architecture_summary(self, source: str, summary: Any) -> None:
        if not isinstance(summary, str):
            summary_str = json.dumps(summary, indent=2)
        else:
            summary_str = summary

        self.conn.execute(
            "INSERT INTO architecture_summaries (source, summary) VALUES (?, ?)",
            (source, summary_str),
        )
        self.conn.execute(
            "INSERT INTO memory_fts (kind, ref, content) VALUES ('architecture', ?, ?)",
            (source, summary_str),
        )
        self.conn.commit()

    def acquire_lease(self, run_id: str, task_id: str, worker_id: str, duration_s: int = 300) -> bool:
        """Attempt to acquire an exclusive lease for a task. Returns True if successful."""
        from datetime import datetime, timedelta, UTC
        now = datetime.now(UTC)
        expires_at = (now + timedelta(seconds=duration_s)).isoformat()
        
        try:
            # 1. Clean up expired leases first
            self.conn.execute(
                "DELETE FROM task_leases WHERE expires_at < ?",
                (now.isoformat(),)
            )
            
            # 2. Try to insert new lease
            self.conn.execute(
                "INSERT INTO task_leases (run_id, task_id, worker_id, expires_at) VALUES (?, ?, ?, ?)",
                (run_id, task_id, worker_id, expires_at)
            )
            self.conn.commit()
            return True
        except sqlite3.IntegrityError:
            # Lease already exists for this (run_id, task_id)
            return False

    def release_lease(self, run_id: str, task_id: str, worker_id: str) -> None:
        """Release a lease owned by the worker."""
        self.conn.execute(
            "DELETE FROM task_leases WHERE run_id = ? AND task_id = ? AND worker_id = ?",
            (run_id, task_id, worker_id)
        )
        self.conn.commit()

    def search(self, query: str, *, kinds: Iterable[str] | None = None, limit: int = 5) -> list[dict[str, Any]]:
        # FTS5 special characters: https://www.sqlite.org/fts5.html#full_text_query_syntax
        forbidden = r'":*^'
        safe_query = query
        for char in forbidden:
            safe_query = safe_query.replace(char, " ")
        
        safe_query = safe_query.strip(". ")
        
        tokens = [token for token in safe_query.split() if token]
        if not tokens:
            return []
        
        match_query = " ".join(f'"{token}"' for token in tokens)
        sql = "SELECT kind, ref, content, rank FROM memory_fts WHERE memory_fts MATCH ?"
        params: list = [match_query]
        if kinds:
            placeholders = ",".join("?" for _ in kinds)
            sql += f" AND kind IN ({placeholders})"
            params.extend(kinds)
        sql += " ORDER BY rank LIMIT ?"
        params.append(limit)
        return [dict(row) for row in self.conn.execute(sql, params)]
