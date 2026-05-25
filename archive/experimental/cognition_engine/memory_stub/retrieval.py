"""Scoped retrieval layer."""

from __future__ import annotations

from .store import SQLiteMemoryStore


class RetrievalLayer:
    def __init__(self, store: SQLiteMemoryStore):
        self.store = store

    def planner_context(self, goal: str, limit: int = 5) -> list[str]:
        rows = self.store.search(
            goal,
            kinds=["artifact", "failure", "architecture"],
            limit=limit,
        )
        return [f"[{row['kind']}:{row['ref']}]\n{row['content'][:1200]}" for row in rows]
