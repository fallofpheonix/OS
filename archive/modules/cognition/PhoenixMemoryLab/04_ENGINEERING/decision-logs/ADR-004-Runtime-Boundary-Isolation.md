# ADR-004: Runtime Boundary Isolation

**Date:** 2026-05-12  
**Status:** ACCEPTED

## Context
The execution substrate has begun to stabilize around a controlled shell runtime. Early work showed that keeping executable code outside `brain/` is not optional; it is a necessary architectural boundary. As additional runtime categories emerge, we need a clear rule for how they are isolated, validated, and composed.

## Decision
All executable runtime behavior will live under `engineering/workspace/`, with each runtime category isolated in its own namespace. Shell execution is centralized in `runtime/shell/`, and filesystem operations will be centralized in `runtime/filesystem/`. Upper layers may orchestrate and serialize results, but they must not contain subprocess logic or direct filesystem primitives.

## Rationale
This separation preserves the vault as cognition and decision infrastructure while keeping execution in the workspace. It prevents boundary collapse as new runtime categories appear. Typed contracts make cross-layer behavior explicit and testable, and centralized policy enforcement reduces accidental exposure of unsafe operations.

## Consequences
Runtime code stays reusable and testable. CLI and orchestrator layers can stabilize around explicit contracts instead of ad hoc dicts. Filesystem access can be guarded with path normalization and read-only policy before write-capable operations are introduced. The tradeoff is that new runtime categories must be added deliberately instead of being embedded directly in higher layers.

## Alternatives Considered
- **Keep runtime logic in the CLI/orchestrator:** rejected because it would blur boundaries and make policy enforcement inconsistent.
- **Use a generic runtime abstraction immediately:** rejected because pressure has not yet justified a shared base result model.
- **Allow direct filesystem access from upper layers:** rejected because it would weaken isolation and make unsafe operations easier to introduce.

## Related
- Related ADR: [[ADR-003-Foundation-First]]
- Related Failure: [[2026-05-executable-code-in-brain]]

## Implementation Notes
- Keep shell execution in `runtime/shell/executor.py`.
- Add filesystem operations in `runtime/filesystem/` with read-only behavior first.
- Preserve typed result contracts per runtime category.
- Validate changes with unit and integration tests before expanding the runtime surface.
