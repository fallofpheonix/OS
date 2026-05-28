# PhoenixMind Modular Architecture Mandate

## Core Rule: Strict Isolation
To ensure PhoenixMind remains stable, maintain hard separation between modules.

## Dependency Graph
core
 ├── agents
 ├── runtime
 ├── validator
 ├── memory
 └── model-router

training
 └── datasets

sandbox
 ├── runtime
 └── evals

observability
 └── reads all

security
 └── protects all

## Forbidden Flows
- PhoenixMind -> Kernel (Use Arbiter/Warden only)
- PhoenixMind -> Containment (Direct interaction prohibited)
- PhoenixMind -> Recovery (Direct interaction prohibited)
