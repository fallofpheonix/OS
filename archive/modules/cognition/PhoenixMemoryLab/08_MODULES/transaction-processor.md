# Module: transaction-processor

## Identity
- **Slug**: transaction-processor
- **Owner**: you
- **Location**: ~/engineering/infrastructure/shared-libraries/transaction-processor/
- **Status**: ACTIVE
- **Version**: 0.1.0
- **Language**: TypeScript
- **Created**: 2026-05-12
- **Last updated**: 2026-05-12

## One-liner
Ensures atomic state mutations and prevents double-spending using idempotent caching and distributed locks.

## API interface
```typescript
class TransactionProcessor {
  processTransfer(idempotencyKey: string, from: string, to: string, amount: number) -> Promise<Receipt>
}
```

## Installation / import
```typescript
import { TransactionProcessor } from '../../infrastructure/shared-libraries/transaction-processor';
const txProcessor = new TransactionProcessor();
```

## Usage example
```typescript
app.post('/transfer', async (req, res) => {
    const key = req.headers['idempotency-key'];
    const receipt = await txProcessor.processTransfer(key, req.body.from, req.body.to, req.body.amount);
    res.json(receipt);
});
```

## Configuration
Requires a persistent Key-Value store (e.g., Redis) for the idempotency cache in production. Currently defaults to in-memory `Map` for sandbox execution.

## Used by
| Project | Since | Notes |
|---------|-------|-------|
| [[05_PROJECTS/ACTIVE/ledger-core]] | 2026-05-12 | uses v0.1.0 |

## Version history
| Version | Date | Change |
|---------|------|--------|
| 0.1.0 | 2026-05-12 | initial extraction (Capstone Phase 2) |

## Test coverage
- Unit tests: yes
- Test file: ~/engineering/infrastructure/shared-libraries/transaction-processor/tests/
- Coverage: complete
- Last test run: 2026-05-12

## Known failure modes
| Failure | Trigger condition | Workaround | Fixed in version |
|---------|-------------------|------------|-----------------|
| Cache Eviction | Redis drops key before retry window closes | Increase TTL for tx keys to 24h | pending |

## Dependencies
- External: `redis` (planned for production)
- Internal modules: none

## Performance characteristics
- Time complexity: O(1) hash lookup
- Space complexity: O(N) where N is number of transactions in TTL window

## Status transition log
- EXPERIMENTAL → ACTIVE: Integrated and verified in ledger-core project.

## Related
- Concept notes: [[03_CORE_KNOWLEDGE/distributed-systems/Idempotency]], [[03_CORE_KNOWLEDGE/databases/Distributed Locks]], [[03_CORE_KNOWLEDGE/databases/ACID Properties]]
- Failure library: [[06_FAILURE_LIBRARY/module-failures/transaction-processor]]
