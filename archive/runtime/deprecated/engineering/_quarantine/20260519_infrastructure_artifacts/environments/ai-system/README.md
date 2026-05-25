# AI System Environment
## ChromaDB + Sentence Transformers + LLaMA Index

**DATED**: 13 May 2026  
**STATUS**: ✅ Operational  
**LOCATION**: `~/engineering/environments/ai-system/`

---

## Quick Start

```bash
cd ~/engineering/environments/ai-system
source venv/bin/activate
```

Or use the activation script:

```bash
bash ~/engineering/environments/ai-system/activate.sh
```

---

## Installed Packages

### Core AI/ML Stack

| Package | Version | Purpose |
|---------|---------|---------|
| **chromadb** | 1.5.9 | Vector database for embeddings |
| **sentence-transformers** | 5.5.0 | Embedding generation |
| **llama-index** | 0.14.21 | RAG (Retrieval-Augmented Generation) framework |

### Dependencies

- **pydantic**: Data validation
- **numpy**: Numerical computing
- **torch**: Deep learning (via sentence-transformers)
- **transformers**: HuggingFace models
- **requests**: HTTP client
- **rich**: Terminal formatting
- **uvicorn**: ASGI server
- Plus 100+ transitive dependencies

---

## Usage Examples

### 1. ChromaDB Vector Store

```python
import chromadb

client = chromadb.Client()
collection = client.create_collection(name="documents")

# Add embeddings
collection.add(
    ids=["doc1", "doc2"],
    documents=["Hello world", "Goodbye world"],
)

# Query
results = collection.query(
    query_texts=["Hello"],
    n_results=2
)
```

### 2. Sentence Embeddings

```python
from sentence_transformers import SentenceTransformer

model = SentenceTransformer('all-MiniLM-L6-v2')
sentences = ["This is a sentence", "This is another sentence"]
embeddings = model.encode(sentences)
```

### 3. LLaMA Index RAG

```python
from llama_index import SimpleDirectoryReader, GPTVectorStoreIndex

documents = SimpleDirectoryReader('data').load_data()
index = GPTVectorStoreIndex.from_documents(documents)
query_engine = index.as_query_engine()
response = query_engine.query("What is in the documents?")
```

---

## Environment Structure

```
ai-system/
├── venv/                    # Python virtual environment
│   ├── bin/
│   │   ├── python3          # Python executable
│   │   └── pip              # Package manager
│   └── lib/python3.14/
│       └── site-packages/   # Installed packages
├── activate.sh              # Activation script
├── requirements.txt         # Pip freeze output (all dependencies)
└── README.md               # This file
```

---

## Install Additional Packages

While activated, use pip:

```bash
source venv/bin/activate
pip install [package-name]
pip freeze > requirements.txt  # Update requirements
```

---

## Troubleshooting

### Import Errors

If you get `ModuleNotFoundError`, ensure you've activated the environment:

```bash
source venv/bin/activate
```

### Check Environment

```bash
which python3                    # Should be in venv/bin/
python3 -c "import chromadb"     # Should work
pip list | grep chromadb         # Should show 1.5.9
```

### Regenerate Environment

```bash
rm -rf venv
python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt
```

---

## Next Steps

1. **Integrate with Event System**: Connect ChromaDB to control-plane/runtime-state/
2. **Build RAG Pipeline**: Use LLaMA Index for knowledge extraction
3. **Embed Event Streams**: Store operational events as embeddings
4. **Add Semantic Search**: Query event history by meaning

---

## Version Info

- Python: 3.14.3
- Created: 13 May 2026
- Environment Type: venv (Python virtual environment)
- Platform: macOS

---

**Ready to use for AI-enhanced event processing and operational intelligence.**
