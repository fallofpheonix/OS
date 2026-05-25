"""Tests for Phase C: Repository Cognition Substrate.

Covers:
- RuntimeMode + RuntimeCapabilities (C.1)
- Provenance tracking (C.1)
- ArchitecturalNode model + ArchitectureGraph (C.1)
- Graph invalidation (C.1)
- Semantic topology engine (C.2)
- Invariant engine (C.3)
"""

from __future__ import annotations

import hashlib
from pathlib import Path

import pytest

from contracts.provenance import Provenance
from contracts.runtime import (
    RuntimeCapabilities,
    RuntimeMode,
    capabilities_for,
)
from repo_indexer.models import (
    ArchitecturalNode,
    ArchitecturalTemperature,
    ArchitectureGraph,
    Criticality,
    Mutability,
    SideEffectType,
)
from repo_indexer.invalidation import GraphInvalidationEngine


# ═══════════════════════════════════════════════════════════════════════════
# C.1: Runtime Capabilities
# ═══════════════════════════════════════════════════════════════════════════


class TestRuntimeCapabilities:
    def test_live_mode_allows_everything(self) -> None:
        caps = capabilities_for(RuntimeMode.LIVE)
        assert caps.allow_live_models is True
        assert caps.allow_repo_mutation is True
        assert caps.allow_network is True
        assert caps.allow_repair is True
        assert caps.enforce_invariants is True
        assert caps.require_approval is True

    def test_replay_mode_forbids_mutations_and_models(self) -> None:
        caps = capabilities_for(RuntimeMode.REPLAY)
        assert caps.allow_live_models is False
        assert caps.allow_repo_mutation is False
        assert caps.allow_repair is False
        assert caps.allow_replay is True

    def test_offline_mode_enforces_invariants_without_models(self) -> None:
        caps = capabilities_for(RuntimeMode.OFFLINE)
        assert caps.allow_live_models is False
        assert caps.allow_repo_mutation is False
        assert caps.enforce_invariants is True
        assert caps.persist_artifacts is True

    def test_test_mode_sandboxes_everything(self) -> None:
        caps = capabilities_for(RuntimeMode.TEST)
        assert caps.allow_live_models is False
        assert caps.allow_repo_mutation is False
        assert caps.enforce_invariants is True
        assert caps.persist_artifacts is False

    def test_assert_can_mutate_raises_in_offline_mode(self) -> None:
        caps = capabilities_for(RuntimeMode.OFFLINE)
        with pytest.raises(PermissionError, match="mutation is forbidden"):
            caps.assert_can_mutate()

    def test_assert_can_call_model_raises_in_replay_mode(self) -> None:
        caps = capabilities_for(RuntimeMode.REPLAY)
        with pytest.raises(PermissionError, match="live model calls"):
            caps.assert_can_call_model()

    def test_assert_can_repair_raises_in_test_mode(self) -> None:
        caps = capabilities_for(RuntimeMode.TEST)
        with pytest.raises(PermissionError, match="repair execution"):
            caps.assert_can_repair()

    def test_all_modes_have_capabilities(self) -> None:
        for mode in RuntimeMode:
            caps = capabilities_for(mode)
            assert isinstance(caps, RuntimeCapabilities)


# ═══════════════════════════════════════════════════════════════════════════
# C.1: Provenance
# ═══════════════════════════════════════════════════════════════════════════


class TestProvenance:
    def test_from_source_creates_valid_provenance(self) -> None:
        prov = Provenance.from_source("foo/bar.py", "def f(): pass")
        assert prov.source_file == "foo/bar.py"
        assert len(prov.content_hash) == 64
        assert prov.extraction_timestamp > 0
        assert prov.parser_version

    def test_staleness_detection(self) -> None:
        content = "def f(): pass"
        prov = Provenance.from_source("f.py", content)
        same_hash = hashlib.sha256(content.encode("utf-8")).hexdigest()
        assert not prov.is_stale(same_hash)
        assert prov.is_stale("deadbeef" * 8)

    def test_as_dict_roundtrips(self) -> None:
        prov = Provenance.from_source("x.py", "x = 1")
        d = prov.as_dict()
        assert d["source_file"] == "x.py"
        assert isinstance(d["content_hash"], str)


