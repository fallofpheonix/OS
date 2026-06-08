# Vault Quick Reference

Your Obsidian vault is now structured as an engineering operating system.

---

## Navigation

### For Project Work (Bookmark These)

- **[[Operating-System]]** — Current phase, active tracks, milestones, operating cycle
- **[[AI Assistant]]** — Your primary project status and milestones
- **[[EXTRACTION_LOG]]** — Modules extracted and candidates tracked
- **[[Session-1-Execution-Plan]]** — Tomorrow's exact task list

### For Decision Making

- **04_ENGINEERING/decision-logs/** — Architecture decisions
  - [[ADR-001-Modular-Monolith]] — Why single repo with clean boundaries
  - [[ADR-002-Python-First]] — Why Python before Rust
  - [[ADR-003-Foundation-First]] — Why no AI until foundation proven

### For Learning From Failures

- **06_FAILURE_LIBRARY/** — Mistakes and how you fixed them
  - Will populate as Phase 1 produces real failures
  - Create with template: [[_TEMPLATES/Failure-Note]]
  
- **04_ENGINEERING/debugging-patterns/** — How you solved hard problems
  - Create with template: [[_TEMPLATES/Debugging-Session]]

### For Capture and Processing

- **01_CAPTURE/inbox.md** — Quick dumps during sessions
- **01_CAPTURE/debugging/debugging-sessions/** — Debugging notes
- **Friday review process** — Process inbox into failure/learning notes

---

## Templates Available

**Location:** `_TEMPLATES/`

- **ADR.md** — Architecture decisions (decisions that affect structure)
- **Failure-Note.md** — Mistakes and recovery (engineering learning)
- **Debugging-Session.md** — Problem-solving sessions (methodology capture)
- **Module.md** — Extracted reusable components (Phase 3+)
- **Project-Active.md** — Active project status (use for secondaryproject too)

**Usage:** Copy template to appropriate location, fill in content

---

## Weekly Operating Cycle

### Monday-Thursday: Execution
- Build in primary project
- Log issues in 01_CAPTURE/inbox.md
- Minimal documentation

### Friday: Reflection
- Process inbox
- Create failure notes from week
- Create debugging pattern notes
- Update [[EXTRACTION_LOG]]
- Update [[Operating-System]] status

### Saturday: Architecture
- Review ADRs
- Refactor code
- Plan next week

### Sunday: Learning
- Study papers/books
- Experiment
- No deadline pressure

---

## Status Indicators

### Project Phases

- **PHASE 0 — FOUNDATION SETUP** ✓ COMPLETE (you are here)
- **PHASE 1 — PRIMARY SYSTEM IMPLEMENTATION** (tomorrow starts)
- **PHASE 2 — FAILURE + PATTERN ACCUMULATION** (week 2+)
- **PHASE 3 — EXTRACTION PHASE** (month 2)
- **PHASE 4 — SECOND SYSTEM REUSE** (month 3)
- **PHASE 5 — ECOSYSTEM CONSOLIDATION** (month 4)

### Project States

- **ACTIVE** — Currently working (AI Assistant)
- **SECONDARY** — Next project (not started)
- **LEARNING** — Research track (OS + Distributed Systems)
- **EXPERIMENTAL** — Playing with ideas
- **REUSABLE_MODULES** — Extracted code (Phase 3+)

---

## File Organization

```
brain/
├── 00_DASHBOARD/
│   ├── Home.md                      # Daily entry point
│   ├── Operating-System.md          # 📌 BOOKMARK THIS
│   ├── Project Dashboard.md
│   └── weekly-review.md
│
├── 01_CAPTURE/
│   ├── inbox.md                     # Session notes dump here
│   ├── debugging/
│   └── ideas/
│
├── 02_ACTIVE_LEARNING/
│   ├── current-books/
│   ├── current-topics/
│   └── experiments/
│
├── 03_CORE_KNOWLEDGE/
│   ├── ai-ml/
│   ├── databases/
│   ├── distributed-systems/
│   ├── os/
│   └── systems/
│
├── 04_ENGINEERING/
│   ├── decision-logs/                # ADR files go here
│   ├── debugging-patterns/           # Debugging session notes
│   ├── EXTRACTION_LOG.md             # 📌 BOOKMARK THIS
│   └── [other engineering notes]
│
├── 05_PROJECTS/
│   ├── ACTIVE/
│   │   ├── AI Assistant.md           # 📌 BOOKMARK THIS
│   │   └── Session-1-Execution-Plan.md
│   ├── SECONDARY/
│   ├── COMPLETED/
│   └── REUSABLE_MODULES/
│
├── 06_FAILURE_LIBRARY/
│   ├── [failure notes populate here]
│   └── [from execution pressure]
│
└── _TEMPLATES/
    ├── ADR.md
    ├── Failure-Note.md
    ├── Debugging-Session.md
    ├── Module.md
    └── Project-Active.md
```

---

## Bookmarks You Should Create

In Obsidian sidebar:

1. [[Operating-System]] → Daily operating system
2. [[AI Assistant]] → Primary project tracking
3. [[EXTRACTION_LOG]] → Module ecosystem
4. [[Session-1-Execution-Plan]] → Tomorrow's tasks
5. `04_ENGINEERING/decision-logs/` → Architecture decisions

These 5 become your operating dashboard.

---

## How to Use During Coding

### During Session

1. **Before starting:** Open [[Session-1-Execution-Plan]] (or next session equivalent)
2. **During work:** Dump issues/observations in 01_CAPTURE/inbox.md
3. **After work:** Note what worked, what didn't

### During Debugging

1. **Create note** with [[_TEMPLATES/Debugging-Session]]
2. **Document** investigation steps
3. **Capture** what you learned
4. **Friday:** Decide if this becomes Failure Note

### When You Make a Decision

1. **Create ADR** with [[_TEMPLATES/ADR.md]]
2. **Place in** `04_ENGINEERING/decision-logs/`
3. **Link** from relevant project note
4. **Decide status:** ACCEPTED | PROPOSED | DEPRECATED

### When Something Breaks

1. **During recovery:** Note in 01_CAPTURE/inbox.md
2. **After recovery:** Create [[_TEMPLATES/Failure-Note]]
3. **Place in** `06_FAILURE_LIBRARY/`
4. **Link** to relevant code or ADR

---

## Key Principles

### Rule 1: Code First, Documentation Later
- Document what you ACTUALLY did, not what you planned to do
- Failure notes come from real failures, not speculation

### Rule 2: Templates Over Custom Notes
- Use the templates provided
- They enforce the right structure
- No blank slate overthinking

### Rule 3: Link Everything
- Every failure links to related code
- Every ADR links to related decisions
- Every project links to relevant architecture

### Rule 4: Friday Processing
- Don't create complex notes during the week
- Dump raw observations in inbox
- Friday: transform into learning notes

### Rule 5: Extraction Log is Sacred
- Only extract modules after 2+ uses
- Track candidates first
- Update when new projects consume

---

## What NOT To Do

❌ Create 50 concept notes before building  
❌ Over-tag everything  
❌ Install 40 Obsidian plugins  
❌ Build giant MOCs (Maps of Content)  
❌ Spend hours styling/organizing  
❌ Extract modules before pattern repeats  

---

## Your Next Steps

**Today:** This setup is complete. Spend last 30 minutes reviewing [[Session-1-Execution-Plan]]

**Tomorrow:** Execute [[Session-1-Execution-Plan]] (2 hour session)

**End of Day 1 of Code:** You have working CLI → Orchestrator → Runtime execution

**End of Week 1:** You have tools (file, git, shell) + tests + failure notes

**End of Week 2:** You have local LLM integration + memory start

---

## Emergency Reference

**Where do I put...?**

| Item | Location | Template |
|------|----------|----------|
| Architecture decision | `04_ENGINEERING/decision-logs/` | ADR.md |
| Failure/mistake | `06_FAILURE_LIBRARY/` | Failure-Note.md |
| Debugging session | `04_ENGINEERING/debugging-patterns/` | Debugging-Session.md |
| New project | `05_PROJECTS/ACTIVE/` | Project-Active.md |
| Module extracted | `05_PROJECTS/REUSABLE_MODULES/` | Module.md |
| Quick notes | `01_CAPTURE/inbox.md` | None (free form) |
| Module tracking | `04_ENGINEERING/EXTRACTION_LOG.md` | Update existing |

---

## Success Metrics for Each Phase

### Phase 0 (Today) ✓
- [x] Vault structured
- [x] Templates created
- [x] Dashboards initialized
- [x] Session plan ready

### Phase 1 (This Week)
- [ ] Code initializing and running
- [ ] CLI + Orchestrator + Runtime working
- [ ] Basic tests passing
- [ ] First failure notes created

### Phase 2 (Next Week)
- [ ] File/Git tools working
- [ ] More failure notes
- [ ] Extraction candidates tracked
- [ ] First ADRs beyond initial 3

### Phase 3 (Month 2)
- [ ] Modules extracted
- [ ] Second project uses extracted modules
- [ ] API boundaries proven stable

---

## Remember

This vault is a **working notebook**, not a finished product.

It becomes valuable as:
- Failures accumulate
- Patterns emerge
- Decisions get questioned
- Modules get extracted
- Architecture proves itself through **real execution pressure**

The vault is only useful because of the code.

Start coding tomorrow. The vault captures what happens next.
