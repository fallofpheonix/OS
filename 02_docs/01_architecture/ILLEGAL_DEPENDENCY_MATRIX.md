# ILLEGAL_DEPENDENCY_MATRIX.md

## 1. Objective
Enforce strict architectural purity by rejecting unauthorized cross-layer dependencies.

## 2. Forbidden Paths

| Source | Target | Reason | Status |
| :--- | :--- | :--- | :--- |
| **research** | **runtime** | Research is non-deterministic and unverified. | **REJECTED** |
| **quantum** | **core** | Quantum primitives are strictly isolated. | **REJECTED** |
| **cognition** | **warden** | AI must never directly actuate. | **REJECTED** |
| **external** | **truth** | External repos cannot write to the evidence chain. | **REJECTED** |
| **external** | **containment** | Containment must be triggered by the Arbiter. | **REJECTED** |
| **memory** | **runtime** | Future memory layers are locked. | **REJECTED** |

## 3. Allowed Flow (Phoenix Matrix)
1. **telemetry** → ingestion
2. **telemetry** → **replay**
3. **replay** → **truth**
4. **truth** → **arbiter**
5. **arbiter** → **warden**
6. **warden** → **containment**
7. **containment** → **recovery**

## 4. Enforcement
- Static analysis checks in CI.
- Manual architectural review for all PRs.
