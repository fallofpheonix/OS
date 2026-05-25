# Import Migration Plan: Stage A (F0 Stabilization)

## Vision
Unify all core PhoenixOS components under the flat `phoenix/` package namespace (e.g., `phoenix/warden`, `phoenix/ledger`).

## Current State
- Fragmented imports: `github.com/fallofpheonix/phoenix-os/phoenix_os/...`
- Replacement hacks in `go.mod` and `go.work`.
- Mixed naming conventions.

## Strategy: Staged Transition
1. **Compatibility**: Keep old imports working via `go.work` and `replace` directives.
2. **Adapter Layer**: Create `phoenix/` wrappers for legacy packages.
3. **Migration**: Update files one by one to use the new namespace.
4. **Freeze**: Remove `go.work` local replacements once all core files are migrated.
5. **Removal**: Delete legacy import paths and folder structures.

## Timeline (F0-F1)
- **Stage 1 (Now)**: Define `IMPORT_MIGRATION_PLAN.md` and identify core targets.
- **Stage 2**: Implement `phoenix/common/logical_clock` as the first unified target.
- **Stage 3**: Migrate `warden`, `arbiter`, and `ledger` to `phoenix/` namespace.

## Enforcement
- `check_illegal_deps.py` will be updated to flag any new legacy imports.
- CI will fail if `experimental` modules are imported by `phoenix/core`.
