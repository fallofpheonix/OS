# RFC: __MODULE_NAME__

## 1. Problem Statement
Concise statement of the problem this module will solve. Include assumptions, scope, and why current solutions are insufficient. Keep to 3–5 sentences.

## 2. Specification
Define the expected behaviour, data formats, configuration knobs, and any invariants. Include example inputs and outputs where helpful.

## 3. Interface
List public APIs, command-line options, file layouts, and RPC/IPC contracts. Provide example calls and expected responses.

## 4. Failure Modes
Enumerate possible failure modes, their severity, and suggested mitigation strategies or recovery plans.

## 5. Performance Budget
Specify latency/throughput targets, memory constraints, and benchmark methods. Reference `bench/` tests used to validate these budgets.

## 6. Validation Criteria
Concrete pass/fail gates for merging and shipping (build, tests, coverage, benchmarks, replay determinism). Link to CI jobs or local commands to run validation.
