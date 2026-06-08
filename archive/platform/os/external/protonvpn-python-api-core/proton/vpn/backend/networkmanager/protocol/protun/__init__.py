"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
"""
This module contains the backend implementing the Protun protocol.


Copyright (c) 2026 Proton AG

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
from .protun import ProtunSmart, ProtunAuto, ProtunUDP, ProtunTCP, ProtunTLS


def register(registry):
    """Registers the Protun protocol implementations in the given registry."""
    registry.register(ProtunUDP)
    registry.register(ProtunTCP)
    registry.register(ProtunTLS)
    registry.register(ProtunSmart)
    registry.register(ProtunAuto)


__all__ = ["register"]
