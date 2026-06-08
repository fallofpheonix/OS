# Weekly Review

## Capture Processing
- [ ] Clear `01_CAPTURE/inbox.md`.
- [ ] Convert spec/plan decisions into ADRs.
- [ ] Convert bugs/errors into Failure notes.
- [ ] Convert discovered concepts into Core Knowledge notes.
- [ ] Update active project notes.

## Spec Kit Sync
- [ ] Every shipped feature has a row in the project `Features Shipped` table.
- [ ] Every shipped feature has at least one ADR.
- [ ] Every failure referenced by a shipped feature exists in `06_FAILURE_LIBRARY/`.
- [ ] No `.specify/` directory exists under `~/engineering/brain/`.

## Module System Recovery Checklist
- [ ] Every module in `08_MODULES/` has at least 1 entry in `Used By`.
- [ ] Every ACTIVE module has at least 1 entry in `Known Failure Modes`.
- [ ] Every module version in Project.md tables matches the version in `08_MODULES/` frontmatter.
- [ ] No module slug appears in a Project.md without a corresponding `08_MODULES/<slug>.md` note.
- [ ] Test coverage for STABLE modules is >= 80%.

## Module Degradation Signals
- Duplicate logic: same function implemented in two project repos without a shared module.
- Orphaned modules: module note exists with no caller in `Used By`.
- Version drift: projects reference different module versions without an ADR.
- Silent failures: module source bug fixed without Failure note and module failure-mode update.
- Phantom dependencies: Project.md references a module missing from `08_MODULES/`.

