"""Multi-language semantic extractors for repository cognition.

Defines the interface and implementations for extracting architectural
data from various source languages.
"""

from __future__ import annotations

import ast
import re
from abc import ABC, abstractmethod
from pathlib import Path

from repo_indexer.models import (
    ArchitecturalNode,
    ClassDefinition,
    FunctionDefinition,
    Parameter,
    SideEffectType,
    Mutability,
    Criticality,
)

class BaseExtractor(ABC):
    """Abstract base class for language-specific semantic extraction."""

    @abstractmethod
    def extract(self, rel_path: str, content: str) -> dict:
        """Extract symbols, imports, and side effects from source."""
        pass

class PythonExtractor(BaseExtractor):
    """Python-specific AST-based extraction."""

    def extract(self, rel_path: str, content: str) -> dict:
        try:
            tree = ast.parse(content)
        except SyntaxError:
            return {"error": "syntax_error"}

        public, internal, classes, funcs = self._extract_symbols(tree)
        imports = self._extract_imports(tree, rel_path)
        side_effects = self._detect_side_effects(tree, imports)

        return {
            "public_api": public,
            "internal_symbols": internal,
            "class_definitions": classes,
            "top_level_functions": funcs,
            "imports": imports,
            "side_effects": list(set(side_effects)),
        }

    def _extract_imports(self, tree: ast.AST, rel_path: str) -> list[str]:
        imports: list[str] = []
        for node in ast.walk(tree):
            if isinstance(node, ast.Import):
                for alias in node.names:
                    imports.append(alias.name)
            elif isinstance(node, ast.ImportFrom):
                if node.module:
                    imports.append(node.module)
        return sorted(set(imports))

    def _extract_symbols(self, tree: ast.AST) -> tuple[list[str], list[str], list[ClassDefinition], list[FunctionDefinition]]:
        public: list[str] = []
        internal: list[str] = []
        classes: list[ClassDefinition] = []
        funcs: list[FunctionDefinition] = []

        def _get_func_def(node: ast.FunctionDef | ast.AsyncFunctionDef) -> FunctionDefinition:
            params = [Parameter(name=arg.arg) for arg in node.args.args]
            return FunctionDefinition(name=node.name, parameters=params, is_async=isinstance(node, ast.AsyncFunctionDef))

        for node in ast.iter_child_nodes(tree):
            if isinstance(node, (ast.FunctionDef, ast.AsyncFunctionDef)):
                (internal if node.name.startswith("_") else public).append(node.name)
                funcs.append(_get_func_def(node))
            elif isinstance(node, ast.ClassDef):
                (internal if node.name.startswith("_") else public).append(node.name)
                bases = [b.id if isinstance(b, ast.Name) else (b.attr if isinstance(b, ast.Attribute) else "") for b in node.bases]
                methods = [_get_func_def(m) for m in node.body if isinstance(m, (ast.FunctionDef, ast.AsyncFunctionDef))]
                classes.append(ClassDefinition(name=node.name, bases=bases, methods=methods))

        return public, internal, classes, funcs

    def _detect_side_effects(self, tree: ast.AST, imports: list[str]) -> list[SideEffectType]:
        effects = []
        source = ast.dump(tree)
        # Heuristics from original semantic.py
        if "write_text" in source or "open" in source: effects.append(SideEffectType.FILESYSTEM_WRITE)
        if "requests" in source or "httpx" in source: effects.append(SideEffectType.NETWORK_CALL)
        if "subprocess" in source or "os.system" in source: effects.append(SideEffectType.PROCESS_SPAWN)
        return effects

class TypeScriptExtractor(BaseExtractor):
    """TypeScript/JavaScript heuristic-based extraction (Regex-based for now)."""

    def extract(self, rel_path: str, content: str) -> dict:
        # 1. Extract imports (import ... from '...')
        imports = re.findall(r"from\s+['\"](.*?)['\"]", content)
        
        # 2. Extract exports (export function/class/const ...)
        public_api = re.findall(r"export\s+(?:async\s+)?(?:function|class|const|let|var|type|interface)\s+([a-zA-Z0-9_]+)", content)
        
        # 3. Detect side effects
        side_effects = []
        if "fs." in content or "writeFile" in content: side_effects.append(SideEffectType.FILESYSTEM_WRITE)
        if "fetch" in content or "axios" in content or "http" in content: side_effects.append(SideEffectType.NETWORK_CALL)
        if "child_process" in content: side_effects.append(SideEffectType.PROCESS_SPAWN)

        return {
            "public_api": public_api,
            "internal_symbols": [], # harder with regex
            "class_definitions": [], # TODO: structured regex
            "top_level_functions": [], # TODO: structured regex
            "imports": imports,
            "side_effects": side_effects,
        }

class RustExtractor(BaseExtractor):
    """Rust heuristic-based extraction."""

    def extract(self, rel_path: str, content: str) -> dict:
        imports = re.findall(r"use\s+([a-zA-Z0-9_:]+)", content)
        public_api = re.findall(r"pub\s+(?:async\s+)?(?:fn|struct|enum|trait|type|const)\s+([a-zA-Z0-9_]+)", content)
        
        side_effects = []
        if "std::fs" in content: side_effects.append(SideEffectType.FILESYSTEM_WRITE)
        if "tokio::net" in content or "reqwest" in content: side_effects.append(SideEffectType.NETWORK_CALL)
        if "std::process" in content: side_effects.append(SideEffectType.PROCESS_SPAWN)

        return {
            "public_api": public_api,
            "internal_symbols": [],
            "class_definitions": [],
            "top_level_functions": [],
            "imports": imports,
            "side_effects": side_effects,
        }
