# PHOENIX MATRIX SOVEREIGN ARCHITECTURE
# [STATUS]: 18-Repository Substrate Consolidated
# [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
# [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
# [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
#!/bin/sh
#
# Stop the parent process (openvpn) gracefully after a small delay

# Determine the OpenVPN PID from its pid file. This works reliably even when
# the OpenVPN process is backgrounded for parallel tests.
MY_PPID=`cat $pid`

# Allow OpenVPN to finish initializing while waiting in the background and then
# killing the process gracefully.
(sleep 5 ; kill -15 $MY_PPID) &
