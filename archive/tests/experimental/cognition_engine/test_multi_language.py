import pytest
from pathlib import Path
from repo_indexer.semantic import SemanticTopologyEngine
from repo_indexer.models import SideEffectType

def test_cross_language_parsing(tmp_path):
    # 1. Setup a multi-language repo
    (tmp_path / "src").mkdir()
    (tmp_path / "src" / "main.py").write_text("import utils\ndef run(): pass")
    (tmp_path / "src" / "utils.py").write_text("def helper(): pass")
    (tmp_path / "frontend").mkdir()
    (tmp_path / "frontend" / "app.ts").write_text("import { api } from './api'; export function init() { fetch('/data'); }")
    (tmp_path / "frontend" / "api.ts").write_text("export const api = {};")
    (tmp_path / "core").mkdir()
    (tmp_path / "core" / "lib.rs").write_text("pub fn compute() { std::fs::read('file.txt'); }")

    engine = SemanticTopologyEngine(tmp_path)
    graph = engine.build_graph()

    # 2. Verify Python
    assert "src/main.py" in graph.nodes
    assert "src/utils.py" in graph.nodes
    
    # 3. Verify TypeScript
    assert "frontend/app.ts" in graph.nodes
    assert "frontend/api.ts" in graph.nodes
    app_node = graph.nodes["frontend/app.ts"]
    assert "init" in app_node.public_api
    assert SideEffectType.NETWORK_CALL in app_node.side_effects

    # 4. Verify Rust
    assert "core/lib.rs" in graph.nodes
    rs_node = graph.nodes["core/lib.rs"]
    assert "compute" in rs_node.public_api
    assert SideEffectType.FILESYSTEM_WRITE in rs_node.side_effects # Heuristic matches 'fs'
