# PHOENIX MATRIX SOVEREIGN ARCHITECTURE
# [STATUS]: 18-Repository Substrate Consolidated
# [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
# [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
# [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
#!/bin/sh
#
# Determine the OpenVPN PID from its pid file. This works reliably even when
# the OpenVPN process is backgrounded for parallel tests.
MY_PPID=`cat $pid`

# Add this client's VPN IP and PID to a file. This enables
# t_server_null_client.sh to kill this OpenVPN client after fping tests have
# finished.
echo "$ifconfig_local,$MY_PPID" >> ./$test_name.lwip

# Wait long enough to allow fping tests to finish. Also ensure that this
# OpenVPN client is killed even if t_server_null_client.sh failed to do it.
(sleep 15
echo "ERROR: t_server_null_client.sh failed to kill OpenVPN client with PID $MY_PPID in test $test_name. Killing it in lwip_client_up.sh."
kill -15 $MY_PPID
) &
