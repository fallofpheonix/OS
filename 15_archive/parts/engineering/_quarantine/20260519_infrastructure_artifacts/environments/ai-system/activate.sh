#!/bin/bash
# AI System Environment Activation
# DATED: 13 May 2026
# PURPOSE: Activate the ChromaDB + LLM environment

cd "$(dirname "$0")" || exit 1
source venv/bin/activate

echo "✓ AI System environment activated"
echo "  Python: $(python3 --version)"
echo "  Location: $(pwd)"
echo ""
echo "Installed packages:"
pip list | grep -E "chromadb|sentence-transformers|llama"
echo ""
echo "Ready to use:"
echo "  - chromadb (vector database)"
echo "  - sentence-transformers (embeddings)"
echo "  - llama-index (RAG framework)"
