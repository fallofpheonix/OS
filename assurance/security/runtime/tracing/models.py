"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
from __future__ import annotations
import os
from dataclasses import dataclass
from datetime import datetime, timezone
from typing import Optional
from runtime.core.event_core.emitter import EventEmitter, Event
from runtime.core.event_core.trace_context import get_trace_id, get_parent_event_id

# GLOBAL EMITTER for the Guard runtime
# In production, this would be initialized by the supervisor
_SCHEMA_PATH = os.path.join(os.path.dirname(os.path.dirname(os.path.dirname(__file__))), "contracts/events/runtime-event.schema.yaml")
_guard_emitter = EventEmitter(
    runtime_id=os.getenv("RUNTIME_ID", "default-guard-runtime"),
    service="control-plane",
    supervisor="default-supervisor",
    component="guard-runtime",
    schema_path=_SCHEMA_PATH
)

@dataclass(frozen=True, slots=True)
class RuntimeTrace:
    trace_id: str
    runtime_category: str
    operation: str
    target: str
    duration_ms: int
    success: bool
    error_type: str
    timestamp: datetime
    # New canonical event reference
    event: Optional[Event] = None


def create_runtime_trace(
    *,
    runtime_category: str,
    operation: str,
    target: str,
    duration_ms: int,
    success: bool,
    error_type: str = "",
    trace_id: str | None = None,
) -> RuntimeTrace:
    # Map raw error_type to a valid failure_class enum
    failure_class = None
    if not success:
        failure_class = "unknown"
        if "timeout" in error_type.lower():
            failure_class = "timeout"
        elif "limit" in error_type.lower() or "too large" in error_type.lower() or "rejected" in error_type.lower():
            failure_class = "resource-exhaustion"
        elif "policy" in error_type.lower() or "violation" in error_type.lower() or "workspace" in error_type.lower():
            failure_class = "authorization"
            
    # 1. Emit canonical event
    now = datetime.now(timezone.utc).isoformat()
    try:
        event = _guard_emitter.emit(
            event_type="phase-transition" if success else "failure",
            severity="info" if success else "error",
            failure_class=failure_class,
            # Mandatory fields for the schema
            lifecycle_phase="ready", 
            degraded=False,
            quarantined=False,
            created_at=now,
            stored_in="memory",
            restart_count=0,
            payload={
                "operation": operation,
                "target": target,
                "duration_ms": duration_ms,
                "raw_error": error_type
            }
        )
    except Exception as e:
        if hasattr(e, 'errors'):
            print("Schema validation errors:", e.errors)
        raise e

    # 2. Return compatible trace object
    return RuntimeTrace(
        trace_id=event.trace_id,
        runtime_category=runtime_category,
        operation=operation,
        target=target,
        duration_ms=max(0, duration_ms),
        success=success,
        error_type=error_type,
        timestamp=datetime.fromisoformat(event.timestamp),
        event=event
    )
