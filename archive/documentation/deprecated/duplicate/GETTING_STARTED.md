# COGNITION ENGINE - GETTING STARTED GUIDE

**This guide gets you running in 30 minutes.**

---

## PREREQUISITES

```bash
# Check you have:
python3 --version        # Should be 3.12+
ollama --version         # Should be installed
sqlite3 --version        # Usually pre-installed

# If not installed:
# macOS: brew install ollama
# Linux: https://ollama.ai
# Windows: https://ollama.ai/download/windows
```

---

## STEP 1: DOWNLOAD MODELS (15 min)

```bash
# In one terminal, start Ollama server
ollama serve

# In another terminal, pull required models
ollama pull phi3:mini              # 2GB - Planning
ollama pull qwen2.5-coder          # 4GB - Code generation  
ollama pull deepseek-coder:6.7b    # 4GB - Debugging
ollama pull mistral                # 4GB - Documentation
ollama pull codellama              # 4GB - Synthesis

# Verify they work
ollama list
# You should see all 5 models listed
```

**This is one-time setup. Takes 15-30 min depending on connection.**

---

## STEP 2: SETUP COGNITION ENGINE (5 min)

```bash
# Clone or navigate to project
cd astraeus-core

# Create virtual environment
python3 -m venv venv
source venv/bin/activate

# Install dependencies
pip install -r requirements.txt

# Verify setup
python scripts/verify_ollama.py

# You should see:
# ✅ Models available
# ✅ Ollama responding
# ✅ Database ready
```

---

## STEP 3: START THE SYSTEM (5 min)

**Terminal 1: API Server**
```bash
cd astraeus-core
source venv/bin/activate
python -m api.main
# You should see: "Uvicorn running on http://localhost:8000"
```

**Terminal 2: Frontend**
```bash
cd astraeus-core/frontend-console
python -m http.server 8000 --directory .
# You should see: "Serving HTTP on port 8000"
```

**Terminal 3: Ollama (if not already running)**
```bash
ollama serve
# You should see: "Listening on 127.0.0.1:11434"
```

---

## STEP 4: OPEN THE WEB INTERFACE (5 min)

Open your browser to: **http://localhost:8000**

You should see:

```
┌─────────────────────────────────────────────────────┐
│           COGNITION ENGINE INTERFACE                │
├──────────────┬──────────────────┬──────────────────┤
│              │                  │                  │
│  Left Panel  │  Center Panel    │  Right Panel     │
│  (Input)     │  (Execution)     │  (Artifacts)     │
│              │                  │                  │
│ Prompt box   │  Task DAG        │  Files           │
│ Session info │  Active model    │  Preview         │
│ Approvals    │  Task queue      │  Diff viewer     │
│              │  Retries         │                  │
└──────────────┴──────────────────┴──────────────────┘
│                   Bottom: Events/Logs/Metrics      │
└────────────────────────────────────────────────────┘
```

---

## STEP 5: RUN YOUR FIRST ORCHESTRATION

### Example 1: Simple REST API

In the **LEFT PANEL**, type:

```
Build a simple REST API with:
- Hello World endpoint
- User data model
- Basic tests
```

Click **[Decompose]**

You should see in the **CENTER PANEL**:
```
Task DAG Generated:
├─ Task 1: Create project structure (READY)
├─ Task 2: Implement hello endpoint (WAITING FOR t1)
├─ Task 3: Add user model (WAITING FOR t1)
├─ Task 4: Write tests (WAITING FOR t2, t3)
└─ Task 5: Summary (WAITING FOR t4)

Ready tasks: 1
Queued tasks: 4
```

Click **[Execute]**

You should see:
```
Task 1: RUNNING
Active Model: phi3:mini
Progress: ████░░░░ 40%
Tokens: 850 / Latency: 2150ms

[After 10-30 seconds, task completes]

Task 1: ✅ COMPLETED
Artifact saved: hello_world.py

Next ready task: Task 2
Task 2: RUNNING
Active Model: qwen2.5-coder
Progress: ██░░░░░░ 15%
```

**In the RIGHT PANEL**, you'll see artifacts generated in real-time:
```
Generated Files:
├─ main.py (REST API scaffold)
├─ models.py (User model)
├─ test_api.py (Tests)
└─ requirements.txt (Dependencies)
```

---

### Example 2: Prompt That Causes Failure & Repair

