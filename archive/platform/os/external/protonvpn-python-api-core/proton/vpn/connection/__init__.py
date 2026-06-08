"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
"""
The public interface and the functionality that's common to all supported
VPN connection backends is defined in this module.


Copyright (c) 2023 Proton AG

This file is part of Proton VPN.

Proton VPN is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Proton VPN is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with ProtonVPN.  If not, see <https://www.gnu.org/licenses/>.
"""

from importlib.metadata import version, PackageNotFoundError

try:
    __version__ = version("proton-vpn-connection")
except PackageNotFoundError:
    __version__ = "development"


# pylint: disable=wrong-import-position
from .vpnconnection import VPNConnection
from .interfaces import (
    VPNServer, ProtocolPorts, VPNCredentials, VPNPubkeyCredentials,
    VPNUserPassCredentials, Settings
)

__all__ = [
    "VPNConnection", "VPNServer", "ProtocolPorts", "VPNCredentials",
    "VPNPubkeyCredentials", "VPNUserPassCredentials", "Settings"
]
