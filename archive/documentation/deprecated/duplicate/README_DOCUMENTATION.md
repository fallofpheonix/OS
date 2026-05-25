# COGNITION ENGINE - DOCUMENTATION SUMMARY

**Three documents. One complete system specification.**

---

## WHAT YOU HAVE

### 📘 Master Specification (COGNITION_ENGINE_MASTER_SPEC.md)
**READ THIS FIRST**
- 25,000+ words
- Complete system design
- Every component detailed
- Execution rules (10 critical rules)
- UI/UX specification
- Long-running operations guide
- Troubleshooting for every common problem
- Deployment checklist

**Use for**: Understanding the system, implementing features, debugging complex issues

**Time to read**: 2-3 hours (skim first, reference later)

---

### 📋 Quick Reference Card (QUICK_REFERENCE_CARD.md)
**PRINT THIS AND KEEP IT NEARBY**
- 15 pages
- Hard constraints (never violate these)
- Daily checklist
- All 10 execution rules at a glance
- Common commands
- Metrics to watch
- Quick troubleshooting lookup table
- Emergency procedures

**Use for**: Daily development, quick lookup, staying on track

**Time to read**: 15 minutes (reference as needed)

---

### 🚀 Getting Started (GETTING_STARTED.md)
**START HERE IF YOU'RE NEW**
- Step-by-step walkthrough
- 30 minutes to first working run
- UI explanation with visuals
- 3 example workflows
- Common mistakes and fixes
- Verification checklist

**Use for**: Onboarding new developers, verifying installation

**Time to read**: 20 minutes (hands-on)

---

## HOW TO USE THESE DOCUMENTS

### Scenario 1: You're Starting From Scratch
```
1. Read: GETTING_STARTED.md (20 min)
   → Get system running
   → Verify first control-plane works

2. Skim: COGNITION_ENGINE_MASTER_SPEC.md sections:
   - ARCHITECTURE OVERVIEW (15 min)
   - HARD CONSTRAINTS (10 min)
   - EXECUTION RULES (15 min)

3. Print: QUICK_REFERENCE_CARD.md
   → Keep at desk
```

### Scenario 2: You're Implementing a Feature
```
1. Reference: QUICK_REFERENCE_CARD.md
   → Check hard constraints
   → Verify you're not breaking rules

2. Read: COGNITION_ENGINE_MASTER_SPEC.md
   → Find relevant architecture section
   → Understand how feature fits in

3. Check: Execution Rules #1-10
   → Ensure feature complies

4. Write tests first
5. Implement
6. Verify metrics stay healthy
```

### Scenario 3: Something's Broken
```
1. Check: QUICK_REFERENCE_CARD.md
   → "Troubleshooting Quick Answers" table
   → "Emergency Procedures"

2. If still stuck, read: COGNITION_ENGINE_MASTER_SPEC.md
   → Search for specific error
   → TROUBLESHOOTING GUIDE section
   → Detailed diagnostic procedures

3. Run debug scripts mentioned in Master Spec
```

### Scenario 4: Long-Running Development
```
Morning (5 min):
  → Print QUICK_REFERENCE_CARD.md checklist
  → Run daily startup commands

During day:
  → Reference QUICK_REFERENCE_CARD.md for quick answers
  → Read specific MASTER_SPEC.md sections as needed

End of day:
  → Review metrics (QUICK_REFERENCE_CARD section)
  → If metrics degraded, debug using MASTER_SPEC.md
```

---

## CRITICAL INFORMATION (READ IMMEDIATELY)

### The 10 Execution Rules (Master Spec, simplified)
```
1. Only ONE model runs at a time (asyncio.Lock required)
2. Offline first (local Ollama only)
3. Validate everything before use
4. Snapshot before mutations
5. Append-only events (never delete)
6. Failures become structured FailureRecords
7. Retry only affected subtrees (not entire DAG)
8. Emit events for every state change
9. Enforce retry budgets (max 3 attempts)
10. Escalate permissions (explicit approval required)

VIOLATING ANY OF THESE BREAKS THE SYSTEM
```

### The 5 Hard Constraints
```
✓ One active model only
✓ Offline first
✓ Append-only events
✓ Bounded retries
✓ Deterministic execution
```

### Current Project Status
```
✅ Implementation: 80% complete
✅ Tested: 20+ test cases passing
✅ Framework: 6 core layers working
✅ Models: 5 specialized models integrated
✅ Memory: 4 separate systems ready

⚠️  Current Focus: Runtime hardening & stability
⚠️  Next Priority: Orchestration stability (Week 1)
```

---

## FILE LOCATIONS

```
astraeus-core/
├── COGNITION_ENGINE_MASTER_SPEC.md    ← READ THIS FIRST
├── QUICK_REFERENCE_CARD.md            ← PRINT THIS
├── GETTING_STARTED.md                 ← START HERE
├── README_DOCUMENTATION.md            ← YOU ARE HERE
│
├── contracts/
│   └── models.py                      ← Data classes
│
├── orchestrator/
│   └── engine.py                      ← Main entry point
│
├── models/
│   ├── base_adapter.py                ← How models work
│   └── {qwen,deepseek,mistral,codellama,phi}_adapter.py
│
├── memory/
│   ├── memory_system.py               ← 4 memory types
│   └── ...
│
├── repair/
│   └── repair_planner.py              ← Self-repair
│
├── cli/
│   └── main.py                        ← Command line
│
├── api/
│   └── main.py                        ← FastAPI server
│
├── frontend-console/
│   └── index.html                     ← Web UI
│
└── tests/
    └── test_*.py                      ← Test suite
```

