# Ecosystem Extraction Strategy

> A module gets its own repo ONLY when it represents compressed understanding, not ongoing confusion.
> Premature repo splitting creates dependency fragmentation, semantic instability, and version chaos.

---

## The Extraction Rule

A module gets its own repo ONLY when ALL of the following are true:

```
✓ stable semantics
✓ reused across projects
✓ independent usefulness
✓ low cognition leakage
✓ bounded topology
✓ understandable in isolation
```

If not: keep it internal.

### The Survival Test

Before extracting a module, ask:
**Can this module survive without the parent system's cognition?**

- **GOOD (`filesystem-runtime`)**: Semantics are local, bounded, and rely on explicit containment semantics.
- **BAD (`control-plane-engine`)**: Depends on evolving topology assumptions, containment semantics, pressure observations, and governance.

---

## Ecosystem Phasing

### Phase 1 — Internal Modules (Current State)

Goal: **Pressure Generation**.

```
forge-agent/
├── runtime/
├── tracing/
├── control-plane/
├── filesystem/
├── shell/
```

*Status: Keep together. Do not turn every folder into a repo.*
*Specific Decision: Keep `control-plane` inside `forge-agent`. It has not earned independence yet.*

### Phase 2 — Stabilized Capabilities

After repeated reuse, bounded semantics emerge:

```
modules/
├── shell-runtime/
├── filesystem-runtime/
├── trace-semantics/
```

### Phase 3 — Ecosystem Extraction

Eventually, bounded capabilities publicize well. Pressure systems usually do not.

**PUBLIC:**
- `github.com/fallofpheonix/shell-runtime`
- `github.com/fallofpheonix/filesystem-runtime`
- `github.com/fallofpheonix/trace-semantics`

**PRIVATE:**
- `github.com/fallofpheonix/control-plane-lab`
- `github.com/fallofpheonix/runtime-governance`

---

## Extraction Candidates

### Good Extraction Candidates (Later)
These will eventually become separate repos because their semantics are bounded, reusable, and have low cognition leakage.

1. `shell-runtime`: Bounded semantics, topology-simple, reusable.
2. `filesystem-runtime`: Clear boundary, explicit containment semantics, reusable trust logic.
3. `trace-semantics`: Immutable local tracing, bounded observability, independently useful.
4. `visual-effects`: Capability-oriented, isolated, reusable (from liquid-glass-js, shadergradient).

### Do NOT Split (Yet or Ever)
These are pressure-generation systems, not stable capabilities. Their semantics are still evolving.

- `control-plane/`
- `governance/`
- `coordination/`
- `semantic-topology/`

### The Centralization Trap
Do NOT create repos named:
- `shared-utils`
- `common-core`
- `base-framework`
- `universal-runtime`

These are premature centralization traps. Ecosystems rot there.

---

## Long-Term Repo Topology

**PRIVATE (Pressure & Cognition)**
- `brain`
- `forge-agent`
- `control-plane-lab`
- `runtime-governance`
- `semantic-topology-lab`

**PUBLIC (Capabilities)**
- `shell-runtime`
- `filesystem-runtime`
- `trace-semantics`
- `visual-effects`
- `render-runtime-primitives`

**ARCHIVE (Dead Weight)**
- `forks`
- `dead experiments`
- `old pressure systems`
- `obsolete architectures`
