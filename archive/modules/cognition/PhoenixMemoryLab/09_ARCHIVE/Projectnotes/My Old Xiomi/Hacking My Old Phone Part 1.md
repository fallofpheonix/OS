## Stage 1: Environment Setup

Set up compiler, filesystem structure, and execution workflow on Mac and mobile.

**Substages:**

- Install toolchain (`clang`, `ssh`)
- Create project structure
- Establish compile + run cycle

---

## Stage 2: Memory Fundamentals (Vector)

Build dynamic array to understand heap allocation, resizing, and pointer control.

**Substages:**

- Struct design (`data`, `size`, `capacity`)
- `malloc` / `realloc` usage
- Bounds-safe operations

---

## Stage 3: Validation & Testing

Verify correctness using structured test cases and failure injection.

**Substages:**

- Functional tests (push, get, set)
- Edge cases (invalid index, empty pop)
- Stress test (large inserts)

---

## Stage 4: Multi-Environment Execution

Run and validate code across systems (Mac → mobile via SSH).

**Substages:**

- Setup Termux
- Transfer files (`scp`)
- Compile and execute remotely

---

## Stage 5: Data Structure Expansion

Move to non-contiguous memory structures (linked list) to deepen pointer reasoning.

**Substages:**

- Node-based design
- Pointer-to-pointer manipulation
- Memory lifecycle management