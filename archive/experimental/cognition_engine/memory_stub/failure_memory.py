"""
Failure memory: debugging and error analysis history.
Long-term store of failures, stack traces, resolutions.
"""

from typing import List, Optional, Dict, Any
from uuid import uuid4
from datetime import datetime

from contracts.models import MemoryRecord, MemoryType


class FailureMemory:
    """
    Stores failure records for debugging reference.
    Enables learning from past failures.
    """
    
    def __init__(self, semantic_store: Any = None, sqlite_store: Any = None) -> None:
        self.semantic_store = semantic_store
        self.sqlite_store = sqlite_store
        self.failures: Dict[str, Dict[str, Any]] = {}
    
    def record_failure(
        self,
        error_type: str,
        stack_trace: str,
        context: Dict[str, Any],
        resolution: Optional[str] = None
    ) -> str:
        """
        Record a failure for future reference.
        
        Returns:
            Failure ID for retrieval
        """
        failure_id = str(uuid4())
        
        failure_record = {
            "id": failure_id,
            "error_type": error_type,
            "stack_trace": stack_trace,
            "context": context,
            "resolution": resolution,
            "timestamp": datetime.now().isoformat(),
            "resolved": bool(resolution)
        }
        
        self.failures[failure_id] = failure_record
        
        if self.sqlite_store:
            import json
            payload = {
                "failure_id": failure_id,
                "run_id": context.get("run_id", "unknown"),
                "task_id": context.get("task_id"),
                "failure_type": error_type,
                "raw_error": stack_trace,
                "structured_context": json.dumps(context),
                "attempt": context.get("attempt", 1),
                "timestamp": failure_record["timestamp"],
                "resolution_status": "success" if resolution else "failed"
            }
            self.sqlite_store.record_failure_record(payload)

        # Store in semantic memory if available
        if self.semantic_store:
            content = f"{error_type}\n{stack_trace}\n{resolution or 'Unresolved'}"
            memory_record = MemoryRecord(
                id=failure_id,
                memory_type=MemoryType.FAILURE,
                content=content,
                metadata={
                    "error_type": error_type,
                    "resolved": bool(resolution),
                    "context": str(context)
                }
            )
            self.semantic_store.store(memory_record)
        
        return failure_id

    def record_repair_outcome(self, failure_id: str, outcome_status: str, notes: str = "") -> None:
        if self.sqlite_store:
            self.sqlite_store.conn.execute(
                "UPDATE failure_records SET resolution_status = ? WHERE failure_id = ?",
                (outcome_status, failure_id)
            )
            self.sqlite_store.conn.commit()

    def get_failure(self, failure_id: str) -> Optional[Dict[str, Any]]:
        """Get failure record by ID."""
        if failure_id in self.failures:
            return self.failures[failure_id]
        if self.sqlite_store:
            row = self.sqlite_store.conn.execute(
                "SELECT * FROM failure_records WHERE failure_id = ?", (failure_id,)
            ).fetchone()
            if row:
                import json
                return {
                    "id": row["failure_id"],
                    "error_type": row["failure_type"],
                    "stack_trace": row["raw_error"],
                    "context": json.loads(row["structured_context"]),
                    "resolution": row["resolution_status"],
                    "timestamp": row["timestamp"],
                    "resolved": row["resolution_status"] == "success"
                }
        return None
    
    def find_similar_failures(
        self,
        error_type: str,
        stack_trace: str
    ) -> List[Dict[str, Any]]:
        """Find similar failures from history that were successfully repaired."""
        if self.sqlite_store:
            rows = self.sqlite_store.conn.execute(
                "SELECT * FROM failure_records WHERE failure_type = ? AND resolution_status = 'success' LIMIT 5",
                (error_type,)
            ).fetchall()
            if rows:
                import json
                return [
                    {
                        "id": row["failure_id"],
                        "error_type": row["failure_type"],
                        "stack_trace": row["raw_error"],
                        "context": json.loads(row["structured_context"]),
                        "resolution": row["resolution_status"],
                        "timestamp": row["timestamp"],
                        "resolved": True
                    }
                    for row in rows
                ]
        
        # Fallback to memory
        similar = [
            f for f in self.failures.values()
            if f["error_type"] == error_type and f.get("resolved")
        ]
        return similar[:5]
