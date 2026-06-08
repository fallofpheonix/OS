# PHOENIX MATRIX SOVEREIGN ARCHITECTURE
# [STATUS]: 18-Repository Substrate Consolidated
# [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
# [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
# [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.

# Builds and runs the tests from this directory.
set -e

SCRIPT_DIR=$(dirname "$(readlink -f "$0")")

pushd $SCRIPT_DIR/../python-proton-vpn-local-agent
cargo build --release
popd

pushd $SCRIPT_DIR
./test_connection.py
./test_playback.py
./test_listener.py
popd
