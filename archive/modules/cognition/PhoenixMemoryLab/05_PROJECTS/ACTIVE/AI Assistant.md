# AI Engineering Assistant

**Status:** ACTIVE  
**Phase:** FOUNDATION (PENDING CODE)  
**Started:** 2026-05-12  
**Repository:** `~/engineering/workspace/forge-agent`

---

## One-Liner
Local development assistant combining code analysis, execution safety, and LLM reasoning with clean architectural boundaries.

---

## Vision

This is your first serious architecture laboratory.

It will force you to solve:
- Process isolation and safety
- Command execution with structured responses
- Memory/session management
- Tool execution and control-plane
- Local LLM integration
- Error handling and recovery

The modules it produces will become foundations for future systems.

---

## Architecture Overview

### System Boundary
**IN SCOPE:**
- Local CLI interface
- Safe command execution
- File operations with validation
- Git operations
- Project analysis
- Session memory
- Structured logging

**OUT OF SCOPE (Phase 2+):**
- Web/API interface
- Distributed execution
- Plugin marketplaces
- Cloud deployment
- Complex RAG systems

### Core Subsystems

```
┌─────────────────────────────────────┐
│      Interface Layer (CLI)          │
│   (User input → Structured I/O)     │
└────────────────┬────────────────────┘
                 │
┌────────────────▼────────────────────┐
│       Agent Orchestrator            │
│   (Routing, Planning, Context)      │
└────────────────┬────────────────────┘
                 │
        ┌────────┴────────┐
        │                 │
┌───────▼──────┐    ┌─────▼────────┐
│ Tool Runtime │    │ Storage      │
│(Execution)   │    │(Sessions)    │
└──────────────┘    └──────────────┘
        │                 │
┌───────▼─────────────────▼─────────┐
│    Infrastructure Layer           │
│  (Config, Logging, Dependencies)  │
└──────────────────────────────────┘
```

### Subsystem Responsibilities

**Interface Layer:**
- CLI argument parsing
- User input validation
- Response formatting
- Error display

**Agent Orchestrator:**
- Receives validated requests
- Selects tools
- Manages context assembly
- Orchestrates runtime calls
- Structures responses

**Tool Runtime:**
- Safe command execution
- Filesystem operations
- Git integration
- Process management
- Timeout/resource limits

**Storage Layer:**
- Session persistence
- Memory/embeddings
- Logs
- Cache

**Infrastructure:**
- Settings and env loading
- Logging infrastructure
- Dependency injection

---

## Current Phase: FOUNDATION

### What's Being Built
1. **Infrastructure Foundation** (Days 1-2)
   - Configuration system
   - Logging system
   - Environment isolation

2. **Execution Core** (Days 3-4)
   - Shell executor with safety
   - Process management
   - Structured responses

3. **Orchestration** (Days 5-6)
   - CLI interface
   - Request routing
   - Tool selection

4. **Testing & Documentation** (Days 7+)
   - Test coverage
   - ADRs
   - Failure notes

### Success Condition
```
User Input (CLI)
    ↓
CLI validates
    ↓
Orchestrator routes
    ↓
Tool Runtime executes safely
    ↓
Structured Response returned
    ↓
Logged for debugging
```

Even if ugly, this flow proves the architecture works.

---

## Immediate Milestones (Week 1)

- [ ] **Day 1:** Repository structure + infrastructure files
- [ ] **Day 2:** Config + logging working
- [ ] **Day 3:** Shell executor in place
- [ ] **Day 4:** First CLI command runs successfully
- [ ] **Day 5:** File tools implemented
- [ ] **Day 6:** Test suite initialized
- [ ] **Day 7:** First ADR + failure notes captured

---

## Architecture Decisions

### Decision 1: Modular Monolith
**ADR:** [[ADR-001-Modular-Monolith]]

Single repository, clean internal boundaries.
- Faster iteration than microservices
- No premature distribution
- Forces good interfaces
- Can extract modules later

### Decision 2: Python for Speed
**ADR:** [[ADR-002-Python-First]]

- Fastest iteration cycle
- Best LLM ecosystem
- Strong subprocess/runtime tooling
- Can migrate pieces to Rust later

### Decision 3: No AI Until Foundation Complete
**ADR:** [[ADR-003-Foundation-First]]

- CLI → Executor → Response works first
- THEN add local LLM
- THEN add memory
- THEN add advanced reasoning

---

## Module Extraction Candidates

### Candidate 1: Logger
**Pattern:** Every subsystem needs structured logging
**Status:** WATCHING (appears in all layers)
**Extraction Target:** Phase 3 (after secondary project uses it)

### Candidate 2: Config System
**Pattern:** Settings needed everywhere
**Status:** WATCHING (centralized config strategy)
**Extraction Target:** Phase 3

### Candidate 3: Shell Executor
**Pattern:** Safe isolated execution needed by tools
**Status:** WATCHING (tool runtime → shell abstraction)
**Extraction Target:** Phase 3

### Candidate 4: Tool Runtime
**Pattern:** Validated tool execution with safety guarantees
**Status:** WATCHING (fundamental abstraction)
**Extraction Target:** Phase 4 (after distributed system uses similar pattern)

See [[EXTRACTION_LOG]] for tracking.

---

## Runtime Requirements

```
Language: Python 3.10+
Framework: Typer (CLI), Pydantic (validation)
Port: 9999 (future API, not Phase 1)
Memory: 256mb minimum
Storage: 1gb for embeddings/models (Phase 2+)

Dependencies (Phase 1):
- typer          (CLI)
- rich           (terminal UI)
- pydantic       (validation)
- python-dotenv  (env loading)

Dependencies (Phase 2):
- ollama or llama.cpp (local LLM)
- sentence-transformers (embeddings)
- sqlite3 (sessions)
```

---

## Known Issues & Constraints

**None yet** — Phase 0, no execution

After first runs, this will populate with:
- Subprocess stdout/stderr timing
- Command escaping edge cases
- Path resolution issues
- Environment variable issues

---

## Related Documents

- [[Operating-System]] — Vault operating system
- [[EXTRACTION_LOG]] — Module tracking
- [[ADR-001-Modular-Monolith]] — Architecture decision
- [[ADR-002-Python-First]] — Language decision

---

## Next Session Agenda

**Session 1: Repository Initialization (Tomorrow)**

1. Create folder structure
2. Initialize git + venv
3. Install dependencies
4. Create config system
5. Create logger system
6. Create shell executor
7. Create first test
8. Document first ADR

**Session 2: Execution Flow**
1. CLI interface
2. Orchestrator routing
3. Integration test
4. Failure capture

**Session 3: Tool Expansion**
1. File tools
2. Git tools
3. Process management
4. Memory system

---

## Last Updated
2026-05-12 (Setup only — no execution yet)

## Next Review
2026-05-13 (After repository initialization)
