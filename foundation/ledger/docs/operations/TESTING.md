---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Ledger — Testing

> Last verified: 2026-06-04

The ledger verifies rollback correctness, validation hash compliance, and resource leak prevention:

```bash
cd foundation/ledger
go test -v ./...
```

## Coverage
- **Rollback Tests**: Asserts correct head tracking and logical tick adjustments.
- **Resource Allocator Tests**: Tests that memory allocation is correctly deallocated during pruning and rollback.
