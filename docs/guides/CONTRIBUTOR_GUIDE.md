---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# PhoenixOS Contributor Guide

## Working Model & Guidelines
All contributions must align with the **4-Directory Fractal Model** and the **6 Core Axioms of the Phoenix Matrix**.

---

## Coding Standards & Rules

### When Modifying Code
1. **Read CLAUDE.md First**: Every main directory contains a `CLAUDE.md` detailing module-specific commands, targets, and styles. Read it before touching code in that directory.
2. **Comment Enforcement**: Every file header must match the mandatory commenting pattern, specifying File, Purpose, Subsystem, Dependencies, and Security considerations.
3. **Strict Typings**: Never write type assertions without checking (`value, ok := ...`).
4. **No Scaled Instability**: Fix unit tests and race conditions locally before proposing distributed integration.
5. **Clean Go Workspace**: Keep `go.work` clean of directories without `go.mod` files.

---

## Workspace Layout
- `Phoenix.Nucleus/`: Low-level system components. No imports from external layers.
- `Phoenix.Cognition/`: Decision engines and advisors. Imports contracts and core.
- `Phoenix.Crucible/`: Simulation and verification. Standalone tests and stress-test engines.
- `Phoenix.Terminus/`: The edge systems and user interfaces. Imports everything.

---

## Commit & PR Checklist
- [ ] Code compiles without warnings (`go build ./...`).
- [ ] Test suite passes with race detection (`go test -race ./...`).
- [ ] Code formatting conforms to `go fmt`.
- [ ] Mandatory comment headers are present.
- [ ] No hardcoded configuration parameters or credentials.
