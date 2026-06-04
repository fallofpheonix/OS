"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
"""
Adversarial orchestration tests: Determinism, Failure Propagation, Auditability.

These tests prove that trust (canonicalization, containment, governance)
survives composition through the orchestration layer.

Test Categories:
1. Deterministic Replay - Same input produces identical execution
2. Failure Propagation - Failures halt deterministically
3. Step Identity - Each step remains individually auditable
4. Trace Immutability - Orchestration traces cannot be modified
"""

import json
import tempfile
import unittest
from pathlib import Path

from runtime.orchestration import (
    CompositeOperation,
    CompositeOperationConfig,
    OrchestrationState,
)


class DeterministicReplayTests(unittest.TestCase):
    """Verify that orchestration produces identical results across runs."""
    
    def test_success_path_deterministic_replay(self):
        """Same successful orchestration always produces identical result."""
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            
            # Create valid configuration
            config_file = root / "deploy.json"
            config_file.write_text(json.dumps({
                "command": "echo",
                "args": ["hello"],
            }))
            
            config = CompositeOperationConfig(workspace_root=str(root))
            orchestrator = CompositeOperation(config)
            
            # Execute twice
            result1 = orchestrator.execute("deploy.json")
            result2 = orchestrator.execute("deploy.json")
            
            # Verify identical execution
            self.assertEqual(result1.final_state, result2.final_state)
            self.assertEqual(len(result1.steps), len(result2.steps))
            self.assertEqual(result1.final_state, OrchestrationState.SUCCESS)
            
            # Verify step sequence is identical
            for step1, step2 in zip(result1.steps, result2.steps):
                self.assertEqual(step1.step_name, step2.step_name)
                self.assertEqual(step1.result.success, step2.result.success)
                # Traces have different timestamps but same operation structure
                self.assertEqual(
                    step1.trace.operation,
                    step2.trace.operation,
                )
    
    def test_read_failure_deterministic_replay(self):
        """Same read failure always produces identical result."""
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            
            config = CompositeOperationConfig(workspace_root=str(root))
            orchestrator = CompositeOperation(config)
            
            # Attempt to read non-existent file twice
            result1 = orchestrator.execute("missing.json")
            result2 = orchestrator.execute("missing.json")
            
            # Verify identical failure
            self.assertEqual(result1.final_state, result2.final_state)
            self.assertEqual(result1.final_state, OrchestrationState.FAILED)
            
            # Verify same step halted
            self.assertEqual(len(result1.steps), len(result2.steps))
            self.assertEqual(result1.steps[0].step_name, "read")
            self.assertFalse(result1.steps[0].result.success)
            
            # Verify error messages are identical
            self.assertEqual(
                result1.steps[0].result.error,
                result2.steps[0].result.error,
            )
    
    def test_validation_failure_deterministic_replay(self):
        """Same validation failure always produces identical result."""
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            
            # Create invalid configuration (missing 'command')
            config_file = root / "invalid.json"
            config_file.write_text(json.dumps({"args": []}))
            
            config = CompositeOperationConfig(workspace_root=str(root))
            orchestrator = CompositeOperation(config)
            
            # Execute twice
            result1 = orchestrator.execute("invalid.json")
            result2 = orchestrator.execute("invalid.json")
            
            # Verify identical failure
            self.assertEqual(result1.final_state, result2.final_state)
            self.assertEqual(result1.final_state, OrchestrationState.FAILED)
            
            # Verify validation step failed
            self.assertEqual(len(result1.steps), 2)  # read + validate
            self.assertEqual(result1.steps[1].step_name, "validate")
            self.assertFalse(result1.steps[1].result.success)
            
            # Verify error messages are identical
            self.assertEqual(
                result1.steps[1].result.error,
                result2.steps[1].result.error,
            )


class FailurePropagationTests(unittest.TestCase):
    """Verify that failures halt deterministically without executing subsequent steps."""
    
    def test_read_failure_prevents_subsequent_steps(self):
        """If read fails, validate/execute/capture are not executed."""
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            
            config = CompositeOperationConfig(workspace_root=str(root))
            orchestrator = CompositeOperation(config)
            
            # Read missing file
            result = orchestrator.execute("missing.json")
            
            # Verify only read step was executed
            self.assertEqual(len(result.steps), 1)
            self.assertEqual(result.steps[0].step_name, "read")
            self.assertFalse(result.steps[0].result.success)
            self.assertEqual(result.final_state, OrchestrationState.FAILED)
    
    def test_validation_failure_prevents_execution(self):
        """If validation fails, execute and capture are not executed."""
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            
            # Create invalid configuration
            config_file = root / "invalid.json"
            config_file.write_text(json.dumps({"args": ["--flag"]}))
            
            config = CompositeOperationConfig(workspace_root=str(root))
            orchestrator = CompositeOperation(config)
            
            # Execute orchestration
            result = orchestrator.execute("invalid.json")
            
            # Verify only read and validate were executed
            self.assertEqual(len(result.steps), 2)
            self.assertEqual(result.steps[0].step_name, "read")
            self.assertEqual(result.steps[1].step_name, "validate")
            self.assertTrue(result.steps[0].result.success)
            self.assertFalse(result.steps[1].result.success)
            self.assertEqual(result.final_state, OrchestrationState.FAILED)
    
    def test_execution_failure_prevents_capture(self):
        """If execution fails, capture is not executed."""
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            
            # Create valid configuration but for non-existent command
            config_file = root / "deploy.json"
            config_file.write_text(json.dumps({
                "command": "this_command_does_not_exist_12345",
                "args": [],
            }))
            
            config = CompositeOperationConfig(workspace_root=str(root))
            orchestrator = CompositeOperation(config)
            
            # Execute orchestration
            result = orchestrator.execute("deploy.json")
            
            # Verify read, validate, and execute were attempted
            # But not capture (no 4th step)
            self.assertLessEqual(len(result.steps), 3)
            self.assertEqual(result.steps[0].step_name, "read")
            self.assertEqual(result.steps[1].step_name, "validate")
            self.assertEqual(result.steps[2].step_name, "execute")
            self.assertFalse(result.steps[2].result.success)
            self.assertEqual(result.final_state, OrchestrationState.FAILED)


