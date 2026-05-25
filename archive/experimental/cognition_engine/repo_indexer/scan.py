"""Repository-aware context scanner."""

from __future__ import annotations

import ast
from pathlib import Path


from typing import Any, Iterator


class RepoIndexer:
    def __init__(self, root: Path | str):
        self.root = Path(root).resolve()

    def scan(self, limit: int = 500) -> dict[str, Any]:
        files = []
        symbols = []
        for path in self._iter_source_files(limit):
            rel = str(path.relative_to(self.root))
            files.append(rel)
            if path.suffix == ".py":
                symbols.extend(self._python_symbols(path, rel))
        return {
            "root": str(self.root),
            "files": files,
            "symbols": symbols,
            "summary": self._summary(files, symbols),
        }

    def _iter_source_files(self, limit: int) -> Iterator[Path]:
        allowed = {".py", ".ts", ".tsx", ".js", ".go", ".rs", ".md"}
        ignored = {".git", "__pycache__", "node_modules", ".venv", "venv", "dist", "build"}
        count = 0
        for path in self.root.rglob("*"):
            if count >= limit:
                break
            if any(part in ignored for part in path.parts):
                continue
            if path.is_file() and path.suffix in allowed:
                count += 1
                yield path

    def _python_symbols(self, path: Path, rel: str) -> list[dict[str, Any]]:
        try:
            tree = ast.parse(path.read_text(encoding="utf-8"))
        except Exception:
            return []
        result = []
        for node in ast.walk(tree):
            if isinstance(node, (ast.FunctionDef, ast.AsyncFunctionDef, ast.ClassDef)):
                result.append({"file": rel, "name": node.name, "line": node.lineno, "kind": type(node).__name__})
        return result

    def _summary(self, files: list[str], symbols: list[dict[str, Any]]) -> str:
        suffix_counts: dict[str, int] = {}
        for file in files:
            suffix = Path(file).suffix or "<none>"
            suffix_counts[suffix] = suffix_counts.get(suffix, 0) + 1
        return (
            f"{len(files)} indexed files; {len(symbols)} Python symbols; "
            f"file types: {suffix_counts}"
        )