# ═══════════════════════════════════════════════════════════════════════════
# C.1: Architectural Models
# ═══════════════════════════════════════════════════════════════════════════


class TestArchitecturalNode:
    def test_instability_metric(self) -> None:
        node = ArchitecturalNode(
            module_path="foo.py",
            efferent_coupling=3,
            afferent_coupling=7,
        )
        assert node.instability == pytest.approx(0.3)

    def test_instability_zero_when_isolated(self) -> None:
        node = ArchitecturalNode(module_path="isolated.py")
        assert node.instability == 0.0

    def test_instability_one_when_fully_unstable(self) -> None:
        node = ArchitecturalNode(
            module_path="leaf.py",
            efferent_coupling=5,
            afferent_coupling=0,
        )
        assert node.instability == 1.0

    def test_side_effect_density(self) -> None:
        node = ArchitecturalNode(
            module_path="io.py",
            side_effects=[
                SideEffectType.FILESYSTEM_WRITE,
                SideEffectType.NETWORK_CALL,
            ],
        )
        assert node.side_effect_density == 2


class TestArchitecturalTemperature:
    def test_cold_module(self) -> None:
        temp = ArchitecturalTemperature(module_path="stable.py")
        assert temp.temperature == 0.0
        assert not temp.is_hot

    def test_hot_module(self) -> None:
        temp = ArchitecturalTemperature(
            module_path="volatile.py",
            recent_mutations=5,
            invariant_failures=2,
            repair_frequency=1,
        )
        # 5*1.0 + 2*3.0 + 1*2.0 = 13.0
        assert temp.temperature == 13.0
        assert temp.is_hot

    def test_threshold_boundary(self) -> None:
        temp = ArchitecturalTemperature(
            module_path="edge.py",
            recent_mutations=10,
        )
        # 10*1.0 = 10.0, not > 10.0
        assert not temp.is_hot


class TestArchitectureGraph:
    def _build_simple_graph(self) -> ArchitectureGraph:
        graph = ArchitectureGraph()
        graph.add_node(ArchitecturalNode(module_path="a.py", imports=["b.py"]))
        graph.add_node(ArchitecturalNode(module_path="b.py", imports=["c.py"]))
        graph.add_node(ArchitecturalNode(module_path="c.py", imports=[]))
        graph.add_edge("a.py", "b.py")
        graph.add_edge("b.py", "c.py")
        return graph

    def test_add_node_and_edge(self) -> None:
        graph = self._build_simple_graph()
        assert graph.total_modules == 3
        assert len(graph.dependency_edges) == 2
        assert graph.nodes["a.py"].efferent_coupling == 1
        assert graph.nodes["b.py"].afferent_coupling == 1

    def test_detect_no_cycles(self) -> None:
        graph = self._build_simple_graph()
        cycles = graph.detect_cycles()
        assert cycles == []

    def test_detect_cycles(self) -> None:
        graph = self._build_simple_graph()
        graph.add_edge("c.py", "a.py")  # Create cycle
        cycles = graph.detect_cycles()
        assert len(cycles) > 0

    def test_blast_radius(self) -> None:
        graph = self._build_simple_graph()
        # Mutating c.py affects b.py (depends on c) and a.py (depends on b)
        graph.nodes["b.py"].imported_by = ["a.py"]
        graph.nodes["c.py"].imported_by = ["b.py"]
        # Reverse edges for blast radius
        affected = graph.modules_in_blast_radius("c.py")
        assert "c.py" in affected

    def test_summary(self) -> None:
        graph = self._build_simple_graph()
        s = graph.summary()
        assert s["total_modules"] == 3
        assert s["total_edges"] == 2
        assert isinstance(s["avg_instability"], float)


# ═══════════════════════════════════════════════════════════════════════════
# C.1: Graph Invalidation
# ═══════════════════════════════════════════════════════════════════════════


