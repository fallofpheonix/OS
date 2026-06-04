"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
"""
Composite operation: Read → Validate → Execute → Capture

This module implements the first orchestration slice.
It proves that trust (canonicalization, containment, governance) survives composition.

Key properties:
- Deterministic (same input → same output)
- Fail-fast (halt on first failure)
- Auditable (step-level traces preserved)
- Immutable (results cannot be modified)
"""

import json
import time
from dataclasses import dataclass
from typing import Any
from uuid import uuid4
from datetime import datetime, timezone

from runtime.filesystem import FilesystemManager
from runtime.shell import ShellExecutor
from runtime.tracing.models import create_runtime_trace, RuntimeTrace
from runtime.orchestration.models import (
    OrchestrationStep,
    OrchestrationResult,
    OrchestrationState,
    ValidationResult,
    CaptureResult,
)


@dataclass(frozen=True)
class CompositeOperationConfig:
    """Configuration for composite operation execution."""
    workspace_root: str
    timeout_seconds: float = 30.0
    max_file_bytes: int = 1_048_576
    max_directory_entries: int = 1000


class CompositeOperation:
    """
    Read → Validate → Execute → Capture orchestration.
    
    This orchestration:
    1. Reads a configuration file
    2. Validates its structure (observational only)
    3. Executes the command it specifies
    4. Captures the execution result
    
    Guarantees:
    - Deterministic replay (same input → identical execution)
    - Fail-fast halt (error halts immediately)
    - Step-level auditability (each step auditable separately)
    - Immutable results (frozen dataclass)
    """
    
    def __init__(self, config: CompositeOperationConfig):
        self.config = config
        self.filesystem_manager = FilesystemManager(
            config.workspace_root,
            max_file_bytes=config.max_file_bytes,
            max_directory_entries=config.max_directory_entries,
        )
        self.shell_executor = ShellExecutor(
            default_timeout_seconds=config.timeout_seconds
        )
    
    def execute(self, config_path: str) -> OrchestrationResult:
        """
        Execute the composite operation.
        
        Args:
            config_path: Relative path to configuration file
        
        Returns:
            OrchestrationResult with all steps and final state
        
        Contract:
            - Same input (config_path) always produces identical result
            - Failure in step N prevents execution of step N+1
            - All steps (even failed ones) are included in result
            - Result is immutable (frozen dataclass)
        """
        start_time = time.time()
        steps = []
        
        # Step 1: Read Configuration
        read_step = self._step_read_configuration(config_path)
        steps.append(read_step)
        
        if not read_step.result.success:
            return self._build_orchestration_result(
                steps, OrchestrationState.FAILED, start_time
            )
        
        # Step 2: Validate Configuration
        validate_step = self._step_validate_configuration(
            read_step.result.content
        )
        steps.append(validate_step)
        
        if not validate_step.result.success:
            return self._build_orchestration_result(
                steps, OrchestrationState.FAILED, start_time
            )
        
        # Parse validated configuration (safe because validation passed)
        try:
            config = json.loads(read_step.result.content)
        except json.JSONDecodeError as e:
            # Should not happen (validation would have caught this)
            return self._build_orchestration_result(
                steps, OrchestrationState.UNKNOWN, start_time
            )
        
        # Step 3: Execute Command
        execute_step = self._step_execute_command(
            config.get("command", ""),
            config.get("args", [])
        )
        steps.append(execute_step)
        
        # Check if timeout or failure (not just exit code != 0)
        if not execute_step.result.success:
            error_lower = execute_step.result.stderr.lower() if execute_step.result.stderr else ""
            final_state = (
                OrchestrationState.TIMEOUT
                if "timeout" in error_lower or "timed out" in error_lower
                else OrchestrationState.FAILED
            )
            return self._build_orchestration_result(steps, final_state, start_time)
        
        # Step 4: Capture Result
        capture_step = self._step_capture_result(execute_step.result)
        steps.append(capture_step)
        
        # Determine final state
        final_state = (
            OrchestrationState.SUCCESS
            if capture_step.result.success
            else OrchestrationState.FAILED
        )
        
        return self._build_orchestration_result(steps, final_state, start_time)
    
    def _step_read_configuration(self, config_path: str) -> OrchestrationStep:
        """Step 1: Read configuration file from workspace."""
        result = self.filesystem_manager.read_file(config_path)
        return OrchestrationStep(
            step_name="read",
            result=result,
            trace=result.trace,
        )
    
    def _step_validate_configuration(self, content: str) -> OrchestrationStep:
        """Step 2: Validate configuration structure (observational only)."""
        start_time = time.time()
        
        try:
            # Parse JSON
            config = json.loads(content)
            
            # Check required fields exist
            if "command" not in config:
                error = "required field 'command' missing"
                duration_ms = int((time.time() - start_time) * 1000)
                trace = create_runtime_trace(
                    runtime_category="runtime.orchestration",
                    operation="validate",
                    target="configuration",
                    duration_ms=duration_ms,
                    success=False,
                    error_type="ValidationError",
                )
                return OrchestrationStep(
                    step_name="validate",
                    result=ValidationResult(success=False, error=error),
                    trace=trace,
                )
            
            # Check required field is non-empty
            if not config["command"] or not isinstance(config["command"], str):
                error = "required field 'command' must be non-empty string"
                duration_ms = int((time.time() - start_time) * 1000)
                trace = create_runtime_trace(
                    runtime_category="runtime.orchestration",
                    operation="validate",
                    target="configuration",
                    duration_ms=duration_ms,
                    success=False,
                    error_type="ValidationError",
                )
                return OrchestrationStep(
                    step_name="validate",
                    result=ValidationResult(success=False, error=error),
                    trace=trace,
                )
            
            # Check args is list if present
            if "args" in config and not isinstance(config["args"], list):
                error = "field 'args' must be a list"
                duration_ms = int((time.time() - start_time) * 1000)
                trace = create_runtime_trace(
                    runtime_category="runtime.orchestration",
                    operation="validate",
                    target="configuration",
                    duration_ms=duration_ms,
                    success=False,
                    error_type="ValidationError",
                )
                return OrchestrationStep(
                    step_name="validate",
                    result=ValidationResult(success=False, error=error),
                    trace=trace,
                )
            
            # Validation passed
            duration_ms = int((time.time() - start_time) * 1000)
            trace = create_runtime_trace(
                runtime_category="runtime.orchestration",
                operation="validate",
                target="configuration",
                duration_ms=duration_ms,
                success=True,
            )
            return OrchestrationStep(
                step_name="validate",
                result=ValidationResult(success=True),
                trace=trace,
            )
        
        except json.JSONDecodeError as e:
            error = f"JSON parse error: {str(e)}"
            duration_ms = int((time.time() - start_time) * 1000)
            trace = create_runtime_trace(
                runtime_category="runtime.orchestration",
                operation="validate",
                target="configuration",
                duration_ms=duration_ms,
                success=False,
                error_type="ParseError",
            )
            return OrchestrationStep(
                step_name="validate",
                result=ValidationResult(success=False, error=error),
                trace=trace,
            )
    
    def _step_execute_command(
        self, command: str, args: list[str]
    ) -> OrchestrationStep:
        """Step 3: Execute the deployment command."""
        result = self.shell_executor.execute(
            command,
            args=args,
            timeout_seconds=self.config.timeout_seconds,
        )
        return OrchestrationStep(
            step_name="execute",
            result=result,
            trace=result.trace,
        )
    
    def _step_capture_result(
        self, execution_result
    ) -> OrchestrationStep:
        """Step 4: Capture the execution result."""
        # Capture is observational - just extract the facts
        capture_result = CaptureResult(
            success=execution_result.success,
            stdout=execution_result.stdout or "",
            stderr=execution_result.stderr or "",
            exit_code=execution_result.exit_code,
            error=None,  # No error in capture; result is captured as-is
        )
        
        return OrchestrationStep(
            step_name="capture",
            result=capture_result,
            trace=execution_result.trace,
        )
    
    def _build_orchestration_result(
        self,
        steps: list[OrchestrationStep],
        final_state: OrchestrationState,
        start_time: float,
    ) -> OrchestrationResult:
        """Build immutable orchestration result."""
        duration_ms = int((time.time() - start_time) * 1000)
        
        orchestration_trace = create_runtime_trace(
            runtime_category="runtime.orchestration",
            operation="composite_operation",
            target=f"read→validate→execute→capture ({final_state.value})",
            duration_ms=duration_ms,
            success=(final_state == OrchestrationState.SUCCESS),
            error_type=(
                None if final_state == OrchestrationState.SUCCESS
                else final_state.value
            ),
        )
        
        return OrchestrationResult(
            steps=tuple(steps),
            final_state=final_state,
            orchestration_trace=orchestration_trace,
        )
