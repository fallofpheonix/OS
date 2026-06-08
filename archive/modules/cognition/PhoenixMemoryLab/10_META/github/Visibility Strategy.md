# GitHub Visibility Strategy

> Repository visibility reflects semantic maturity, not whether code "works."
> Public repos = compressed engineering understanding. Not raw experimentation history.
> Optimize for **clarity density**, not complexity density.

---

## Lifecycle Model

Visibility is NOT a binary private→public transition.

Four states:

```
Research          → exploring unknowns, no stability guarantees
    ↓
Pressure          → architecture under real load, boundaries forming
    ↓
Stabilized        → boundaries survived pressure, reusable value proven
    ↓
Public Artifact   → compressed, independently understandable
```

Many repos should NEVER leave Pressure phase. That is healthy, not a failure.
Premature exposure encourages premature polish instead of engineering evolution.

---

## Ecosystem Taxonomy

Five layers:

### 1. Cognition Layer — PRIVATE ONLY

Purpose: engineering memory.

| Repo | Content |
|---|---|
| `brain` | ADRs, failure topology, semantic maps, control-plane analysis, comparative topology |
| Ecosystem audit docs | Governance decisions (like this document) |

Never public. This is institutional memory, not portfolio.

### 2. Runtime Layer — PRIVATE DURING PRESSURE

Purpose: execution pressure generation.

| Repo | Stage | Notes |
|---|---|---|
| `forge-agent` | Pressure | Evolving runtime substrate |
| `control-plane-lab` | Research/Pressure | Orchestration experiments |
| `governance-lab` | Research/Pressure | Infrastructure governance |

These generate architectural learning. They may never stabilize enough to publish — that is fine.

### 3. Capability Layer — POTENTIALLY PUBLIC

Purpose: compressed reusable capability.

| Repo | Stage | Public Gate |
|---|---|---|
| `SmartAPILimiter` | Stabilizing | Needs: clean C API, bindings, benchmarks |
| `agentskill` | Stabilizing | Needs: ownership clarity, versioning |
| `shell-runtime` | Research | Needs: stable interface, docs |
| `filesystem-runtime` | Research | Needs: stable interface, docs |
| `visual-effects` | Research | Needs: reimplementation from fork analysis |
| `trace-semantics` | Research | Needs: bounded scope, examples |

These are the strongest future public candidates — bounded scope, understandable semantics, low control-plane dependency.

### 4. Research Layer — SELECTIVELY PUBLIC

Purpose: original systems-thinking output.

| Repo | Stage | Public Gate |
|---|---|---|
| `semantic-topology-study` | Research | Needs: reproducible methodology |
| `runtime-locality-study` | Research | Needs: clear thesis + evidence |
| `rendering-runtime-analysis` | Research | Needs: extracted from fork analysis |

Can become public when methodology is independently understandable.

### 5. Archive Layer — MOSTLY PRIVATE

Purpose: institutional memory (or deletion).

| Repo | Decision |
|---|---|
| `codes` | Archive immediately — coursework dump, negative signal |
| `truenotes` | Archive immediately — upstream badges, negative originality |
| `cv` | Make private or delete — not a software repo |
| `constellation_of_us` | Private until rebuilt |
| `myxiomi` | Private — no software content |
| `Noesis` | Split or archive — two unrelated projects |
| `my-portfolio` | Rebuild as curated project index |
| `cognitron-game` | Split or archive — two products in one repo |

---

## Domain Repo Lifecycle Positions

| Repo | Layer | Lifecycle Stage | Readiness Gate |
|---|---|---|---|
| `UDIE` | Capability | Pressure | Migration squashing, API contracts, deployment SLOs |
| `LAMP` | Capability | Pressure | GDAL reproducibility, benchmark datasets, ADRs |
| `healingstone` | Capability | Pressure | Remove backup tree, publish sample datasets |
| `sira` | Research | Pressure | Remove duplicated submission tree, experiment cards |
| `ParticleStimulator` | Capability | Pressure | Split archive, typed event schema, job queue |
| `LifeTrack` | Capability | Research/Pressure | Threat model, encrypted storage, real integrations |
| `AI4MH` | Capability | Research | Model cards, bias eval, auth, audit logs |
| `TrustLab` | Research | Research | Remove stale trees, schema versioning |
| `AI-PFI` | Capability | Research | Tests, CI, crawler contracts |
| `AutoTRandHD` | Research | Research | Externalize model artifacts, reproducibility |
| `ArtExtract` | Research | Research | Consolidate, experiment registry |
| `ChoreoAI` | Research | Research | Narrow scope, reproducible metrics |
| `TerraHerb` | Capability | Research | Fix CI, model provenance |
| `SecureForg` | Research | Research | Sandboxing, false-positive corpus |
| `AutoMation-Engine` | Archive | Research | Fundamental rebuild or archive |
| `AutoEIT Suite` (to merge) | Capability | Stabilizing | Merge transcription + scoring, package API |

