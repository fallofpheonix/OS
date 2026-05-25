# COGNITION ENGINE - QUICK REFERENCE CARD

**Print this. Use daily.**

---

## HARD CONSTRAINTS (Never Violate)

```
🔴 ONE MODEL AT A TIME
   Use: asyncio.Lock on model inference

🔴 OFFLINE FIRST
   Local Ollama only, no cloud APIs

🔴 BOUNDED RETRIES
   task_retries: 2
   repair_retries: 1
   total: 3 max

🔴 SNAPSHOT BEFORE MUTATION
   Every file edit → snapshot first

🔴 APPEND-ONLY EVENTS
   Never delete from event log

🔴 VALIDATION REQUIRED
   Every output validated before use

🔴 FAILURE → STRUCTURE
   All failures become FailureRecords

🔴 LOCALIZED RETRIES ONLY
   Don't rerun entire DAG
```

---

## DAILY CHECKLIST

```bash
# 🌅 MORNING (5 min)
python scripts/verify_ollama.py
python scripts/check_db_integrity.py

# 👨‍💻 DURING DEVELOPMENT
python -m pytest tests/ -v
uv run ruff check .
uv run mypy cognition_engine/

# 🧪 BEFORE COMMIT
python -m pytest tests/
python scripts/replay_recent_run.py

# 🌙 END OF DAY
sqlite3 data/cognition.db "SELECT COUNT(*) as run_count FROM runs WHERE DATE(started_at) = DATE('now')"
python scripts/collect_daily_metrics.py
```

---

## THE 10 EXECUTION RULES (Reference)

| # | Rule | Violation = |
|---|------|-----------|
| 1 | Execute tasks sequentially, one model active | Nondeterministic behavior |
| 2 | Use asyncio.Lock on model access | Concurrent inference (crashes) |
| 3 | Persist results immediately | Lost work + corruption |
| 4 | Failures become FailureRecords | Silent failures |
| 5 | Retry only affected subtree | Wasted computation + staleness |
| 6 | Snapshot before dangerous tasks | Unrecoverable corruption |
| 7 | Validate before acceptance | Bad artifacts in chain |
| 8 | Emit events for everything | Unreplayable runs |
| 9 | Enforce retry budgets | Infinite loops |
| 10 | Escalate permissions | Unauthorized actions |

---

## ORCHESTRATION FLOW

```
INPUT
  ↓
Planner (phi3:mini)
  ↓ decompose to tasks
DAG Created
  ↓
Router assigns models
  ↓
Execute Queue (sequential)
  ├─ Task 1 (model A) → Validate → ✓ Store
  ├─ Task 2 (model B) → Validate → ✓ Store
  ├─ Task 3 (model C) → Validate → ✗ Failure
  │   ↓
  │   Repair Planner
  │   ↓
  │   Repair Tasks
  │   ↓ → Success → Resume Task 3
  │
  └─ Remaining Tasks → ...
  ↓
Final Synthesis
  ↓
OUTPUT + ARTIFACTS
```

---

## FAILURE HANDLING FLOW

```
Task Fails
  ↓
Classify Failure Type
  ├─ SyntaxError
  ├─ ImportError
  ├─ TestFailure
  ├─ HallucinatedAPI
  ├─ RuntimeException
  ├─ Timeout
  └─ Unclassified
  ↓
Create FailureRecord → Persist to DB
  ↓
Is it repairable?
  ├─ Yes → Repair Planner
  │   ├─ Generate repair DAG
  │   ├─ Execute repairs
  │   └─ Rerun affected tasks
  │
  └─ No → Request Help
      └─ Wait for human guidance
```

---

## MEMORY ARCHITECTURE

```
Don't use monolithic memory!

Use separated systems:

┌─────────────────────────────────────┐
│  SEMANTIC MEMORY (ChromaDB)         │
│  Purpose: Repo code understanding   │
│  Lifetime: Permanent                │
│  Size: ~100MB per repo              │
└─────────────────────────────────────┘

┌─────────────────────────────────────┐
│  SESSION MEMORY (In-memory dict)    │
│  Purpose: Active run context        │
│  Lifetime: Per session              │
│  Size: <1MB                         │
└─────────────────────────────────────┘

┌─────────────────────────────────────┐
│  FAILURE MEMORY (SQLite)            │
│  Purpose: Debugging history         │
│  Lifetime: Permanent                │
│  Size: ~50MB after 1000 runs        │
└─────────────────────────────────────┘

┌─────────────────────────────────────┐
│  ARCHITECTURE MEMORY (SQLite)       │
│  Purpose: ADRs + design decisions   │
│  Lifetime: Permanent                │
│  Size: <10MB                        │
└─────────────────────────────────────┘
```

