"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
import os
import json
import dataclasses
from pathlib import Path
from datetime import datetime
from typing import Dict, Set
from modules.core.event_core.event import Event

class IngestionLayer:
    def __init__(self, workspace_root: str):
        self.workspace_root = Path(workspace_root)
        self.state_dir = self.workspace_root / "runtime-state"
        self.state_dir.mkdir(parents=True, exist_ok=True)
        
        self._emitted_event_ids: Set[str] = set()
        self._last_timestamps: Dict[str, str] = {} # service -> last_timestamp

    def ingest(self, event: Event):
        # 1. Deduplication
        if event.event_id in self._emitted_event_ids:
            return
            
        # 2. Timestamp monotonicity validation
        last_ts = self._last_timestamps.get(event.service)
        if last_ts and event.timestamp < last_ts:
            # In a real system, we might quarantine this event or raise a critical error
            print(f"CRITICAL: Monotonicity violation for {event.service}: {event.timestamp} < {last_ts}")
            
        # 3. Persistence
        service_events_dir = self.state_dir / event.service / "events"
        now = datetime.fromisoformat(event.timestamp)
        path = service_events_dir / f"{now.year}/{now.month:02d}/{now.day:02d}"
        path.mkdir(parents=True, exist_ok=True)
        
        filename = path / "events.jsonl"
        
        # Append-only
        with open(filename, 'a') as f:
            f.write(json.dumps(dataclasses.asdict(event)) + "\n")
            
        self._emitted_event_ids.add(event.event_id)
        self._last_timestamps[event.service] = event.timestamp
        
        # 4. Rotation/Indexing (Stubs for MVP)
        # TODO: Implement 1GB rotation and checksum index files
