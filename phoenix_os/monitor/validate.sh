#!/bin/bash
set -e
echo "Validating Entropy Engine..."
FAILS=0

# RFC check
if [ -s RFC.md ]; then echo "[PASS] RFC.md exists"; else echo "[FAIL] RFC.md missing"; FAILS=$((FAILS+1)); fi

# Schema check  
if python3 -c "import json; json.load(open('schema.json'))" 2>/dev/null; then echo "[PASS] schema.json valid"; else echo "[FAIL] schema.json invalid"; FAILS=$((FAILS+1)); fi

# Build artifact
if [ -f artifacts/build/entropy_engine ]; then echo "[PASS] Binary exists"; else echo "[FAIL] Binary missing"; FAILS=$((FAILS+1)); fi

# Coverage
if [ -f artifacts/coverage.txt ]; then echo "[PASS] Coverage report exists"; else echo "[FAIL] Coverage missing"; FAILS=$((FAILS+1)); fi

# Benchmark
if [ -f artifacts/benchmark.log ]; then echo "[PASS] Benchmark log exists"; else echo "[FAIL] Benchmark missing"; FAILS=$((FAILS+1)); fi

# Replay
if [ -f artifacts/replay.json ]; then echo "[PASS] Replay output exists"; else echo "[FAIL] Replay missing"; FAILS=$((FAILS+1)); fi

# Hashes
if [ -f artifacts/build/hashes.sha256 ]; then echo "[PASS] Hashes exist"; else echo "[FAIL] Hashes missing"; FAILS=$((FAILS+1)); fi

echo ""
if [ $FAILS -eq 0 ]; then
  echo "VALIDATION: PASS"
  echo '{"status": "PASS", "checks": 7, "failures": 0}' > artifacts/validation.json
else
  echo "VALIDATION: FAIL ($FAILS failures)"
  echo '{"status": "FAIL", "checks": 7, "failures": '$FAILS'}' > artifacts/validation.json
  exit 1
fi
