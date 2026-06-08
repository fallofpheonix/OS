# PHOENIX MATRIX SOVEREIGN ARCHITECTURE
# [STATUS]: 18-Repository Substrate Consolidated
# [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
# [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
# [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
#!/bin/bash
IMAGE_NAME="phoenix-ebpf-builder"
CONTAINER_DIR="/Users/fallofpheonix/os/pheonixos/phoenix_os/telemetry/ebpf"

# Build the docker image if it doesn't exist
docker build -t $IMAGE_NAME -f build.dockerfile .

# Run the compilation inside the container
docker run --rm \
    -v "$CONTAINER_DIR:/build" \
    -v "/sys/kernel/btf:/sys/kernel/btf:ro" \
    $IMAGE_NAME \
    make
