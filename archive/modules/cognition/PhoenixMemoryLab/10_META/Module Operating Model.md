# Module Operating Model

## Definition
A module is reusable code imported by two or more standalone projects. Module code lives in `~/engineering/infrastructure/shared-libraries/<module-name>/`. Module thinking lives in `08_MODULES/<module-slug>.md`.

## Non-Project Rule
Modules are not projects. They have no Spec Kit lifecycle, no `spec.md`, and no shipping milestone. They track interface contracts, usage graph, version state, and known failure modes.

## Status State Machine
- EXPERIMENTAL: public API unstable; not for production projects.
- ACTIVE: used by at least one project; API stable for minor versions.
- STABLE: 3+ months in use, no open bugs, 80%+ coverage.
- DEPRECATED: phased out; note must include a Migration section.

Normal transition: `EXPERIMENTAL -> ACTIVE -> STABLE`. `DEPRECATED` may apply from any state.

## Creation Trigger
Create a module note only when code is imported by a second project. Before that, keep it inline inside the first project.

## Placement
- Template: `10_META/templates/Module.md`
- Notes: `08_MODULES/<module-slug>.md`
- Code: `~/engineering/infrastructure/shared-libraries/<module-name>/`
- Archive: `09_ARCHIVE/modules/`

## Naming
- Lowercase, hyphen-separated slug.
- Format: `<domain>-<function>`.
- Filename slug must match `module-id` frontmatter exactly.
- No underscores.

## Hard Completion Rules
- Minimum two links in `Linked Concepts` to `03_CORE_KNOWLEDGE/`.
- Minimum one entry in `Used By`.
- ACTIVE modules require at least one `Known Failure Modes` entry.
- `Public API` table must contain at least one real signature.

## Soft Review Rules
- Coverage below 70% requires `needs-tests` tag.
- `0.x.x` with STABLE status is contradictory.
- 3+ callers with no linked ADR requires architecture review.

## Project Integration
When a project uses a module:
- Update that project's `Project.md` `Modules Used` table.
- Update `.specify/memory/constitution.md` under `## Dependencies (shared libraries)`.
- Keep versions synchronized between brain and Spec Kit constitution.

## Archive Rule
Never delete module notes directly. Set `status: DEPRECATED`, add Migration if known, then move to `09_ARCHIVE/modules/`.

