"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
import unittest
import os
from pathlib import Path
from modules.core.event_core.emitter import EventEmitter
from modules.core.event_core.trace_context import TraceContext, get_trace_id, get_parent_event_id
from modules.core.event_core.schema_validator import SchemaValidationError

class TestEventEmitter(unittest.TestCase):
    def setUp(self):
        self.schema_path = str(Path(__file__).parents[4] / "contracts" / "events" / "runtime-event.schema.yaml")
        self.runtime_id = "test-runtime"
        self.emitter = EventEmitter(
            runtime_id=self.runtime_id,
            service="supervisor",
            supervisor="test-supervisor",
            component="test-component",
            schema_path=self.schema_path
        )

    def test_basic_emission(self):
        event = self.emitter.emit(
            event_type="startup",
            severity="info",
            lifecycle_phase="bootstrap"
        )
        self.assertEqual(event.runtime_id, self.runtime_id)
        self.assertEqual(event.event_type, "startup")
        self.assertEqual(event.severity, "info")
        self.assertIsNotNone(event.event_id)
        self.assertIsNotNone(event.trace_id)

    def test_trace_propagation(self):
        trace_id = "custom-trace-id"
        with TraceContext(trace_id=trace_id):
            try:
                e1 = self.emitter.emit(
                    event_type="startup",
                    severity="info",
                    lifecycle_phase="bootstrap"
                )
                self.assertEqual(e1.trace_id, trace_id)
                
                e2 = self.emitter.emit(
                    event_type="phase-transition",
                    severity="info",
                    lifecycle_phase="ready"
                )
                self.assertEqual(e2.trace_id, trace_id)
                self.assertEqual(e2.parent_event, e1.event_id)
            except SchemaValidationError as e:
                print(f"\nValidation errors: {e.errors}")
                raise

    def test_validation_failure(self):
        with self.assertRaises(SchemaValidationError):
            self.emitter.emit(
                event_type="invalid-type", # Not in enum
                severity="info",
                lifecycle_phase="bootstrap"
            )

if __name__ == "__main__":
    unittest.main()
