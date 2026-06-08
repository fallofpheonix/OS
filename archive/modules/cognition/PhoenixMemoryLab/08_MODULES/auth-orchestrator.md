# Module: auth-orchestrator

## Identity
- **Slug**: auth-orchestrator
- **Owner**: you
- **Location**: ~/engineering/infrastructure/shared-libraries/auth-orchestrator/
- **Status**: ACTIVE
- **Version**: 0.1.0
- **Language**: TypeScript
- **Created**: 2026-05-12
- **Last updated**: 2026-05-12

## One-liner
A unified Zero-Trust authentication module that orchestrates Bcrypt hashing, TOTP verification, and HS256 JWT validation into a single secure interface.

## API interface
```typescript
class AuthOrchestrator {
  registerUser(password: string) -> Promise<string>
  loginPhase1(password: string, hash: string) -> Promise<boolean>
  setup2FA() -> string
  loginPhase2(token: string, secret: string) -> string | null
  validateSession(jwt: string) -> boolean
}
```

## Installation / import
```typescript
import { AuthOrchestrator } from '../../infrastructure/shared-libraries/auth-orchestrator';
const auth = new AuthOrchestrator();
```

## Usage example
```typescript
const isValidPw = await auth.loginPhase1(user.pw, db.hash);
if (isValidPw) {
    const jwt = auth.loginPhase2(userInput.token, db.mfaSecret);
    if (jwt) res.cookie('session', jwt);
}
```

## Configuration
Requires the individual micro-modules (`auth-bcrypt`, `totp-2fa`, `jwt-analyzer`) to be present in the shared-libraries registry.

## Used by
| Project | Since | Notes |
|---------|-------|-------|
| [[05_PROJECTS/ACTIVE/aegis-auth]] | 2026-05-12 | uses v0.1.0 |
| [[05_PROJECTS/ACTIVE/ledger-core]] | 2026-05-12 | placeholder |

## Version history
| Version | Date | Change |
|---------|------|--------|
| 0.1.0 | 2026-05-12 | initial extraction (Capstone Phase 1) |

## Test coverage
- Unit tests: yes
- Test file: ~/engineering/infrastructure/shared-libraries/auth-orchestrator/tests/
- Coverage: complete
- Last test run: 2026-05-12

## Known failure modes
| Failure | Trigger condition | Workaround | Fixed in version |
|---------|-------------------|------------|-----------------|
| Partial Session Lockout | User loses MFA device during Phase 2 | Implement MFA recovery codes | pending |

## Dependencies
- External: none natively (wraps internal modules)
- Internal modules: `auth-bcrypt`, `totp-2fa`, `jwt-analyzer`

## Performance characteristics
- Time complexity: O(1) for JWT/TOTP, heavily dependent on Bcrypt work factor for Phase 1.
- Space complexity: O(1)

## Status transition log
- EXPERIMENTAL → ACTIVE: Integrated and verified in aegis-auth project.

## Related
- Concept notes: [[03_CORE_KNOWLEDGE/security/auth/Zero Trust Architecture]], [[03_CORE_KNOWLEDGE/security/auth/Multi-Factor Authentication]]
- Failure library: [[06_FAILURE_LIBRARY/module-failures/auth-orchestrator]]
