---\nStatus: Planned\nImplementation: 5%\nConfidence: Conceptual\n---\n# ARCHITECTURE RULES

## Allowed Project Types
- app, module, service, sdk, infra, tooling, research, experiment, engine, framework, library

## Extraction Rules
- Rule of 2: Only extract after a pattern repeats in 2+ projects.
- Modules contain no app logic.
- Forks are temporary extraction mines.

## Naming Standards
- Directories: snake_case or kebab-case.
- Repos: match directory name.

## Runtime Isolation
- No global dependencies.
- Every project has its own environment in environments/.

## Contract Governance
- No service interface may exist without explicit contract definition.
- Interoperability depends on stable contracts, not internal implementation.
- All shared contracts must be registered in control-plane/contracts/.

## Python Tooling Governance
- All Python projects inherit Ruff configuration from infrastructure/python/ruff/.
- Ecosystem code policy is enforced at Layer 2 (Build Validation).
