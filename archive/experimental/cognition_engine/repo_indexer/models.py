"""Architectural node and graph models for repository cognition.

These are NOT parsing utilities. They represent the architectural
state model that the engine reasons against. The repository itself
is the cognitive substrate, not merely the target of generation.
"""

from __future__ import annotations

from dataclasses import dataclass, field
from enum import Enum
from typing import Any

from contracts.provenance import Provenance


# ---------------------------------------------------------------------------
# Side-effect taxonomy
# ---------------------------------------------------------------------------

class SideEffectType(Enum):
    """Typed side-effect categories for operational architecture cognition.

    Graphed separately from import topology so that mutation safety,
    rollback risk, and trust scoring can be computed independently.
    """

    FILESYSTEM_READ = "filesystem_read"
    FILESYSTEM_WRITE = "filesystem_write"
    NETWORK_CALL = "network_call"
    PROCESS_SPAWN = "process_spawn"
    DATABASE_WRITE = "database_write"
    EXTERNAL_API = "external_api"
    ENVIRONMENT_MUTATION = "environment_mutation"


# ---------------------------------------------------------------------------
# Mutability / Criticality taxonomy
# ---------------------------------------------------------------------------

class Mutability(Enum):
    """How volatile a module is expected to be."""

    FROZEN = "frozen"       # contracts, interfaces — must never change without review
    STABLE = "stable"       # core logic — changes rarely, high blast radius
    EVOLVING = "evolving"   # feature modules — changes regularly, bounded blast radius
    VOLATILE = "volatile"   # experiments, scripts — changes constantly, low blast radius


class Criticality(Enum):
    """Operational criticality of a module."""

    FOUNDATION = "foundation"   # system won't start without it
    CORE = "core"               # primary functionality depends on it
    SUPPORTING = "supporting"   # used but not load-bearing
    PERIPHERAL = "peripheral"   # optional, removable without breakage


# ---------------------------------------------------------------------------
# Architectural Node
# ---------------------------------------------------------------------------

@dataclass
class Parameter:
    name: str
    type_hint: str | None = None
    default_value: str | None = None


@dataclass
class FunctionDefinition:
    name: str
    parameters: list[Parameter] = field(default_factory=list)
    is_async: bool = False
    return_type: str | None = None


@dataclass
class ClassDefinition:
    name: str
    bases: list[str] = field(default_factory=list)
    methods: list[FunctionDefinition] = field(default_factory=list)


@dataclass
class ArchitecturalNode:
    """A module's identity in the architectural graph.

    Goes beyond mechanical parsing to encode:
    - what subsystem/layer/bounded context it belongs to
    - its semantic role in the system
    - its side-effect profile
    - its public vs internal API surface
    """

    module_path: str

    # ── Architectural identity ──
    subsystem: str = ""
    layer: str = ""
    bounded_context: str = ""

    # ── Semantic role ──
    role: str = ""
    mutability: Mutability = Mutability.EVOLVING
    criticality: Criticality = Criticality.SUPPORTING

    # ── Topology ──
    imports: list[str] = field(default_factory=list)
    imported_by: list[str] = field(default_factory=list)

    # ── Side effects ──
    side_effects: list[SideEffectType] = field(default_factory=list)
    filesystem_access: bool = False
    network_access: bool = False
    process_spawning: bool = False

    # ── Cognition metadata ──
    public_api: list[str] = field(default_factory=list)
    internal_symbols: list[str] = field(default_factory=list)
    class_definitions: list[ClassDefinition] = field(default_factory=list)
    top_level_functions: list[FunctionDefinition] = field(default_factory=list)
    architectural_tags: list[str] = field(default_factory=list)

    # ── Provenance ──
    provenance: Provenance | None = None

    # ── Coupling metrics (computed) ──
    efferent_coupling: int = 0   # outgoing dependencies (Ce)
    afferent_coupling: int = 0   # incoming dependencies (Ca)

    @property
    def instability(self) -> float:
        """Martin's instability metric: I = Ce / (Ca + Ce).

        Range [0, 1]. 0 = maximally stable, 1 = maximally unstable.
        Returns 0.0 if the module has no couplings (isolated).
        """
        total = self.afferent_coupling + self.efferent_coupling
        if total == 0:
            return 0.0
        return self.efferent_coupling / total

    @property
    def side_effect_density(self) -> int:
        """Number of distinct side-effect types this module exhibits."""
        return len(self.side_effects)


# ---------------------------------------------------------------------------
# Architectural Temperature
# ---------------------------------------------------------------------------

