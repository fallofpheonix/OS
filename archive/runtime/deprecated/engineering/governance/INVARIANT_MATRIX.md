# Invariant Matrix

| ID | Category | Type | Enforcement Layer | Status | Severity |
|---|---|---|---|---|---|
| `modules_no_apps` | Structural | `forbidden_imports` | CI / InvariantEngine | Enforced | Error |
| `infra_no_experiments` | Structural | `forbidden_imports` | CI / InvariantEngine | Enforced | Error |
| `research_isolation` | Structural | `forbidden_imports` | CI / InvariantEngine | Enforced | Error |
| `sdks_only_contracts` | Structural | `allowed_imports` | CI / InvariantEngine | Enforced | Error |
| `no_validator_orchestrator_cycle` | Structural | `forbidden_imports` | CI / InvariantEngine | Enforced | Error |
| `no_global_cycles` | Structural | `cycle_free` | CI / InvariantEngine | Enforced | Error |
| `all_model_adapters_implement_protocol`| Behavioral | `interface` | CI / InvariantEngine | Enforced | Warning |
| `async_handlers_must_timeout` | Behavioral | `pattern` | CI / InvariantEngine | Stubbed | Warning |
| `mutations_require_snapshot` | Operational | `trigger` | Runtime Event Log | Deferred | Error |
| `dangerous_commands_require_approval` | Operational | `trigger` | Runtime Validator | Deferred | Error |
| `repair_budget_enforced` | Operational | `trigger` | Runtime Executor | Deferred | Error |
| `live_mode_requires_ollama` | Operational | `trigger` | Runtime Loader | Deferred | Error |

*Note: Behavioral and Operational invariants require deeper static AST parsing or runtime integration, which are prioritized in later evolution phases.*