### Forks — Temporary Analysis Specimens

```
fork → analyze → compress semantics → build original → delete fork
```

| Fork | Extraction Status | Action |
|---|---|---|
| `hermes-agent` | Active analysis | Extract agent patterns to brain → delete |
| `OsdagBridge` | Active analysis | Extract structural engineering learnings → delete |
| `MolecularNodes` | Active analysis | Extract 3D rendering patterns → delete |
| `gemini-cli` | Active analysis | Extract CLI agent patterns → delete |

Delete fork when:
- ✓ Analysis extracted to brain
- ✓ Patterns documented
- ✓ Useful primitives reimplemented as original capability
- ✓ No benchmarking need remains
- ✓ No upstream tracking needed

---

## What Should NEVER Be Pushed

### Runtime artifacts (any repo)

```
.venv / venv / node_modules
__pycache__ / *.pyc
build / dist / target
*.o / *.a / *.so / *.dylib / *.exe
embeddings / vector DBs
runtime caches
subprocess outputs
large binaries
AI model files / checkpoints
generated traces
local environment configs
.DS_Store / *.log
```

### Secrets (any repo, ever)

```
API keys
.env files with real values
runtime secrets
control-plane state
automation credentials
local indexing data
filesystem snapshots
```

---

## Structural Invariants

### The Separation

```
brain/      → cognition (markdown, analysis, patterns, ADRs)
workspace/  → execution (code, builds, deployments)
```

Never mix. Protect aggressively.

### Public Readiness Test

A repo should become public ONLY if:

```
Someone else can understand its boundaries
without your internal cognition system.
```

If a repo requires:
- Your brain vault to make sense
- Your local conventions
- Hidden operational context
- Undocumented control-plane semantics

Then it is NOT mature enough to be public.

### Clarity Density Rule

Public repos optimize for **clarity density** — not complexity density.

```
✓ Clear boundaries, obvious scope, independent value
✗ Interesting complexity that requires insider knowledge
```

This determines whether GitHub becomes:
- Respected engineering systems (clarity)
- Chaotic experimental accumulation (complexity)

### Visibility Transition Gates

```
Research → Pressure:
  Architecture exists. Boundaries forming.

Pressure → Stabilized:
  1. Architecture survived real load
  2. Boundaries are stable
  3. No operational residue
  4. Abstractions mature

Stabilized → Public Artifact:
  5. Someone external can understand the boundaries
  6. Documentation meets baseline contract
  7. CI/tests/security pass
  8. Independent usefulness proven
  9. Optimized for clarity density
```

Many repos stay in Pressure permanently. That is healthy engineering, not failure.

---

## Recommended Final GitHub Topology

### Phase 1 — Immediate (now)

- [x] `brain` → PRIVATE, git-synced, cognition only ✅
- [ ] Archive: `codes`, `truenotes`, `cv` (make private or delete)
- [ ] Remove committed binaries from: `AutoTRandHD` (best.pt), `audio_transcription` (model/audio), `SmartAPILimiter` (compiled .o)
- [ ] Rewrite `fallofpheonix` profile README

### Phase 2 — Consolidation (30 days)

- [ ] Merge `audio_transcription` + `AutoEIT-STS` → `autoeit-suite`
- [ ] Split or archive `Noesis`
- [ ] Split `cognitron-game`
- [ ] Delete extracted forks
- [ ] Add CI to: `AI-PFI`, `TrustLab`, `SmartAPILimiter`

### Phase 3 — Publication (60-90 days)

- [ ] Publish `SmartAPILimiter` (with C API, bindings, benchmarks)
- [ ] Publish `agentskill` (if ownership resolved)
- [ ] Promote primary capstones: `UDIE`, `LAMP`, `healingstone`, `sira`, `ParticleStimulator`
- [ ] Publish `autoeit-suite` as package

### Long-term Target

```
private cognition engine          → brain
    ↓
private pressure systems          → runtime + research repos
    ↓
compressed stable capabilities    → capability modules
    ↓
public research/capability        → published artifacts
```

```
5-6 primary public repos   → compressed engineering signal
brain (private forever)    → institutional memory
2-3 reusable packages      → published capability modules
0 stale forks              → semantic compression complete
profile README             → systems/platform engineering identity
```

---

## Commit Philosophy (All Repos)

Commits reflect semantic evolution:

```
✓ Add control-plane locality analysis
✓ Document coordination gravity signals
✓ Record runtime containment invariants
✓ Add comparative topology observations
```

Not:

```
✗ misc updates
✗ brain changes
✗ notes
✗ fix stuff
```

---

## Strategic Transition

```
Current:  fork → analyze (curator of external systems)

Target:   analyze → compress → build original → delete fork
          (builder of original systems)
```

This transition is critical. Without it, the ecosystem becomes a museum of other people's architectures instead of a source of original engineering capability.
