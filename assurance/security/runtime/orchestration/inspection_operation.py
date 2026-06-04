"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
import json
import time
from dataclasses import dataclass

from runtime.filesystem import FilesystemManager
from runtime.tracing.models import create_runtime_trace
from runtime.orchestration.models import (
    OrchestrationStep,
    OrchestrationResult,
    OrchestrationState,
    InspectionResult,
    RejectionResult,
)


@dataclass(frozen=True)
class InspectionOperationConfig:
    workspace_root: str
    max_file_bytes: int = 1_048_576
    max_directory_entries: int = 1000


class InspectionOperation:
    """
    Read → Inspect → Reject orchestration.
    
    This is an independent orchestration chain designed to prove
    that multiple semantic pipelines can coexist without framework collapse
    or shared semantic gravity.
    """
    def __init__(self, config: InspectionOperationConfig):
        self.config = config
        self.filesystem_manager = FilesystemManager(
            config.workspace_root,
            max_file_bytes=config.max_file_bytes,
            max_directory_entries=config.max_directory_entries,
        )

    def execute(self, target_path: str) -> OrchestrationResult:
        start_time = time.time()
        steps = []
        
        # Step 1: Read target
        read_step = self._step_read_target(target_path)
        steps.append(read_step)
        
        if not read_step.result.success:
            return self._build_result(steps, OrchestrationState.FAILED, start_time)
            
        # Step 2: Inspect Content
        inspect_step = self._step_inspect_content(read_step.result.content)
        steps.append(inspect_step)
        
        if not inspect_step.result.success:
            return self._build_result(steps, OrchestrationState.FAILED, start_time)
            
        # Step 3: Reject based on inspection findings
        # In this chain, a successful inspection results in a rejection
        reject_step = self._step_reject(inspect_step.result.findings)
        steps.append(reject_step)
        
        return self._build_result(steps, OrchestrationState.SUCCESS, start_time)

    def _step_read_target(self, target_path: str) -> OrchestrationStep:
        result = self.filesystem_manager.read_file(target_path)
        return OrchestrationStep(
            step_name="read",
            result=result,
            trace=result.trace,
        )

    def _step_inspect_content(self, content: str) -> OrchestrationStep:
        start_time = time.time()
        try:
            data = json.loads(content)
            findings = []
            if "malicious" in data:
                findings.append("malicious payload detected")
            if "dangerous_flag" in data and data["dangerous_flag"] is True:
                findings.append("dangerous flag is set")
                
            duration = int((time.time() - start_time) * 1000)
            trace = create_runtime_trace(
                runtime_category="runtime.orchestration",
                operation="inspect",
                target="configuration_payload",
                duration_ms=duration,
                success=True,
            )
            return OrchestrationStep(
                step_name="inspect",
                result=InspectionResult(success=True, findings=findings),
                trace=trace,
            )
        except json.JSONDecodeError as e:
            duration = int((time.time() - start_time) * 1000)
            trace = create_runtime_trace(
                runtime_category="runtime.orchestration",
                operation="inspect",
                target="configuration_payload",
                duration_ms=duration,
                success=False,
                error_type="ParseError",
            )
            return OrchestrationStep(
                step_name="inspect",
                result=InspectionResult(success=False, findings=[], error=str(e)),
                trace=trace,
            )

    def _step_reject(self, findings: list[str]) -> OrchestrationStep:
        start_time = time.time()
        
        reason = "rejection based on findings: " + ", ".join(findings) if findings else "default rejection (inspect -> reject chain)"
        
        duration = int((time.time() - start_time) * 1000)
        trace = create_runtime_trace(
            runtime_category="runtime.orchestration",
            operation="reject",
            target="configuration_payload",
            duration_ms=duration,
            success=True,
        )
        return OrchestrationStep(
            step_name="reject",
            result=RejectionResult(success=True, reason=reason),
            trace=trace,
        )

    def _build_result(self, steps: list[OrchestrationStep], final_state: OrchestrationState, start_time: float) -> OrchestrationResult:
        duration_ms = int((time.time() - start_time) * 1000)
        
        orchestration_trace = create_runtime_trace(
            runtime_category="runtime.orchestration",
            operation="inspection_operation",
            target=f"read→inspect→reject ({final_state.value})",
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
