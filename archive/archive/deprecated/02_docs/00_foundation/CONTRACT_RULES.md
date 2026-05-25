# PhoenixOS: Contract Rules

1. **Semantic Versioning:** `contracts.Version`.
2. **API Levels:** Breaking changes increment `APILevel`.
3. **No Direct Telemetry:** High-level components use `PolicyContext` or `EvidenceWrapper`.
4. **Immutability:** Finalized contracts are immutable.
5. **Validation:** `ContractValidator` required for all policies.
