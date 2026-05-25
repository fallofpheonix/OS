"""Incremental graph invalidation engine.

Never rebuild the entire architecture graph from scratch.
Invalidate locally, recompute affected topology, and propagate.
Without this, repository-scale cognition becomes computationally unstable.
"""

from __future__ import annotations

import hashlib
from pathlib import Path
from typing import Any

from repo_indexer.models import ArchitectureGraph


class GraphInvalidationEngine:
    """Tracks content hashes and invalidates stale nodes incrementally."""

    def __init__(self, graph: ArchitectureGraph):
        self.graph = graph
        self._content_hashes: dict[str, str] = {}

    def register_hash(self, module_path: str, content_hash: str) -> None:
        """Record the current content hash for a module."""
        self._content_hashes[module_path] = content_hash

    def check_staleness(self, module_path: str, current_content: str) -> bool:
        """Return True if the module's content has changed since last registration."""
        current_hash = hashlib.sha256(current_content.encode("utf-8")).hexdigest()
        previous_hash = self._content_hashes.get(module_path)
        if previous_hash is None:
            return True  # never seen → treat as stale
        return current_hash != previous_hash

    def invalidate(self, module_path: str) -> InvalidationResult:
        """Remove a stale node and recompute local topology.

        Steps:
        1. Remove the stale node from the graph.
        2. Remove all edges involving the stale node.
        3. Update coupling counts on affected neighbors.
        4. Identify invariants that need re-evaluation.
        5. Return the set of affected modules for downstream re-extraction.
        """
        affected_modules: set[str] = set()
        removed_edges: list[tuple[str, str]] = []

        # Collect affected neighbors before removal
        if module_path in self.graph.nodes:
            node = self.graph.nodes[module_path]
            affected_modules.update(node.imports)
            affected_modules.update(node.imported_by)

        # Remove node
        self.graph.nodes.pop(module_path, None)
        self.graph.total_modules = len(self.graph.nodes)

        # Remove edges involving this module and update coupling
        surviving_edges: list[tuple[str, str]] = []
        for src, tgt in self.graph.dependency_edges:
            if src == module_path or tgt == module_path:
                removed_edges.append((src, tgt))
                # Decrement coupling on surviving neighbor
                if src == module_path and tgt in self.graph.nodes:
                    self.graph.nodes[tgt].afferent_coupling = max(
                        0, self.graph.nodes[tgt].afferent_coupling - 1
                    )
                if tgt == module_path and src in self.graph.nodes:
                    self.graph.nodes[src].efferent_coupling = max(
                        0, self.graph.nodes[src].efferent_coupling - 1
                    )
            else:
                surviving_edges.append((src, tgt))

        self.graph.dependency_edges = surviving_edges

        # Remove from imported_by lists of neighbors
        for neighbor_path in affected_modules:
            neighbor = self.graph.nodes.get(neighbor_path)
            if neighbor and module_path in neighbor.imported_by:
                neighbor.imported_by.remove(module_path)
            if neighbor and module_path in neighbor.imports:
                neighbor.imports.remove(module_path)

        # Clear cached content hash
        self._content_hashes.pop(module_path, None)

        # Remove stale cycles and violations
        self.graph.cyclic_edges = [
            (s, t) for s, t in self.graph.cyclic_edges
            if s != module_path and t != module_path
        ]
        self.graph.boundary_violations = [
            v for v in self.graph.boundary_violations
            if v.source_module != module_path and v.target_module != module_path
        ]

        return InvalidationResult(
            invalidated_module=module_path,
            affected_modules=affected_modules,
            removed_edges=removed_edges,
        )

    def scan_for_staleness(self, root: Path) -> list[str]:
        """Scan all tracked modules and return paths that are stale."""
        stale: list[str] = []
        for module_path, stored_hash in self._content_hashes.items():
            full_path = root / module_path
            if not full_path.exists():
                stale.append(module_path)
                continue
            try:
                content = full_path.read_text(encoding="utf-8")
                current_hash = hashlib.sha256(content.encode("utf-8")).hexdigest()
                if current_hash != stored_hash:
                    stale.append(module_path)
            except Exception:
                stale.append(module_path)
        return stale


class InvalidationResult:
    """Result of a graph invalidation operation."""

    def __init__(
        self,
        invalidated_module: str,
        affected_modules: set[str],
        removed_edges: list[tuple[str, str]],
    ):
        self.invalidated_module = invalidated_module
        self.affected_modules = affected_modules
        self.removed_edges = removed_edges

    @property
    def requires_re_extraction(self) -> bool:
        """Whether downstream modules need their topology recomputed."""
        return len(self.affected_modules) > 0

    def as_dict(self) -> dict[str, Any]:
        return {
            "invalidated_module": self.invalidated_module,
            "affected_modules": sorted(self.affected_modules),
            "removed_edges": self.removed_edges,
            "requires_re_extraction": self.requires_re_extraction,
        }
