# Module: jwt-analyzer

## Identity
- **Slug**: jwt-analyzer
- **Owner**: you
- **Location**: ~/engineering/infrastructure/shared-libraries/jwt-analyzer/
- **Status**: ACTIVE
- **Version**: 0.1.0
- **Language**: TypeScript
- **Created**: 2026-05-12
- **Last updated**: 2026-05-12

## One-liner
Securely inspects, signs, and verifies JSON Web Tokens, specifically preventing algorithm confusion and "none" algorithm attacks.

## API interface
```typescript
JwtAnalyzer.inspect(token: string) → any
Decodes the token without verification to inspect headers (e.g. 'alg').

JwtAnalyzer.verifyHS256(token: string, secret: string) → any
Verifies the token explicitly enforcing the HS256 algorithm.

JwtAnalyzer.signHS256(payload: object, secret: string, expiresIn?: string) → string
Signs a payload securely using HS256.
```

## Installation / import
```bash
npm install jsonwebtoken
npm install -D @types/jsonwebtoken

# Import
import { JwtAnalyzer } from '../../infrastructure/shared-libraries/jwt-analyzer';
```

## Usage example
```typescript
const payload = JwtAnalyzer.verifyHS256(reqToken, process.env.SECRET);
if (!payload) throw new Error("Invalid or insecure token!");
```

## Configuration
None.

## Used by
| Project | Since | Notes |
|---------|-------|-------|
| [[05_PROJECTS/ACTIVE/jwt-security-analyzer]] | 2026-05-12 | uses v0.1.0 |
| [[05_PROJECTS/ACTIVE/secure-rest-api-rbac]] | 2026-05-12 | uses v0.1.0 |

## Version history
| Version | Date | Change |
|---------|------|--------|
| 0.1.0 | 2026-05-12 | initial extraction |

## Test coverage
- Unit tests: yes
- Test file: ~/engineering/infrastructure/shared-libraries/jwt-analyzer/tests/
- Coverage: partial
- Last test run: 2026-05-12

## Known failure modes
| Failure | Trigger condition | Workaround | Fixed in version |
|---------|-------------------|------------|-----------------|
| Key rotation missing | When secret is changed, all tokens invalidate | Support JWKS | pending |

## Dependencies
- External: jsonwebtoken
- Internal modules: none

## Performance characteristics
- Time complexity: O(N) where N is length of payload
- Space complexity: O(N)

## Status transition log
- EXPERIMENTAL → ACTIVE: Proven across two projects in Month 2.

## Related
- Concept notes: [[03_CORE_KNOWLEDGE/security/auth/JWT Structure]], [[03_CORE_KNOWLEDGE/security/auth/Algorithm Confusion]], [[03_CORE_KNOWLEDGE/security/auth/None Algorithm Attack]]
- Failure library: [[06_FAILURE_LIBRARY/module-failures/jwt-analyzer]]
