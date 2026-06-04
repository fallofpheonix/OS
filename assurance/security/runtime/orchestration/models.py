"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
"""
Orchestration result models.

These models preserve step-level auditability and immutability.
Each step's trace is independent and non-collapsible.
"""

from dataclasses import dataclass
from enum import Enum
from typing import Any
from runtime.filesystem.models import FileOperationResult
from runtime.shell.models import ExecutionResult
from runtime.tracing.models import RuntimeTrace


class OrchestrationState(Enum):
    """Final state of orchestration execution."""
    SUCCESS = "success"      # all steps executed successfully
    FAILED = "failed"        # a step failed, orchestration halted
    TIMEOUT = "timeout"      # a step timed out, orchestration halted
    UNKNOWN = "unknown"      # orchestration could not complete (internal error)


@dataclass(frozen=True, slots=True)
class ValidationResult:
    """Result of configuration validation step."""
    success: bool
    error: str | None = None
    
    def __post_init__(self):
        if self.success and self.error is not None:
            raise ValueError("success=True cannot have error message")
        if not self.success and self.error is None:
            raise ValueError("success=False requires error message")


@dataclass(frozen=True, slots=True)
class CaptureResult:
    """Result of execution capture step."""
    success: bool
    stdout: str = ""
    stderr: str = ""
    exit_code: int | None = None
    error: str | None = None
    
    def __post_init__(self):
        if self.success and self.error is not None:
            raise ValueError("success=True cannot have error message")


@dataclass(frozen=True, slots=True)
class InspectionResult:
    """Result of configuration inspection step."""
    success: bool
    findings: list[str]
    error: str | None = None


@dataclass(frozen=True, slots=True)
class RejectionResult:
    """Result of configuration rejection step."""
    success: bool
    reason: str
    error: str | None = None


@dataclass(frozen=True, slots=True)
class VerificationResult:
    """Result of artifact verification step."""
    success: bool
    verified_type: str
    error: str | None = None


@dataclass(frozen=True, slots=True)
class AnalysisResult:
    """Result of advanced artifact analysis."""
    success: bool
    mime_type: str
    error: str | None = None


@dataclass(frozen=True, slots=True)
class SymbolResult:
    """Result of symbol inspection."""
    success: bool
    symbols_found: list[str]
    error: str | None = None


@dataclass(frozen=True, slots=True)
class OrchestrationStep:
    """
    Single step in an orchestration.
    
    Properties:
    - step_name: identifier
    - result: operation result (domain-specific)
    - trace: immutable runtime trace
    
    Each step remains individually auditable.
    No step information is collapsed or summarized.
    """
    step_name: str  # "read", "validate", "execute", "capture", "inspect", "reject", "verify", "classify", "analyze", "nm_inspect", "symbol_verify"
    result: FileOperationResult | ValidationResult | ExecutionResult | CaptureResult | InspectionResult | RejectionResult | VerificationResult | AnalysisResult | SymbolResult
    trace: RuntimeTrace


@dataclass(frozen=True, slots=True)
class OrchestrationResult:
    """
    Complete orchestration result.
    
    Properties:
    - steps: all steps executed (including failed ones)
    - final_state: SUCCESS, FAILED, TIMEOUT, UNKNOWN
    - orchestration_trace: aggregated trace
    
    Contracts:
    - All steps are traceable
    - Partial failure traces are complete
    - No step information is lost
    - Deterministic replay possible (same input → same steps)
    """
    steps: tuple[OrchestrationStep, ...]
    final_state: OrchestrationState
    orchestration_trace: RuntimeTrace
    
    def __post_init__(self):
        # Validate step structure
        if not self.steps:
            raise ValueError("orchestration result must have at least one step")
        
        # Validate final state consistency
        if self.final_state == OrchestrationState.SUCCESS:
            # We no longer strictly require all steps to have succeeded,
            # as some orchestration chains (like ArtifactVerification) 
            # might intentionally consume and handle a step failure.
            pass
        
        if self.final_state == OrchestrationState.FAILED:
            # Failed requires at least one step to have failed
            has_failure = any(not step.result.success for step in self.steps)
            if not has_failure:
                raise ValueError(
                    "final_state=FAILED but all steps succeeded"
                )
