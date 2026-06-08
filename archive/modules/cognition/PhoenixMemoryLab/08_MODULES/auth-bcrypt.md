# Module: auth-bcrypt

## Identity
- **Slug**: auth-bcrypt
- **Owner**: you
- **Location**: ~/engineering/infrastructure/shared-libraries/auth-bcrypt/
- **Status**: ACTIVE
- **Version**: 0.1.0
- **Language**: TypeScript
- **Created**: 2026-05-12
- **Last updated**: 2026-05-12

## One-liner
Provides secure password hashing, salting, and verification using bcrypt for node.js systems.

## API interface
```typescript
AuthBcrypt.hashPassword(password: string) → Promise<string>
Hashes a plaintext password using bcrypt with 12 salt rounds.

AuthBcrypt.verifyPassword(password: string, hash: string) → Promise<boolean>
Verifies a plaintext password against a stored hash.
```

## Installation / import
```bash
# In your node project
npm install bcrypt
npm install -D @types/bcrypt

# Import from shared library
import { AuthBcrypt } from '../../infrastructure/shared-libraries/auth-bcrypt';
```

## Usage example
```typescript
const hash = await AuthBcrypt.hashPassword('my-secret-pass');
const isValid = await AuthBcrypt.verifyPassword('my-secret-pass', hash);
```

## Configuration
| Key | Type | Default | Description |
|-----|------|---------|-------------|
| SALT_ROUNDS | number | 12 | The number of bcrypt hashing rounds |

## Used by
| Project | Since | Notes |
|---------|-------|-------|
| [[05_PROJECTS/ACTIVE/login-system-bcrypt]] | 2026-05-12 | uses v0.1.0 |
| [[05_PROJECTS/ACTIVE/ledger-core]] | 2026-05-12 | placeholder for future integration |

## Version history
| Version | Date | Change |
|---------|------|--------|
| 0.1.0 | 2026-05-12 | initial extraction |

## Test coverage
- Unit tests: partial
- Test file: ~/engineering/infrastructure/shared-libraries/auth-bcrypt/tests/
- Coverage: N/A
- Last test run: 2026-05-12

## Known failure modes
| Failure | Trigger condition | Workaround | Fixed in version |
|---------|-------------------|------------|-----------------|
| CPU starvation | High concurrent login requests | Rate limiting (brute-force-guard) | pending |

## Dependencies
- External: bcrypt@5.1.1
- Internal modules: none

## Performance characteristics
- Time complexity: O(2^cost) per hash
- Space complexity: O(1)
- Benchmarks: Approx ~10 hashes per second per core at cost=12

## Status transition log
- EXPERIMENTAL → ACTIVE: extracted after successful implementation in login-system-bcrypt and anticipated use in ledger-core

## Related
- Concept notes: [[03_CORE_KNOWLEDGE/security/auth/Password Hashing]], [[03_CORE_KNOWLEDGE/security/auth/bcrypt]]
- Failure library: [[06_FAILURE_LIBRARY/module-failures/auth-bcrypt]]