---

## RECOMMENDED READING ORDER

### For New Developers (Complete Onboarding)
```
1. README_DOCUMENTATION.md (this file)        [5 min]
   → Understand what you're working with

2. GETTING_STARTED.md                         [20 min]
   → Get system running + first control-plane

3. QUICK_REFERENCE_CARD.md                    [15 min]
   → Print it, understand daily ops

4. COGNITION_ENGINE_MASTER_SPEC.md:
   a. Executive Overview                      [10 min]
   b. Architecture Overview                   [20 min]
   c. Core Components                         [30 min]
   d. Execution Rules section                 [20 min]

5. Rest of MASTER_SPEC as needed              [Reference]
```

**Total time**: ~2 hours (then reference as needed)

### For Experienced Developers (Quick Start)
```
1. GETTING_STARTED.md                         [10 min]
2. QUICK_REFERENCE_CARD.md                    [10 min]
3. Specific MASTER_SPEC.md sections           [As needed]
```

**Total time**: ~20 min to get running

### For Architecture Reviews
```
1. COGNITION_ENGINE_MASTER_SPEC.md:
   - Executive Overview
   - Current State Assessment
   - Architecture Overview
   - Hard Constraints
   - Implementation Specification
   - Execution Rules
```

**Total time**: 1-2 hours

---

## KEY COMMANDS

### Verify System is Healthy
```bash
python scripts/verify_ollama.py
python scripts/check_db_integrity.py
```

### Run a Simple Orchestration
```bash
python -m cli.main "Build a simple REST API" \
  --artifacts /tmp/artifacts \
  --data /tmp/data
```

### Test Everything
```bash
python -m pytest tests/ -v
```

### Replay a Failed Run
```bash
python scripts/replay_run.py run_123
```

### Monitor a Long-Running Session
```bash
watch -n 2 'sqlite3 data/cognition.db "SELECT COUNT(*) as completed_tasks FROM tasks WHERE run_id=\"run_123\" AND status=\"completed\""'
```

---

## METRICS CHECKLIST

After every run, verify:
```
✓ Task completion rate: >70%
✓ Repair success rate: >50%
✓ Avg task latency: <3000ms
✓ Token usage: <5000/task
✓ Failure count: <3
✓ Approval requests: 0-2 (appropriate)
✓ Memory usage: stable
✓ Event log complete
```

If any metric degrades: **INVESTIGATE IMMEDIATELY**

---

## WHAT SUCCESS LOOKS LIKE

### Week 1
- Orchestration stable (100+ runs)
- DAG decomposition working
- Model routing working
- Artifacts generated
- Tests passing

### Week 2-3
- Repair loops working
- Common failures auto-fixed
- Retries localized
- Failure taxonomy growing

### Week 4+
- Long sessions stable (8+ hours)
- Repository intelligence working
- Safe mutations guaranteed
- Permission system enforced
- Help system functional

### Month 2
- Self-sustaining system
- Minimal human intervention
- Observable + recoverable
- Production-ready

---

## COMMUNICATION PROTOCOL

When handing off work or asking for help, use this template:

```markdown
## Current Status

- Running since: [timestamp]
- Last successful run: [run_id]
- Current issue: [description]
- Metrics impact: [what changed]

## What I Tried

1. [Action 1] → Result
2. [Action 2] → Result
3. [Action 3] → Result

## What's Blocked

- [Specific blocker]
- Evidence: [log excerpt, error, metric]

## Next Steps (My Recommendation)

1. [What should happen next]
2. [Why this approach]
3. [Expected outcome]

## Questions

- [What I need help with]
```

---

## WHAT'S ALREADY WORKING

```
✅ Planner (task decomposition)
✅ DAG control-plane
✅ Model routing (5 models)
✅ Sequential execution
✅ Artifact system
✅ Event bus
✅ Memory systems (4 types)
✅ Validator + failure capture
✅ Repair planner
✅ Snapshots + rollback
✅ Help request system
✅ Transaction framework
✅ Repo indexer
✅ Test suite (20+ tests)
```

Your job: **Harden, verify, extend incrementally**

---

## WHAT NEEDS WORK

```
🔄 In Progress:
  - Orchestration stability (focus: this week)
  - Repair loop maturity
  - Repository awareness completion
  - Transaction safety

⏳ Next:
  - Long-running session stability
  - Permission escalation UI
  - Help system integration
  - Advanced repair strategies
```

---

## GOLDEN RULES

```
✅ Code changes in small increments
✅ Test after every change
✅ Keep logs append-only
✅ Metrics tell the truth
✅ Replay reveals problems
✅ When in doubt, ask metrics
✅ When stuck, check events log
✅ Failures are features (not bugs)
✅ Stability > Features
✅ Observable > Magic
```

