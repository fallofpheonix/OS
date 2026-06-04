"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
"""
Orchestration module: Deterministic bounded composition of runtime operations.

This module implements controlled multi-step orchestration that preserves
trust boundaries, determinism, and auditability from the runtime substrate.

Guarantees:
- Deterministic replay (same input → identical execution)
- Fail-fast semantics (halt on first failure)
- Step-level auditability (each step traceable)
- Immutable results (frozen dataclasses)
- No runtime boundary bypass (uses FilesystemManager and ShellExecutor only)
"""

from runtime.orchestration.models import (
    OrchestrationState,
    OrchestrationStep,
    OrchestrationResult,
    ValidationResult,
    CaptureResult,
    InspectionResult,
    RejectionResult,
    VerificationResult,
    AnalysisResult,
    SymbolResult,
)
from runtime.orchestration.composite_operation import (
    CompositeOperationConfig,
    CompositeOperation,
)
from runtime.orchestration.inspection_operation import (
    InspectionOperationConfig,
    InspectionOperation,
)
from runtime.orchestration.artifact_operation import (
    ArtifactVerificationConfig,
    ArtifactVerificationOperation,
)
from runtime.orchestration.analysis_operation import (
    AdvancedAnalysisConfig,
    AdvancedArtifactAnalysisOperation,
)
from runtime.orchestration.symbol_operation import (
    SymbolInspectionConfig,
    SymbolInspectionOperation,
)

__all__ = [
    "OrchestrationState",
    "OrchestrationStep",
    "OrchestrationResult",
    "ValidationResult",
    "CaptureResult",
    "InspectionResult",
    "RejectionResult",
    "VerificationResult",
    "AnalysisResult",
    "SymbolResult",
    "CompositeOperationConfig",
    "CompositeOperation",
    "InspectionOperationConfig",
    "InspectionOperation",
    "ArtifactVerificationConfig",
    "ArtifactVerificationOperation",
    "AdvancedAnalysisConfig",
    "AdvancedArtifactAnalysisOperation",
    "SymbolInspectionConfig",
    "SymbolInspectionOperation",
]
