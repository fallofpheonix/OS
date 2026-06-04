"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
from pathlib import Path
from time import perf_counter_ns

from infrastructure.logging.logger import get_logger
from runtime.filesystem.models import FileOperationResult
from runtime.filesystem import policy
from runtime.filesystem.resolver import Resolver
from runtime.filesystem import exceptions
from runtime.tracing import create_runtime_trace

LOGGER = get_logger("runtime.filesystem")


class FilesystemManager:
    def __init__(self, root_path: str | Path, *, max_file_bytes: int = 1_048_576, max_directory_entries: int = 1_000):
        self.root_path = Path(root_path).resolve()
        self.resolver = Resolver(self.root_path)
        self.resource_limits = policy.ResourceLimits(
            max_file_bytes=max_file_bytes,
            max_directory_entries=max_directory_entries,
        )

    def read_file(self, requested_path: str) -> FileOperationResult:
        started_at = perf_counter_ns()
        operation = "read_file"
        try:
            path = self.resolver.resolve(requested_path)
            self.resolver.ensure_within_workspace(path)
            if not path.is_file():
                return self._result(operation, requested_path, False, started_at, error="file not found")

            content = policy.read_text_file(path, self.resource_limits)
            return self._result(operation, requested_path, True, started_at, content=content, exists=True)
        except exceptions.RuntimeErrorBase as error:
            return self._result(operation, requested_path, False, started_at, error=str(error))
        except Exception as error:
            return self._result(operation, requested_path, False, started_at, error=str(error))

    def list_directory(self, requested_path: str = ".") -> FileOperationResult:
        started_at = perf_counter_ns()
        operation = "list_directory"
        try:
            path = self.resolver.resolve(requested_path)
            self.resolver.ensure_within_workspace(path)
            if not path.exists() or not path.is_dir():
                return self._result(operation, requested_path, False, started_at, error="directory not found")

            entries = policy.list_directory_entries(path, self.resource_limits)
            return self._result(operation, requested_path, True, started_at, entries=entries, exists=True)
        except exceptions.RuntimeErrorBase as error:
            return self._result(operation, requested_path, False, started_at, error=str(error))
        except Exception as error:
            return self._result(operation, requested_path, False, started_at, error=str(error))

    def exists(self, requested_path: str) -> FileOperationResult:
        started_at = perf_counter_ns()
        operation = "exists"
        try:
            path = self.resolver.resolve(requested_path)
            self.resolver.ensure_within_workspace(path)
            exists = path.exists()
            return self._result(operation, requested_path, True, started_at, exists=exists)
        except exceptions.RuntimeErrorBase as error:
            return self._result(operation, requested_path, False, started_at, error=str(error))
        except Exception as error:
            return self._result(operation, requested_path, False, started_at, error=str(error))

    def _result(
        self,
        operation: str,
        requested_path: str,
        success: bool,
        started_at: int,
        *,
        exists: bool = False,
        content: str = "",
        entries: list[str] | None = None,
        error: str = "",
    ) -> FileOperationResult:
        duration_ms = max(0, (perf_counter_ns() - started_at) // 1_000_000)
        result = FileOperationResult(
            operation=operation,
            path=requested_path,
            success=success,
            duration_ms=duration_ms,
            exists=exists,
            content=content,
            entries=entries or [],
            error=error,
            trace=create_runtime_trace(
                runtime_category="filesystem",
                operation=operation,
                target=requested_path,
                duration_ms=duration_ms,
                success=success,
                error_type=error or "",
            ),
        )
        LOGGER.info(
            "filesystem operation",
            extra={
                "subsystem": "runtime.filesystem",
                "operation": operation,
                "path": requested_path,
                "duration_ms": result.duration_ms,
                "success": result.success,
            },
        )
        return result
