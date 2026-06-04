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
from runtime.shell import ShellExecutor
from runtime.tracing.models import create_runtime_trace
from runtime.orchestration.models import (
    OrchestrationStep,
    OrchestrationResult,
    OrchestrationState,
    VerificationResult,
)


@dataclass(frozen=True)
class AdvancedAnalysisConfig:
    workspace_root: str
    timeout_seconds: float = 5.0


@dataclass(frozen=True, slots=True)
class AnalysisResult:
    """Result of advanced artifact analysis."""
    success: bool
    mime_type: str
    error: str | None = None


class AdvancedArtifactAnalysisOperation:
    """
    Read → Classify → Accept/Reject orchestration.
    
    This chain proves interpretation asymmetry: it requires a much higher density 
    of interpretation than other chains. Instead of just accepting a binary failure, 
    it orchestrates a secondary shell execution to extract metadata and classify the binary.
    
    CRITICALLY: It does this locally, without enriching the FilesystemManager with 
    mime-type extraction or creating global 'semantic adapters'.
    """
    def __init__(self, config: AdvancedAnalysisConfig):
        self.config = config
        self.filesystem_manager = FilesystemManager(config.workspace_root)
        self.shell_executor = ShellExecutor(default_timeout_seconds=config.timeout_seconds)

    def execute(self, target_path: str) -> OrchestrationResult:
        start_time = time.time()
        steps = []
        
        # Step 1: Read Attempt
        read_step = self._step_read_target(target_path)
        steps.append(read_step)
        
        if read_step.result.success:
            # We wanted a binary, this is text.
            analysis_step = self._step_analyze(success=False, mime_type="text/plain")
            steps.append(analysis_step)
            return self._build_result(steps, OrchestrationState.FAILED, start_time)
            
        if "binary content rejected" not in str(read_step.result.error).lower():
            # A true filesystem failure (e.g. missing file)
            analysis_step = self._step_analyze(success=False, mime_type="unknown")
            steps.append(analysis_step)
            return self._build_result(steps, OrchestrationState.FAILED, start_time)
            
        # Step 2: Complex Interpretation via Shell
        # We know it's binary, but we need to know WHAT KIND of binary.
        # We do not ask the filesystem for this; we orchestrate a shell command.
        classify_step = self._step_classify_binary(target_path)
        steps.append(classify_step)
        
        if not classify_step.result.success:
            analysis_step = self._step_analyze(success=False, mime_type="unclassifiable")
            steps.append(analysis_step)
            return self._build_result(steps, OrchestrationState.FAILED, start_time)
            
        output = classify_step.result.stdout.lower()
        
        # Step 3: Deep Semantic Analysis
        # We ONLY accept Mach-O or ELF executables.
        if "mach-o" in output or "elf" in output:
            analysis_step = self._step_analyze(success=True, mime_type="application/x-executable")
            steps.append(analysis_step)
            return self._build_result(steps, OrchestrationState.SUCCESS, start_time)
        else:
            analysis_step = self._step_analyze(success=False, mime_type="application/octet-stream")
            steps.append(analysis_step)
            return self._build_result(steps, OrchestrationState.FAILED, start_time)

    def _step_read_target(self, target_path: str) -> OrchestrationStep:
        result = self.filesystem_manager.read_file(target_path)
        return OrchestrationStep(
            step_name="read",
            result=result,
            trace=result.trace,
        )
        
    def _step_classify_binary(self, target_path: str) -> OrchestrationStep:
        # We must use absolute path for the shell execution
        abs_path = self.filesystem_manager.resolver.resolve(target_path)
        result = self.shell_executor.execute("file", args=["-b", str(abs_path)])
        return OrchestrationStep(
            step_name="classify",
            result=result,
            trace=result.trace,
        )

    def _step_analyze(self, success: bool, mime_type: str) -> OrchestrationStep:
        start_time = time.time()
        
        duration = int((time.time() - start_time) * 1000)
        trace = create_runtime_trace(
            runtime_category="runtime.orchestration",
            operation="advanced_analysis",
            target=mime_type,
            duration_ms=duration,
            success=success,
            error_type=None if success else "ClassificationRejected",
        )
        return OrchestrationStep(
            step_name="analyze",
            result=AnalysisResult(
                success=success, 
                mime_type=mime_type,
                error=None if success else f"rejected mime type: {mime_type}"
            ),
            trace=trace,
        )

    def _build_result(self, steps: list[OrchestrationStep], final_state: OrchestrationState, start_time: float) -> OrchestrationResult:
        duration_ms = int((time.time() - start_time) * 1000)
        
        orchestration_trace = create_runtime_trace(
            runtime_category="runtime.orchestration",
            operation="advanced_analysis_operation",
            target=f"read→classify→analyze ({final_state.value})",
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
