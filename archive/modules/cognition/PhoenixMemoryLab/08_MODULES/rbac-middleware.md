# Module: rbac-middleware

## Identity
- **Slug**: rbac-middleware
- **Owner**: you
- **Location**: ~/engineering/infrastructure/shared-libraries/rbac-middleware/
- **Status**: ACTIVE
- **Version**: 0.1.0
- **Language**: TypeScript
- **Created**: 2026-05-12
- **Last updated**: 2026-05-12

## One-liner
Express middleware that enforces Role-Based Access Control (RBAC) by inspecting authenticated request payloads.

## API interface
```typescript
RbacMiddleware.requireRole(requiredRole: string) → NextFunction
Allows only the exact role provided. Returns 403 Forbidden otherwise.

RbacMiddleware.requireAnyRole(allowedRoles: string[]) → NextFunction
Allows if the user possesses at least one of the roles in the array.
```

## Installation / import
```bash
# Import
import { RbacMiddleware } from '../../infrastructure/shared-libraries/rbac-middleware';
```

## Usage example
```typescript
app.get('/admin', RbacMiddleware.requireRole('admin'), adminHandler);
app.get('/manager', RbacMiddleware.requireAnyRole(['admin', 'manager']), managerHandler);
```

## Configuration
Relies on `req.user.roles` being populated as a `string[]` by prior authentication middleware.

## Used by
| Project | Since | Notes |
|---------|-------|-------|
| [[05_PROJECTS/ACTIVE/secure-rest-api-rbac]] | 2026-05-12 | uses v0.1.0 |

## Version history
| Version | Date | Change |
|---------|------|--------|
| 0.1.0 | 2026-05-12 | initial extraction |

## Test coverage
- Unit tests: pending
- Test file: ~/engineering/infrastructure/shared-libraries/rbac-middleware/tests/
- Coverage: N/A
- Last test run: 2026-05-12

## Known failure modes
| Failure | Trigger condition | Workaround | Fixed in version |
|---------|-------------------|------------|-----------------|
| Missing user object | Middleware placed before authentication | Ensure strict route order | N/A |

## Dependencies
- External: express (types)
- Internal modules: expects jwt-analyzer or similar to populate `req.user`

## Performance characteristics
- Time complexity: O(R) where R is number of roles the user has
- Space complexity: O(1)

## Status transition log
- EXPERIMENTAL → ACTIVE: Successfully deployed in the secure REST API project.

## Related
- Concept notes: [[03_CORE_KNOWLEDGE/security/auth/RBAC]], [[03_CORE_KNOWLEDGE/security/auth/Principle of Least Privilege]]
- Failure library: [[06_FAILURE_LIBRARY/module-failures/rbac-middleware]]
