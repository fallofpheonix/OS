---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Contracts — Testing Strategy

> Last verified: 2026-06-04

## Compatibility Testing

Contracts are validated using type assertions in downstream test packages:

```bash
# Verify compatibility against contract specifications
cd assurance/validation/contract-tests
go test -v ./...
```

The compatibility tests assert that:
1. Downstream implementations satisfy `ILedger` and `EventEnvelope`.
2. JSON marshaling matches expected schemas.
