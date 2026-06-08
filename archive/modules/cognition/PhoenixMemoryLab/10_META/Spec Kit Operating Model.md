# Spec Kit Operating Model

## Invariant
Brain stores thinking. Workspace stores execution. Spec Kit lives inside each workspace project, never inside the vault.

## Placement
- Global tool: `specify`, installed once with `uv tool install specify-cli --from git+https://github.com/github/spec-kit.git`
- Per-project execution root: `~/engineering/workspace/active/project-name/`
- Per-project Spec Kit artifacts: `~/engineering/workspace/active/project-name/.specify/`
- Brain references Spec Kit artifacts by path. Brain does not copy `spec.md`, `plan.md`, `tasks.md`, `research.md`, or `data-model.md`.

## Project Cycle
```text
Brain idea -> Spec Kit spec -> Spec Kit plan/tasks -> implementation -> Brain extraction
```

## New Project Sequence
1. Capture raw idea in `01_CAPTURE/inbox.md`.
2. Create `05_PROJECTS/ACTIVE/project-name/Project.md` from `10_META/templates/Project.md`.
3. Create workspace directory under `~/engineering/workspace/active/project-name/`.
4. Run independence checklist.
5. Initialize Spec Kit once:

```bash
cd ~/engineering/workspace/active/project-name
specify init . --integration claude
```

## Feature Sequence
1. `/speckit.constitution`
2. `/speckit.specify`
3. `/speckit.clarify`
4. `/speckit.plan`
5. Review `plan.md`, `research.md`, and `data-model.md`.
6. `/speckit.tasks`
7. Review `tasks.md`.
8. `/speckit.implement`
9. Verify with tests, lint, Docker, and health checks.
10. Extract ADRs, failures, and concept updates back into the brain.

## Extraction Requirements
- Minimum one ADR per shipped feature.
- Failure notes for every bug or unexpected behavior during implementation or verification.
- Concept notes strengthened from `research.md` and `plan.md`.
- Project `Features Shipped` row updated before marking a feature `SHIPPED`.

## Anti-Patterns
- Do not run `specify init` more than once per project.
- Do not install `specify-cli` from PyPI.
- Do not place `.specify/` anywhere under `~/engineering/brain/`.
- Do not copy Spec Kit generated files into the brain.
- Do not implement before spec, clarification, plan, and tasks are reviewed.

## Verification
```bash
specify version
find ~/engineering/brain -name .specify -type d
find ~/engineering/workspace -maxdepth 3 -name .specify -type d
```