Try this prompt:
```
Build a FastAPI app that:
- Has JWT authentication  
- Uses PostgreSQL
- Includes Docker setup
- Has full test coverage
```

You'll likely see:
```
Task 3: FAILED
Error: ImportError: No module named 'pydantic'

REPAIR INITIATED:
Suggested fix: pip install pydantic fastapi
Running repair...

Task 3: REPAIR COMPLETED ✅

Affected downstream (Task 4, 5):
↻ Recalculating...
↻ Task 4: RUNNING (using repaired output)
```

**What happened**:
1. ❌ Task failed with ImportError
2. 📝 System classified it as FailureType.IMPORT_ERROR
3. 🔧 Repair planner created fix: "install dependency"
4. ✅ Dependency installed, task retried
5. ⚡ Only affected downstream tasks rerun (not entire DAG)
6. ✅ All completed successfully

---

### Example 3: Approval Gate

Try this prompt:
```
Setup a PostgreSQL database and create tables
```

You might see:

```
⚠️  APPROVAL REQUIRED

Action: Install system package: postgresql-dev
Permission Level: 4 (install_and_configure)
Reason: PostgreSQL integration required

[Approve] [Deny] [Modify]
```

Click **[Approve]**

```
✅ APPROVED (by you)
Installation proceeds...
Task completed successfully
```

**This shows**: Dangerous actions require explicit human approval.

---

## STEP 6: UNDERSTAND THE INTERFACE

### LEFT PANEL: Input & Control

```
┌─────────────────────────────────────┐
│  Engineering Goal                   │
├─────────────────────────────────────┤
│ [Text area for your prompt]          │
│                                     │
│ Example:                            │
│ "Build a REST API with JWT auth,    │
│  tests, and Docker setup"           │
│                                     │
├─────────────────────────────────────┤
│  [Decompose] [Execute] [Pause]      │
│  [Resume]    [Reset]                │
├─────────────────────────────────────┤
│  Session Info                       │
│  ID: sess_abc123                    │
│  Runtime: 00:12:34                  │
│  Status: EXECUTING                  │
├─────────────────────────────────────┤
│  Approval Requests (if any)         │
│  ⚠️  Install postgres?              │
│  [Approve] [Deny]                   │
└─────────────────────────────────────┘
```

**Key Buttons**:
- **Decompose**: Parse prompt → Create task DAG
- **Execute**: Run tasks sequentially
- **Pause**: Pause execution (can resume)
- **Reset**: Clear current run, start over

### CENTER PANEL: Execution Visualization

```
┌─────────────────────────────────────┐
│  Task DAG                           │
│                                     │
│  ┌──────┐     ┌──────┐             │
│  │Task1 │────→│Task2 │─┐           │
│  └──────┘     └──────┘ │           │
│                        ↓           │
│                    ┌──────┐        │
│                    │Task3 │        │
│                    └──────┘        │
├─────────────────────────────────────┤
│  Active Model: qwen2.5-coder        │
│  Progress: ████████░░ 80%           │
│  Tokens: 2,450 / Latency: 3,200ms   │
├─────────────────────────────────────┤
│  Task Queue                         │
│  1. Task1 ✅ COMPLETED              │
│  2. Task2 🔄 RUNNING                │
│  3. Task3 ⏳ READY                  │
│  4. Task4 ⌛ WAITING (depends: t3)  │
├─────────────────────────────────────┤
│  Retries & Repairs                  │
│  Task2: Retry 1/2                   │
│  Repair: ImportError → pip install  │
│  Status: ✅ Repaired successfully   │
└─────────────────────────────────────┘
```

**What you're watching**:
- DAG structure (task dependencies)
- Which model is currently working
- Progress of current task
- Queue of pending tasks
- Any repairs in progress

### RIGHT PANEL: Artifacts & Results

```
┌──────────────────────────────────────┐
│  Generated Artifacts                 │
├──────────────────────────────────────┤
│  ✓ main.py (REST API)                │
│  ✓ models.py (Data models)           │
│  ✓ tests.py (Unit tests)             │
│  ✓ docker-compose.yml                │
│  ✓ README.md (Documentation)         │
├──────────────────────────────────────┤
│  Preview: [main.py selected]         │
├──────────────────────────────────────┤
│  from fastapi import FastAPI         │
│  from pydantic import BaseModel      │
│                                      │
│  app = FastAPI()                     │
│                                      │
│  @app.get("/")                       │
│  async def hello():                  │
│      return {"message": "Hello"}     │
│                                      │
├──────────────────────────────────────┤
│  Diff Viewer                         │
│  (Shows changes made to files)       │
│  + 45 lines added                    │
│  - 2 lines removed                   │
│  ~ 12 lines modified                 │
└──────────────────────────────────────┘
```

