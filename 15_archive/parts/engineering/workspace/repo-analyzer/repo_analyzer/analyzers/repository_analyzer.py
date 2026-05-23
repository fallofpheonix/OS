"""
Integrated repository analysis system.
Coordinates scanning, embedding, storage, and search.
"""

from pathlib import Path
from dataclasses import dataclass

from ..ingest.scanner import RepositoryScanner, ScanConfig
from ..embeddings.generator import EmbeddingGenerator, EmbeddingConfig
from ..vector_store.chroma_store import ChromaVectorStore
from ..contracts.models import RepositoryAnalysis, CodeChunk, SemanticSearchResult, Language
from ..parsers import ASTParser, DependencyGraphBuilder


@dataclass
class RepoAnalysisConfig:
    """Configuration for complete repository analysis."""
    scan_config: ScanConfig = None
    embedding_config: EmbeddingConfig = None
    vector_db_path: Path = None
    
    def __post_init__(self):
        if self.scan_config is None:
            self.scan_config = ScanConfig()
        if self.embedding_config is None:
            self.embedding_config = EmbeddingConfig()
        if self.vector_db_path is None:
            self.vector_db_path = Path.home() / '.repo-analyzer' / 'db'


class RepositoryAnalyzer:
    """
    Integrated analyzer: scans, embeds, stores, and searches code.
    Separates concerns: ingest → embeddings → storage → search
    """
    
    def __init__(self, config: RepoAnalysisConfig = None):
        self.config = config or RepoAnalysisConfig()
        self.scanner = RepositoryScanner(self.config.scan_config)
        self.embedder = EmbeddingGenerator(self.config.embedding_config)
        self.vector_store = ChromaVectorStore(self.config.vector_db_path)
        self.parser = ASTParser()
        self.graph_builder = DependencyGraphBuilder()
        self.collection_initialized = False

    def initialize_search(self, collection_name: str = "repo") -> None:
        """Initialize the vector store for search by connecting to an existing collection."""
        self.vector_store.create_collection(name=collection_name)
        self.collection_initialized = True

    def analyze(self, repo_path: Path, collection_name: str = "repo") -> RepositoryAnalysis:
        """
        Analyze a repository: scan, embed, store.
        """
        repo_path = Path(repo_path).resolve()
        
        # Create collection
        self.vector_store.create_collection(
            name=collection_name,
            metadata={"repo": str(repo_path)}
        )
        self.collection_initialized = True
        
        # Scan repository
        files = list(self.scanner.scan(repo_path))
        print(f"Scanned {len(files)} files")
        
        # Collect symbol-level chunks and repository metadata
        chunks = []
        chunk_ids = []
        chunk_embeddings = []
        chunk_metadatas = []
        symbol_index = {}
        dependency_results = []
        discovered_languages = set()
        
        for file_path in files:
            try:
                symbols = self.parser.extract_symbols(file_path)
                symbol_index[file_path] = symbols
                dependency_results.append(symbols)
                discovered_languages.add(symbols.language)

                content = file_path.read_text(encoding='utf-8', errors='ignore')
                lines = content.split('\n')
                rel_path = file_path.relative_to(repo_path)

                symbol_chunks = symbols.functions + symbols.classes
                if not symbol_chunks:
                    symbol_chunks = []

                for symbol in symbol_chunks:
                    start_index = max(symbol.start_line - 1, 0)
                    end_index = max(symbol.end_line, symbol.start_line)
                    chunk_text = '\n'.join(lines[start_index:end_index]).strip()
                    if not chunk_text:
                        continue

                    chunk = CodeChunk(
                        file_path=file_path,
                        start_line=symbol.start_line,
                        end_line=symbol.end_line,
                        content=chunk_text,
                        chunk_type=symbol.kind.value,
                        language=symbols.language,
                        metadata={
                            "repo": str(repo_path),
                            "symbol": symbol.name,
                            "parent": symbol.parent,
                            "imports": symbols.imports,
                            "exports": symbols.exports,
                        },
                    )
                    chunks.append(chunk)
                    chunk_ids.append(f"{str(rel_path).replace('/', '_')}_{symbol.name}_{symbol.start_line}")
                    chunk_metadatas.append({
                        "file": str(rel_path),
                        "start_line": symbol.start_line,
                        "end_line": symbol.end_line,
                        "symbol": symbol.name,
                        "symbol_kind": symbol.kind.value,
                        "imports": symbols.imports,
                    })

                if not symbol_chunks:
                    chunk = CodeChunk(
                        file_path=file_path,
                        start_line=1,
                        end_line=len(lines),
                        content=content,
                        chunk_type="module",
                        language=symbols.language,
                        metadata={
                            "repo": str(repo_path),
                            "imports": symbols.imports,
                            "exports": symbols.exports,
                        },
                    )
                    chunks.append(chunk)
                    chunk_ids.append(f"{str(rel_path).replace('/', '_')}_module")
                    chunk_metadatas.append({
                        "file": str(rel_path),
                        "start_line": 1,
                        "end_line": len(lines),
                        "symbol": None,
                        "symbol_kind": "module",
                        "imports": symbols.imports,
                    })
            except Exception as e:
                print(f"Error processing {file_path}: {e}")
                continue
        
        # Generate embeddings
        print(f"Embedding {len(chunks)} code chunks...")
        texts_to_embed = [chunk.content for chunk in chunks]
        chunk_embeddings = self.embedder.embed(texts_to_embed)
        
        # Store in vector DB
        documents = [f"{chunk.file_path.name}:{chunk.start_line}" for chunk in chunks]
        self.vector_store.add_embeddings(
            chunk_ids,
            chunk_embeddings,
            chunk_metadatas,
            documents
        )
        print(f"Stored {len(chunks)} embeddings")

        dependency_graph = self.graph_builder.build_import_graph(dependency_results)
        module_graph = self.graph_builder.build_module_graph(dependency_results)
        
        return RepositoryAnalysis(
            root_path=repo_path,
            discovered_languages=discovered_languages or {Language.UNKNOWN},
            total_files=len(files),
            total_lines=sum(len(f.read_text(encoding='utf-8', errors='ignore').split('\n')) 
                           for f in files),
            symbol_index=symbol_index,
            dependency_graph=self.graph_builder.summarize(dependency_graph),
            architecture_patterns=["symbol_level_chunks", "import_graph", "module_graph"],
            metrics={"modules": len(module_graph.nodes) if hasattr(module_graph, "nodes") else len(module_graph)},
        )
    
    def search(self, query: str, repo_path: Path, n_results: int = 10) -> SemanticSearchResult:
        """
        Semantic search over analyzed repository.
        """
        if not self.collection_initialized:
            raise ValueError("Repository not analyzed yet. Call analyze() first.")
        
        # Embed query
        query_embedding = self.embedder.embed_single(query)
        
        # Search
        doc_ids, distances, metadatas = self.vector_store.search(
            query_embedding,
            n_results=n_results
        )
        
        # Build results
        chunks = []
        repo_base = Path(repo_path).resolve()
        for doc_id, metadata in zip(doc_ids, metadatas):
            content = ""
            try:
                # Resolve relative path strictly against the repository root
                file_path = (repo_base / metadata['file']).resolve()
                if not str(file_path).startswith(str(repo_base)):
                    content = "[Security Error: Path traversal attempt blocked]"
                elif file_path.exists():
                    lines = file_path.read_text(encoding='utf-8', errors='ignore').split('\n')
                    start = metadata['start_line'] - 1
                    end = metadata['end_line']
                    content = '\n'.join(lines[start:end])
                else:
                    content = "[Content not found - file may have moved]"
            except Exception:
                file_path = Path(metadata['file'])
                content = "[Error resolving file path]"

            chunks.append(CodeChunk(
                file_path=file_path,
                start_line=metadata['start_line'],
                end_line=metadata['end_line'],
                content=content,
                chunk_type=metadata.get('symbol_kind', 'code'),
                language=Language.UNKNOWN # TODO: map from file ext
            ))
        
        return SemanticSearchResult(
            query=query,
            chunks=chunks,
            relevance_scores=[float(d) for d in distances],
            summary=None
        )
