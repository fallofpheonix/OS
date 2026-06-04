"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
import shlex
import subprocess
from time import perf_counter_ns

from infrastructure.logging.logger import get_logger
from runtime.shell.models import ExecutionResult
from runtime.tracing import create_runtime_trace

LOGGER = get_logger("runtime.shell")

COMMAND_POLICY = {
    "pwd": {"max_args": 0},
    "ls": {"max_args": 2},
    "echo": {"max_length": 200},
    "sleep": {"max_args": 1},
    "file": {"max_args": 5},
    "nm": {"max_args": 5},
}


class ShellExecutor:
    """
    Shell command executor with configurable timeouts.
    
    Uses functional execution underneath but provides a class-based
    interface compatible with control-plane layer.
    """
    
    def __init__(self, default_timeout_seconds: float = 30.0):
        self.default_timeout_seconds = default_timeout_seconds
    
    def execute(
        self,
        command: str,
        args: list[str] | None = None,
        timeout_seconds: float | None = None,
    ) -> ExecutionResult:
        """
        Execute a shell command.
        
        Args:
            command: Command name (e.g., "echo", "ls", "pwd")
            args: Command arguments (optional)
            timeout_seconds: Command timeout in seconds (uses default if None)
        
        Returns:
            ExecutionResult with exit code, stdout, stderr
        """
        if args is None:
            args = []
        
        if timeout_seconds is None:
            timeout_seconds = self.default_timeout_seconds
        
        # Build command string from command + args
        command_str = command
        if args:
            command_str = f"{command} {' '.join(shlex.quote(arg) for arg in args)}"
        
        # Use functional executor
        return execute(command_str, timeout=timeout_seconds)


def _build_result(command: str, args: list[str], success: bool, stdout: str, stderr: str, exit_code: int, duration_ns: int) -> ExecutionResult:
    duration_ms = max(0, duration_ns // 1_000_000)
    return ExecutionResult(
        command=command,
        args=args,
        success=success,
        stdout=stdout,
        stderr=stderr,
        exit_code=exit_code,
        duration_ms=duration_ms,
        trace=create_runtime_trace(
            runtime_category="shell",
            operation="execute",
            target=command,
            duration_ms=duration_ms,
            success=success,
            error_type=stderr if not success else "",
        ),
    )


def execute(command: str, timeout: float = 10.0) -> ExecutionResult:
    """Execute a controlled shell command and return a structured response."""
    started_at = perf_counter_ns()
    if not command or not command.strip():
        return _build_result(command, [], False, "", "empty command", 127, perf_counter_ns() - started_at)

    parts = shlex.split(command)
    if not parts:
        return _build_result(command, [], False, "", "failed to parse command", 127, perf_counter_ns() - started_at)

    command_name = parts[0]
    args = parts[1:]
    policy = COMMAND_POLICY.get(command_name)
    if policy is None:
        return _build_result(command, args, False, "", f"command not allowed: {command_name}", 126, perf_counter_ns() - started_at)

    if "max_args" in policy and len(args) > policy["max_args"]:
        return _build_result(command, args, False, "", f"too many arguments for {command_name}", 126, perf_counter_ns() - started_at)

    if "max_length" in policy and len(command) > policy["max_length"]:
        return _build_result(command, args, False, "", f"command too long for {command_name}", 126, perf_counter_ns() - started_at)

    try:
        proc = subprocess.run(parts, capture_output=True, text=True, timeout=timeout)
        result = _build_result(
            command,
            args,
            proc.returncode == 0,
            proc.stdout or "",
            proc.stderr or "",
            proc.returncode,
            perf_counter_ns() - started_at,
        )
        LOGGER.info(
            "execution completed",
            extra={
                "subsystem": "runtime.shell",
                "command": command_name,
                "command_args": args,
                "duration_ms": result.duration_ms,
                "exit_code": result.exit_code,
                "success": result.success,
            },
        )
        return result
    except subprocess.TimeoutExpired as e:
        stdout = e.stdout or ""
        stderr = (e.stderr or "") + "process timed out"
        result = _build_result(command, args, False, stdout, stderr, -1, perf_counter_ns() - started_at)
        LOGGER.info(
            "execution timed out",
            extra={
                "subsystem": "runtime.shell",
                "command": command_name,
                "command_args": args,
                "duration_ms": result.duration_ms,
                "exit_code": result.exit_code,
                "success": result.success,
            },
        )
        return result
    except Exception as e:
        result = _build_result(command, args, False, "", str(e), 1, perf_counter_ns() - started_at)
        LOGGER.info(
            "execution failed",
            extra={
                "subsystem": "runtime.shell",
                "command": command_name,
                "command_args": args,
                "duration_ms": result.duration_ms,
                "exit_code": result.exit_code,
                "success": result.success,
            },
        )
        return result