**What you can do**:
- Click artifacts to preview them
- See diffs of changes
- Download files
- Copy code snippets

### BOTTOM PANEL: Detailed Logs

Three tabs:

**[Events] [Logs] [Metrics] [Help]**

```
╔════════════════════════════════════════════╗
║               EVENTS TAB                   ║
╠════════════════════════════════════════════╣
║ 14:30:45 | session_created | sess_xyz    ║
║ 14:30:46 | task_created | t1             ║
║ 14:30:47 | task_created | t2             ║
║ 14:30:48 | task_ready | t1               ║
║ 14:30:49 | task_started | t1 | phi       ║
║ 14:30:52 | task_completed | t1           ║
║ 14:30:53 | task_ready | t2               ║
║ 14:30:54 | task_started | t2 | qwen      ║
║ ...                                       ║
╚════════════════════════════════════════════╝

╔════════════════════════════════════════════╗
║               METRICS TAB                  ║
╠════════════════════════════════════════════╣
║ Task Completion: 85%                      ║
║ Repair Success: 72%                       ║
║ Avg Latency: 2,350ms                      ║
║ Tokens Used: 8,450                        ║
║ Model Usage:                              ║
║   phi: 1 task                             ║
║   qwen: 3 tasks                           ║
║   deepseek: 2 tasks                       ║
║   mistral: 0 tasks                        ║
║   codellama: 0 tasks                      ║
╚════════════════════════════════════════════╝
```

---

## STEP 7: COMMON WORKFLOWS

### Workflow A: Generate Code Artifacts

```
1. Type prompt in LEFT PANEL
   "Generate a FastAPI user service"

2. Click [Decompose]
   → See task DAG in CENTER

3. Click [Execute]
   → Watch tasks execute in CENTER
   → See files appear in RIGHT

4. Download artifacts from RIGHT PANEL
   → main.py, models.py, tests.py, etc.
```

### Workflow B: Handle a Failure & Watch Repair

```
1. Prompt has an error
   "Use library X" (X not installed)

2. Task fails (see RED in CENTER)

3. System auto-repairs
   → Repair task created
   → Dependency installed
   → Original task retried
   → Status: ✅ COMPLETED

4. Failed task never happened in final output
```

### Workflow C: Approve Dangerous Action

```
1. Prompt needs elevated permission
   "Install PostgreSQL and initialize database"

2. APPROVAL REQUEST appears (LEFT PANEL)
   ⚠️  Install system package: postgresql-dev
   [Approve] [Deny]

3. You click [Approve]

4. Action proceeds with audit log recorded
   (Visible in EVENTS tab)
```

### Workflow D: Long-Running Session

```
1. Start control-plane
   "Build a complete backend service"
   (Multiple complex tasks)

2. System shows progress
   Estimated time: 15 minutes

3. Can pause with [Pause] button
   → All state saved

4. Can come back 1 hour later
   Click [Resume]
   → System resumes from checkpoint
   → No lost work

5. Even if browser/server crashes:
   → State persisted to SQLite
   → Run: python scripts/resume_session.py <session_id>
```

---

## STEP 8: VERIFY IT'S WORKING

After your first successful run:

```bash
# Check database
sqlite3 data/cognition.db "SELECT COUNT(*) as runs FROM runs"
# Should show: 1

# Check artifacts were saved
ls -la artifacts/
# Should show: run_001/ with task outputs

# Check events were logged
sqlite3 data/cognition.db "SELECT COUNT(*) as events FROM events WHERE run_id = (SELECT id FROM runs LIMIT 1)"
# Should show: >10

# Verify replay works
python scripts/replay_run.py run_001
# Should complete without errors

# Check metrics
python scripts/collect_metrics.py
# Should show: task_completion_rate, etc.
```

---

## STEP 9: TROUBLESHOOTING YOUR FIRST RUN

### Problem: "Models don't appear"
```
Symptom: Task shows "WAITING FOR MODEL" forever
Fix:
  1. Check: ollama list
  2. If empty: run ollama pull phi3:mini etc.
  3. Check: python scripts/verify_ollama.py
```

