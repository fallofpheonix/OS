---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Security — Testing

> Last verified: 2026-06-04

We test warden policy reactions and namespace isolation guarantees:

```bash
cd assurance/security
go test -v ./...
```

## Python Daemon Tests
If testing the python telemetry daemon:

```bash
poetry run pytest assurance/security/tests/
```
