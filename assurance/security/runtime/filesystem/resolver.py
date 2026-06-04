"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.

Canonicalization and containment resolver for the runtime filesystem.

Responsibilities:
- Provide a stable, idempotent `resolve()` contract that returns a fully
  canonical absolute path for a raw input (relative, absolute, or symlink).
- Expose `ensure_within_workspace()` to enforce a single trusted root.
- Avoid leaking low-level OS exceptions; surface domain exceptions.

Design notes:
- Relative paths are interpreted relative to the configured workspace root.
- `resolve()` performs: normalize -> follow symlinks -> re-canonicalize.
- The module is intentionally small and focused; resource policy is kept
  separate.
"""
from __future__ import annotations

import os
from pathlib import Path
from typing import Union

from runtime.filesystem import exceptions


class Resolver:
    """Resolver anchored to a single trusted workspace root."""

    def __init__(self, workspace_root: Union[str, Path]):
        self.workspace_root = Path(workspace_root).resolve()

    def resolve(self, raw_path: Union[str, Path]) -> Path:
        """Resolve `raw_path` to a canonical absolute Path.

        Accepts relative paths (interpreted against the workspace root)
        and absolute paths.

        Raises domain exceptions on invalid inputs or detected symlink loops.
        """
        try:
            p = Path(raw_path)
        except Exception as exc:  # pragma: no cover - defensive
            raise exceptions.InvalidPath(str(exc)) from None

        if not p.is_absolute():
            p = self.workspace_root.joinpath(p)

        normalized = Path(os.path.normpath(str(p)))
        resolved = self._resolve_symlink_chain(normalized)
        canonical = Path(os.path.normpath(str(resolved)))

        second_resolved = self._resolve_symlink_chain(canonical)
        second_canonical = Path(os.path.normpath(str(second_resolved)))
        if str(second_canonical) != str(canonical):
            raise exceptions.SymlinkLoop("non-idempotent resolution detected")

        return canonical

    def ensure_within_workspace(self, resolved_path: Union[str, Path]) -> None:
        """Raise `WorkspaceBoundaryViolation` if `resolved_path` is outside root."""
        p = Path(resolved_path)
        try:
            p_abs = p.resolve(strict=False)
        except RuntimeError:
            raise exceptions.SymlinkLoop("symlink loop during containment check")

        try:
            common = Path(os.path.commonpath([str(self.workspace_root), str(p_abs)]))
        except ValueError:
            raise exceptions.WorkspaceBoundaryViolation(
                f"{p_abs!s} is not within workspace {self.workspace_root!s}"
            )

        if common != self.workspace_root:
            raise exceptions.WorkspaceBoundaryViolation(
                f"{p_abs!s} is not within workspace {self.workspace_root!s}"
            )

    def _resolve_symlink_chain(self, path: Path) -> Path:
        """Resolve symlinks in every path component and detect loops."""
        if not path.is_absolute():
            raise exceptions.InvalidPath("resolver expects an absolute path at this stage")

        seen: set[Path] = set()
        pending_parts = list(path.parts[1:])
        current = Path(path.anchor)

        while pending_parts:
            next_part = pending_parts.pop(0)
            candidate = current / next_part

            if candidate.is_symlink():
                if candidate in seen:
                    raise exceptions.SymlinkLoop(f"symlink loop detected at {candidate!s}")
                seen.add(candidate)

                target = Path(os.readlink(candidate))
                if target.is_absolute():
                    expanded = Path(os.path.normpath(str(target)))
                else:
                    expanded = Path(os.path.normpath(str(candidate.parent / target)))

                pending_parts = list(expanded.parts[1:]) + pending_parts
                current = Path(expanded.anchor)
                continue

            current = candidate

        return current
