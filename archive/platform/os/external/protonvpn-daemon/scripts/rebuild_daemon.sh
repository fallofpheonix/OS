# PHOENIX MATRIX SOVEREIGN ARCHITECTURE
# [STATUS]: 18-Repository Substrate Consolidated
# [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
# [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
# [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
#!/usr/bin/env bash

set -e

source /etc/os-release
if ! [[ $ID == "debian" || $ID_LIKE == "debian" ]]
then
    echo "This script only works on debian[-based] systems"
    exit 1
fi


./scripts/create_changelogs.py
sudo apt remove -y proton-vpn-daemon
sudo mk-build-deps -ir --tool "apt-get -y"
dpkg-buildpackage -b --no-sign
pkg_version=$(dpkg-parsechangelog --show-field Version)
sudo dpkg -i "../proton-vpn-daemon_${pkg_version}_all.deb"
journalctl -u me.proton.vpn.split_tunneling.service -f
