# PhoenixOS: Contract Rules

1. **Semantic Versioning:** All core components use `contracts.Version`.
2. **API Levels:** Breaking interface changes require an increment to `APILevel`.
3. **No Direct Telemetry:** High-level components (Arbiter, Warden) must use `PolicyContext` or `EvidenceWrapper`, never raw telemetry structs from the monitor.
4. **Immutability:** Once a contract is finalized for a phase, it cannot be modified without a formal deprecation cycle and an entry in `contracts.DeprecationEntry`.
5. **Validation:** Policy implementations must pass `ContractValidator` checks.
