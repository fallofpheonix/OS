# Build Pipeline & Versioning

PhoenixOS build automation guarantees the integrity of code from compilation to standalone appliance rollout.

## Pipeline Steps

```
[Lint & Static Analysis] -> [Unit & Race Tests] -> [Go Compile] -> [Appliance Packaging]
```

1. **Lint & Static Analysis:** Run `golangci-lint` to verify code hygiene and find potential unhandled map locks.
2. **Race-Detector Tests:** Execute `go test -race ./...` to prevent concurrency regressions.
3. **Compilation:** Build main binary targeting static compilation (`CGO_ENABLED=1` for SQLite support).
4. **Appliance Packaging:** Package the binary as an initrd/busybox bootable appliance (Stage 3).

## Version Policy
PhoenixOS adheres to semantic versioning (Major.Minor.Patch). Documentation modifications mapping to state schema changes must increment the Minor version.
