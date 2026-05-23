"""
ChromaDB integration for vector storage and semantic search.
"""

from pathlib import Path
from typing import List, Dict, Optional, Tuple
from dataclasses import asdict


class ChromaVectorStore:
    """
    ChromaDB-backed vector store for code embeddings.
    Handles persistence and semantic search.
    """
    
    def __init__(self, persist_dir: Path = None):
        self.persist_dir = persist_dir or Path.home() / '.repo-analyzer' / 'chromadb'
        self.persist_dir.mkdir(parents=True, exist_ok=True)
        self.client = None
        self.collection = None
        self._initialize()
    
    def _initialize(self):
        """Lazy-load ChromaDB client."""
        try:
            import chromadb
            self.client = chromadb.PersistentClient(
                path=str(self.persist_dir)
            )
        except ImportError:
            raise ImportError("chromadb required: pip install chromadb")
    
    def create_collection(self, name: str, metadata: Dict = None) -> None:
        """Create a new collection for storing embeddings."""
        if self.client is None:
            self._initialize()
        
        self.collection = self.client.get_or_create_collection(
            name=name,
            metadata=metadata or {"hnsw:space": "cosine"}
        )
    
    def add_embeddings(
        self,
        ids: List[str],
        embeddings: List[List[float]],
        metadatas: List[Dict],
        documents: List[str]
    ) -> None:
        """Add code chunks with embeddings to the collection."""
        if self.collection is None:
            raise ValueError("No collection created. Call create_collection first.")
        
        self.collection.add(
            ids=ids,
            embeddings=embeddings,
            metadatas=metadatas,
            documents=documents
        )
    
    def search(
        self,
        query_embedding: List[float],
        n_results: int = 10
    ) -> Tuple[List[str], List[float], List[Dict]]:
        """
        Search for similar code chunks.
        
        Returns:
            (document_ids, distances, metadatas)
        """
        if self.collection is None:
            raise ValueError("No collection created. Call create_collection first.")
        
        results = self.collection.query(
            query_embeddings=[query_embedding],
            n_results=n_results,
            include=["documents", "metadatas", "distances"]
        )
        
        return (
            results['ids'][0],
            results['distances'][0],
            results['metadatas'][0]
        )
    
    def delete_collection(self, name: str) -> None:
        """Delete a collection."""
        if self.client is None:
            self._initialize()
        self.client.delete_collection(name=name)
