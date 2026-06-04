# PHOENIX MATRIX SOVEREIGN ARCHITECTURE
# [STATUS]: 18-Repository Substrate Consolidated
# [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
# [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
# [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
#!/bin/bash
set -e
echo "--- Running Multi-Stage Build & Test ---"
echo "[STAGE 1]: Building Core Validation Logic..."
go build ./...

echo "[STAGE 2]: Running Unit Tests..."
go test -v ./unit/...

echo "[STAGE 3]: Running Proof Tests..."
go test -v ./proofs/...

echo "[STAGE 4]: Running Integration Tests..."
go test -v ./integration/...

echo "[STAGE 5]: Running Python Guard Tests..."
PYTHONPATH=../PhoenixGuard python3 test_guard_daemon.py
PYTHONPATH=../PhoenixGuard python3 test_orchestrator.py
PYTHONPATH=../PhoenixGuard python3 test_service.py

echo "--- ALL STAGES COMPLETE ---"
