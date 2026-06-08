# Module: [Name]

**Status:** EXPERIMENTAL | EVOLVING | STABLE  
**Extracted:** [Date] from [[Parent Project]]  
**Location:** `modules/[module-name]/`

## Purpose
[One-paragraph description of what this module does]

## Source Project
[[Project Name]] — [version extracted]

## Extraction Rationale
[Why was this extracted?]
[How many projects needed it?]
[What patterns emerged?]

## Public API Surface

### Core Functions/Classes
```python
# What is exported?
def primary_interface():
    pass

class MainAbstraction:
    pass
```

### Input/Output Contracts
```python
# Type signatures and requirements
```

## Dependencies
- dependency-1
- dependency-2
- [none if self-contained]

## Coupling Removed
[What was decoupled by extracting this?]
- Before: [Entangled pattern]
- After: [Clean boundary]

## Coupling Remaining
[What still depends on this being in a certain place?]

## Test Coverage
- Unit tests: X%
- Integration tests: Present?
- Edge cases covered?

## Known Limitations
- Limitation 1
- Limitation 2

## Consumers
- [[Project A]] — usage context
- [[Project B]] — usage context
- [In development]

## API Stability
**Current:** STABLE | EVOLVING | EXPERIMENTAL

If EVOLVING:
- Breaking changes possible in: [areas]
- Stabilization target: [date]

## Future Improvements
- Improvement 1: [description]
- Improvement 2: [description]
- Performance consideration: [area]

## Related
- Extracted from: [[Project Name]]
- Sibling module: [[Other Module]]
- [[Related ADR]]
