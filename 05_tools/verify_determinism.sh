#!/bin/bash
set -e

echo "=== PhoenixOS Determinism Verification ==="

# 1. Build
echo "[1/4] Building runtime..."
cd phoenix_os && go build -o ../build/phoenix-os . && cd ..

# 2. First Run
echo "[2/4] Running first execution..."
OUTPUT1=$(./build/phoenix-os)
HASH1=$(echo "$OUTPUT1" | grep "Authoritative Output Hash:" | awk '{print $5}')
echo "      Hash 1: $HASH1"

# 3. Second Run
echo "[3/4] Running second execution..."
OUTPUT2=$(./build/phoenix-os)
HASH2=$(echo "$OUTPUT2" | grep "Authoritative Output Hash:" | awk '{print $5}')
echo "      Hash 2: $HASH2"

# 4. Compare
echo "[4/4] Comparing hashes..."
if [ "$HASH1" == "$HASH2" ] && [ -n "$HASH1" ]; then
    echo "SUCCESS: Replay is 100% deterministic."
    exit 0
else
    echo "FAILURE: Replay divergence detected or hash empty!"
    echo "Hash 1: $HASH1"
    echo "Hash 2: $HASH2"
    exit 1
fi