---

## EMERGENCY CONTACTS

### Models timeout forever
→ See: QUICK_REFERENCE_CARD.md > EMERGENCY PROCEDURES > "If Models All Timeout"

### Database corrupted
→ See: QUICK_REFERENCE_CARD.md > EMERGENCY PROCEDURES > "If Database Corrupted"

### WebSocket hangs
→ See: QUICK_REFERENCE_CARD.md > EMERGENCY PROCEDURES > "If WebSocket Hangs"

### Stuck in infinite loop
→ See: QUICK_REFERENCE_CARD.md > EMERGENCY PROCEDURES > "If Stuck in Repair Loop"

### Memory explodes
→ See: QUICK_REFERENCE_CARD.md > EMERGENCY PROCEDURES > "If Memory Explodes"

---

## FINAL CHECKLIST BEFORE STARTING

```
Before you begin development:

□ Have you read GETTING_STARTED.md?
□ Can you start Ollama? (ollama serve)
□ Can you start API? (python -m api.main)
□ Can you access UI? (http://localhost:8000)
□ Did first control-plane work?
□ Can you see metrics? (Bottom panel)
□ Can you replay a run? (python scripts/replay_run.py)
□ Do you have QUICK_REFERENCE_CARD.md printed?
□ Do you understand the 10 Execution Rules?
□ Do you know where to find MASTER_SPEC.md?

If you answered YES to all: YOU'RE READY TO BUILD
If you answered NO to any: Read corresponding guide above
```

---

## DOCUMENT REFERENCE MAP

```
Question                                   → Document
─────────────────────────────────────────────────────────
"How do I get started?"                    → GETTING_STARTED.md
"What's the system architecture?"          → MASTER_SPEC.md: Architecture
"How does execution work?"                 → MASTER_SPEC.md: Execution Rules
"What are hard constraints?"               → QUICK_REF.md or MASTER_SPEC.md: Hard Constraints
"How do I add a feature?"                  → QUICK_REF.md: Feature Addition Checklist
"Something broke, help!"                   → QUICK_REF.md: Troubleshooting
"How do I monitor it?"                     → MASTER_SPEC.md: Verification & Monitoring
"How do I deploy this?"                    → MASTER_SPEC.md: Deployment Checklist
"What's the memory model?"                 → MASTER_SPEC.md: Layer 6 + QUICK_REF.md
"How do repairs work?"                     → MASTER_SPEC.md: Layer 5
"How do I debug?"                          → MASTER_SPEC.md: Troubleshooting Guide
"What metrics matter?"                     → QUICK_REF.md: Metrics to Watch
"I need help NOW"                          → QUICK_REF.md: Emergency Procedures
"Long sessions, how?"                      → MASTER_SPEC.md: Long-Running Operation
"Is this production-ready?"                → MASTER_SPEC.md: Deployment Checklist
```

---

## NEXT STEPS RIGHT NOW

### If you're new:
1. Open GETTING_STARTED.md
2. Follow steps 1-5 (should take 20 minutes)
3. Run your first control-plane
4. Come back here

### If you're extending:
1. Skim QUICK_REFERENCE_CARD.md (10 min)
2. Find your feature in MASTER_SPEC.md
3. Check: Does it violate hard constraints?
4. Write test first
5. Implement
6. Verify metrics

### If you're debugging:
1. Note the symptom
2. Check QUICK_REFERENCE_CARD.md > Troubleshooting
3. If not there, search MASTER_SPEC.md > Troubleshooting Guide
4. Follow diagnostic steps
5. Apply fix
6. Verify metrics recover

---

## DOCUMENT MAINTENANCE

These docs should stay in sync:

```
When you:                          Update these docs:
─────────────────────────────────────────────────
Add a new failure type            MASTER_SPEC.md: Failure Types
Add a new metric                  QUICK_REF.md: Metrics to Watch
Add a new rule                    QUICK_REF.md: Hard Constraints
Change execution flow             MASTER_SPEC.md: Orchestration Flow
Add emergency procedure           QUICK_REF.md: Emergency Procedures
Change UI layout                  MASTER_SPEC.md: UI Specification
Add common troubleshooting issue  MASTER_SPEC.md: Troubleshooting
```

---

## SUMMARY

**You now have everything needed to:**

✅ Understand the system (MASTER_SPEC.md)
✅ Operate it daily (QUICK_REF.md)
✅ Get started in 30 minutes (GETTING_STARTED.md)
✅ Debug problems (Troubleshooting in both)
✅ Add features safely (Rules + patterns)
✅ Deploy confidently (Checklist)
✅ Hand off to others (This protocol)

**Three documents. One complete system. Ready to build.**

---

**Start here**: GETTING_STARTED.md → Get running in 30 minutes

**Refer here**: QUICK_REFERENCE_CARD.md → Keep at desk

**Deep dive here**: COGNITION_ENGINE_MASTER_SPEC.md → Understand everything

---

**Questions?** All answers are in these three documents. Use the Reference Map above to find them.

**Ready?** Open GETTING_STARTED.md and begin.
