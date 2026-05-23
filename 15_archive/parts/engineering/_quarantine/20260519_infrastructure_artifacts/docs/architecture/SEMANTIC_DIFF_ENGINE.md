# Semantic Diff Engine

## 1. Overview
The Semantic Diff Engine (SDE) shifts Astraeus from "text-editing" to "logic-mutating." Instead of calculating line-by-line diffs, the SDE calculates the **Delta of Intent** by comparing Abstract Syntax Trees (ASTs) and Repository Topology before and after a mutation.

## 2. AST-Aware Diffing
The SDE decomposes every file mutation into discrete AST operations:
- **Symbol Additions/Removals:** Detects if new classes, functions, or variables are introduced.
- **Logic Alterations:** Compares the `body` of functions. It ignores changes in whitespace or comments and focuses on control flow (if/else, loops) and state changes (assignments).
- **Attribute Scoping:** Monitors changes to visibility (e.g., changing `_private` to `public`).

## 3. Topology-Aware Diffing
Mutations are mapped to the `ArchitectureGraph` to evaluate their "Blast Radius":
- **Edge Mutation:** Detects if a mutation introduces a new dependency or usage edge.
- **Boundary Crossings:** Flags if a mutation touches code in a higher `Criticality` tier than the task's scope.
- **Dependency Strengthening:** Detects if a module becomes more tightly coupled to another module post-mutation.

## 4. Interface-Aware Diffing
The SDE specifically monitors the "Contracts" between modules:
- **Signature Parity:** Detects changes in method parameters, default values, or return type hints.
- **Contract Breakage:** If `Module A` depends on `Module B.foo()`, and a mutation in `Module B` changes `foo()`'s signature or return type, the SDE flags a **Contract Violation** even if the syntax remains valid.

## 5. Behavior-Aware Diffing
- **In-Memory Execution Simulation:** (Experimental) The SDE uses symbolic execution or light-weight mocking to verify that a mutation to `x > 0` does not break the invariant `x != 0`.
- **Semantic Hash:** Every function is assigned a "Semantic Hash" based on its AST structure. If a mutation changes the text but the Semantic Hash remains identical (e.g., refactoring variable names), the SDE marks it as a "Zero-Impact Refactor."

## 6. Implementation Strategy
- **Extraction:** Use Python's `ast` module to generate canonical representations of code blocks.
- **Comparison:** Use a tree-differencing algorithm (e.g., Zhang-Shasha or similar) to find the minimum set of AST transformations.
- **Integration:** The SDE output is passed to the `VerificationPipeline` and the `RepairPlanner` to provide high-fidelity feedback.
