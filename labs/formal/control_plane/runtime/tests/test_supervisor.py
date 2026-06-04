"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
import unittest
import tempfile
import shutil
import json
from pathlib import Path
from control_plane.runtime.supervisor import Supervisor

class TestSupervisor(unittest.TestCase):
    def setUp(self):
        self.temp_dir = tempfile.mkdtemp()
        self.schema_path = str(Path(__file__).parents[3] / "contracts" / "events" / "runtime-event.schema.yaml")
        self.supervisor = Supervisor(
            workspace_root=self.temp_dir,
            supervisor_name="test-supervisor",
            schema_path=self.schema_path
        )

    def tearDown(self):
        shutil.rmtree(self.temp_dir)

    def test_lifecycle_events(self):
        self.supervisor.startup()
        self.supervisor.ready()
        self.supervisor.degraded("high cpu")
        self.supervisor.shutdown()
        
        # Verify events were persisted
        events_base = Path(self.temp_dir) / "runtime-state" / "supervisor" / "events"
        event_files = list(events_base.rglob("events.jsonl"))
        self.assertTrue(len(event_files) > 0)
        
        with open(event_files[0], 'r') as f:
            events = [json.loads(line) for line in f]
            
        self.assertEqual(len(events), 4)
        self.assertEqual(events[0]['event_type'], "startup")
        self.assertEqual(events[1]['event_type'], "phase-transition")
        self.assertEqual(events[2]['event_type'], "degradation")
        self.assertEqual(events[3]['event_type'], "shutdown")

    def test_restart_causality(self):
        old_runtime_id = self.supervisor.runtime_id
        self.supervisor.restart("manual update")
        
        new_runtime_id = self.supervisor.runtime_id
        self.assertNotEqual(old_runtime_id, new_runtime_id)
        self.assertEqual(self.supervisor.restart_count, 1)
        
        # Check persisted events
        events_base = Path(self.temp_dir) / "runtime-state" / "supervisor" / "events"
        event_files = list(events_base.rglob("events.jsonl"))
        
        restart_event = None
        for file in event_files:
            with open(file, 'r') as f:
                for line in f:
                    e = json.loads(line)
                    if e['event_type'] == "restart":
                        restart_event = e
                        break
        
        self.assertIsNotNone(restart_event)
        self.assertEqual(restart_event['metadata']['previous_runtime_id'], old_runtime_id)
        self.assertEqual(restart_event['restart_count'], 1)

if __name__ == "__main__":
    unittest.main()
