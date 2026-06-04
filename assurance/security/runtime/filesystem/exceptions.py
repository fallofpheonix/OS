"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.

Domain-level filesystem runtime exceptions.

These exceptions form a flat, domain-oriented surface that callers
can depend on. They intentionally avoid leaking low-level OS or
library exception types.
"""
from __future__ import annotations


class RuntimeErrorBase(Exception):
    """Base class for runtime domain errors."""


class PathTraversalBlocked(RuntimeErrorBase):
    """Raised when path traversal (e.g., ..) is blocked by policy."""


class WorkspaceBoundaryViolation(RuntimeErrorBase):
    """Raised when a resolved path lies outside the trusted workspace root."""


class SymlinkLoop(RuntimeErrorBase):
    """Raised when a symlink loop is detected during resolution."""


class InvalidPath(RuntimeErrorBase):
    """Raised when the input path is syntactically invalid."""


class UnsupportedTarget(RuntimeErrorBase):
    """Raised when the resolved target is of an unsupported type (e.g., device)."""


class FileTooLarge(RuntimeErrorBase):
    """Raised when an attempted operation would touch a file exceeding size limits."""


class BinaryFileRejected(RuntimeErrorBase):
    """Raised when a binary file is rejected by text-only enforcement."""


class UnsupportedEncoding(RuntimeErrorBase):
    """Raised when a file cannot be decoded using the supported text encoding."""


class DirectoryTooLarge(RuntimeErrorBase):
    """Raised when directory enumeration exceeds the configured limit."""
