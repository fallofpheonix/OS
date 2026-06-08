# Brain Vault Analysis

Date: 2026-05-13
Scope: Entire `~/engineering` ecosystem

## 1. Executive Reality
You are no longer organizing projects. You are organizing capabilities.

Current transition:
- Architecture engineering -> runtime systems engineering
- Structural governance -> survivability validation
- Integrity checks -> trust-aware control plane (early)

Current bottleneck:
- Not design
- Not module count
- Not new repos
- Operational proof under hostile runtime conditions

## 2. Canonical Ecosystem Structure
Use this as the long-term invariant filesystem model:

```text
~/engineering/
├── brain/             # cognition, notes, ADRs, failure library (non-runtime)
├── workspace/         # active executable projects
├── modules/           # reusable capabilities extracted from projects/forks
├── infrastructure/    # templates, CI/CD, observability, shared deployment assets
├── control-plane/     # control plane, governance, sync engine, policy
├── forks/             # temporary upstream mining lifecycle only
├── environments/      # runtime isolation and environment definitions
├── services/          # deployable service units
├── sdk/               # generated/manual client SDKs and bindings
├── apps/              # user-facing applications
├── research/          # experiments/spikes (promotion candidate zone)
└── archive/           # retired systems/artifacts
```

Boundary rule:
- `brain/` is for cognition and governance documents.
- Executable runtime systems belong under `workspace/`, `services/`, `apps/`, or `modules/`.

## 3. Repository Taxonomy (Single Ownership)
Each repo must belong to one class only.

| Class | Purpose | Deployable | Reusable |
|---|---|---:|---:|
| core | foundation primitives | Yes | Very High |
| module | reusable subsystem | Optional | High |
| infra | deployment/platform ops | Yes | High |
| sdk | client bindings | No | High |
| app | end-user product | Yes | Low |
| service | network runtime component | Yes | Medium |
| research | experiments | Maybe | Uncertain |
| fork | extraction source (temporary) | No | No |
| tooling | dev tooling/automation | Optional | Medium |

## 4. Module Layering Model
Canonical module topology:

```text
modules/
├── core/
├── shared/
├── systems/
├── networking/
├── automation/
├── security/
├── ai/
├── data/
├── observability/
├── control-plane/
└── experimental/
```

Extraction threshold:
- Extract only when reused >= 2 times and interface stability exists.

## 5. Language and Tool Compatibility Matrix
Choose language by failure domain and runtime profile.

| Domain | Primary | Secondary | Notes |
|---|---|---|---|
| Runtime substrate | Rust | Zig/C | safety + concurrency first |
| Backend APIs/services | Go | Rust | fast iteration + operational clarity |
| AI control-plane/evals | Python | Go | ecosystem leverage |
| AI inference core | Rust | C++ | performance-critical path |
| Networking systems | Rust | Go | latency + memory safety |
| Security/crypto | Rust | C | avoid dynamic runtime-heavy stacks |
| CLI/tooling | Rust | Go | strong distribution and reliability |
| Workflow automation | Go | Python | operational scripts/workers |
| Frontend | TypeScript | Rust (Tauri shell) | UI and integrations |
| Data pipelines | Python | Rust | productivity then optimize hotspots |
| Infra as code | Terraform | Pulumi | avoid ad-hoc shell-only infra |

Compatibility rules:
- Cross-language communication must use contracts, not shared internal structs.
- Generate SDKs from contracts for TS/Python/Go/Rust.

## 6. Contract-First Interoperability
Required standards:

- Internal RPC: Protobuf + gRPC
- External APIs: OpenAPI REST
- Async events: Protobuf/Avro event schema
- Config/validation: JSON Schema
- Telemetry: OpenTelemetry semantic conventions

Never share:
- Internal runtime state
- ORM entities across boundaries
- Direct DB coupling between services

## 7. Sync Architecture (Local + Remote + Ecosystem)
Use `control-plane/sync-engine/` as source of sync truth:

```text
control-plane/sync-engine/
├── sync-config.yaml
├── repo-registry.yaml
├── dependency-graph.yaml
└── sync-history.log
```

Sync directions:
- GitHub -> local mirror
- local -> GitHub push
- module version propagation -> dependent repos
- contract updates -> centrally governed rollout
- CI template updates -> inherited/reused workflows

Recommended implementation split:
- Core sync engine: Go
- Filesystem/watcher acceleration: Rust
- High-level control-plane glue: Python/Go

## 8. Fork Governance and Extraction Lifecycle
Treat forks as temporary mining zones only.

```text
forks/
├── active/       # newly tracked upstream forks
├── extracting/   # currently mining reusable capabilities
├── mined/        # extraction complete, pending archive/remove
└── archived/     # retired forks, no active dependency
```

Lifecycle:
1. Fork upstream
2. Analyze reusable capabilities
3. Extract into `modules/` with rewritten interfaces
4. Decouple from upstream assumptions
5. Validate independent tests/contracts
6. Mark as `mined`
7. Archive/remove fork

Hard rule:
- No permanent production dependency on fork internals.

## 9. Current Maturity Map
| Domain | Maturity |
|---|---|
| Repo structure and governance | Advanced |
| ABI/contract intent | Advanced |
| Runtime fabric | Intermediate |
| Lifecycle control-plane | Intermediate |
| Observability convergence | Intermediate |
| Supervision correctness | Weak-Intermediate |
| Stress validation | Weak |
| Trust anchoring/authenticity chain | Weak |
| Automated remediation/self-healing | Absent |

## 10. Highest-Risk Open Issues
1. Supervisor correctness under panic/restart storms remains fragile.
2. Shared receiver lock (`Arc<Mutex<Receiver>>`) can throttle concurrency.
3. Runtime eventing is incomplete (panic/restart/saturation/degradation visibility).
4. Empirical stress evidence is insufficient (soak/flood/fuzz/chaos).
5. Authenticity model is likely local-root bound; provenance/append-only trust chain is missing.

## 11. Immediate Execution Sequence
### Phase A - Runtime Truth First
- Build `runtime-lab/` for flood, soak, fuzz, starvation, sink-failure, lifecycle chaos.
- Define pass/fail SLOs for survivability.

### Phase B - Supervisor Hardening
- Formal failure taxonomy (transient/persistent/degraded/fatal/quarantine).
- Restart budgets + exponential backoff + escalation semantics.
- Worker liveness guarantees and sink-failure isolation.

### Phase C - Runtime Event Bus
- Emit structured events: panic, restart, saturation, degradation, quarantine.
- Pipe into control plane state and alerting.

### Phase D - Observability Convergence
- Prometheus metrics + OTEL traces + Grafana dashboards.
- Causal links: runtime event -> service impact -> policy decision.

### Phase E - Sync and Extraction Discipline
- Finalize sync-engine policy and registry enforcement.
- Continue fork extraction and retire mined forks aggressively.

## 12. What Not To Do Yet
- Add many new modules before runtime validation closes.
- Expand autonomous agent capabilities without trusted invariant enforcement.
- Introduce distributed-runtime complexity before local substrate survivability is proven.
- Keep long-lived forks as hidden dependencies.

## 13. Foundational Capability Build Order
Tier 1:
- `logging-core`
- `config-core`
- `telemetry-core`
- `auth-core`
- `runtime-core`
- `event-core`

Tier 2:
- `scheduler`
- `workflow-engine`
- `cache-layer`
- `database-abstraction`
- `message-broker`

Tier 3:
- `ai-control-plane`
- `distributed-runtime`
- `plugin-system`
- `simulation-engine`

## 14. Final Strategic Rule
- Project = temporary composition
- Module = permanent capability

Scale outcome depends on enforcing that distinction consistently.
