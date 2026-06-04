# PhoenixOS System Versioning Policy

To guarantee long-term determinism, replayability, and backward compatibility across all 18 repositories, the PhoenixOS ecosystem enforces strict versioning constraints on all shared contracts and interfaces.

## Version Classes

We categorize versioning into six distinct system classes:

| Class | Scope | Format | Breaking Policy |
|:---|:---|:---|:---|
| **Contract** | Data structures in `proto/v1/` and event bus definitions. | `v[Major]` (e.g., `v1`, `v2`) | Major bump required for deleted or modified fields. |
| **Schema** | JSON and YAML schemas in `schemas/v1/`. | SemVer `[Major].[Minor].[Patch]` | Minor for optional fields; Major for required fields or structural changes. |
| **Protocol** | Wire behaviors, message sequences, consensus protocols. | SemVer `[Major].[Minor].[Patch]` | Requires a protocol negotiation step if Major version mismatch. |
| **API** | REST/gRPC endpoints (defined in OpenAPI specifications). | `v[Major].[Minor]` (e.g., `v1.0`, `v1.1`) | Deprecation cycles must last at least two major client releases. |
| **FSM** | State transition tables and YAML state graphs. | SemVer `[Major].[Minor].[Patch]` | Any change to allowable states or forbidden transitions requires a Major bump. |
| **Proof** | TLA+ models and correctness verification properties. | SemVer `[Major].[Minor].[Patch]` | Must match the corresponding contract or FSM version it verifies. |

## Versioning Rules

1. **Replay Compatibility Rule**: Changes to any Contract or Schema must not break replay validation of historical events. If a structure changes in a breaking manner, a new major version directory must be introduced (e.g., `proto/v2/`), leaving historical versions untouched.
2. **Fail-Closed on Mismatch**: If an event or API request is received with an unsupported major version, the system must immediately reject the message and transition to `SUSPICIOUS` or `CRITICAL` (in Warden FSM) to prevent state divergence.
3. **Immutability of Historical Definitions**: Once a protobuf stub or yaml contract is pushed to `main` on `PhoenixCore`, it is immutable. Future enhancements must be appended via fields with default zero-values or defined in new versions.
4. **Deterministic Version Binding**: All serialized event envelopes must explicitly state their `schema_version` as a string and match the protobuf package version.
