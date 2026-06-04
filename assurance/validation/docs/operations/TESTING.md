---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Validation — Testing

> Last verified: 2026-06-04

## Running Staged Verification Pipeline

The entire pipeline is run via the orchestrator shell script:

```bash
cd assurance/validation
./staged_verification.sh
```

## Running Soak Tests (24-hour verification)

Soak tests check long-term memory leaks and clock drift:

```bash
cd assurance/validation/soak
go test -v -run=TestReplay24h -timeout=25h
```