@dataclass
class ArchitecturalTemperature:
    """Subsystem instability over time.

    Hot subsystems should reject autonomous mutation,
    require approval, trigger more validation, and reduce repair confidence.
    Cold subsystems are safer for autonomous evolution.
    """

    module_path: str
    recent_mutations: int = 0
    invariant_failures: int = 0
    repair_frequency: int = 0
    dependency_churn: int = 0

    @property
    def temperature(self) -> float:
        """Composite temperature score. Higher = hotter = less safe."""
        return (
            self.recent_mutations * 1.0
            + self.invariant_failures * 3.0
            + self.repair_frequency * 2.0
            + self.dependency_churn * 1.5
        )

    @property
    def is_hot(self) -> bool:
        """Whether this module should reject autonomous mutation."""
        return self.temperature > 10.0


# ---------------------------------------------------------------------------
# Boundary Violation
# ---------------------------------------------------------------------------

@dataclass(frozen=True)
class BoundaryViolation:
    """A detected violation of an architectural boundary contract."""

    invariant_id: str
    source_module: str
    target_module: str
    violation_type: str
    severity: str = "error"
    message: str = ""


# ---------------------------------------------------------------------------
# Architecture Graph
# ---------------------------------------------------------------------------

@dataclass
class ArchitectureGraph:
    """The complete architectural state model of a repository.

    This is the cognitive substrate the engine reasons against.
    It encodes not just what exists, but what is allowed,
    what is forbidden, and what is dangerous.
    """

    nodes: dict[str, ArchitecturalNode] = field(default_factory=dict)
    dependency_edges: list[tuple[str, str]] = field(default_factory=list)
    cyclic_edges: list[tuple[str, str]] = field(default_factory=list)
    forbidden_edges: list[tuple[str, str]] = field(default_factory=list)
    boundary_violations: list[BoundaryViolation] = field(default_factory=list)
    temperatures: dict[str, ArchitecturalTemperature] = field(default_factory=dict)

    # ── Graph-level metrics ──
    total_modules: int = 0
    total_symbols: int = 0
    total_violations: int = 0

    def add_node(self, node: ArchitecturalNode) -> None:
        """Register a module in the graph."""
        self.nodes[node.module_path] = node
        self.total_modules = len(self.nodes)

    def add_edge(self, source: str, target: str) -> None:
        """Register a dependency edge and update coupling counts."""
        self.dependency_edges.append((source, target))
        if source in self.nodes:
            self.nodes[source].efferent_coupling += 1
        if target in self.nodes:
            self.nodes[target].afferent_coupling += 1

    def detect_cycles(self) -> list[tuple[str, str]]:
        """Find and record cyclic dependency edges using DFS."""
        adjacency: dict[str, list[str]] = {}
        for src, tgt in self.dependency_edges:
            adjacency.setdefault(src, []).append(tgt)

        visited: set[str] = set()
        in_stack: set[str] = set()
        cycles: list[tuple[str, str]] = []

        def _dfs(node: str) -> None:
            visited.add(node)
            in_stack.add(node)
            for neighbor in adjacency.get(node, []):
                if neighbor in in_stack:
                    cycles.append((node, neighbor))
                elif neighbor not in visited:
                    _dfs(neighbor)
            in_stack.discard(node)

        for node in adjacency:
            if node not in visited:
                _dfs(node)

        self.cyclic_edges = cycles
        return cycles

    def modules_in_blast_radius(self, module_path: str) -> set[str]:
        """Compute the transitive set of modules affected by a mutation."""
        adjacency: dict[str, list[str]] = {}
        for src, tgt in self.dependency_edges:
            adjacency.setdefault(tgt, []).append(src)  # reverse: who depends on target

        affected: set[str] = set()
        frontier = [module_path]
        while frontier:
            current = frontier.pop()
            if current in affected:
                continue
            affected.add(current)
            frontier.extend(adjacency.get(current, []))
        return affected

    def hot_modules(self) -> list[str]:
        """Return module paths with elevated architectural temperature."""
        return [path for path, temp in self.temperatures.items() if temp.is_hot]

    def summary(self) -> dict[str, Any]:
        """Human-readable graph summary for diagnostics."""
        layers: dict[str, int] = {}
        for node in self.nodes.values():
            layers[node.layer] = layers.get(node.layer, 0) + 1
            
        return {
            "total_modules": self.total_modules,
            "total_symbols": self.total_symbols,
            "total_edges": len(self.dependency_edges),
            "cyclic_edges": len(self.cyclic_edges),
            "boundary_violations": len(self.boundary_violations),
            "hot_modules": self.hot_modules(),
            "layers": layers,
            "avg_instability": (
                sum(n.instability for n in self.nodes.values()) / len(self.nodes)
                if self.nodes
                else 0.0
            ),
        }
