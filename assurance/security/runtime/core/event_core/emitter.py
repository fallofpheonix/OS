"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
import uuid
import time
from datetime import datetime, timezone
from typing import Optional, Any, Callable, Dict, Set
from .event import Event
from .schema_validator import EventSchemaValidator, SchemaValidationError
from .trace_context import get_trace_id, get_parent_event_id, set_parent_event_id

class EventEmitter:
    def __init__(self, 
                 runtime_id: str, 
                 service: str, 
                 supervisor: str, 
                 component: str,
                 schema_path: str,
                 sink: Optional[Callable[[Event], None]] = None):
        self.runtime_id = runtime_id
        self.service = service
        self.supervisor = supervisor
        self.component = component
        self.validator = EventSchemaValidator(schema_path)
        self.sink = sink
        self._emitted_event_ids: Set[str] = set()
        self._sequence_counter = 0

    def emit(self, 
             event_type: str, 
             severity: str, 
             payload: Dict[str, Any] = None,
             failure_class: Optional[str] = None,
             lifecycle_phase: str = "unknown",
             restart_count: int = 0,
             degraded: bool = False,
             quarantined: bool = False,
             **kwargs) -> Event:
        
        event_id = str(uuid.uuid4())
        trace_id = get_trace_id() or str(uuid.uuid4())
        parent_event = get_parent_event_id()
        
        now = datetime.now(timezone.utc).isoformat()
        
        # Default field mapping
        event_fields = {
            "event_id": event_id,
            "trace_id": trace_id,
            "parent_event": parent_event,
            "runtime_id": self.runtime_id,
            "service": self.service,
            "supervisor": self.supervisor,
            "component": self.component,
            "event_type": event_type,
            "severity": severity,
            "failure_class": failure_class,
            "lifecycle_phase": lifecycle_phase,
            "restart_count": restart_count,
            "timestamp": now,
            "occurred_at": now,
            "degraded": degraded,
            "quarantined": quarantined,
            "created_at": now,
            "stored_in": "memory",
            "sequence_number": self._sequence_counter,
            "metadata": payload or {}
        }
        
        # Allow overrides from kwargs
        event_fields.update(kwargs)
        
        event = Event(**event_fields)
        
        # Deduplication check
        if event_id in self._emitted_event_ids:
            return event # Already emitted

        # Validation
        try:
            self.validator.validate_event_object(event)
        except SchemaValidationError as e:
            # Re-raise or handle as needed for machine-parseability
            raise
            
        self._sequence_counter += 1
        self._emitted_event_ids.add(event_id)
        
        # Update parent context for subsequent events in this thread/task
        set_parent_event_id(event_id)
        
        if self.sink:
            self.sink(event)
            
        return event
