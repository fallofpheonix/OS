#!/bin/bash
set -e

# PhoenixOS Master Build Script
# Builds all core services in the 7-layer stack

echo "------------------------------------------------"
echo "PHOENIX OS: MASTER BUILD"
echo "------------------------------------------------"

SERVICES=(
    "phoenix_os/bus"
    "phoenix_os/monitor"
    "phoenix_os/trace"
    "phoenix_os/sentinel"
    "phoenix_os/arbiter"
    "phoenix_os/warden"
    "phoenix_os/ledger"
    "phoenix_os/guard"
    "phoenix_os/nexus"
)

mkdir -p build/bin

for service in "${SERVICES[@]}"; do
    echo "Building $service..."
    binary_name=$(basename "$service")
    if [ -f "$service/build.sh" ]; then
        cd "$service" && bash build.sh && cd - > /dev/null
        # Move from local artifacts to central build/bin
        cp "$service/artifacts/phoenix_$binary_name" "build/bin/" 2>/dev/null || \
        cp "$service/artifacts/$binary_name" "build/bin/phoenix_$binary_name" 2>/dev/null || \
        echo "Warning: Could not find binary for $service in artifacts/"
    else
        # Default build if no build.sh exists
        go build -o "build/bin/phoenix_$binary_name" "./$service/src"/*.go
    fi
    echo "Done."
done

echo "------------------------------------------------"
echo "PHOENIX OS: BUILD COMPLETE"
echo "Binaries located in: build/bin/"
ls -lh build/bin/
echo "------------------------------------------------"
