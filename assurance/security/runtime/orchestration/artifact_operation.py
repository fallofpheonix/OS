"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
import time
from dataclasses import dataclass

from runtime.filesystem import FilesystemManager
from runtime.tracing.models import create_runtime_trace
from runtime.orchestration.models import (
    OrchestrationStep,
    OrchestrationResult,
    OrchestrationState,
    VerificationResult,
)


@dataclass(frozen=True)
class ArtifactVerificationConfig:
    workspace_root: str


class ArtifactVerificationOperation:
    """
    Read → Verify → Accept orchestration.
    
    This chain exists to prove semantic divergence: it expects a binary artifact.
    It consumes the exact same FilesystemManager, but interprets its failures
    differently than the CompositeOperation.
    """
    def __init__(self, config: ArtifactVerificationConfig):
        self.config = config
        self.filesystem_manager = FilesystemManager(config.workspace_root)

    def execute(self, target_path: str) -> OrchestrationResult:
        start_time = time.time()
        steps = []
        
        # Step 1: Attempt to read the artifact
        read_step = self._step_read_target(target_path)
        steps.append(read_step)
        
        # Divergent Semantics:
        # The filesystem natively rejects binary files.
        # If the read succeeded, it's a text file (Invalid for this chain).
        if read_step.result.success:
            verify_step = self._step_verify(success=False, artifact_type="text")
            steps.append(verify_step)
            return self._build_result(steps, OrchestrationState.FAILED, start_time)
            
        # If the read failed due to binary rejection, it's exactly what we wanted!
        # We do NOT mutate the read_step's trace. It remains a runtime failure.
        if "binary content rejected" in str(read_step.result.error).lower():
            verify_step = self._step_verify(success=True, artifact_type="binary")
            steps.append(verify_step)
            return self._build_result(steps, OrchestrationState.SUCCESS, start_time)
            
        # Any other filesystem failure (not found, escape, etc.) is a true failure
        verify_step = self._step_verify(success=False, artifact_type="unknown")
        steps.append(verify_step)
        return self._build_result(steps, OrchestrationState.FAILED, start_time)

    def _step_read_target(self, target_path: str) -> OrchestrationStep:
        result = self.filesystem_manager.read_file(target_path)
        return OrchestrationStep(
            step_name="read",
            result=result,
            trace=result.trace,
        )

    def _step_verify(self, success: bool, artifact_type: str) -> OrchestrationStep:
        start_time = time.time()
        
        duration = int((time.time() - start_time) * 1000)
        trace = create_runtime_trace(
            runtime_category="runtime.orchestration",
            operation="verify_artifact",
            target=artifact_type,
            duration_ms=duration,
            success=success,
            error_type=None if success else "InvalidArtifactType",
        )
        return OrchestrationStep(
            step_name="verify",
            result=VerificationResult(
                success=success, 
                verified_type=artifact_type,
                error=None if success else f"expected binary, got {artifact_type}"
            ),
            trace=trace,
        )

    def _build_result(self, steps: list[OrchestrationStep], final_state: OrchestrationState, start_time: float) -> OrchestrationResult:
        duration_ms = int((time.time() - start_time) * 1000)
        
        orchestration_trace = create_runtime_trace(
            runtime_category="runtime.orchestration",
            operation="artifact_verification",
            target=f"read→verify ({final_state.value})",
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
