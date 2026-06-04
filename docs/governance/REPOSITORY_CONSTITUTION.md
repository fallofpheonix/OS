---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# PhoenixOS Repository Constitution

This document defines the mandatory invariants and governance rules for all code and files within the PhoenixOS ecosystem.

## 1. Naming Standards
- **Modules/Subsystems:** Must use PascalCase with the `Phoenix.` prefix for monorepo module names (e.g., `Phoenix.Nucleus`).
- **Contracts:** Must use lowercase, slash-separated boundary names under `contracts/` (e.g., `contracts/events/v1`).
- **Packages:** Must be short, single-word, lowercase, and descriptive (e.g., `runtime`, `ledger`, `guard`).
- **Files:** Must use `snake_case.go` for Go files and uppercase markdown names for canonical architecture and governance documents.

## 2. Packages

### PACKAGE-001: No Empty Package Shells
A package may not exist solely to satisfy an import path. Every non-test package must contain production logic, domain models, or interfaces. Empty packages, placeholder structures, and import-only packages are prohibited.

### PACKAGE-002: No Silent Exported Stubs
An exported function, method, or constructor may not unconditionally return placeholder success values (e.g., `return nil`, `return true`). Every exported API must contain real production logic or return an explicit `ErrNotImplemented`.

### PACKAGE-003: No Fake Security Controls
Security-related exported APIs (Authentication, Authorization, Validation, Encryption) may not contain placeholder implementations. Security stubs must fail closed and return explicit errors.

## 3. Dependencies
- **Strict Hierarchy:** The Dependency Graph is enforced. `Phoenix.Nucleus` is the canonical root. No package in Nucleus may depend on `Cognition`, `Crucible`, or `Terminus`.
- **No Circular Imports:** Circular dependencies are structurally forbidden and will fail the build.
- **Third-Party:** External dependencies must be minimized, audited, and isolated within `Phoenix.Terminus` or explicitly authorized modules.

## 4. Testing
- **Determinism:** All tests must be 100% deterministic. Flaky tests are considered broken invariants.
- **Coverage:** Core invariants and state transitions require exhaustive test coverage.
- **Chaos & Fuzzing:** Chaos testing and fuzzing must be executed in bounded environments within `Phoenix.Crucible`.

## 5. Documentation

### The Core Invariant
**NO FILE EXISTS WITHOUT DOCUMENTATION.** Canonical architecture and governance files must remain synchronized with implementation changes.

### Mandatory File Metadata
Every file MUST contain a standardized header block detailing:
1. File Header (Name, Path, Purpose)
2. Subsystem (Owner Domain, Sub-domain)
3. Dependencies & Dependents
4. Security & Performance Notes
5. Status Metadata: `[STATUS: <STATE>]`, `[OWNER: <DOMAIN>]`

### Documentation Flow
A change is INVALID if implementation changes without documentation changes. Any agent or developer modifying a canonical boundary file MUST update the file header, dependency graph, architecture map, and any affected contract or debt register entries.

## 6. Versioning
- **Semantic Versioning:** All exported APIs and artifacts must use strict semantic versioning.
- **Immutability:** Once a contract version is published, it is immutable. Breaking changes require a new major version.
- **Ledger Compatibility:** Changes to event schema must include migration paths or maintain backward compatibility to guarantee replay.

## 7. Contract Ownership
- Only contract packages may define public interfaces.
- Implementations may only satisfy contracts; they may not redefine them.
- Contract compatibility tests are required before extraction of any boundary into a new repository.

## 8. The Arbiter Mandate
The `Phoenix.Arbiter` subsystem is the sovereign authority for repository integrity. It provides the `make check-invariants` tooling and has the power to block merges, fail builds, and quarantine undocumented code.
