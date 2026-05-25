# Module Boundaries

## 1. Allowable Module Interactions
- **Apps** may compose **Modules** and call **Services**.
- **Services** may compose **Modules**.
- **Modules** must be leaf nodes; they cannot import **Apps** or **Services**.
- **Infrastructure** operates independently and cannot depend on **Apps**, **Services**, or **Modules** (unless strictly for deployment tooling).

## 2. The "Shared" Violation
The presence of a top-level `shared/` directory alongside `modules/` creates ambiguity. 
- All domain-agnostic library code must reside in `modules/`.
- All domain-specific reusable code must reside within the respective domain's bounded context.
- `shared/` is deprecated and marked for refactoring.

## 3. Tooling Boundaries
- CI and Build Tools must reside in `infrastructure/`.
- Tooling cannot bleed into `apps/` or `modules/`.
