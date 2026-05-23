# Domain Documentation

This project uses a multi-context documentation layout.

## Layout

- **Root:** `CONTEXT-MAP.md` points to per-context `CONTEXT.md` files.
- **ADRs:** Architectural Decision Records are located in `docs/adr/` or per-subsystem `docs/*/adr/`.

## Consumer Rules

- Always check `CONTEXT-MAP.md` first to find the relevant context for a file or task.
- Respect the boundaries defined in per-context `CONTEXT.md` files.
