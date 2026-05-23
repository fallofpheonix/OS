"""
Code chunk embedding generation using sentence-transformers.
"""

from typing import List, Optional
from dataclasses import dataclass


@dataclass
class EmbeddingConfig:
    """Configuration for embedding generation."""
    model_name: str = "all-MiniLM-L6-v2"  # Fast, suitable for M3
    batch_size: int = 32
    normalize: bool = True
    device: str = "mps"  # Apple Metal acceleration


class EmbeddingGenerator:
    """Generates embeddings for code using sentence-transformers."""
    
    def __init__(self, config: EmbeddingConfig = None):
        self.config = config or EmbeddingConfig()
        self.model = None
        self._initialize()
    
    def _initialize(self):
        """Lazy-load embedding model."""
        try:
            from sentence_transformers import SentenceTransformer
            self.model = SentenceTransformer(
                self.config.model_name,
                device=self.config.device
            )
        except ImportError:
            raise ImportError("sentence-transformers required: pip install sentence-transformers")
    
    def embed(self, texts: List[str]) -> List[List[float]]:
        """
        Generate embeddings for a list of texts.
        Lazy loads model on first call.
        """
        if self.model is None:
            self._initialize()
        
        return self.model.encode(
            texts,
            batch_size=self.config.batch_size,
            normalize_embeddings=self.config.normalize
        ).tolist()
    
    def embed_single(self, text: str) -> List[float]:
        """Generate embedding for a single text."""
        return self.embed([text])[0]
