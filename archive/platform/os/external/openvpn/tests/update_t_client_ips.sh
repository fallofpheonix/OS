# PHOENIX MATRIX SOVEREIGN ARCHITECTURE
# [STATUS]: 18-Repository Substrate Consolidated
# [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
# [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
# [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
#!/bin/sh
#
# This --up script caches the IPs handed out by the test VPN server to a file
# for later use.

RC="$TOP_BUILDDIR/t_client_ips.rc"

grep EXPECT_IFCONFIG4_$TESTNUM= $RC > /dev/null 2>&1
if [ $? -ne 0 ]; then
    echo "EXPECT_IFCONFIG4_$TESTNUM=$ifconfig_local" >> $RC
fi

grep EXPECT_IFCONFIG6_$TESTNUM= $RC > /dev/null 2>&1
if [ $? -ne 0 ]; then
    echo "EXPECT_IFCONFIG6_$TESTNUM=$ifconfig_ipv6_local" >> $RC
fi
