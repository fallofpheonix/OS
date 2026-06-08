# Module: iptables-manager

## Identity
- **Slug**: iptables-manager
- **Owner**: you
- **Location**: ~/engineering/infrastructure/shared-libraries/iptables-manager/
- **Status**: ACTIVE
- **Version**: 0.1.0
- **Language**: Python
- **Created**: 2026-05-12
- **Last updated**: 2026-05-12

## One-liner
Python wrapper for executing and managing `iptables` firewall rules programmatically.

## API interface
```python
IptablesManager.block_ip(ip_address: str)
Appends a DROP rule to the INPUT chain for the given source IP.

IptablesManager.allow_port(port: int, protocol: str = 'tcp')
Appends an ACCEPT rule to the INPUT chain for a specific destination port.

IptablesManager.list_rules() -> str
Returns the verbose, numeric listing of all current iptables rules.

IptablesManager.flush_rules()
Flushes all chains.
```

## Installation / import
```python
# Import from shared library
import sys
sys.path.append('../../infrastructure/shared-libraries/iptables-manager')
from iptables_manager import IptablesManager
```

## Usage example
```python
IptablesManager.block_ip('203.0.113.50')
IptablesManager.allow_port(22)
print(IptablesManager.list_rules())
```

## Configuration
Requires `sudo` privileges to run.

## Used by
| Project | Since | Notes |
|---------|-------|-------|
| [[05_PROJECTS/ACTIVE/iptables-firewall]] | 2026-05-12 | uses v0.1.0 |
| [[05_PROJECTS/ACTIVE/web-application-firewall]] | 2026-05-12 | placeholder for future integration |

## Version history
| Version | Date | Change |
|---------|------|--------|
| 0.1.0 | 2026-05-12 | initial extraction |

## Test coverage
- Unit tests: mock only (subprocess mocked)
- Test file: ~/engineering/infrastructure/shared-libraries/iptables-manager/tests/
- Coverage: N/A
- Last test run: 2026-05-12

## Known failure modes
| Failure | Trigger condition | Workaround | Fixed in version |
|---------|-------------------|------------|-----------------|
| Privilege Error | Script not run as root | Check UID at startup | pending |

## Dependencies
- External: `iptables` CLI must be installed on the host OS
- Internal modules: none

## Performance characteristics
- Time complexity: Dependent on subprocess invocation overhead (~10-20ms)

## Status transition log
- EXPERIMENTAL → ACTIVE: extracted from iptables-firewall for use in WAF and Zero-Trust systems.

## Related
- Concept notes: [[03_CORE_KNOWLEDGE/security/network/Netfilter]], [[03_CORE_KNOWLEDGE/security/network/iptables chains]]
- Failure library: [[06_FAILURE_LIBRARY/module-failures/iptables-manager]]
