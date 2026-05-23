"""
AST-aware symbol extraction for repository analysis.
"""

from __future__ import annotations

import ast
import re
from dataclasses import dataclass
from pathlib import Path
from typing import Dict, Iterable, List, Optional, Sequence, Set, Tuple

from ..contracts.models import CodeSymbol, Language, SymbolExtractionResult, SymbolKind

try:
    from tree_sitter_languages import get_parser as get_tree_sitter_parser
except ImportError:  # pragma: no cover - optional dependency
    get_tree_sitter_parser = None


_PYTHON_IMPORT_RE = re.compile(r"^(?:from\s+([\w\.]+)\s+import|import\s+([\w\.,\s]+))")


class ASTParser:
    """Extract symbols and dependencies from source files."""

    def detect_language(self, file_path: Path) -> Language:
        suffix = file_path.suffix.lower()
        mapping = {
            ".py": Language.PYTHON,
            ".ts": Language.TYPESCRIPT,
            ".tsx": Language.TYPESCRIPT,
            ".js": Language.JAVASCRIPT,
            ".jsx": Language.JAVASCRIPT,
            ".go": Language.GO,
            ".rs": Language.RUST,
            ".java": Language.JAVA,
            ".cs": Language.CSHARP,
            ".rb": Language.RUBY,
        }
        return mapping.get(suffix, Language.UNKNOWN)

    def extract_symbols(self, file_path: Path) -> SymbolExtractionResult:
        """Extract structured symbols from a file."""
        file_path = Path(file_path)
        language = self.detect_language(file_path)

        if language == Language.PYTHON:
            return self._extract_python(file_path)

        parser = self._load_tree_sitter_parser(language)
        if parser is not None:
            return self._extract_with_tree_sitter(file_path, language, parser)

        return self._extract_with_regex(file_path, language)

    def extract_functions(self, file_path: Path) -> List[CodeSymbol]:
        """Extract function-level symbols."""
        return self.extract_symbols(file_path).functions

    def extract_classes(self, file_path: Path) -> List[CodeSymbol]:
        """Extract class-level symbols."""
        return self.extract_symbols(file_path).classes

    def extract_imports(self, file_path: Path) -> List[str]:
        """Extract import statements and module dependencies."""
        return self.extract_symbols(file_path).imports

    def extract_exports(self, file_path: Path) -> List[str]:
        """Extract exported symbols."""
        return self.extract_symbols(file_path).exports

    def _extract_python(self, file_path: Path) -> SymbolExtractionResult:
        source = file_path.read_text(encoding="utf-8", errors="ignore")
        tree = ast.parse(source)
        lines = source.splitlines()
        parent_map = self._build_parent_map(tree)

        functions: List[CodeSymbol] = []
        classes: List[CodeSymbol] = []
        imports: List[str] = []
        calls: List[str] = []
        top_level_exports: List[str] = []

        for node in ast.walk(tree):
            if isinstance(node, ast.Import):
                imports.extend(self._collect_import_names(node.names))
            elif isinstance(node, ast.ImportFrom):
                module_name = node.module or ""
                imports.extend(self._collect_import_from_names(module_name, node.names))
            elif isinstance(node, ast.ClassDef):
                classes.append(
                    CodeSymbol(
                        name=node.name,
                        kind=SymbolKind.CLASS,
                        path=file_path,
                        start_line=node.lineno,
                        end_line=getattr(node, "end_lineno", node.lineno),
                        signature=f"class {node.name}",
                        metadata={"bases": [self._safe_unparse(base) for base in node.bases]},
                    )
                )
                if self._is_top_level(node, parent_map):
                    top_level_exports.append(node.name)
            elif isinstance(node, (ast.FunctionDef, ast.AsyncFunctionDef)):
                parent = self._nearest_class(node, parent_map)
                kind = SymbolKind.METHOD if parent else SymbolKind.FUNCTION
                functions.append(
                    CodeSymbol(
                        name=node.name,
                        kind=kind,
                        path=file_path,
                        start_line=node.lineno,
                        end_line=getattr(node, "end_lineno", node.lineno),
                        parent=parent,
                        signature=self._signature_from_function(node),
                        metadata={"decorators": [self._safe_unparse(dec) for dec in node.decorator_list]},
                    )
                )
                if self._is_top_level(node, parent_map) and not node.name.startswith("_"):
                    top_level_exports.append(node.name)
            elif isinstance(node, ast.Call):
                calls.append(self._call_name(node))

        exports = self._extract_python_exports(tree, top_level_exports)
        dependencies = {self._normalize_dependency(item) for item in imports}

        return SymbolExtractionResult(
            file=file_path,
            language=Language.PYTHON,
            functions=functions,
            classes=classes,
            imports=sorted(set(imports)),
            exports=exports,
            calls=sorted({call for call in calls if call}),
            dependencies={dependency for dependency in dependencies if dependency},
        )

    def _extract_python_exports(self, tree: ast.AST, defaults: Sequence[str]) -> List[str]:
        explicit_exports: List[str] = []
        for node in tree.body:
            if not isinstance(node, ast.Assign):
                continue
            for target in node.targets:
                if isinstance(target, ast.Name) and target.id == "__all__":
                    value = getattr(node, "value", None)
                    if isinstance(value, (ast.List, ast.Tuple, ast.Set)):
                        for element in value.elts:
                            if isinstance(element, ast.Constant) and isinstance(element.value, str):
                                explicit_exports.append(element.value)
        if explicit_exports:
            return explicit_exports
        return list(defaults)

    def _extract_with_tree_sitter(
        self,
        file_path: Path,
        language: Language,
        parser,
    ) -> SymbolExtractionResult:
        source = file_path.read_text(encoding="utf-8", errors="ignore")
        tree = parser.parse(source.encode("utf-8"))
        root = tree.root_node

        functions: List[CodeSymbol] = []
        classes: List[CodeSymbol] = []
        imports: List[str] = []
        calls: List[str] = []

        def walk(node, class_name: Optional[str] = None) -> None:
            node_type = getattr(node, "type", "")
            if node_type in {"function_definition", "method_definition"}:
                identifier = self._first_named_child_text(node, source)
                if identifier:
                    functions.append(
                        CodeSymbol(
                            name=identifier,
                            kind=SymbolKind.METHOD if class_name else SymbolKind.FUNCTION,
                            path=file_path,
                            start_line=node.start_point[0] + 1,
                            end_line=node.end_point[0] + 1,
                            parent=class_name,
                        )
                    )
            elif node_type in {"class_definition", "class_declaration"}:
                identifier = self._first_named_child_text(node, source)
                if identifier:
                    classes.append(
                        CodeSymbol(
                            name=identifier,
                            kind=SymbolKind.CLASS,
                            path=file_path,
                            start_line=node.start_point[0] + 1,
                            end_line=node.end_point[0] + 1,
                        )
                    )
                    class_name = identifier

            if node_type in {"import_statement", "import_from_statement"}:
                extracted = self._extract_tree_sitter_import_text(node, source)
                if extracted:
                    imports.append(extracted)

            if node_type == "call_expression":
                call_name = self._extract_tree_sitter_call_text(node, source)
                if call_name:
                    calls.append(call_name)

            for child in getattr(node, "children", []):
                walk(child, class_name=class_name)

        walk(root)
        exports = [symbol.name for symbol in classes + functions if not symbol.name.startswith("_")]
        return SymbolExtractionResult(
            file=file_path,
            language=language,
            functions=functions,
            classes=classes,
            imports=sorted(set(imports)),
            exports=exports,
            calls=sorted(set(calls)),
            dependencies={self._normalize_dependency(item) for item in imports if item},
        )

    def _extract_with_regex(self, file_path: Path, language: Language) -> SymbolExtractionResult:
        source = file_path.read_text(encoding="utf-8", errors="ignore")
        imports: List[str] = []
        functions: List[CodeSymbol] = []
        classes: List[CodeSymbol] = []

        for line_number, line in enumerate(source.splitlines(), start=1):
            match = _PYTHON_IMPORT_RE.match(line.strip())
            if match:
                imports.extend(filter(None, match.groups()))

        return SymbolExtractionResult(
            file=file_path,
            language=language,
            functions=functions,
            classes=classes,
            imports=sorted(set(imports)),
            exports=[],
            calls=[],
            dependencies={self._normalize_dependency(item) for item in imports if item},
        )

    def _load_tree_sitter_parser(self, language: Language):
        if get_tree_sitter_parser is None:
            return None

        language_name = {
            Language.TYPESCRIPT: "typescript",
            Language.JAVASCRIPT: "javascript",
            Language.GO: "go",
            Language.RUST: "rust",
            Language.JAVA: "java",
            Language.CSHARP: "c_sharp",
            Language.RUBY: "ruby",
        }.get(language)

        if not language_name:
            return None

        try:
            return get_tree_sitter_parser(language_name)
        except Exception:
            return None

    def _collect_import_names(self, aliases: Iterable[ast.alias]) -> List[str]:
        names: List[str] = []
        for alias in aliases:
            names.append(alias.name)
            if alias.asname:
                names.append(alias.asname)
        return names

    def _collect_import_from_names(self, module_name: str, aliases: Iterable[ast.alias]) -> List[str]:
        names: List[str] = []
        for alias in aliases:
            if module_name:
                names.append(f"{module_name}.{alias.name}")
            else:
                names.append(alias.name)
        return names

    def _build_parent_map(self, tree: ast.AST) -> Dict[ast.AST, ast.AST]:
        parent_map: Dict[ast.AST, ast.AST] = {}
        for parent in ast.walk(tree):
            for child in ast.iter_child_nodes(parent):
                parent_map[child] = parent
        return parent_map

    def _is_top_level(self, node: ast.AST, parent_map: Dict[ast.AST, ast.AST]) -> bool:
        parent = parent_map.get(node)
        return isinstance(parent, ast.Module)

    def _nearest_class(self, node: ast.AST, parent_map: Dict[ast.AST, ast.AST]) -> Optional[str]:
        parent = parent_map.get(node)
        while parent is not None:
            if isinstance(parent, ast.ClassDef):
                return parent.name
            parent = parent_map.get(parent)
        return None

    def _signature_from_function(self, node: ast.AST) -> str:
        args = getattr(node, "args", None)
        if args is None:
            return getattr(node, "name", "function")

        positional = [arg.arg for arg in args.posonlyargs] + [arg.arg for arg in args.args]
        if args.vararg:
            positional.append(f"*{args.vararg.arg}")
        positional.extend(arg.arg for arg in args.kwonlyargs)
        if args.kwarg:
            positional.append(f"**{args.kwarg.arg}")
        return f"{getattr(node, 'name', 'function')}({', '.join(positional)})"

    def _safe_unparse(self, node: ast.AST) -> str:
        try:
            return ast.unparse(node)
        except Exception:
            return node.__class__.__name__

    def _call_name(self, node: ast.Call) -> str:
        func = node.func
        if isinstance(func, ast.Name):
            return func.id
        if isinstance(func, ast.Attribute):
            return self._attribute_name(func)
        return ""

    def _attribute_name(self, node: ast.Attribute) -> str:
        parts: List[str] = [node.attr]
        current = node.value
        while isinstance(current, ast.Attribute):
            parts.append(current.attr)
            current = current.value
        if isinstance(current, ast.Name):
            parts.append(current.id)
        return ".".join(reversed(parts))

    def _first_named_child_text(self, node, source: str) -> str:
        for child in getattr(node, "children", []):
            if getattr(child, "is_named", False):
                return source[child.start_byte:child.end_byte]
        return ""

    def _extract_tree_sitter_import_text(self, node, source: str) -> str:
        return " ".join(
            part.strip() for part in source[node.start_byte:node.end_byte].split() if part.strip()
        )

    def _extract_tree_sitter_call_text(self, node, source: str) -> str:
        for child in getattr(node, "children", []):
            if getattr(child, "type", "") in {"identifier", "attribute"}:
                return source[child.start_byte:child.end_byte]
        return ""

    def _normalize_dependency(self, item: str) -> str:
        item = item.strip()
        if not item:
            return ""
        return item.split()[0].split(".")[0]