class StepIdentityTests(unittest.TestCase):
    """Verify that each step remains individually auditable."""
    
    def test_success_path_preserves_step_identity(self):
        """All steps in success path are individually identifiable."""
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            
            config_file = root / "deploy.json"
            config_file.write_text(json.dumps({
                "command": "echo",
                "args": ["test"],
            }))
            
            config = CompositeOperationConfig(workspace_root=str(root))
            orchestrator = CompositeOperation(config)
            
            result = orchestrator.execute("deploy.json")
            
            # Verify all 4 steps are present
            self.assertEqual(len(result.steps), 4)
            
            # Verify step sequence
            expected_steps = ["read", "validate", "execute", "capture"]
            for step, expected_name in zip(result.steps, expected_steps):
                self.assertEqual(step.step_name, expected_name)
                self.assertIsNotNone(step.result)
                self.assertIsNotNone(step.trace)
    
    def test_partial_failure_preserves_step_identity(self):
        """Even in failure, all executed steps are individually auditable."""
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            
            # Create invalid configuration
            config_file = root / "invalid.json"
            config_file.write_text(json.dumps({"description": "no command"}))
            
            config = CompositeOperationConfig(workspace_root=str(root))
            orchestrator = CompositeOperation(config)
            
            result = orchestrator.execute("invalid.json")
            
            # Verify 2 steps executed: read (success), validate (failure)
            self.assertEqual(len(result.steps), 2)
            self.assertEqual(result.steps[0].step_name, "read")
            self.assertEqual(result.steps[1].step_name, "validate")
            
            # Verify each step's result is preserved
            self.assertTrue(result.steps[0].result.success)
            self.assertFalse(result.steps[1].result.success)
            
            # Verify each step's trace is independent
            self.assertIsNotNone(result.steps[0].trace)
            self.assertIsNotNone(result.steps[1].trace)
            self.assertNotEqual(
                result.steps[0].trace.trace_id,
                result.steps[1].trace.trace_id,
            )


class TraceImmutabilityTests(unittest.TestCase):
    """Verify that orchestration traces cannot be modified after creation."""
    
    def test_orchestration_result_immutable(self):
        """OrchestrationResult is frozen and cannot be modified."""
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            
            config_file = root / "deploy.json"
            config_file.write_text(json.dumps({
                "command": "echo",
                "args": ["frozen"],
            }))
            
            config = CompositeOperationConfig(workspace_root=str(root))
            orchestrator = CompositeOperation(config)
            
            result = orchestrator.execute("deploy.json")
            
            # Attempt to modify result (should fail)
            with self.assertRaises((AttributeError, TypeError)):
                result.final_state = OrchestrationState.FAILED
            
            with self.assertRaises((AttributeError, TypeError)):
                result.steps = ()
    
    def test_orchestration_step_immutable(self):
        """OrchestrationStep is frozen and cannot be modified."""
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            
            config_file = root / "deploy.json"
            config_file.write_text(json.dumps({
                "command": "echo",
                "args": ["immutable"],
            }))
            
            config = CompositeOperationConfig(workspace_root=str(root))
            orchestrator = CompositeOperation(config)
            
            result = orchestrator.execute("deploy.json")
            step = result.steps[0]
            
            # Attempt to modify step (should fail)
            with self.assertRaises((AttributeError, TypeError)):
                step.step_name = "invalid"
            
            with self.assertRaises((AttributeError, TypeError)):
                step.result = None
    
    def test_orchestration_trace_immutable(self):
        """RuntimeTrace in orchestration is frozen."""
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            
            config_file = root / "deploy.json"
            config_file.write_text(json.dumps({
                "command": "echo",
                "args": ["trace"],
            }))
            
            config = CompositeOperationConfig(workspace_root=str(root))
            orchestrator = CompositeOperation(config)
            
            result = orchestrator.execute("deploy.json")
            trace = result.orchestration_trace
            
            # Attempt to modify trace (should fail)
            with self.assertRaises((AttributeError, TypeError)):
                trace.success = False


if __name__ == "__main__":
    unittest.main()
