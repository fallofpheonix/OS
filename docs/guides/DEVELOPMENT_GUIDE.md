---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# DEVELOPMENT_GUIDE.md — PhoenixOS Development Guide

## Prerequisites

| Requirement | Version | Purpose |
|-------------|---------|---------|
| Go | 1.25+ | Core language |
| Linux | 5.x+ | eBPF support |
| Docker | 24.x+ | Container deployment |
| Docker Compose | 2.x+ | Multi-node cluster |
| Git | 2.x+ | Version control |

## Repository Setup

```bash
# Clone the repository
git clone https://github.com/fallofpheonix/os
cd os

# Sync Go workspace
go work sync

# Install dependencies
go mod download
```

## Building

```bash
# Build all modules
make build

# Build a specific module
go build ./Phoenix.Nucleus/PhoenixCore/...

# Build the main binary
go build ./Phoenix.Terminus/PhoenixOS/cmd/phoenixd
```

## Testing

```bash
# Run all tests
go test ./...

# Run with race detector
go test -race ./...

# Run specific package tests
go test ./Phoenix.Nucleus/PhoenixGuard/...

# Run benchmarks
go test -bench=. ./Phoenix.Nucleus/PhoenixCore/...
```

## Code Style

### Go
- Follow `gofmt` formatting
- Use interfaces for testability
- All functions must have workflow position comments
- No hardcoded credentials
- Use `crypto/rand` for cryptographic operations

### Python
- Follow PEP 8
- Use type hints
- Run `ruff` linter

### TypeScript
- Follow ESLint rules
- Use strict mode

## Architecture Guidelines

### 1. The 6 Axioms
Every code change must respect:
1. **Determinism is sacred** — No non-deterministic primitives
2. **Replay is authoritative** — State must be reconstructable
3. **AI is advisory** — AI never directly controls
4. **Control must be bounded** — All actions have rollback
5. **Telemetry correctness > AI sophistication** — Precision first
6. **Never scale instability** — Stability before scaling

### 2. Dependency Rules
- Contract packages are the canonical contract source
- No cross-boundary types outside contract packages
- Dependencies flow through adapters, not direct implementation imports
- No circular dependencies

### 3. Security Guidelines
- All input must be validated
- Use signature verification for external data
- Implement defense in depth
- Log all security-relevant events
- Follow the principle of least privilege

### 4. Testing Requirements
- Unit tests for all public functions
- Integration tests for cross-module interactions
- Race detector must pass
- Fuzz testing for parsing functions
- Chaos testing for reliability

## Deployment

### Local Development
```bash
# Start single node
go run ./Phoenix.Terminus/PhoenixOS/cmd/phoenixd

# Start 3-node cluster
make ignite
```

### Docker
```bash
# Build image
docker build -t phoenixos -f Dockerfile.node .

# Run container
docker run -it --privileged phoenixos
```

### Production
```bash
# Build optimized binary
go build -ldflags="-s -w" ./Phoenix.Terminus/PhoenixOS/cmd/phoenixd

# Deploy with systemd
cp phoenixd /usr/local/bin/
systemctl enable phoenixd
systemctl start phoenixd
```

## Debugging

### Logs
```bash
# View container logs
docker compose logs -f

# View specific node logs
docker compose logs node-alpha
```

### eBPF
```bash
# Check eBPF programs
bpftool prog list

# Check eBPF maps
bpftool map list
```

### Metrics
```bash
# Prometheus metrics available at
curl http://localhost:8080/metrics
```

## Contributing

1. Create a feature branch
2. Make changes following the guidelines
3. Run `make test` to verify
4. Update documentation if needed
5. Submit a pull request

## Related Documents

- [CONTEXT.md](CONTEXT.md) — System context
- [WORKING_MODEL.md](WORKING_MODEL.md) — Architecture details
- [PHASE_4_PROTOCOL_SPECIFICATION.md](PHASE_4_PROTOCOL_SPECIFICATION.md) — Wire contracts

## Research & Reference

Third-party reference projects are located in `Phoenix.Terminus/PhoenixExternal/research/` (8.6GB, 72+ repos). These are NOT part of PhoenixOS core and should not be modified.
