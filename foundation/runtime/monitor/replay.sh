# PHOENIX MATRIX SOVEREIGN ARCHITECTURE
# [STATUS]: 18-Repository Substrate Consolidated
# [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
# [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
# [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
#!/bin/bash
set -e
mkdir -p artifacts
echo "Running Replay (3 deterministic runs)..."
for i in 1 2 3; do
  go run ./src/... --replay replay/sample.json > artifacts/replay_run_${i}.json
  echo "Run $i complete."
done

# Check determinism
if diff artifacts/replay_run_1.json artifacts/replay_run_2.json > /dev/null && diff artifacts/replay_run_2.json artifacts/replay_run_3.json > /dev/null; then
  echo "[PASS] All 3 replay runs produced identical output."
  cp artifacts/replay_run_1.json artifacts/replay.json
else
  echo "[FAIL] Replay runs produced different output!"
  exit 1
fi
