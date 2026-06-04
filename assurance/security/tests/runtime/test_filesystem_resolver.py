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
from runtime.filesystem.resolver import Resolver


class FilesystemResolverTests(unittest.TestCase):
    def test_resolve_is_idempotent_for_relative_and_absolute_paths(self):
        with tempfile.TemporaryDirectory() as tempdir:
            workspace = Path(tempdir) / "workspace"
            workspace.mkdir()
            (workspace / "dir").mkdir()
            (workspace / "dir" / "file.txt").write_text("x", encoding="utf-8")
            os.symlink(str(workspace / "dir" / "file.txt"), str(workspace / "dir" / "link"))

            resolver = Resolver(workspace)

            resolved_relative = resolver.resolve(Path("dir") / "link")
            resolved_absolute = resolver.resolve(resolved_relative)

            self.assertEqual(resolved_relative, resolved_absolute)
            self.assertTrue(resolved_relative.is_absolute())

    def test_workspace_boundary_rejects_symlink_escape(self):
        with tempfile.TemporaryDirectory() as tempdir:
            root = Path(tempdir)
            workspace = root / "workspace"
            workspace.mkdir()
            outside = root / "outside"
            outside.mkdir()
            (outside / "secret.txt").write_text("secret", encoding="utf-8")
            os.symlink(str(outside / "secret.txt"), str(workspace / "escape"))

            resolver = Resolver(workspace)
            resolved = resolver.resolve("escape")

            with self.assertRaises(exceptions.WorkspaceBoundaryViolation):
                resolver.ensure_within_workspace(resolved)

    def test_circular_symlink_fails(self):
        with tempfile.TemporaryDirectory() as tempdir:
            workspace = Path(tempdir) / "workspace"
            workspace.mkdir()
            os.symlink(str(workspace / "loop_b"), str(workspace / "loop_a"))
            os.symlink(str(workspace / "loop_a"), str(workspace / "loop_b"))

            resolver = Resolver(workspace)

            with self.assertRaises(exceptions.SymlinkLoop):
                resolver.resolve("loop_a")


if __name__ == "__main__":
    unittest.main()
