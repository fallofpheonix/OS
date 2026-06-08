# Module: packet-sniffer

## Identity
- **Slug**: packet-sniffer
- **Owner**: you
- **Location**: ~/engineering/infrastructure/shared-libraries/packet-sniffer/
- **Status**: ACTIVE
- **Version**: 0.1.0
- **Language**: Python
- **Created**: 2026-05-12
- **Last updated**: 2026-05-12

## One-liner
Wrapper around Scapy for easily configurable packet capture and payload inspection.

## API interface
```python
PacketSniffer(interface: str = None)
Initializes the sniffer on a specific interface.

PacketSniffer.start_capture(packet_count: int = 10, filter_bpf: str = "", callback=None)
Synchronously captures packets matching the BPF filter and executes a callback for each.
```

## Installation / import
```bash
pip install scapy
```

```python
# Import from shared library
import sys
sys.path.append('../../infrastructure/shared-libraries/packet-sniffer')
from sniffer import PacketSniffer
```

## Usage example
```python
sniffer = PacketSniffer()
sniffer.start_capture(10, "tcp port 80")
```

## Configuration
Requires root privileges for promiscuous mode and raw socket access.

## Used by
| Project | Since | Notes |
|---------|-------|-------|
| [[05_PROJECTS/ACTIVE/network-packet-sniffer]] | 2026-05-12 | uses v0.1.0 |
| [[05_PROJECTS/ACTIVE/mitm-simulation]] | 2026-05-12 | placeholder |
| [[05_PROJECTS/ACTIVE/intrusion-detection-system]] | 2026-05-12 | placeholder |

## Version history
| Version | Date | Change |
|---------|------|--------|
| 0.1.0 | 2026-05-12 | initial extraction |

## Test coverage
- Unit tests: mock only
- Test file: ~/engineering/infrastructure/shared-libraries/packet-sniffer/tests/
- Coverage: N/A
- Last test run: 2026-05-12

## Known failure modes
| Failure | Trigger condition | Workaround | Fixed in version |
|---------|-------------------|------------|-----------------|
| Dropped packets | High traffic volume >1Gbps | Use PF_RING or separate thread | pending |

## Dependencies
- External: Scapy, libpcap (OS level)
- Internal modules: none

## Performance characteristics
- Time complexity: O(1) per packet
- Space complexity: O(P) if storing packets in memory

## Status transition log
- EXPERIMENTAL → ACTIVE: completed implementation in packet-sniffer project.

## Related
- Concept notes: [[03_CORE_KNOWLEDGE/security/network/Raw Sockets]], [[03_CORE_KNOWLEDGE/security/network/TCP-IP Stack]]
- Failure library: [[06_FAILURE_LIBRARY/module-failures/packet-sniffer]]
