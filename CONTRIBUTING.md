# PhoenixOS Contribution Guidelines

Welcome to the PhoenixOS project! We are building an autonomous, security-first Operating System. We welcome contributions from both humans and AI agents.

## 1. Code of Conduct
We adhere to a standard professional Code of Conduct:
- Be respectful, professional, and collaborative.
- Use inclusive language.
- Assume positive intent.
- Harassment or demeaning behavior will not be tolerated.

## 2. Contribution Workflow (Agents & Humans)

### Step 1: Claiming Issues
1. Search the [GitHub Issue Tracker](https://github.com/fallofpheonix/PheonixOS/issues) for tasks labeled `ready-for-agent`.
2. Comment on the issue stating intent to claim it (e.g., "I am claiming this issue").

### Step 2: Implementation
1. Create a branch: `feature/<issue-number>-<brief-title>` or `fix/<issue-number>-<brief-title>`.
2. Implement your changes, prioritizing:
   - **Idiomatic Go/Python:** Follow existing project conventions.
   - **Test-First:** Always add a corresponding test file (e.g., `_test.go` or `test_*.py`).
   - **Verification:** Run existing CI checks (`make all`, `pytest`).

### Step 3: Pull Request & Merge
1. Create a PR with a description referencing the issue ID: `Resolves #<issue-id>`.
2. Ensure the PR is **MERGEABLE** (no conflicts).
3. Once passed, the Orchestrator agent will automatically trigger the merge if criteria are met.

## 3. Style & Standards
- **Documentation:** Every architectural change must update `PHOENIX_TASKS.md` and relevant `docs/` files.
- **Testing:** 100% test coverage is preferred; minimum 90% pass rate is enforced.
- **CI/CD:** Pipelines are automatically triggered on PR creation.
