# Module: impossible-travel

## Identity
- **Slug**: impossible-travel
- **Owner**: you
- **Location**: ~/engineering/infrastructure/shared-libraries/impossible-travel/
- **Status**: ACTIVE
- **Version**: 0.1.0
- **Language**: Python
- **Created**: 2026-05-12
- **Last updated**: 2026-05-12

## One-liner
Calculates the physical speed between two login locations using the Haversine formula to detect session hijacking via VPNs.

## API interface
```python
ImpossibleTravelDetector()
Initializes the detector with a GeoIP database mapping.

ImpossibleTravelDetector.analyze_travel(ip_1: str, time_1: datetime, ip_2: str, time_2: datetime) -> dict
Calculates the distance and required speed between the two events. Returns a status of 'POSSIBLE', 'IMPOSSIBLE', or 'UNKNOWN'.
```

## Installation / import
```python
# Import from shared library
import sys
sys.path.append('../../infrastructure/shared-libraries/impossible-travel')
from detector import ImpossibleTravelDetector
```

## Usage example
```python
detector = ImpossibleTravelDetector()
result = detector.analyze_travel(old_ip, old_time, new_ip, new_time)
if result['status'] == 'IMPOSSIBLE':
    trigger_account_lockout()
```

## Configuration
Requires an underlying GeoIP Database. The current implementation uses a mock dictionary `GEO_DB` which must be replaced with `geoip2` bindings in production.

## Used by
| Project | Since | Notes |
|---------|-------|-------|
| [[05_PROJECTS/ACTIVE/impossible-travel-detector]] | 2026-05-12 | uses v0.1.0 |
| [[05_PROJECTS/ACTIVE/ledger-core]] | 2026-05-12 | placeholder |

## Version history
| Version | Date | Change |
|---------|------|--------|
| 0.1.0 | 2026-05-12 | initial extraction |

## Test coverage
- Unit tests: yes
- Test file: ~/engineering/infrastructure/shared-libraries/impossible-travel/tests/
- Coverage: partial
- Last test run: 2026-05-12

## Known failure modes
| Failure | Trigger condition | Workaround | Fixed in version |
|---------|-------------------|------------|-----------------|
| False Positives | Corporate VPN routing changes IP drastically | Whitelist known corporate VPN blocks | pending |

## Dependencies
- External: none natively (needs `geoip2` for production)
- Internal modules: none

## Performance characteristics
- Time complexity: O(1) for Haversine math
- Space complexity: O(1)

## Status transition log
- EXPERIMENTAL → ACTIVE: verified in the impossible-travel-detector project.

## Related
- Concept notes: [[03_CORE_KNOWLEDGE/security/auth/GeoIP Tracking]], [[03_CORE_KNOWLEDGE/security/auth/Session Hijacking]], [[03_CORE_KNOWLEDGE/ml-quant/algorithms/Haversine Formula]]
- Failure library: [[06_FAILURE_LIBRARY/module-failures/impossible-travel]]
