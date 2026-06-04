"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
from __future__ import annotations

import tempfile
import unittest
from pathlib import Path

from runtime.filesystem.manager import FilesystemManager
from runtime.shell.executor import execute
from runtime.tracing import RuntimeTrace


class RuntimeTracingTests(unittest.TestCase):
    def test_shell_execution_emits_trace(self):
        result = execute("pwd")

        self.assertIsInstance(result.trace, RuntimeTrace)
        self.assertEqual(result.trace.runtime_category, "shell")
        self.assertEqual(result.trace.operation, "execute")
        self.assertEqual(result.trace.target, "pwd")
        self.assertTrue(result.trace.success)
        self.assertGreaterEqual(result.trace.duration_ms, 0)

    def test_filesystem_failure_emits_trace(self):
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            manager = FilesystemManager(root, max_file_bytes=1)
            result = manager.read_file("missing.txt")

            self.assertIsInstance(result.trace, RuntimeTrace)
            self.assertEqual(result.trace.runtime_category, "filesystem")
            self.assertEqual(result.trace.operation, "read_file")
            self.assertEqual(result.trace.target, "missing.txt")
            self.assertFalse(result.trace.success)

    def test_trace_is_immutable(self):
        trace = execute("pwd").trace
        assert trace is not None

        with self.assertRaises(Exception):
            trace.duration_ms = 10  # type: ignore[misc]


if __name__ == "__main__":
    unittest.main()
