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
from unittest.mock import patch

from runtime.filesystem.manager import FilesystemManager
from runtime.orchestration import CompositeOperation, CompositeOperationConfig, OrchestrationState


class FailureOrderPerturbationTests(unittest.TestCase):
    """
    Proves that semantic containment remains identical even if invalidity
    discovery order changes internally. 
    
    Verifies there is no "semantic precedence leakage" (e.g. orchestrator
    treating "size limit" fundamentally differently than "containment breach").
    """

    def setUp(self):
        self.tempdir = tempfile.TemporaryDirectory()
        self.root = Path(self.tempdir.name) / "workspace"
        self.root.mkdir()
        
        self.outside = Path(self.tempdir.name) / "outside"
        self.outside.mkdir()
        
        # Create a payload that violates THREE boundaries simultaneously:
        # 1. Outside workspace (containment escape)
        # 2. Oversized (> 1024 bytes)
        # 3. Binary (contains null bytes)
        self.ultimate_invalid_payload = self.outside / "bad.bin"
        self.ultimate_invalid_payload.write_bytes(b"hello\x00world" + b"x" * 2000)
        
        # Target symlink in workspace
        self.target = self.root / "target"
        os.symlink(str(self.ultimate_invalid_payload), str(self.target))

    def tearDown(self):
        self.tempdir.cleanup()

    def test_semantic_containment_survives_perturbation(self):
        """
        Orchestrator must handle all perturbed detection orders identically.
        No special casing or dynamic trace mutation based on WHICH filesystem
        error was discovered first.
        """
        config = CompositeOperationConfig(workspace_root=str(self.root))
        
        # We need a low size limit so our payload trips the size check
        original_init = FilesystemManager.__init__
        def custom_init(self_mgr, root_path, **kwargs):
            original_init(self_mgr, root_path, max_file_bytes=1024)
            
        with patch.object(FilesystemManager, '__init__', custom_init):
            # -------------------------------------------------------------
            # Case A: Standard detection order (typically Containment first)
            # -------------------------------------------------------------
            orchestrator_a = CompositeOperation(config)
            result_a = orchestrator_a.execute("target")
            
            # -------------------------------------------------------------
            # Case B: Size detected first
            # -------------------------------------------------------------
            original_read = FilesystemManager.read_file
            
            def read_detect_size_first(self_mgr, requested_path: str):
                raw_path = Path(self_mgr.root_path) / requested_path
                if raw_path.exists() and raw_path.stat().st_size > self_mgr.resource_limits.max_file_bytes:
                    return self_mgr._result("read_file", requested_path, False, 0, error="perturbed: size limit exceeded")
                return original_read(self_mgr, requested_path)

            orchestrator_b = CompositeOperation(config)
            with patch.object(FilesystemManager, 'read_file', read_detect_size_first):
                result_b = orchestrator_b.execute("target")
                
            # -------------------------------------------------------------
            # Case C: Binary detected first
            # -------------------------------------------------------------
            def read_detect_binary_first(self_mgr, requested_path: str):
                raw_path = Path(self_mgr.root_path) / requested_path
                if raw_path.exists() and b"\x00" in raw_path.read_bytes()[:1024]:
                    return self_mgr._result("read_file", requested_path, False, 0, error="perturbed: binary content rejected")
                return original_read(self_mgr, requested_path)

            orchestrator_c = CompositeOperation(config)
            with patch.object(FilesystemManager, 'read_file', read_detect_binary_first):
                result_c = orchestrator_c.execute("target")
            
            # =============================================================
            # ASSERTIONS
            # =============================================================
            
            # First, verify the perturbation actually worked (errors must differ)
            error_a = result_a.steps[0].result.error
            error_b = result_b.steps[0].result.error
            error_c = result_c.steps[0].result.error
            
            self.assertNotEqual(error_a, error_b)
            self.assertNotEqual(error_b, error_c)
            self.assertNotEqual(error_a, error_c)
            
            # Now, verify semantic topology remains STRICTLY IDENTICAL
            for result in [result_a, result_b, result_c]:
                # 1. Same halt layer
                self.assertEqual(result.final_state, OrchestrationState.FAILED)
                self.assertEqual(len(result.steps), 1)
                
                step = result.steps[0]
                self.assertEqual(step.step_name, "read")
                self.assertFalse(step.result.success)
                
                # 2. Same trace topology (operation, category, and success status)
                self.assertEqual(step.trace.operation, "read_file")
                self.assertEqual(step.trace.runtime_category, "filesystem")
                self.assertFalse(step.trace.success)


if __name__ == "__main__":
    unittest.main()
