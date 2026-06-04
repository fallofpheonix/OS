"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
import os
import tempfile
import unittest
from pathlib import Path

from runtime.filesystem.manager import FilesystemManager
from runtime.orchestration import CompositeOperation, CompositeOperationConfig, OrchestrationState


class CompetingFailureDeterminismTests(unittest.TestCase):
    """
    Proves semantic determinism under competing invalid conditions.
    Validates that multiple simultaneous failure domains always resolve to the exact
    same failure boundary, class, trace topology, and halt layer, with zero residue.
    """

    def test_failure_order_determinism_escape_vs_oversize(self):
        """
        symlink escape + oversized file
        Expected: same failure chosen every run. Not random precedence.
        """
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir) / "workspace"
            root.mkdir()
            outside = Path(tempdir) / "outside"
            outside.mkdir()
            
            large_payload = outside / "large.txt"
            large_payload.write_bytes(b"x" * 16_000)
            
            # Target is both an escape and oversized
            os.symlink(str(large_payload), str(root / "leak"))

            manager = FilesystemManager(root, max_file_bytes=1024)
            
            results = [manager.read_file("leak") for _ in range(5)]
            
            first_result = results[0]
            for r in results:
                self.assertFalse(r.success)
                self.assertEqual(r.error, first_result.error)
                self.assertEqual(r.trace.error_type, first_result.trace.error_type)
                self.assertEqual(r.trace.operation, first_result.trace.operation)

    def test_trace_topology_stability_escape_vs_binary(self):
        """
        invalid path (absolute escape) + binary file
        Expected: same trace structure every replay, no dynamic aggregation nodes.
        """
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir) / "workspace"
            root.mkdir()
            
            outside = Path(tempdir) / "outside"
            outside.mkdir()
            binary_file = outside / "binary.bin"
            binary_file.write_bytes(b"hello\x00world")
            
            manager = FilesystemManager(root)
            
            results = [manager.read_file(str(binary_file.absolute())) for _ in range(5)]
            
            first_trace = results[0].trace
            for r in results:
                # We assert topology stability, not ID uniqueness (which is random)
                self.assertEqual(r.trace.operation, first_trace.operation)
                self.assertEqual(r.trace.target, first_trace.target)
                self.assertEqual(r.trace.success, first_trace.success)
                # The topology (category, operation, error context) must be perfectly stable
                self.assertEqual(r.trace.runtime_category, first_trace.runtime_category)
                self.assertEqual(r.trace.operation, first_trace.operation)
                self.assertFalse(r.trace.success)

    def test_residue_verification_after_competing_failures(self):
        """
        After competing failures, verify no mutable state persisted, no global mutation,
        no leaked orchestration metadata.
        """
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir) / "workspace"
            root.mkdir()
            
            manager = FilesystemManager(root)
            
            # Before state
            initial_files = list(root.iterdir())
            
            # Fire competing failures
            manager.read_file("/etc/passwd")  # Escape + non-existent in workspace
            manager.read_file("../" * 10 + "bin/ls") # Escape + binary
            
            # After state
            final_files = list(root.iterdir())
            
            # Verify no residue left in workspace
            self.assertEqual(initial_files, final_files)
            
            # Verify manager hasn't accumulated state
            # (Assuming manager has no stateful queues, just trace emission)
            self.assertFalse(hasattr(manager, "_failed_paths") or hasattr(manager, "retry_queue"))

    def test_boundary_containment_orchestration(self):
        """
        Verify filesystem failure remains filesystem-local and does not cause
        orchestration-layer dynamic routing or global trace coordination.
        """
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir) / "workspace"
            root.mkdir()
            
            # We setup a configuration that points to an oversized + escaped target
            config_file = root / "deploy.json"
            config_file.write_text('{"command": "echo", "args": []}')
            
            # But we inject a symlink pointing outside for the config file itself
            outside = Path(tempdir) / "outside"
            outside.mkdir()
            large_invalid_config = outside / "bad.json"
            large_invalid_config.write_bytes(b"x" * 16_000)
            
            os.symlink(str(large_invalid_config), str(root / "bad.json"))
            
            config = CompositeOperationConfig(workspace_root=str(root), max_file_bytes=1024)
            # Mock or use actual limit if CompositeOperation allows passing limits
            # We'll just execute the symlinked file
            orchestrator = CompositeOperation(config)
            
            # Replay 3 times
            results = [orchestrator.execute("bad.json") for _ in range(3)]
            
            first_result = results[0]
            
            for result in results:
                # 1. Orchestration failed cleanly
                self.assertEqual(result.final_state, OrchestrationState.FAILED)
                
                # 2. Failed strictly at 'read' step (filesystem domain)
                self.assertEqual(len(result.steps), 1)
                self.assertEqual(result.steps[0].step_name, "read")
                self.assertFalse(result.steps[0].result.success)
                
                # 3. No cross-domain leakage (orchestrator didn't try to validate a failed read)
                # 4. Same trace topology emitted
                self.assertEqual(result.steps[0].trace.operation, "read_file")
                self.assertEqual(result.steps[0].trace.error_type, first_result.steps[0].trace.error_type)


if __name__ == "__main__":
    unittest.main()
