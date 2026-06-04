"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
import os
import json
import uuid
from pathlib import Path
from typing import Optional, List
from modules.core.event_core.emitter import EventEmitter, Event
from modules.core.event_core.trace_context import TraceContext
from control_plane.runtime_state.ingestion.ingestor import IngestionLayer

class Supervisor:
    def __init__(self, 
                 workspace_root: str,
                 supervisor_name: str,
                 schema_path: str):
        self.workspace_root = Path(workspace_root)
        self.supervisor_name = supervisor_name
        self.schema_path = schema_path
        
        self.state_dir = self.workspace_root / "runtime-state" / "supervisor"
        self.checkpoint_file = self.state_dir / "checkpoint.json"
        self.state_dir.mkdir(parents=True, exist_ok=True)
        
        self.ingestor = IngestionLayer(workspace_root)
        self.runtime_id, self.restart_count, self.previous_runtime_id = self._load_checkpoint()
        
        self.emitter = EventEmitter(
            runtime_id=self.runtime_id,
            service="supervisor",
            supervisor=self.supervisor_name,
            component="core",
            schema_path=self.schema_path,
            sink=self.ingestor.ingest
        )

    def _load_checkpoint(self):
        if self.checkpoint_file.exists():
            with open(self.checkpoint_file, 'r') as f:
                data = json.load(f)
                return data['runtime_id'], data.get('restart_count', 0), data.get('previous_runtime_id')
        
        return str(uuid.uuid4()), 0, None

    def _save_checkpoint(self):
        with open(self.checkpoint_file, 'w') as f:
            json.dump({
                'runtime_id': self.runtime_id,
                'restart_count': self.restart_count,
                'previous_runtime_id': self.previous_runtime_id
            }, f)

    def startup(self):
        self.emitter.emit(
            event_type="startup",
            severity="info",
            lifecycle_phase="bootstrap",
            restart_count=self.restart_count
        )

    def ready(self):
        self.emitter.emit(
            event_type="phase-transition", # Roadmap says 'ready', but schema says 'phase-transition'
            severity="info",
            lifecycle_phase="ready",
            restart_count=self.restart_count
        )

    def restart(self, reason: str):
        self.previous_runtime_id = self.runtime_id
        self.runtime_id = str(uuid.uuid4())
        self.restart_count += 1
        self._save_checkpoint()
        
        self.emitter.emit(
            event_type="restart",
            severity="warning",
            lifecycle_phase="bootstrap",
            restart_count=self.restart_count,
            payload={"reason": reason, "previous_runtime_id": self.previous_runtime_id}
        )

    def degraded(self, reason: str):
        self.emitter.emit(
            event_type="degradation",
            severity="warning",
            lifecycle_phase="degraded",
            restart_count=self.restart_count,
            degraded=True,
            payload={"reason": reason}
        )

    def shutdown(self):
        self.emitter.emit(
            event_type="shutdown",
            severity="info",
            lifecycle_phase="shutdown",
            restart_count=self.restart_count
        )

    def fatal(self, error: str):
        self.emitter.emit(
            event_type="failure",
            severity="critical",
            lifecycle_phase="shutdown",
            restart_count=self.restart_count,
            failure_class="unknown",
            payload={"error": error}
        )
