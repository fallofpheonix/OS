"""Memory package initialization and memory system coordinator."""

from pathlib import Path
from contracts.models import MemoryType
from .session_memory import SessionMemory
from .failure_memory import FailureMemory
from .architecture_memory import ArchitectureMemory
from .store import SQLiteMemoryStore
from .retrieval import RetrievalLayer


from typing import Any


class MemorySystem:
    """
    Unified memory system with separate subsystems by type.
    Coordinates all memory operations.
    """
    
    def __init__(self, persist_dir: Path | str | None = None) -> None:
        if persist_dir and str(persist_dir).endswith(".sqlite3"):
            self.sqlite = SQLiteMemoryStore(persist_dir)
        else:
            persist_dir = Path(persist_dir or "data")
            self.sqlite = SQLiteMemoryStore(persist_dir / "memory.sqlite3")
            
        self.retrieval = RetrievalLayer(self.sqlite)

        # Legacy semantic-like facade. Avoid requiring ChromaDB for core execution.
        self.semantic = self.sqlite
        
        # Initialize session memory (short-term, active task)
        self.session = SessionMemory()
        
        # Initialize specialized memory subsystems
        self.failures = FailureMemory(self.sqlite)
        self.architecture = ArchitectureMemory(self.sqlite)
    
    def record_failure(self, error_type: str, stack_trace: str, context: dict[str, Any], resolution: str | None = None) -> str:
        """Record a failure through failure memory."""
        return self.failures.record_failure(error_type, stack_trace, context, resolution)
    
    def record_decision(self, title: str, context: str, decision: str, consequences: str, alternatives: str | None = None) -> str:
        """Record an architecture decision through architecture memory."""
        return self.architecture.record_decision(title, context, decision, consequences, alternatives)
    
    def get_relevant_memories(
        self,
        query_embedding: list[float],
        memory_type: MemoryType,
        n_results: int = 5
    ) -> list[Any]:
        """Retrieve relevant memories by type."""
        return []
