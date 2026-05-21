#!/bin/bash
set -e
echo "Validating __MODULE_NAME__..."
FAILS=0
if [ -s RFC.md ]; then echo "[PASS] RFC.md exists"; else echo "[FAIL] RFC.md missing"; FAILS=$((FAILS+1)); fi
if python3 -c "import json; json.load(open('schema.json'))" 2>/dev/null; then echo "[PASS] schema.json valid"; else echo "[FAIL] schema.json invalid"; FAILS=$((FAILS+1)); fi
if [ -f artifacts/build/hashes.sha256 ]; then echo "[PASS] Build hashes exist"; else echo "[FAIL] Build hashes missing"; FAILS=$((FAILS+1)); fi
if [ -f artifacts/coverage.txt ]; then echo "[PASS] Coverage report exists"; else echo "[FAIL] Coverage missing"; FAILS=$((FAILS+1)); fi
if [ -f artifacts/benchmark.log ]; then echo "[PASS] Benchmark log exists"; else echo "[FAIL] Benchmark missing"; FAILS=$((FAILS+1)); fi
echo ""
if [ $FAILS -eq 0 ]; then
  echo "VALIDATION: PASS"
  echo '{"status": "PASS", "failures": 0}' > artifacts/validation.json
else
  echo "VALIDATION: FAIL ($FAILS failures)"
  echo '{"status": "FAIL", "failures": '$FAILS'}' > artifacts/validation.json
  exit 1
fi
