"""
Core data models for repository analysis.
"""

from dataclasses import dataclass, field
from typing import List, Dict, Set, Optional, Any
from pathlib import Path
from enum import Enum


class Language(str, Enum):
    """Supported programming languages."""
    PYTHON = "python"
    TYPESCRIPT = "typescript"
    JAVASCRIPT = "javascript"
    GO = "go"
    RUST = "rust"
    JAVA = "java"
    CSHARP = "csharp"
    RUBY = "ruby"
    UNKNOWN = "unknown"


class SymbolKind(str, Enum):
    """High-level symbol categories extracted from source files."""

    FUNCTION = "function"
    METHOD = "method"
    CLASS = "class"
    IMPORT = "import"
    EXPORT = "export"
    CALL = "call"


@dataclass
class CodeSymbol:
    """A named symbol extracted from source code."""

    name: str
    kind: SymbolKind
    path: Path
    start_line: int
    end_line: int
    parent: Optional[str] = None
    signature: Optional[str] = None
    metadata: Dict[str, Any] = field(default_factory=dict)


@dataclass
class SymbolExtractionResult:
    """Structured output from AST-aware symbol extraction."""

    file: Path
    language: Language
    functions: List[CodeSymbol] = field(default_factory=list)
    classes: List[CodeSymbol] = field(default_factory=list)
    imports: List[str] = field(default_factory=list)
    exports: List[str] = field(default_factory=list)
    calls: List[str] = field(default_factory=list)
    dependencies: Set[str] = field(default_factory=set)


@dataclass
class FileMetadata:
    """Metadata about a source file."""
    path: Path
    language: Language
    size_bytes: int
    lines_of_code: int
    functions: int
    classes: int
    imports: List[str] = field(default_factory=list)
    exports: List[str] = field(default_factory=list)
    complexity: float = 0.0
    embedding_vector: Optional[List[float]] = None


@dataclass
class RepositoryAnalysis:
    """High-level repository analysis result."""
    root_path: Path
    discovered_languages: Set[Language]
    total_files: int
    total_lines: int
    files: Dict[Path, FileMetadata] = field(default_factory=dict)
    symbol_index: Dict[Path, SymbolExtractionResult] = field(default_factory=dict)
    dependency_graph: Dict[str, List[str]] = field(default_factory=dict)
    architecture_patterns: List[str] = field(default_factory=list)
    metrics: Dict[str, float] = field(default_factory=dict)


@dataclass
class CodeChunk:
    """A chunk of code with embeddings and metadata."""
    file_path: Path
    start_line: int
    end_line: int
    content: str
    chunk_type: str  # "function", "class", "module", "comment"
    language: Language
    embedding: Optional[List[float]] = None
    metadata: Dict[str, any] = field(default_factory=dict)


@dataclass
class SemanticSearchResult:
    """Result from semantic code search."""
    query: str
    chunks: List[CodeChunk]
    relevance_scores: List[float]
    summary: Optional[str] = None


@dataclass
class DuplicationReport:
    """Report of code duplication patterns."""
    duplicates: Dict[str, List[Path]]  # content hash -> list of file paths
    similarity_threshold: float
    total_duplicated_lines: int


@dataclass
class ExtractedModule:
    """A suggested reusable module."""
    name: str
    files: Set[Path]
    exports: List[str]
    dependencies: Set[str]
    cohesion_score: float
    coupling_score: float
