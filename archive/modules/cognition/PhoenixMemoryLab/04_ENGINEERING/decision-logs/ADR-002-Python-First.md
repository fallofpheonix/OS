# ADR-002: Python First for Rapid Iteration

**Date:** 2026-05-12  
**Status:** ACCEPTED  
**Authors:** Engineering System

---

## Context

The AI Assistant is a greenfield architecture laboratory. Success depends on:
- Fast iteration
- Clear architectural boundaries  
- Rich ecosystem tooling
- Rapid feedback on design decisions

Language choice affects all of these.

Candidates: Python, Rust, Go, TypeScript

---

## Decision

**Python for Phase 1-2.**

Once architecture stabilizes and modules are extracted, specific high-performance subsystems can migrate to Rust or Go. But initially: Python.

---

## Rationale

### For AI Assistant Specifically

**Python Strengths:**
- LLM ecosystem is Python-native (ollama, llama.cpp, sentence-transformers)
- Subprocess/process management is first-class (easy to build safe execution layer)
- Runtime introspection enables tool execution patterns
- Fastest iteration speed
- Rich standard library
- REPL debugging (invaluable when exploring architecture)

**Python Weaknesses (acceptable for Phase 1-2):**
- Not compiled (but we're not shipping binary yet)
- Runtime overhead (acceptable for development system)
- GIL (fine for I/O-bound task control-plane)

**Rust Weaknesses (at this stage):**
- Slow to write
- Borrow checker slows experimentation
- Ecosystem for LLMs is smaller
- Async/concurrency learning curve delays architecture clarity

**Go Weaknesses:**
- Less natural for agent/control-plane patterns
- No strong LLM ecosystem
- Goroutines are great but overkill for Phase 1

---

## Consequences

### POSITIVE

1. **Fast Architecture Exploration**
   - Errors caught quickly
   - Patterns tested immediately
   - Boundaries validated through code

2. **LLM Integration Simple**
   - ollama/llama.cpp integration easier
   - Sentence-transformers for embeddings
   - Direct API integration simpler

3. **Process Management Elegant**
   - subprocess module is mature
   - Easy to wrap safely
   - STDOUT/STDERR handling clean

4. **Reusable Module Extraction**
   - Clean modules from Python
   - Can wrap in other languages later
   - Clear interface extraction

### NEGATIVE

1. **Will Need Migration for Performance**
   - If tool runtime is bottleneck → migrate to Rust
   - If control-plane is CPU-bound → migrate critical path
   - If inference is needed locally → optimization needed
   - Mitigated by: clean module boundaries enable selective migration

2. **Deployment Requires Python**
   - Not a problem for local systems
   - Docker handles this
   - Fine until Phase 6 (infrastructure)

3. **Runtime Performance Overhead**
   - Acceptable for control-plane
   - Not acceptable for tight loops (why we extract)
   - Mitigated by: keep hot paths in extracted modules

---

## Alternatives Considered

### Alternative 1: Rust
**Why rejected:**

Rust's strengths (memory safety, performance) aren't needed yet.
Rust's weaknesses (slow to write, complex async) are painful now.

Trade: slow iteration speed for performance we don't need yet.

**Revisit after:** Phase 2, when we know what needs to be fast.

### Alternative 2: Go
**Why rejected:**

Go is good for distributed systems. We're not distributed yet.
Go for local agent control-plane is overkill.

**Revisit after:** Phase 4, when distributed execution is needed.

### Alternative 3: TypeScript
**Why rejected:**

Node.js subprocess management is awkward.
LLM ecosystem is JS-heavy but backend-weak.
Good for frontend, wrong for system engineering.

---

## Migration Strategy

### Phase 1-2: Python
Everything in Python.

### Phase 3+: Strategic Rust
When we identify hot paths:
- Migrate critical subsystems to Rust
- Expose via Python FFI or subprocess
- Keep boundaries clean

Example:
```python
# core/orchestrator stays Python
# runtime/tool_executor → Rust + subprocess (isolated)
# storage/vector_db → Rust binary (subprocess)
```

### Packaging
- Phase 1-2: single Python package
- Phase 3: Python + embedded Rust binaries
- Phase 4: Mixed language module ecosystem

---

## Implementation Notes

### Dependencies (Phase 1)
```
typer           # CLI framework
rich            # Terminal UI
pydantic        # Validation & config
python-dotenv   # Environment
pytest          # Testing (add early)
```

### Future Additions (Phase 2)
```
ollama/llama-cpp    # Local LLM
sentence-transformers  # Embeddings
sqlalchemy          # Storage
```

### Performance Monitoring
- Add timing instrumentation early
- Profile hot paths by Phase 2
- Identify Rust migration candidates by Month 2

---

## Related Decisions

- [[ADR-001-Modular-Monolith]] — Enables clean language boundaries
- [[ADR-003-Foundation-First]] — Python speeds up foundation phase

---

## Review Schedule

- **Week 1:** Validate that Python is fast enough
- **Week 2:** Check if async/concurrency needed
- **Month 1:** Identify performance bottlenecks
- **Month 2:** Plan any Rust migrations

---

## Status History
- 2026-05-12: ACCEPTED (language selection for Phase 1)
