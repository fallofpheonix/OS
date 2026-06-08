# Canonical Numeric Model Specification

**Status:** DRAFT (Vertical Slice Step 4)
**Confidence:** High
**Owner:** Phoenix.Nucleus Team

## 1. Requirement
Bit-perfect determinism across all platforms and execution sessions. 
IEEE-754 `float64` is prohibited for any value that influences system state, consensus, or state hashes.

## 2. Representation
The system uses **Fixed-Point Arithmetic** backed by `int64`.

- **Type:** `FixedPoint` (struct-wrapped `int64`)
- **Base:** Decimal (Base-10)
- **Precision:** 6 decimal places
- **Scaling Factor:** 1,000,000 ($10^6$)
- **Range:** approximately $\pm 9.22 \times 10^{12}$ (sufficient for system metrics, reputation, and coordinates).

## 3. Arithmetic Behavior
- **Addition/Subtraction:** Standard integer operations on scaled values.
- **Multiplication:** $(A \times B) / 1,000,000$ to maintain scaling.
- **Division:** $(A \times 1,000,000) / B$ to maintain scaling.
- **Overflow:** All operations MUST use saturating semantics (`SaturatingAdd`, etc.) to prevent wrap-around corruption.
- **Rounding:** Truncation (floor toward zero) is the default for efficiency. If symmetric rounding is required, it must be explicitly named (e.g., `MulRounded`).

## 4. Serialization
- **Format:** JSON
- **Representation:** Integer (scaled value) or String.
- **Canonical Decision:** The value SHALL be serialized as a JSON object `{"v": <int64>}` or a String to ensure no JSON parser attempts to interpret it as a float.
- **Vertical Slice Choice:** `{"v": <int64>}` for clear typing.

## 5. Implementation Path
1. Consolidate `foundation/ledger/fixedpoint.go` and `foundation/runtime/common/math/fixed_point.go` into a single canonical type in `foundation/runtime/common/math/fixedpoint.go`.
2. Update all importers (`game`, `ledger`, `security`).
3. Add `determinism_test.go` to prove cross-platform compatibility (using pre-computed hex values).

## 6. Prohibited Primitives
- `float32`
- `float64`
- `math.Sin`, `math.Cos` (unless using deterministic lookup tables)
- Non-deterministic `map` iterations during state hashing
