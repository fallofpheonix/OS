# PhoenixOS Documentation

> **Authority Model**: Canonical-Reference-Audit-Research-Archive (CRARA)
> **Governance Phase**: 4A.6 Documentation Authority Cleanup

This directory serves as the centralized repository for all system documentation, governed by the [DOCUMENT_AUTHORITY_MATRIX.md](./DOCUMENT_AUTHORITY_MATRIX.md).

## 1. Documentation Hierarchy

| Directory | Purpose | Authority |
| :--- | :--- | :--- |
| **[`canonical/`](./canonical/)** | The single source of truth for architecture, specs, and laws. | CANONICAL |
| **[`reference/`](./reference/)** | Derived data, reports, guides, and API documentation. | REFERENCE |
| **[`audit/`](./audit/)** | Verification reports, coverage matrices, and implementation audits. | AUDIT |
| **[`research/`](./research/)** | Conceptual, speculative, and exploratory work. | RESEARCH |
| **[`archive/`](./archive/)** | Superseded, historical, or deprecated documentation. | ARCHIVE |

## 2. Governance Rules

1. **Single Source of Truth**: Every concept must have exactly one Canonical document.
2. **Derived Authority**: Reference and Audit documents must point back to Canonical sources.
3. **Isolation**: Research and Archive documents must not be cited as authoritative for implementation.
4. **Move Policy**: As research matures, it may be promoted to Canonical via an ADR.

---
*For a detailed mapping of all files, see [DOCUMENT_AUTHORITY_MATRIX.md](./DOCUMENT_AUTHORITY_MATRIX.md).*