### Problem: "WebSocket connection failed"
```
Symptom: "Failed to connect to ws://localhost:8000"
Fix:
  1. Check: Is API server running? (Terminal 1)
  2. Check: http://localhost:8000/health (should return 200)
  3. Restart: Kill API server, restart it
```

### Problem: "Tasks get stuck"
```
Symptom: Task shows "RUNNING" for 5+ minutes
Fix:
  1. Check: ps aux | grep ollama
  2. Check: Does Ollama have timeout? (default: 180s)
  3. Increase: In models/ollama_client.py, timeout=300
  4. Restart API server
```

### Problem: "Artifacts don't appear in RIGHT PANEL"
```
Symptom: Task completed but no files shown
Fix:
  1. Check: artifacts/ folder exists
  2. Check: Browser refresh (F5)
  3. Check: WebSocket still connected?
  4. Restart frontend: pkill -f http.server
```

---

## STEP 10: NEXT STEPS

After first successful run:

```
✅ You've verified:
   - Ollama working
   - Models responsive
   - Database functioning
   - WebSocket communication
   - Task execution
   - Artifact generation

Now explore:
   [ ] Review QUICK_REFERENCE_CARD.md
   [ ] Read ARCHITECTURE section of Master Spec
   [ ] Try different prompt types
   [ ] Monitor metrics dashboard
   [ ] Test failure + repair cycle
   [ ] Test approval gates
   [ ] Understand memory systems
```

---

## QUICK REFERENCE: UI KEYBOARD SHORTCUTS

(Coming soon - implement in frontend)

```
Ctrl+Enter    Decompose current prompt
Ctrl+Shift+E  Execute DAG
Space         Toggle pause/resume
Ctrl+L        Clear logs
Ctrl+S        Save session
Ctrl+?        Open help
```

---

## EXAMPLE PROMPTS TO TRY

### Simple (2-3 tasks)
```
Create a Python module that:
- Implements a linked list
- Has unit tests
- Has documentation
```

### Medium (5-7 tasks)
```
Build a CLI tool that:
- Parses CSV files
- Filters by criteria
- Exports to JSON
- Has --help and --version flags
- Includes integration tests
```

### Complex (10+ tasks)
```
Create a full stack web application:
- Backend: FastAPI with user authentication
- Frontend: React with hooks
- Database: PostgreSQL with migrations
- Testing: unit + integration tests
- Deployment: Docker + docker-compose
- Documentation: API docs + setup guide
- CI/CD: GitHub Actions workflow
```

---

## THINGS TO NOTICE

After each run, look for:

```
✅ Task DAG is sensible
   (Tasks are properly ordered with dependencies)

✅ Model assignment makes sense
   (Code → Qwen, Debugging → DeepSeek, etc.)

✅ Validation works
   (Bad outputs are caught)

✅ Artifacts are useful
   (Can actually use the generated files)

✅ Metrics improve
   (Success rate trending up)

✅ Repairs work when needed
   (Failures auto-resolve)

⚠️  If any of these don't happen,
   something needs fixing
```

---

## WHERE TO GET HELP

```
For questions about:
  - Architecture        → Read: MASTER_SPEC.md, ARCHITECTURE section
  - Failures            → Read: Troubleshooting section
  - Metrics             → Read: QUICK_REFERENCE_CARD.md
  - Specific error      → Search: sqlite3 data/cognition.db
  - How system works    → Review: orchestrator/control_plane.py

For problems:
  1. Check browser console (F12)
  2. Check backend logs (Terminal 1)
  3. Check database: sqlite3 data/cognition.db
  4. Run: python scripts/verify_ollama.py
  5. Run: python scripts/check_db_integrity.py
  6. Ask: Review MASTER_SPEC.md Troubleshooting section
```

---

## SUMMARY

You now have a working **self-repairing engineering cognition runtime** that:

```
✅ Accepts mixed natural-language prompts
✅ Decomposes them into task DAGs
✅ Routes to specialized models (one at a time)
✅ Validates outputs
✅ Repairs failures automatically
✅ Retries only affected subtrees
✅ Maintains persistent state
✅ Asks for help when blocked
✅ Persists artifacts
✅ Keeps everything observable

All in a web interface you can monitor in real-time.
```

---

**Ready to dive deeper?**
→ Read: `COGNITION_ENGINE_MASTER_SPEC.md`

**Need quick answers?**
→ Read: `QUICK_REFERENCE_CARD.md`

**Ready to build the system further?**
→ Start: `python -m pytest tests/ -v`
