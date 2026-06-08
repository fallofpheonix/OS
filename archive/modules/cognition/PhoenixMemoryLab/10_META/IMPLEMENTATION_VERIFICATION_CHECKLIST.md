# IMPLEMENTATION VERIFICATION CHECKLIST

## 0. When To Use This
Use this while writing code. Not before. Not after.
After every file creation, validate against this list.

## 1. Domain Layer Verification
For every file inside `lib/domain/`:
* Contains zero Flutter imports.
* Contains zero database imports.
* Uses only Dart SDK.
* Has no async IO.
* No serialization annotations.
* No JSON methods.
* No framework types.
If any appear, delete and rewrite.

## 2. Model Integrity Check
Each model must behave like a record, not an object graph.
Must NOT contain:
* UI formatting
* Validation logic
* Repository calls
* Constructors doing work
Models are data only.

## 3. Repository Contract Check
Inside `domain/repositories`:
* Only abstract classes allowed.
* No implements.
* No persistence details.
* No DTO references.
If you mention Hive, Isar, SQLite here, you broke the boundary.

## 4. Infrastructure Check
Inside `lib/infrastructure/`:
* This is the only place allowed to import database packages.
* DTOs must never escape this layer.
* Mapping Domain ↔ DTO must be explicit.
Search entire project: `grep -R "VitalRecord" lib/`
It must only exist in infrastructure.

## 5. Controller Check
Controllers must:
* Import domain services.
* Import repository interfaces.
* NOT import database.
* NOT import Flutter widgets.
* NOT contain business rules.
Controllers orchestrate. Nothing else.

## 6. Provider Check
Providers must:
* Call controller methods only.
* Not compute anything.
* Not validate anything.
* Not transform data.
If a provider contains logic, move it down a layer.

## 7. UI Check
Widgets must:
* Read ViewState only.
* Never call repository.
* Never call services.
* Never parse data.
* Never spawn isolates.
UI is a renderer, not a participant.

## 8. Data Flow Validation
You must be able to trace every write like this:
Widget → Provider → Controller → Domain ValidationService → Repository Interface → Infrastructure Implementation → Database
If any arrow is skipped, architecture is already rotting.

## 9. Determinism Test (Run Daily)
1. Insert sample data.
2. Kill app.
3. Delete in-memory state.
4. Reload from DB.
5. Recompute baseline.
Output must be identical.
If not identical, you introduced hidden state.

## 10. Query Safety Check
Every fetch must include a bounded range.
Forbidden: `SELECT * FROM vitals;`
Required pattern: `WHERE timestamp BETWEEN X AND Y`
No unbounded scans allowed.

## 11. Soft Delete Enforcement
Verify repository filters deleted records internally.
Caller must NOT pass flags like: `includeDeleted = false`
That leaks persistence concern upward.

## 12. Background Work Check
All heavy computation must go through: `core/background/background_service.dart`
Search for `compute(` or isolate usage outside this file. There must be none.

## 13. File Creation Discipline
If adding a feature required editing more than:
* 1 domain file
* 1 infrastructure adapter
* 1 controller
* 1 mapper
* 1 UI file
You created coupling. Stop and reassess.

## 14. Build Should Work With UI Deleted
Temporarily comment out `features/`. Domain tests must still pass. If they fail, UI leaked into logic.

## 15. Sprint 1 Exit Criteria
You are allowed to move forward only when:
* You can log weight.
* You can fetch by date range.
* Baseline recomputes correctly.
* App works fully offline.
* No architectural violations found above.
No extra features permitted before this state is clean.
