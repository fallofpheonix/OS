# Module: brute-force-guard

## Identity
- **Slug**: brute-force-guard
- **Owner**: you
- **Location**: ~/engineering/infrastructure/shared-libraries/brute-force-guard/
- **Status**: ACTIVE
- **Version**: 0.1.0
- **Language**: TypeScript
- **Created**: 2026-05-12
- **Last updated**: 2026-05-12

## One-liner
Provides express-rate-limit wrappers to protect API and Login endpoints against credential stuffing and brute force attacks.

## API interface
```typescript
BruteForceGuard.createLoginLimiter() → RateLimitRequestHandler
Returns a strict middleware limit (5 requests per 15 minutes) for authentication routes.

BruteForceGuard.createApiLimiter() → RateLimitRequestHandler
Returns a standard middleware limit (100 requests per 15 minutes) for general API endpoints.
```

## Installation / import
```bash
npm install express-rate-limit

# Import
import { BruteForceGuard } from '../../infrastructure/shared-libraries/brute-force-guard';
```

## Usage example
```typescript
app.post('/login', BruteForceGuard.createLoginLimiter(), loginHandler);
```

## Configuration
None yet. Defaults hardcoded for simple plug-and-play.

## Used by
| Project | Since | Notes |
|---------|-------|-------|
| [[05_PROJECTS/ACTIVE/brute-force-protection]] | 2026-05-12 | uses v0.1.0 |
| [[05_PROJECTS/ACTIVE/secure-rest-api-rbac]] | 2026-05-12 | uses v0.1.0 |

## Version history
| Version | Date | Change |
|---------|------|--------|
| 0.1.0 | 2026-05-12 | initial extraction |

## Test coverage
- Unit tests: pending
- Test file: ~/engineering/infrastructure/shared-libraries/brute-force-guard/tests/
- Coverage: N/A
- Last test run: 2026-05-12

## Known failure modes
| Failure | Trigger condition | Workaround | Fixed in version |
|---------|-------------------|------------|-----------------|
| In-memory reset | Memory is cleared on server restart | Use Redis store | planned |

## Dependencies
- External: express-rate-limit
- Internal modules: none

## Performance characteristics
- Time complexity: O(1)
- Space complexity: O(N) where N is number of active IP addresses

## Status transition log
- EXPERIMENTAL → ACTIVE: Integrated effectively into the secure rest API rbac project.

## Related
- Concept notes: [[03_CORE_KNOWLEDGE/security/auth/Rate Limiting]], [[03_CORE_KNOWLEDGE/security/auth/Account Lockout]], [[03_CORE_KNOWLEDGE/security/auth/IP Blocking]]
- Failure library: [[06_FAILURE_LIBRARY/module-failures/brute-force-guard]]