---

## MODEL ASSIGNMENT ROUTING

```
TaskType                  → Model
─────────────────────────────────────
CODE_GENERATION           → QWEN
DEBUGGING                 → DEEPSEEK
DOCUMENTATION             → MISTRAL
ARCHITECTURE              → QWEN
EXTRACTION                → QWEN
PLANNING                  → PHI
SYNTHESIS                 → CODELLAMA
TESTING                   → QWEN
```

---

## PERMISSION LEVELS

```
Level 0: observe_only
  └─ Read artifacts, no execution

Level 1: artifact_write
  └─ Generate and save files (no live repo)

Level 2: project_write (DEFAULT)
  └─ Edit repo files + git operations

Level 3: sandbox_execute
  └─ Run tests + build commands

Level 4: network_execute
  └─ Download + install dependencies

Level 5: install_and_configure
  └─ System package installation

⚠️  Levels 3+ require explicit approval
```

---

## COMMON COMMANDS

```bash
# 🚀 CORE SUBSTRATE CLI
python -m cli.main run "Engineering Goal" --repo /path
python -m cli.main replay run_ID
python -m cli.main verify
python -m cli.main audit --verbose

# 🛠️ DEVELOPMENT & OPS
ollama serve &
python -m api.main &
python -m pytest tests/ -v
uv run ruff check .

# 🔍 HEALTH CHECKS
python scripts/verify_ollama.py
python scripts/check_db_integrity.py
```

---

## METRICS TO WATCH

```
Per Run:
  ✓ Task completion rate (target: >70%)
  ✓ Repair success rate (target: >50%)
  ✓ Avg task latency (target: <3000ms)
  ✓ Token usage (target: <5000/task)
  ✓ Failure count (target: <3)

Cumulative:
  ✓ Run success rate (target: >80%)
  ✓ Recovery time (target: <5min)
  ✓ Session length (target: >8hr)
  ✓ Memory stability (target: no growth >2x)

If metrics degrade:
  STOP → Investigate → Fix → Verify metrics recover
```

---

## TROUBLESHOOTING QUICK ANSWERS

| Problem | First Check | If Still Broken |
|---------|------------|-----------------|
| Task hangs | `ollama ps` | Kill + restart ollama, increase timeout |
| ImportError loop | `pip list` or `requirements.txt` | Check repair strategy exists |
| Memory grows | `ps aux \| grep python` | Implement memory pruning |
| Events missing | `sqlite3 events table` | Check event_bus.emit() called |
| DAG invalid | `visualize_dag.py` | Validate: no cycles, all deps exist |
| Artifacts corrupt | `verify_artifacts.py` | Restore from snapshot |
| WebSocket hangs | Browser console: `ws.readyState` | Check backend: `tail -f logs/api.log` |
| Approval timeout | `sqlite3 approvals table` | Auto-deny + escalate to help |

---

## FILE STRUCTURE AT A GLANCE

```
astraeus-core/
├── contracts/          ← Data models (DO NOT MODIFY)
├── planner/           ← Decomposition + planning
├── orchestrator/      ← Core execution loop
├── models/            ← Model adapters (5 models)
├── validator/         ← Validation + failure capture
├── repair/            ← Repair planning + strategies
├── memory/            ← 4 separate memory systems
├── runtime/           ← Sandbox + snapshots
├── transactions/      ← Multi-file safety
├── repo_indexer/      ← AST + code understanding
├── cli/               ← Command line interface
├── api/               ← FastAPI backend
├── frontend-console/  ← Web UI
├── tests/             ← Test suite
├── scripts/           ← Utility scripts
├── artifacts/         ← Generated files (don't commit)
├── data/              ← SQLite databases
└── docs/              ← Documentation
```

---

## EXECUTION PHASES (Master Roadmap)

Execution is governed by the 12-domain hierarchy in **[TODO.md](./TODO.md)**.

- **Phase A-B**: Foundation, Sanitization, Runtime, and Safety Substrate.
- **Phase C-D**: Event Sourcing, Replay, and Repository Cognition.
- **Phase E-F**: Semantic Verification and Temporal Cognition.
- **Phase G-H**: Concurrency, Distributed Governance, and Zero Trust.
- **Phase I-J**: Observability, Operations, and CI/CD.
- **Phase K-L**: Adaptive Repair, Intelligence, and Productionization.

