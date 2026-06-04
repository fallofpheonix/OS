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


class FilesystemGovernanceTests(unittest.TestCase):
    def test_rejects_large_file(self):
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            (root / "huge.txt").write_bytes(b"x" * 32)

            manager = FilesystemManager(root, max_file_bytes=8)
            result = manager.read_file("huge.txt")

            self.assertFalse(result.success)
            self.assertIn("size limit", result.error.lower())

    def test_rejects_binary_file(self):
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            (root / "binary.bin").write_bytes(b"abc\x00def")

            manager = FilesystemManager(root)
            result = manager.read_file("binary.bin")

            self.assertFalse(result.success)
            self.assertIn("binary content rejected", result.error.lower())

    def test_rejects_unsupported_encoding(self):
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            (root / "invalid.txt").write_bytes(b"\xff\xfe\xfa")

            manager = FilesystemManager(root)
            result = manager.read_file("invalid.txt")

            self.assertFalse(result.success)
            self.assertIn("unsupported encoding", result.error.lower())

    def test_rejects_large_directory_listing(self):
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            directory = root / "dir"
            directory.mkdir()
            for index in range(5):
                (directory / f"file_{index}.txt").write_text(str(index), encoding="utf-8")

            manager = FilesystemManager(root, max_directory_entries=3)
            result = manager.list_directory("dir")

            self.assertFalse(result.success)
            self.assertIn("entry limit", result.error.lower())


if __name__ == "__main__":
    unittest.main()
