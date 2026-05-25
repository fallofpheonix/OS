# DEPENDENCY_RULES.md

## 1. Objective
Enforce the unidirectional flow of the Phoenix Matrix and reject unauthorized dependencies.

## 2. Required Flow
Telemetry → Replay → Truth → Arbiter → Warden → Containment → Recovery

## 3. Forbidden Paths
- **Research** → **Runtime** (No experimental logic in core)
- **Quantum** → **Core** (Strict isolation)
- **Cognition** → **Warden** (AI must not actuate)
- **External** → **Truth** (No third-party writes to evidence)
- **Memory** → **Containment** (L5.5/L5 only)
- **Archive** → **Runtime** (No ghost imports)