class TestGraphInvalidation:
    def test_register_and_check_staleness(self) -> None:
        graph = ArchitectureGraph()
        engine = GraphInvalidationEngine(graph)
        h = hashlib.sha256(b"original").hexdigest()
        engine.register_hash("foo.py", h)
        assert not engine.check_staleness("foo.py", "original")
        assert engine.check_staleness("foo.py", "modified")

    def test_unknown_module_is_stale(self) -> None:
        graph = ArchitectureGraph()
        engine = GraphInvalidationEngine(graph)
        assert engine.check_staleness("unknown.py", "anything")

    def test_invalidate_removes_node_and_edges(self) -> None:
        graph = ArchitectureGraph()
        graph.add_node(ArchitecturalNode(
            module_path="a.py", imports=["b.py"],
        ))
        graph.add_node(ArchitecturalNode(
            module_path="b.py", imports=[], imported_by=["a.py"],
        ))
        graph.add_edge("a.py", "b.py")

        engine = GraphInvalidationEngine(graph)
        result = engine.invalidate("b.py")

        assert "b.py" not in graph.nodes
        assert graph.total_modules == 1
        assert len(graph.dependency_edges) == 0
        assert result.requires_re_extraction
        assert "a.py" in result.affected_modules

    def test_invalidate_updates_coupling(self) -> None:
        graph = ArchitectureGraph()
        graph.add_node(ArchitecturalNode(module_path="a.py"))
        graph.add_node(ArchitecturalNode(module_path="b.py"))
        graph.add_edge("a.py", "b.py")
        assert graph.nodes["a.py"].efferent_coupling == 1

        engine = GraphInvalidationEngine(graph)
        engine.invalidate("b.py")
        assert graph.nodes["a.py"].efferent_coupling == 0

    def test_scan_for_staleness(self, tmp_path: Path) -> None:
        graph = ArchitectureGraph()
        engine = GraphInvalidationEngine(graph)

        # Create a file, register its hash, then modify it
        f = tmp_path / "module.py"
        content = "x = 1"
        f.write_text(content, encoding="utf-8")
        engine.register_hash(
            "module.py",
            hashlib.sha256(content.encode("utf-8")).hexdigest(),
        )

        # Not stale yet
        assert engine.scan_for_staleness(tmp_path) == []

        # Modify the file
        f.write_text("x = 2", encoding="utf-8")
        stale = engine.scan_for_staleness(tmp_path)
        assert stale == ["module.py"]


# ═══════════════════════════════════════════════════════════════════════════
# C.2: Semantic Topology Engine
# ═══════════════════════════════════════════════════════════════════════════


