"""
Semantic memory: long-term engineering knowledge retrieval.
Uses embeddings for semantic search over code, ADRs, failures, architecture.
"""

from typing import List
from pathlib import Path

from contracts.models import MemoryRecord, MemoryType


from typing import Any


class SemanticMemoryStore:
    """
    Stores and retrieves semantic memory.
    Separate collections for different memory types.
    """
    
    def __init__(self, persist_dir: Path | None = None) -> None:
        self.persist_dir = persist_dir or Path.home() / '.astraeus-core' / 'memory'
        self.persist_dir.mkdir(parents=True, exist_ok=True)
        
        # Initialize ChromaDB client
        self.client: Any | None = None
        self.collections: dict[MemoryType, Any] = {}
        self._initialize()
    
    def _initialize(self) -> None:
        """Lazy-load ChromaDB."""
        try:
            import chromadb
            self.client = chromadb.PersistentClient(
                path=str(self.persist_dir)
            )
            if self.client:
                # Create collections for each memory type
                for memory_type in MemoryType:
                    self.collections[memory_type] = self.client.get_or_create_collection(
                        name=f"memory_{memory_type.value}",
                        metadata={"memory_type": memory_type.value}
                    )
        except ImportError:
            raise ImportError("chromadb required: pip install chromadb")
    
    def store(self, record: MemoryRecord) -> None:
        """Store a memory record."""
        if not self.client:
            self._initialize()
        
        collection = self.collections.get(record.memory_type)
        if not collection:
            raise ValueError(f"Unknown memory type: {record.memory_type}")
        
        collection.add(
            ids=[record.id],
            embeddings=[record.embedding] if record.embedding else None,
            metadatas=[record.metadata],
            documents=[record.content]
        )
    
    def retrieve(
        self,
        query_embedding: List[float],
        memory_type: MemoryType,
        n_results: int = 5
    ) -> List[MemoryRecord]:
        """
        Retrieve memory records by semantic similarity.
        
        Returns:
            List of MemoryRecord with relevance scores
        """
        collection = self.collections.get(memory_type)
        if not collection:
            raise ValueError(f"Unknown memory type: {memory_type}")
        
        results = collection.query(
            query_embeddings=[query_embedding],
            n_results=n_results,
            include=["documents", "metadatas", "distances"]
        )
        
        records = []
        for i, (doc_id, doc, metadata, distance) in enumerate(zip(
            results['ids'][0],
            results['documents'][0],
            results['metadatas'][0],
            results['distances'][0]
        )):
            record = MemoryRecord(
                id=doc_id,
                memory_type=memory_type,
                content=doc,
                metadata=metadata,
                relevance_score=1.0 - distance  # Convert distance to relevance
            )
            records.append(record)
        
        return records
    
    def list_by_type(self, memory_type: MemoryType) -> List[MemoryRecord]:
        """List all memories of a type."""
        collection = self.collections.get(memory_type)
        if not collection:
            return []
        
        # Get all documents from collection
        result = collection.get()
        records = []
        for doc_id, doc, metadata in zip(
            result['ids'],
            result['documents'],
            result['metadatas']
        ):
            record = MemoryRecord(
                id=doc_id,
                memory_type=memory_type,
                content=doc,
                metadata=metadata
            )
            records.append(record)
        
        return records
