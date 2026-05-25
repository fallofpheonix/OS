# CI Enforcement Plan

## 1. Governance Engine Execution
- **Trigger**: The `.github/workflows/architecture_governance.yml` runs on every `push` and `pull_request` to the `main` branch.
- **Job**: `invariant-enforcement` runs on an Ubuntu runner. It parses the `invariants.yaml` file across the repository architecture.

## 2. Invariant Rules Tested
The CI checks structural invariants. Examples include:
- `modules_no_apps`: modules cannot depend on apps.
- `infra_no_experiments`: infrastructure cannot depend on experiments.
- `research_isolation`: research cannot depend on apps or services.
- `sdks_only_contracts`: SDKs can only import contracts and standards.
- `no_global_cycles`: Global cycle detection ensures no circular dependencies.

## 3. Failure Behavior
- If any imported dependencies violate the `forbidden_imports` or fail to match `allowed_imports`, the CI script `ci_invariant_check.py` returns exit code 1.
- PRs violating the architecture boundaries will be rejected by GitHub checks, preventing merge.

## 4. Future Additions (Next Iterations)
- **Mutation CI**: Hooking into git diffs to detect if forbidden repositories are mutated without proper approval.
- **Replay CI**: Automatically verifying deterministic behavior for PRs touching state machines.