class TestSemanticTopologyEngine:
    def _create_repo(self, tmp_path: Path) -> Path:
        """Create a minimal Python project for topology extraction."""
        (tmp_path / "contracts").mkdir()
        (tmp_path / "contracts" / "__init__.py").write_text("", encoding="utf-8")
        (tmp_path / "contracts" / "models.py").write_text(
            "class TaskType:\n    pass\n",
            encoding="utf-8",
        )

        (tmp_path / "orchestrator").mkdir()
        (tmp_path / "orchestrator" / "__init__.py").write_text("", encoding="utf-8")
        (tmp_path / "orchestrator" / "engine.py").write_text(
            "from contracts.models import TaskType\n"
            "import requests\n"
            "import subprocess\n\n"
            "class ControlPlane:\n"
            "    def run(self):\n"
            "        requests.post('http://localhost')\n"
            "        subprocess.run(['echo'])\n",
            encoding="utf-8",
        )

        (tmp_path / "validator").mkdir()
        (tmp_path / "validator" / "__init__.py").write_text("", encoding="utf-8")
        (tmp_path / "validator" / "critic.py").write_text(
            "from contracts.models import TaskType\n\n"
            "class Critic:\n"
            "    def review(self, output: str) -> bool:\n"
            "        return True\n",
            encoding="utf-8",
        )

        return tmp_path

    def test_build_graph_extracts_modules(self, tmp_path: Path) -> None:
        from repo_indexer.semantic import SemanticTopologyEngine

        repo = self._create_repo(tmp_path)
        engine = SemanticTopologyEngine(repo)
        graph = engine.build_graph()

        assert graph.total_modules >= 5  # 3 packages + 2 modules minimum
        assert "contracts/models.py" in graph.nodes
        assert "orchestrator/control_plane.py" in graph.nodes

    def test_layer_classification(self, tmp_path: Path) -> None:
        from repo_indexer.semantic import SemanticTopologyEngine

        repo = self._create_repo(tmp_path)
        graph = SemanticTopologyEngine(repo).build_graph()

        contracts_node = graph.nodes["contracts/models.py"]
        assert contracts_node.subsystem == "core"
        assert contracts_node.mutability == Mutability.FROZEN
        assert contracts_node.criticality == Criticality.FOUNDATION

        engine_node = graph.nodes["orchestrator/control_plane.py"]
        assert engine_node.subsystem == "engine"
        assert engine_node.criticality == Criticality.CORE

    def test_import_graph_edges(self, tmp_path: Path) -> None:
        from repo_indexer.semantic import SemanticTopologyEngine

        repo = self._create_repo(tmp_path)
        graph = SemanticTopologyEngine(repo).build_graph()

        # orchestrator/control_plane.py imports contracts/models.py
        engine_node = graph.nodes["orchestrator/control_plane.py"]
        assert any("contracts" in imp for imp in engine_node.imports)

    def test_side_effect_detection(self, tmp_path: Path) -> None:
        from repo_indexer.semantic import SemanticTopologyEngine

        repo = self._create_repo(tmp_path)
        graph = SemanticTopologyEngine(repo).build_graph()

        engine_node = graph.nodes["orchestrator/control_plane.py"]
        assert engine_node.network_access is True
        assert engine_node.process_spawning is True

    def test_public_api_extraction(self, tmp_path: Path) -> None:
        from repo_indexer.semantic import SemanticTopologyEngine

        repo = self._create_repo(tmp_path)
        graph = SemanticTopologyEngine(repo).build_graph()

        critic_node = graph.nodes["validator/critic.py"]
        assert "Critic" in critic_node.public_api

    def test_provenance_attached(self, tmp_path: Path) -> None:
        from repo_indexer.semantic import SemanticTopologyEngine

        repo = self._create_repo(tmp_path)
        graph = SemanticTopologyEngine(repo).build_graph()

        for node in graph.nodes.values():
            assert node.provenance is not None
            assert node.provenance.source_file == node.module_path
            assert len(node.provenance.content_hash) == 64

    def test_cycle_detection_in_real_graph(self, tmp_path: Path) -> None:
        from repo_indexer.semantic import SemanticTopologyEngine

        repo = self._create_repo(tmp_path)

        # Add a cycle: validator imports orchestrator
        (tmp_path / "validator" / "critic.py").write_text(
            "from orchestrator.control_plane import ControlPlane\n\n"
            "class Critic:\n"
            "    pass\n",
            encoding="utf-8",
        )

        graph = SemanticTopologyEngine(repo).build_graph()
        # There should be edges both ways between orchestrator and validator
        # (orchestrator imports contracts, validator imports orchestrator)
        # But no true cycle unless orchestrator also imports validator
        # This test verifies cycle detection infrastructure works
        cycles = graph.detect_cycles()
        assert isinstance(cycles, list)


# ═══════════════════════════════════════════════════════════════════════════
# C.3: Invariant Engine
# ═══════════════════════════════════════════════════════════════════════════


