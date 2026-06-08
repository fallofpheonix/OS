# PHOENIX MATRIX SOVEREIGN ARCHITECTURE
# [STATUS]: 18-Repository Substrate Consolidated
# [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
# [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
# [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
#!/usr/bin/env bash

# ==============================================================================
# PHOENIX MATRIX: Ignition Sequence
# ==============================================================================
# This script enforces the deterministic pre-flight checks before the Matrix boots.

set -euo pipefail

echo "=================================================="
echo "🦅 INITIATING PHOENIX SUBSTRATE BOOT SEQUENCE"
echo "=================================================="

# 1. Verification of Contract Reality
echo ">>> [STAGE 1] Verifying Contract Reality..."
if [ ! -d "../PhoenixCore/proto/v1" ]; then
    echo "ERROR: PhoenixCore contracts not found. The Matrix has no communication spine."
    echo "Execute Vector 1 (Contract Realization) first."
    exit 1
fi
echo "✓ Contracts located."

# 2. Verification of Environment Determinism
echo ">>> [STAGE 2] Checking Environmental Determinism..."
if [ ! -f ".env" ]; then
    echo "WARNING: .env file not found. Cloning from .env.example..."
    if [ ! -f ".env.example" ]; then
        echo "ERROR: .env.example contract is missing. Halting ignition."
        exit 1
    fi
    cp .env.example .env
    echo "✓ Environment instantiated from defaults. (Review .env before production)"
else
    echo "✓ Local environment configuration detected."
fi

# 3. Validation of Network Isolation (Dry Run)
echo ">>> [STAGE 3] Validating Spatial Geometry..."
docker compose config > /dev/null
if [ $? -eq 0 ]; then
    echo "✓ Matrix geometry is mathematically valid."
else
    echo "ERROR: Invalid docker-compose configuration. Halting ignition."
    exit 1
fi

echo "=================================================="
echo ">>> ALL INVARIANTS SATISFIED."
echo ">>> IGNITING THE SUBSTRATE..."
echo "=================================================="

# Start the cluster in detached mode
docker compose up -d

echo ""
echo "🦅 PhoenixOS Substrate is active. Awaiting advisory intelligence."
