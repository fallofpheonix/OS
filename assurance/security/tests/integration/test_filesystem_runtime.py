"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
from pathlib import Path

from runtime.filesystem.manager import FilesystemManager


WORKSPACE_ROOT = Path("/Users/fallofpheonix/engineering/workspace/forge-agent")


def test_list_workspace_root():
    manager = FilesystemManager(WORKSPACE_ROOT)
    result = manager.list_directory(".")

    assert result.success is True
    assert "README.md" in result.entries
    assert result.duration_ms >= 0


def test_read_readme():
    manager = FilesystemManager(WORKSPACE_ROOT)
    result = manager.read_file("README.md")

    assert result.success is True
    assert "forge-agent" in result.content.lower()
    assert result.exists is True
    assert result.duration_ms >= 0


def test_exists_true_and_false():
    manager = FilesystemManager(WORKSPACE_ROOT)

    existing = manager.exists("README.md")
    missing = manager.exists("definitely-not-here.txt")

    assert existing.success is True
    assert existing.exists is True
    assert missing.success is True
    assert missing.exists is False


def test_path_traversal_is_blocked():
    manager = FilesystemManager(WORKSPACE_ROOT)
    result = manager.read_file("../brain/06_FAILURE_LIBRARY/2026-05-executable-code-in-brain.md")

    assert result.success is False
    assert "escapes filesystem root" in result.error.lower()