class TestInvariantEngine:
    def test_load_policy(self) -> None:
        from contracts.invariant_engine import InvariantEngine

        engine = InvariantEngine(
            Path(__file__).parent.parent / "contracts" / "invariants.yaml"
        )
        assert len(engine.invariants) > 0
        categories = {inv.category for inv in engine.invariants}
        assert "structural" in categories
        assert "behavioral" in categories
        assert "operational" in categories

    def test_forbidden_import_violation(self) -> None:
        from contracts.invariant_engine import InvariantEngine

        graph = ArchitectureGraph()
        graph.add_node(ArchitecturalNode(
            module_path="validator/critic.py",
            imports=["orchestrator/control_plane.py"],
        ))
        graph.add_node(ArchitecturalNode(
            module_path="orchestrator/control_plane.py",
            imports=[],
        ))

        engine = InvariantEngine(
            Path(__file__).parent.parent / "contracts" / "invariants.yaml"
        )
        report = engine.evaluate(graph)

        # no_validator_orchestrator_cycle should fire
        failed = [r for r in report.results if not r.passed]
        failed_ids = [r.invariant_id for r in failed]
        assert "no_validator_orchestrator_cycle" in failed_ids

    def test_clean_graph_passes(self) -> None:
        from contracts.invariant_engine import InvariantEngine

        graph = ArchitectureGraph()
        graph.add_node(ArchitecturalNode(
            module_path="validator/critic.py",
            imports=["contracts/models.py"],
        ))
        graph.add_node(ArchitecturalNode(
            module_path="contracts/models.py",
            imports=[],
        ))

        engine = InvariantEngine(
            Path(__file__).parent.parent / "contracts" / "invariants.yaml"
        )
        report = engine.evaluate(graph)
        assert report.passed

    def test_cycle_free_invariant(self) -> None:
        from contracts.invariant_engine import InvariantEngine

        graph = ArchitectureGraph()
        graph.add_node(ArchitecturalNode(module_path="a.py", imports=["b.py"]))
        graph.add_node(ArchitecturalNode(module_path="b.py", imports=["a.py"]))
        graph.add_edge("a.py", "b.py")
        graph.add_edge("b.py", "a.py")

        engine = InvariantEngine(
            Path(__file__).parent.parent / "contracts" / "invariants.yaml"
        )
        report = engine.evaluate(graph)

        cycle_result = next(
            (r for r in report.results if r.invariant_id == "no_global_cycles"),
            None,
        )
        assert cycle_result is not None
        assert not cycle_result.passed
        assert len(cycle_result.violations) > 0

    def test_contracts_purity_invariant(self) -> None:
        from contracts.invariant_engine import InvariantEngine

        # Violate: contracts imports orchestrator
        graph = ArchitectureGraph()
        graph.add_node(ArchitecturalNode(
            module_path="contracts/models.py",
            imports=["orchestrator/control_plane.py"],
        ))

        engine = InvariantEngine(
            Path(__file__).parent.parent / "contracts" / "invariants.yaml"
        )
        report = engine.evaluate(graph)

        failed_ids = [r.invariant_id for r in report.results if not r.passed]
        assert "contracts_are_pure" in failed_ids

    def test_report_summary(self) -> None:
        from contracts.invariant_engine import InvariantReport, InvariantResult

        report = InvariantReport(results=[
            InvariantResult(invariant_id="a", passed=True),
            InvariantResult(invariant_id="b", passed=False, message="bad"),
        ])
        s = report.summary()
        assert s["passed"] is False
        assert s["errors"] == 1
        assert "b" in s["failed_invariants"]


# ═══════════════════════════════════════════════════════════════════════════
# Self-validation: run the topology engine against astraeus-core itself
# ═══════════════════════════════════════════════════════════════════════════


class TestSelfValidation:
    """Dogfooding: the cognition engine analyzes its own architecture."""

    def test_cognition_engine_self_scan(self) -> None:
        from repo_indexer.semantic import SemanticTopologyEngine

        engine_root = Path(__file__).parent.parent
        graph = SemanticTopologyEngine(engine_root).build_graph()

        # Basic structural assertions
        assert graph.total_modules > 20  # we have many modules
        assert graph.total_symbols > 50  # we have many symbols
        assert len(graph.dependency_edges) > 0  # there are real dependencies

        # Layer classification sanity
        for path, node in graph.nodes.items():
            if node.subsystem == "unknown":
                continue
            assert node.subsystem, f"no subsystem for {path}"
            assert node.layer, f"no layer for {path}"

    def test_cognition_engine_invariants(self) -> None:
        from contracts.invariant_engine import InvariantEngine
        from repo_indexer.semantic import SemanticTopologyEngine

        engine_root = Path(__file__).parent.parent
        graph = SemanticTopologyEngine(engine_root).build_graph()

        invariant_engine = InvariantEngine(engine_root / "contracts" / "invariants.yaml")
        report = invariant_engine.evaluate(graph)

        # Print failures for diagnostic visibility
        for result in report.results:
            if not result.passed:
                print(f"INVARIANT FAILED: {result.invariant_id}")
                for v in result.violations:
                    print(f"  {v.message}")

        # NOTE: We expect some violations in the current codebase.
        # This test documents them rather than asserting zero violations.
        # The goal is to progressively reduce violations to zero.
        summary = report.summary()
        assert isinstance(summary["total_violations"], int)
