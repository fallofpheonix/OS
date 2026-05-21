# Red Team Environment

## Purpose

Provide authorized offensive testing tools inside controlled execution modes.

## Included Tool Areas

- Recon.
- Enumeration.
- Web testing.
- Wireless analysis.
- Binary analysis.
- Exploit research.
- Reverse engineering.
- Password-audit tooling.

## Execution Modes

| Mode | Constraint |
|---|---|
| Lab mode | Isolated internal targets only |
| Sandbox mode | No access to production network |
| Air-gapped mode | No external network |

## Safety

- Disabled by default.
- Requires explicit enablement.
- Requires operator acknowledgment.
- Logs execution.
- No default target lists.
- No embedded credentials.
- No autonomous exploitation.

## Policy Gate

Red-team tooling must not be enabled in Blue Team, Forensics, or Training mode unless the session policy explicitly permits it.

## Packaging

Kali-derived builds may include tool packages.

Arch/LFS builds should use minimal curated tool groups instead of broad offensive metapackages.

