"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.

Filesystem resource governance helpers.

Trust validation lives in `runtime.filesystem.resolver`. This module owns
resource limits and text-only enforcement for filesystem operations.
"""
from __future__ import annotations

from dataclasses import dataclass
from pathlib import Path

from runtime.filesystem import exceptions


@dataclass(frozen=True)
class ResourceLimits:
	max_file_bytes: int = 1_048_576
	max_directory_entries: int = 1_000
	encoding: str = "utf-8"


def ensure_file_within_size_limit(path: Path, limits: ResourceLimits) -> None:
	size_bytes = path.stat().st_size
	if size_bytes > limits.max_file_bytes:
		raise exceptions.FileTooLarge(
			f"file exceeds size limit: {size_bytes} > {limits.max_file_bytes} bytes"
		)


def read_text_file(path: Path, limits: ResourceLimits) -> str:
	ensure_file_within_size_limit(path, limits)
	data = path.read_bytes()
	if b"\x00" in data:
		raise exceptions.BinaryFileRejected(f"binary content rejected: {path!s}")

	try:
		return data.decode(limits.encoding, errors="strict")
	except UnicodeDecodeError as error:
		raise exceptions.UnsupportedEncoding(
			f"unsupported encoding for {path!s}: {limits.encoding}"
		) from error


def list_directory_entries(path: Path, limits: ResourceLimits) -> list[str]:
	entries = sorted(entry.name for entry in path.iterdir())
	if len(entries) > limits.max_directory_entries:
		raise exceptions.DirectoryTooLarge(
			f"directory exceeds entry limit: {len(entries)} > {limits.max_directory_entries}"
		)
	return entries

