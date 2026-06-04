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

from runtime.orchestration import (
    CompositeOperation, 
    CompositeOperationConfig,
    ArtifactVerificationOperation,
    ArtifactVerificationConfig,
    OrchestrationState,
)


class SemanticDivergenceTests(unittest.TestCase):
    """
    Proves that semantic divergence can exist under shared runtime pressure.
    Two orchestration chains consume identical filesystem trace failures
    but resolve them differently without requiring a global interpretation layer.
    """

    def setUp(self):
        self.tempdir = tempfile.TemporaryDirectory()
        self.root = Path(self.tempdir.name) / "workspace"
        self.root.mkdir()
        
        # We need a text file and a binary file
        self.text_file = self.root / "config.json"
        self.text_file.write_text('{"command": "echo", "args": ["test"]}')
        
        self.binary_file = self.root / "artifact.o"
        self.binary_file.write_bytes(b"\x00\x01\x02\x03binarydata")

    def tearDown(self):
        self.tempdir.cleanup()

    def test_semantic_divergence_on_binary_rejection(self):
        """
        Verify that Chain A (Composite) and Chain B (ArtifactVerification)
        consume the EXACT same runtime failure trace, but interpret it divergently.
        """
        comp_config = CompositeOperationConfig(workspace_root=str(self.root))
        comp_op = CompositeOperation(comp_config)
        
        artifact_config = ArtifactVerificationConfig(workspace_root=str(self.root))
        artifact_op = ArtifactVerificationOperation(artifact_config)
        
        # Both read the binary file
        comp_result = comp_op.execute("artifact.o")
        artifact_result = artifact_op.execute("artifact.o")
        
        # 1. The filesystem MUST remain semantically neutral (it fails both reads)
        comp_read_trace = comp_result.steps[0].trace
        artifact_read_trace = artifact_result.steps[0].trace
        
        # Both traces must be identical filesystem rejections
        self.assertFalse(comp_read_trace.success)
        self.assertFalse(artifact_read_trace.success)
        self.assertEqual(comp_read_trace.runtime_category, "filesystem")
        self.assertEqual(artifact_read_trace.runtime_category, "filesystem")
        self.assertEqual(comp_read_trace.error_type, artifact_read_trace.error_type)
        
        # Both must explicitly cite binary rejection in their result error
        self.assertIn("binary", str(comp_result.steps[0].result.error).lower())
        self.assertIn("binary", str(artifact_result.steps[0].result.error).lower())
        
        # 2. Semantic Divergence without framework mutation
        # Chain A interprets the read failure as an Orchestration Failure
        self.assertEqual(comp_result.final_state, OrchestrationState.FAILED)
        self.assertEqual(len(comp_result.steps), 1)
        
        # Chain B interprets the read failure as an Orchestration Success!
        # (Because it verified the artifact is indeed binary)
        self.assertEqual(artifact_result.final_state, OrchestrationState.SUCCESS)
        self.assertEqual(len(artifact_result.steps), 2)
        
        # Chain B's verify trace must exist and be successful
        verify_trace = artifact_result.steps[1].trace
        self.assertTrue(verify_trace.success)
        self.assertEqual(verify_trace.operation, "verify_artifact")

    def test_semantic_divergence_on_text_success(self):
        """
        Verify the inverse case: a text file succeeds in the runtime,
        but fails in Chain B's semantic interpretation.
        """
        comp_config = CompositeOperationConfig(workspace_root=str(self.root))
        comp_op = CompositeOperation(comp_config)
        
        artifact_config = ArtifactVerificationConfig(workspace_root=str(self.root))
        artifact_op = ArtifactVerificationOperation(artifact_config)
        
        # Both read the text file
        comp_result = comp_op.execute("config.json")
        artifact_result = artifact_op.execute("config.json")
        
        # 1. The filesystem natively accepts it
        comp_read_trace = comp_result.steps[0].trace
        artifact_read_trace = artifact_result.steps[0].trace
        
        self.assertTrue(comp_read_trace.success)
        self.assertTrue(artifact_read_trace.success)
        
        # 2. Semantic Divergence
        # Chain A proceeds successfully to completion
        self.assertEqual(comp_result.final_state, OrchestrationState.SUCCESS)
        
        # Chain B interprets text success as an Orchestration Failure!
        self.assertEqual(artifact_result.final_state, OrchestrationState.FAILED)
        self.assertEqual(len(artifact_result.steps), 2)
        
        verify_trace = artifact_result.steps[1].trace
        self.assertFalse(verify_trace.success)
        self.assertEqual(verify_trace.error_type, "InvalidArtifactType")


if __name__ == "__main__":
    unittest.main()
