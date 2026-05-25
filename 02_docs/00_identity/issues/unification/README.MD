# DOC-001 through DOC-010: Documentation Unification Pass

**Label**: documentation, consolidation
**Status**: OPEN

## Objective
Reduce documentation count from 150+ to 25 canonical files.

## Issues
- **DOC-001**: Merge runtime docs
- **DOC-002**: Merge replay docs
- **DOC-003**: Merge truth docs
- **DOC-004**: Merge validation docs
- **DOC-005**: Merge roadmap variants
- **DOC-006**: Merge future systems
- **DOC-007**: Remove generated artifacts
- **DOC-008**: Research relevance scan
- **DOC-009**: Duplicate folder collapse
- **DOC-010**: Final doc freeze

## Implementation Plan
1. Generate `DOC_SIMILARITY_MATRIX.md` to identify overlaps.
2. Draft `DOC_KEEP_LIST.md` and `DOC_REMOVE_LIST.md`.
3. Create `DOC_MERGE_PLAN.md` for multi-source files.
4. Execute merges and move redundant docs to `archive/documentation/version_history/`.
