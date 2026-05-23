#!/bin/bash
set -e
ROOT_DIR=$(pwd)
SERVICES=("phoenix_os/bus" "phoenix_os/monitor" "phoenix_os/trace" "07_security/physics" "phoenix_os/arbiter" "phoenix_os/warden" "phoenix_os/ledger" "phoenix_os/guard" "10_kernel/sandbox" "phoenix_os/nexus")
mkdir -p build/bin
for service in "${SERVICES[@]}"; do
    cd "$ROOT_DIR/$service"
    go build -o "$ROOT_DIR/build/bin/$(basename $service)" .
done
cd "$ROOT_DIR"
./12_infrastructure/deployment/deploy.sh
cd "$ROOT_DIR/05_tools/ui"
go build -o phoenix_ui
cp phoenix_ui "$ROOT_DIR/build/bin/"
