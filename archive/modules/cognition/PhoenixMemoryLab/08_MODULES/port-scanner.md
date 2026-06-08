# Module: port-scanner

## Identity
- **Slug**: port-scanner
- **Owner**: you
- **Location**: ~/engineering/infrastructure/shared-libraries/port-scanner/
- **Status**: ACTIVE
- **Version**: 0.1.0
- **Language**: Python
- **Created**: 2026-05-12
- **Last updated**: 2026-05-12

## One-liner
Multi-threaded TCP port scanner for rapid network reconnaissance and vulnerability assessment.

## API interface
```python
PortScanner(target_ip: str)
Initializes a scanner for a specific target.

PortScanner.scan_range(start_port: int, end_port: int, max_threads: int = 100) -> list[int]
Concurrently scans the port range and returns a sorted list of open ports.
```

## Installation / import
```python
# Import from shared library
import sys
sys.path.append('../../infrastructure/shared-libraries/port-scanner')
from scanner import PortScanner
```

## Usage example
```python
scanner = PortScanner('192.168.1.1')
open_ports = scanner.scan_range(1, 1024)
print(f"Open ports: {open_ports}")
```

## Configuration
None.

## Used by
| Project | Since | Notes |
|---------|-------|-------|
| [[05_PROJECTS/ACTIVE/port-scanner-mini-nmap]] | 2026-05-12 | uses v0.1.0 |
| [[05_PROJECTS/ACTIVE/vulnerability-assessment-lab]] | 2026-05-12 | placeholder |
| [[05_PROJECTS/ACTIVE/bug-bounty-recon-tool]] | 2026-05-12 | placeholder |

## Version history
| Version | Date | Change |
|---------|------|--------|
| 0.1.0 | 2026-05-12 | initial extraction |

## Test coverage
- Unit tests: yes
- Test file: ~/engineering/infrastructure/shared-libraries/port-scanner/tests/
- Coverage: partial
- Last test run: 2026-05-12

## Known failure modes
| Failure | Trigger condition | Workaround | Fixed in version |
|---------|-------------------|------------|-----------------|
| Too many open files | Thread limit exceeds OS fd limits | Enforce max_threads cap | 0.1.0 |

## Dependencies
- External: none (uses standard library `socket` and `concurrent.futures`)
- Internal modules: none

## Performance characteristics
- Time complexity: O(N) where N is number of ports, heavily reduced by max_threads concurrency.
- Space complexity: O(T) where T is max_threads.

## Status transition log
- EXPERIMENTAL → ACTIVE: verified with port-scanner-mini-nmap project.

## Related
- Concept notes: [[03_CORE_KNOWLEDGE/security/network/TCP SYN Scanning]], [[03_CORE_KNOWLEDGE/security/network/Service Fingerprinting]]
- Failure library: [[06_FAILURE_LIBRARY/module-failures/port-scanner]]