---

## WHEN TO PANIC (Red Flags)

```
🔴 RED FLAGS (Fix immediately):

1. Task success rate drops <50%
   → Something broke in validation or routing

2. Same failure repeating 3+ times
   → Repair strategy doesn't work, remove it

3. Memory growing >2x in 1 hour
   → Implement cleanup immediately

4. Events missing from log
   → Verify event_bus.emit() called everywhere

5. Replay gives different result
   → System has nondeterministic behavior, find it

6. Approval never resolves
   → WebSocket issue, check backend

7. DAG has cycles
   → Planner broke, verify decomposition

8. Snapshot restore fails
   → Corruption risk, stop mutations immediately
```

---

## FEATURE ADDITION CHECKLIST

Before adding any new feature:

- [ ] Write test first (TDD)
- [ ] Does it violate any hard constraint?
- [ ] Does it use monolithic memory? (Bad!)
- [ ] Does it emit events?
- [ ] Does it handle failures?
- [ ] Does it respect retry budgets?
- [ ] Does it work in replay?
- [ ] Do metrics stay healthy?
- [ ] Is it documented?

---

## DEPLOYMENT READINESS

System is ready when it meets all **Project Completion Conditions** in the Master Roadmap:

✅ All mutations are reversible  
✅ All replay is deterministic  
✅ All architecture rules are enforced  
✅ All cognition is repository-grounded  
✅ All memory is temporally reconstructable  
✅ All operations are observable  
✅ All failures are recoverable  
✅ All concurrency is governed  
✅ All mutations are semantically verified  
✅ All dangerous operations are sandboxed  
✅ All autonomy is bounded  
✅ All state is lineage-traceable  

---

## CRITICAL PHONE NUMBER NUMBERS

(These are your actual safe limits)

```
Max tasks per control-plane: 100
Max task duration: 300 seconds
Max retry count: 3
Max repair count: 1
Max session duration: 24 hours
Max memory per session: 2GB
Max artifact size: 100MB
Max database size: 10GB (then archive)

Exceed these = system design violated
```

---

## QUESTIONS TO ASK YOURSELF

**Before committing code:**
- Does this respect the 10 execution rules?
- Will this change be visible in metrics?
- Can the system replay this?
- What happens if this fails halfway?
- Is rollback safe?
- Does a human need to approve this?

**Before pushing to production:**
- Did I run the full test suite?
- Did metrics stay healthy?
- Can I replay the last 10 runs?
- Is the documentation updated?
- Will the next developer understand this?

**When something breaks:**
- What metrics changed?
- What events are missing?
- Can I replay the failing run?
- When did it start?
- What changed recently?

---

## EMERGENCY PROCEDURES

### If Models All Timeout
```bash
kill -9 $(pgrep -f ollama)
sleep 5
ollama serve &
sleep 10
python scripts/verify_ollama.py
```

### If Database Corrupted
```bash
cp data/cognition.db data/cognition.db.backup
sqlite3 data/cognition.db "PRAGMA integrity_check"
# If broken:
rm data/cognition.db
# System will recreate schema
python -m api.main
```

### If WebSocket Hangs
```bash
# In browser console:
ws.close()
location.reload()
# Or restart backend:
pkill -f "api.main"
python -m api.main
```

### If Memory Explodes
```bash
ps aux | grep cognition
# Note process ID
kill -15 <PID>  # Graceful shutdown
# Wait for checkpoints to save
sleep 10
# Resume session
python scripts/resume_session.py <session_id>
```

### If Stuck in Repair Loop
```bash
# Identify the task
sqlite3 data/cognition.db "SELECT task_id, repair_count FROM tasks WHERE status='running'"
# Force stop
python scripts/force_terminate.py <task_id>
# It will request help
# Or disable repair strategy
# Edit repair/repair_planner.py
```

---

## FINAL WISDOM

```
✅ The hardest part isn't code generation.
   It's RELIABLE RECOVERY after failure.

✅ The system succeeds when:
   - Every failure is captured
   - Every repair can be evaluated
   - Every run can be replayed
   - Every decision is observable

✅ When in doubt, add a metric.
   Metrics never lie.

✅ When stuck, replay the run.
   Replay reveals everything.

✅ When confused, read this card again.
   All answers are here.
```

---

**Last Updated**: May 2026  
**For Questions**: Review the Master Specification or existing code  
**To Learn**: Run tests first: `python -m pytest tests/ -v`  
**To Deploy**: Follow the Deployment Checklist above
