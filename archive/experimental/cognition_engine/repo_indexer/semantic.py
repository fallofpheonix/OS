"""Semantic topology engine for repository cognition.

Extracts:
- Import graph with module-level resolution
- Side-effect profile per module
- Public vs internal API surface
- Architectural layer/subsystem classification

This is Phase D: Repository Cognition completion.
Supports multiple languages via specialized extractors.
"""

from __future__ import annotations

from pathlib import Path
from typing import Any

from contracts.provenance import Provenance
from repo_indexer.models import (
    ArchitecturalNode,
    ArchitectureGraph,
    Criticality,
    Mutability,
    SideEffectType,
)
from repo_indexer.extractors import (
    PythonExtractor,
    TypeScriptExtractor,
    RustExtractor,
)


# ---------------------------------------------------------------------------
# Layer classification heuristics
# ---------------------------------------------------------------------------

_LAYER_MAP: dict[str, tuple[str, str, Mutability, Criticality]] = {
    "contracts": ("core", "contracts", Mutability.FROZEN, Criticality.FOUNDATION),
    "events": ("core", "events", Mutability.STABLE, Criticality.FOUNDATION),
    "models": ("infrastructure", "model_adapters", Mutability.STABLE, Criticality.CORE),
    "orchestrator": ("engine", "orchestration", Mutability.EVOLVING, Criticality.CORE),
    "planner": ("engine", "planning", Mutability.EVOLVING, Criticality.CORE),
    "validator": ("engine", "validation", Mutability.EVOLVING, Criticality.CORE),
    "repair": ("engine", "repair", Mutability.EVOLVING, Criticality.CORE),
    "runtime": ("infrastructure", "runtime", Mutability.STABLE, Criticality.CORE),
    "memory": ("infrastructure", "memory", Mutability.STABLE, Criticality.CORE),
    "tools": ("infrastructure", "tools", Mutability.EVOLVING, Criticality.SUPPORTING),
    "help": ("engine", "help", Mutability.EVOLVING, Criticality.SUPPORTING),
    "repo_indexer": ("cognition", "indexer", Mutability.EVOLVING, Criticality.CORE),
    "shared_context": ("core", "shared_state", Mutability.STABLE, Criticality.FOUNDATION),
    "sandbox": ("infrastructure", "sandbox", Mutability.STABLE, Criticality.SUPPORTING),
    "transactions": ("infrastructure", "transactions", Mutability.STABLE, Criticality.SUPPORTING),
    "agents": ("engine", "agents", Mutability.VOLATILE, Criticality.PERIPHERAL),
    "api": ("infrastructure", "api", Mutability.EVOLVING, Criticality.SUPPORTING),
    "cli": ("infrastructure", "cli", Mutability.EVOLVING, Criticality.PERIPHERAL),
    "scripts": ("infrastructure", "scripts", Mutability.VOLATILE, Criticality.PERIPHERAL),
    "tests": ("testing", "tests", Mutability.VOLATILE, Criticality.PERIPHERAL),
    "metrics": ("infrastructure", "metrics", Mutability.STABLE, Criticality.SUPPORTING),
}


class SemanticTopologyEngine:
    """Builds an ArchitectureGraph from a repository root.

    Extracts import topology, side-effect profiles, coupling metrics,
    and architectural classification for multiple languages.
    """

    def __init__(self, root: Path | str):
        self.root = Path(root).resolve()
        self.extractors = {
            ".py": PythonExtractor(),
            ".ts": TypeScriptExtractor(),
            ".tsx": TypeScriptExtractor(),
            ".js": TypeScriptExtractor(),
            ".jsx": TypeScriptExtractor(),
            ".rs": RustExtractor(),
        }

    def build_graph(self, limit: int = 1000) -> ArchitectureGraph:
        graph = ArchitectureGraph()

        # Phase 1: Extract all nodes
        for path in self._iter_source_files(limit):
            rel = str(path.relative_to(self.root))
            try:
                content = path.read_text(encoding="utf-8")
            except Exception:
                continue

            node = self._extract_node(rel, content, path.suffix)
            graph.add_node(node)

        # Phase 2: Build edges
        for module_path, node in graph.nodes.items():
            for imported in node.imports:
                # Basic resolution: check if imported module exists in graph
                if imported in graph.nodes:
                    graph.add_edge(module_path, imported)
                else:
                    # Best-effort resolution for extensionless imports (JS/TS)
                    for ext in [".py", ".ts", ".tsx", ".js", ".rs"]:
                        if f"{imported}{ext}" in graph.nodes:
                            graph.add_edge(module_path, f"{imported}{ext}")
                            break

        graph.detect_cycles()
        graph.total_symbols = sum(len(n.public_api) + len(n.internal_symbols) for n in graph.nodes.values())
        return graph

    def _extract_node(self, rel_path: str, content: str, suffix: str) -> ArchitecturalNode:
        print(f"TRACE: extracting {rel_path}")
        provenance = Provenance.from_source(rel_path, content)
        
        extractor = self.extractors.get(suffix)
        if not extractor:
             return ArchitecturalNode(module_path=rel_path, provenance=provenance)

        data = extractor.extract(rel_path, content)
        if "error" in data:
            return ArchitecturalNode(module_path=rel_path, provenance=provenance, architectural_tags=["syntax_error"])

        subsystem, layer, mutability, criticality = self._classify_layer(rel_path)

        return ArchitecturalNode(
            module_path=rel_path,
            subsystem=subsystem,
            layer=layer,
            bounded_context=subsystem,
            mutability=mutability,
            criticality=criticality,
            imports=data.get("imports", []),
            side_effects=data.get("side_effects", []),
            public_api=data.get("public_api", []),
            internal_symbols=data.get("internal_symbols", []),
            class_definitions=data.get("class_definitions", []),
            top_level_functions=data.get("top_level_functions", []),
            provenance=provenance,
        )

    def _classify_layer(self, rel_path: str) -> tuple[str, str, Mutability, Criticality]:
        res = ("unknown", "unknown", Mutability.VOLATILE, Criticality.PERIPHERAL)
        parts = Path(rel_path).parts
        if parts:
            top = parts[0]
            if top in _LAYER_MAP:
                res = _LAYER_MAP[top]
        print(f"TRACE: classified {rel_path} as {res[0]}")
        return res

    def _iter_source_files(self, limit: int):
        extensions = set(self.extractors.keys())
        ignored = {".git", "__pycache__", "node_modules", ".venv", "venv", "dist", "build"}
        count = 0
        # Sort by suffix to process .py first (original behavior)
        for path in sorted(self.root.rglob("*"), key=lambda p: p.suffix):
            if count >= limit: break
            if path.suffix not in extensions: continue
            if any(part in ignored for part in path.parts): continue
            count += 1
            yield path
