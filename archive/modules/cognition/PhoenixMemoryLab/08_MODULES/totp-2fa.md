# Module: totp-2fa

## Identity
- **Slug**: totp-2fa
- **Owner**: you
- **Location**: ~/engineering/infrastructure/shared-libraries/totp-2fa/
- **Status**: ACTIVE
- **Version**: 0.1.0
- **Language**: TypeScript
- **Created**: 2026-05-12
- **Last updated**: 2026-05-12

## One-liner
Generates TOTP secrets, produces QR codes, and verifies 6-digit tokens according to RFC 6238.

## API interface
```typescript
Totp2fa.generateSecret() → string
Generates a random base32 encoded secret.

Totp2fa.generateQRCode(user: string, service: string, secret: string) → Promise<string>
Generates a base64 Data URL for a QR code.

Totp2fa.verifyToken(token: string, secret: string) → boolean
Validates the token against the secret and current time.
```

## Installation / import
```bash
npm install otplib qrcode
npm install -D @types/qrcode

# Import from shared library
import { Totp2fa } from '../../infrastructure/shared-libraries/totp-2fa';
```

## Usage example
```typescript
const secret = Totp2fa.generateSecret();
const qr = await Totp2fa.generateQRCode('user@example.com', 'App', secret);
const isValid = Totp2fa.verifyToken('123456', secret);
```

## Configuration
| Key | Type | Default | Description |
|-----|------|---------|-------------|
| Step | number | 30 | Window of time in seconds token is valid |

## Used by
| Project | Since | Notes |
|---------|-------|-------|
| [[05_PROJECTS/ACTIVE/totp-2fa-system]] | 2026-05-12 | uses v0.1.0 |
| [[05_PROJECTS/ACTIVE/ledger-core]] | 2026-05-12 | placeholder for future integration |
| [[05_PROJECTS/ACTIVE/aegis-auth]] | 2026-05-12 | placeholder |

## Version history
| Version | Date | Change |
|---------|------|--------|
| 0.1.0 | 2026-05-12 | initial extraction |

## Test coverage
- Unit tests: partial
- Test file: ~/engineering/infrastructure/shared-libraries/totp-2fa/tests/
- Coverage: N/A
- Last test run: 2026-05-12

## Known failure modes
| Failure | Trigger condition | Workaround | Fixed in version |
|---------|-------------------|------------|-----------------|
| Time drift failure | Client clock is >30s out of sync | Allow window tolerance | pending |

## Dependencies
- External: otplib, qrcode
- Internal modules: none

## Performance characteristics
- Time complexity: O(1)
- Space complexity: O(1)
- Benchmarks: Microseconds to verify.

## Status transition log
- EXPERIMENTAL → ACTIVE: completed and verified in totp-2fa-system.

## Related
- Concept notes: [[03_CORE_KNOWLEDGE/security/auth/TOTP RFC 6238]], [[03_CORE_KNOWLEDGE/security/auth/QR Code Enrollment]]
- Failure library: [[06_FAILURE_LIBRARY/module-failures/totp-2fa]]
