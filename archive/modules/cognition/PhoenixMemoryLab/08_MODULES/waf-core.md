# Module: waf-core

## Identity
- **Slug**: waf-core
- **Owner**: you
- **Location**: ~/engineering/infrastructure/shared-libraries/waf-core/
- **Status**: ACTIVE
- **Version**: 0.1.0
- **Language**: Python
- **Created**: 2026-05-12
- **Last updated**: 2026-05-12

## One-liner
Inspects payloads for XSS and SQL injection, dynamically dropping connections and automatically banning IPs via iptables.

## API interface
```python
WAFCore(auto_ban: bool = False)
Initializes the firewall logic, linking it to iptables-manager if auto_ban is enabled.

WAFCore.inspect_payload(payload: str, source_ip: str) -> bool
Evaluates a payload using strict regex definitions. Returns True if safe, False if malicious.
```

## Installation / import
```python
# Import from shared library
import sys
sys.path.append('../../infrastructure/shared-libraries/waf-core')
from waf_core import WAFCore
```

## Usage example
```python
waf = WAFCore(auto_ban=True)
is_safe = waf.inspect_payload(request_body, client_ip)
if not is_safe:
    abort(403)
```

## Configuration
Enabling `auto_ban=True` requires `sudo` privileges because it invokes `iptables_manager.py` under the hood to write physical firewall rules.

## Used by
| Project | Since | Notes |
|---------|-------|-------|
| [[05_PROJECTS/ACTIVE/web-application-firewall]] | 2026-05-12 | uses v0.1.0 |
| [[05_PROJECTS/ACTIVE/ledger-core]] | 2026-05-12 | placeholder |
| [[05_PROJECTS/ACTIVE/aegis-auth]] | 2026-05-12 | placeholder |

## Version history
| Version | Date | Change |
|---------|------|--------|
| 0.1.0 | 2026-05-12 | initial extraction |

## Test coverage
- Unit tests: yes
- Test file: ~/engineering/infrastructure/shared-libraries/waf-core/tests/
- Coverage: partial
- Last test run: 2026-05-12

## Known failure modes
| Failure | Trigger condition | Workaround | Fixed in version |
|---------|-------------------|------------|-----------------|
| False Positives | HTML formatting submitted in standard forms | Whitelist specific routes | pending |

## Dependencies
- External: `re` (Standard library)
- Internal modules: `iptables-manager` (optional, for auto-banning)

## Performance characteristics
- Time complexity: Dependent on regex compilation and payload length O(N)
- Space complexity: O(1)

## Status transition log
- EXPERIMENTAL → ACTIVE: Integrated and verified in web-application-firewall project.

## Related
- Concept notes: [[03_CORE_KNOWLEDGE/security/web/SQL Injection Prevention]], [[03_CORE_KNOWLEDGE/security/web/XSS Filtering]], [[03_CORE_KNOWLEDGE/security/web/Regex Payloads]]
- Failure library: [[06_FAILURE_LIBRARY/module-failures/waf-core]]
