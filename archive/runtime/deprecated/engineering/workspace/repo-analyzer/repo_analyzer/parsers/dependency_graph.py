"""
Dependency graph construction for repository analysis.
"""

from __future__ import annotations

from pathlib import Path
from typing import Dict, Iterable, List, Optional

from ..contracts.models import SymbolExtractionResult

try:
    import networkx as nx
except ImportError:  # pragma: no cover - optional dependency
    nx = None


class DependencyGraphBuilder:
    """Build import, call, and module graphs from symbol extraction results."""

    def __init__(self):
        self.networkx = nx

    def build_import_graph(self, analyses: Iterable[SymbolExtractionResult]):
        """Build a directed import graph.

        If networkx is not installed, a plain adjacency mapping is returned.
        """
        if self.networkx is None:
            return self._build_plain_graph(analyses)

        graph = self.networkx.DiGraph()
        for analysis in analyses:
            source = str(analysis.file)
            graph.add_node(source, language=analysis.language.value)
            for dependency in analysis.dependencies:
                graph.add_edge(source, dependency, kind="imports")
        return graph

    def build_call_graph(self, analyses: Iterable[SymbolExtractionResult]):
        """Build a best-effort call graph.

        Nodes are fully qualified file/symbol identifiers.
        """
        if self.networkx is None:
            return self._build_plain_call_graph(analyses)

        graph = self.networkx.DiGraph()
        for analysis in analyses:
            file_prefix = str(analysis.file)
            for symbol in analysis.functions + analysis.classes:
                node_id = f"{file_prefix}::{symbol.name}"
                graph.add_node(node_id, kind=symbol.kind.value)
            for caller in analysis.functions:
                caller_id = f"{file_prefix}::{caller.name}"
                for callee in analysis.calls:
                    graph.add_edge(caller_id, callee, kind="calls")
        return graph

    def build_module_graph(self, analyses: Iterable[SymbolExtractionResult]):
        """Alias for import graph with module-oriented metadata."""
        return self.build_import_graph(analyses)

    def summarize(self, graph) -> Dict[str, List[str]]:
        """Summarize either a networkx graph or a plain adjacency mapping."""
        if self.networkx is not None and hasattr(graph, "nodes"):
            return {
                "nodes": list(graph.nodes),
                "edges": [f"{source}->{target}" for source, target in graph.edges],
            }
        return {
            "nodes": list(graph.keys()),
            "edges": [f"{source}->{target}" for source, targets in graph.items() for target in targets],
        }

    def _build_plain_graph(self, analyses: Iterable[SymbolExtractionResult]) -> Dict[str, List[str]]:
        graph: Dict[str, List[str]] = {}
        for analysis in analyses:
            graph[str(analysis.file)] = sorted(set(analysis.dependencies))
        return graph

    def _build_plain_call_graph(self, analyses: Iterable[SymbolExtractionResult]) -> Dict[str, List[str]]:
        graph: Dict[str, List[str]] = {}
        for analysis in analyses:
            caller_prefix = str(analysis.file)
            for symbol in analysis.functions:
                graph[f"{caller_prefix}::{symbol.name}"] = sorted(set(analysis.calls))
        return graph