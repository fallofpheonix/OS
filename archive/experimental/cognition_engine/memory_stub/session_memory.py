"""
Session memory: active task context and working memory.
Lightweight, session-scoped storage for current work.
"""

from typing import Dict, Any, Optional, List


class SessionMemory:
    """
    Session-scoped memory for active task.
    Cleared at end of session.
    """
    
    def __init__(self) -> None:
        self.state: Dict[str, Any] = {}
        self.task_context: Dict[str, Any] = {}
        self.working_memory: List[str] = []  # FIFO buffer
        self.max_working_items = 50
    
    def set_context(self, key: str, value: Any) -> None:
        """Set task context variable."""
        self.task_context[key] = value
    
    def get_context(self, key: str) -> Optional[Any]:
        """Get task context variable."""
        return self.task_context.get(key)
    
    def add_to_working(self, item: str) -> None:
        """Add to working memory buffer."""
        self.working_memory.append(item)
        if len(self.working_memory) > self.max_working_items:
            self.working_memory.pop(0)  # FIFO
    
    def get_working_memory(self) -> List[str]:
        """Get current working memory."""
        return self.working_memory[-10:]  # Return last 10 items
    
    def clear(self) -> None:
        """Clear session memory."""
        self.state.clear()
        self.task_context.clear()
        self.working_memory.clear()
