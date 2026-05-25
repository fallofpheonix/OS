# Operations: Debugging Workflow

Use this document to troubleshoot and resolve common failures in the runtime environment.

## 1. Build Failures
Clear the local compile caches and regenerate dependency locks:
```bash
go clean -cache
go mod tidy
go work sync
go build ./...
```

## 2. Race Conditions
Identify thread-safety issues under concurrent execution:
```bash
go test -race ./...
```
*Resolution:* Wrap shared state updates in `sync.Mutex` or `sync.RWMutex` (e.g. process lists in process audits).

## 3. Replay Divergence
If trace replay fails with a hash or cursor mismatch:
```bash
go test -run Replay -count=100 ./...
```
*Resolution:* Check that all timestamps are cleared via a `Normalize()` method before hashing. Confirm fields are marshalled in consistent alphabetical order.
