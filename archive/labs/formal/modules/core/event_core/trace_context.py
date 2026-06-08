"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
import threading
import contextvars
from typing import Optional

# Using contextvars for thread-local and async-compatible context management
_trace_id_var: contextvars.ContextVar[Optional[str]] = contextvars.ContextVar("trace_id", default=None)
_parent_event_id_var: contextvars.ContextVar[Optional[str]] = contextvars.ContextVar("parent_event_id", default=None)

def get_trace_id() -> Optional[str]:
    return _trace_id_var.get()

def set_trace_id(trace_id: str) -> contextvars.Token:
    return _trace_id_var.set(trace_id)

def reset_trace_id(token: contextvars.Token):
    _trace_id_var.reset(token)

def get_parent_event_id() -> Optional[str]:
    return _parent_event_id_var.get()

def set_parent_event_id(event_id: str) -> contextvars.Token:
    return _parent_event_id_var.set(event_id)

def reset_parent_event_id(token: contextvars.Token):
    _parent_event_id_var.reset(token)

class TraceContext:
    """Context manager for managing trace_id and parent_event_id."""
    def __init__(self, trace_id: str, parent_event_id: Optional[str] = None):
        self.trace_id = trace_id
        self.parent_event_id = parent_event_id
        self._tokens = []

    def __enter__(self):
        self._tokens.append(_trace_id_var.set(self.trace_id))
        if self.parent_event_id:
            self._tokens.append(_parent_event_id_var.set(self.parent_event_id))
        return self

    def __exit__(self, exc_type, exc_val, exc_tb):
        for token in reversed(self._tokens):
            if token.var == _trace_id_var:
                _trace_id_var.reset(token)
            elif token.var == _parent_event_id_var:
                _parent_event_id_var.reset(token)
        self._tokens.clear()

def preserve_trace():
    """Returns a TraceContext for the currently active trace."""
    return TraceContext(get_trace_id(), get_parent_event_id())
