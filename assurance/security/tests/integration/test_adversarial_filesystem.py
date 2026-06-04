"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
from __future__ import annotations

import os
import tempfile
import unittest
from pathlib import Path

from runtime.filesystem import exceptions
from runtime.filesystem.manager import FilesystemManager
from runtime.tracing import RuntimeTrace


class AdversarialContainmentTests(unittest.TestCase):
    """Tests that prove post-resolution containment always holds."""

    def test_relative_traversal_is_blocked(self):
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir) / "workspace"
            root.mkdir()
            outside = Path(tempdir) / "outside"
            outside.mkdir()
            (outside / "secret.txt").write_text("secret", encoding="utf-8")

            manager = FilesystemManager(root)
            result = manager.read_file("../outside/secret.txt")

            self.assertFalse(result.success)
            self.assertIsInstance(result.trace, RuntimeTrace)
            self.assertFalse(result.trace.success)

    def test_absolute_path_escape_is_blocked(self):
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir) / "workspace"
            root.mkdir()
            manager = FilesystemManager(root)

            result = manager.read_file("/etc/passwd")

            self.assertFalse(result.success)

    def test_symlink_to_outside_root_is_blocked(self):
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir) / "workspace"
            root.mkdir()
            outside = Path(tempdir) / "outside"
            outside.mkdir()
            (outside / "secret.txt").write_text("secret", encoding="utf-8")

            os.symlink(str(outside / "secret.txt"), str(root / "leak"))

            manager = FilesystemManager(root)
            result = manager.read_file("leak")

            self.assertFalse(result.success)

    def test_nested_symlink_escape_is_blocked(self):
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir) / "workspace"
            root.mkdir()
            outside = Path(tempdir) / "outside"
            outside.mkdir()
            (outside / "secret.txt").write_text("secret", encoding="utf-8")

            (root / "dir").mkdir()
            os.symlink(str(outside / "secret.txt"), str(root / "dir" / "nested_leak"))

            manager = FilesystemManager(root)
            result = manager.read_file("dir/nested_leak")

            self.assertFalse(result.success)

    def test_circular_symlink_is_rejected(self):
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            os.symlink(str(root / "loop_b"), str(root / "loop_a"))
            os.symlink(str(root / "loop_a"), str(root / "loop_b"))

            manager = FilesystemManager(root)
            result = manager.read_file("loop_a")

            self.assertFalse(result.success)


class AdversarialResourceExhaustionTests(unittest.TestCase):
    """Tests that prove resource governance survives pressure."""

    def test_oversized_file_rejected(self):
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            (root / "large.txt").write_bytes(b"x" * 16_000)

            manager = FilesystemManager(root, max_file_bytes=1024)
            result = manager.read_file("large.txt")

            self.assertFalse(result.success)
            self.assertIn("size limit", result.error.lower())

    def test_oversized_directory_listing_rejected(self):
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            (root / "dir").mkdir()
            for i in range(20):
                (root / "dir" / f"file_{i}.txt").write_text(str(i), encoding="utf-8")

            manager = FilesystemManager(root, max_directory_entries=5)
            result = manager.list_directory("dir")

            self.assertFalse(result.success)
            self.assertIn("entry limit", result.error.lower())

    def test_binary_payload_rejected(self):
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            (root / "binary.bin").write_bytes(b"hello\x00world")

            manager = FilesystemManager(root)
            result = manager.read_file("binary.bin")

            self.assertFalse(result.success)
            self.assertIn("binary content rejected", result.error.lower())

    def test_malformed_encoding_rejected(self):
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            (root / "bad.txt").write_bytes(b"\x80\x81\x82\x83")

            manager = FilesystemManager(root)
            result = manager.read_file("bad.txt")

            self.assertFalse(result.success)
            self.assertIn("unsupported encoding", result.error.lower())


class AdversarialSemanticStabilityTests(unittest.TestCase):
    """Tests that prove runtime semantics remain deterministic under pressure."""

    def test_repeated_containment_violation_consistent(self):
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir) / "workspace"
            root.mkdir()
            manager = FilesystemManager(root)

            results = [manager.read_file("../outside.txt") for _ in range(3)]

            self.assertTrue(all(not r.success for r in results))
            self.assertTrue(all(isinstance(r.trace, RuntimeTrace) for r in results))
            self.assertTrue(all(not r.trace.success for r in results))

    def test_no_raw_os_exceptions_escape(self):
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            manager = FilesystemManager(root)

            result = manager.read_file("../../../etc/passwd")

            self.assertFalse(result.success)
            self.assertIsInstance(result.trace, RuntimeTrace)

    def test_deterministic_error_classification(self):
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            (root / "huge.txt").write_bytes(b"x" * 8192)

            manager = FilesystemManager(root, max_file_bytes=1024)
            result1 = manager.read_file("huge.txt")

            manager2 = FilesystemManager(root, max_file_bytes=1024)
            result2 = manager2.read_file("huge.txt")

            self.assertFalse(result1.success)
            self.assertFalse(result2.success)
            self.assertIn("size limit", result1.error.lower())
            self.assertIn("size limit", result2.error.lower())


class AdversarialTraceIntegrityTests(unittest.TestCase):
    """Tests that prove runtime introspection remains trustworthy."""

    def test_trace_emitted_on_success(self):
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            (root / "file.txt").write_text("content", encoding="utf-8")

            manager = FilesystemManager(root)
            result = manager.read_file("file.txt")

            self.assertIsNotNone(result.trace)
            self.assertEqual(result.trace.runtime_category, "filesystem")
            self.assertEqual(result.trace.operation, "read_file")
            self.assertTrue(result.trace.success)

    def test_trace_emitted_on_containment_failure(self):
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir) / "workspace"
            root.mkdir()
            manager = FilesystemManager(root)

            result = manager.read_file("../outside.txt")

            self.assertIsNotNone(result.trace)
            self.assertEqual(result.trace.runtime_category, "filesystem")
            self.assertFalse(result.trace.success)

    def test_trace_emitted_on_governance_failure(self):
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            (root / "huge.txt").write_bytes(b"x" * 2048)

            manager = FilesystemManager(root, max_file_bytes=512)
            result = manager.read_file("huge.txt")

            self.assertIsNotNone(result.trace)
            self.assertFalse(result.trace.success)
            self.assertNotEqual(result.trace.error_type, "")

    def test_trace_duration_consistent_with_result(self):
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            (root / "file.txt").write_text("x" * 100, encoding="utf-8")

            manager = FilesystemManager(root)
            result = manager.read_file("file.txt")

            self.assertEqual(result.duration_ms, result.trace.duration_ms)

    def test_trace_is_immutable_under_all_paths(self):
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            (root / "file.txt").write_text("content", encoding="utf-8")

            manager = FilesystemManager(root)
            result = manager.read_file("file.txt")
            trace = result.trace

            with self.assertRaises(Exception):
                trace.success = False  # type: ignore[misc]


if __name__ == "__main__":
    unittest.main()
