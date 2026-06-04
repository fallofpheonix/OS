---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Runtime — Testing Strategy

> Last verified: 2026-06-04

We run a suite of tests to guarantee consensus stability and replay consistency:

```bash
# Run runtime tests
cd foundation/runtime
go test -v ./...
```

## Key Test Profiles
- **Determinism Tests**: Checks that replaying identical logs on different cores generates identical state hash outputs.
- **Chaos Tests**: Injects network drops and node partition scenarios into the BFT consensus engine.
