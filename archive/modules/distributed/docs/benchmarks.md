---\nStatus: Partial\nImplementation: 70%\nConfidence: Tested\n---\n# Benchmark Plan

## Scope

This file tracks benchmark methodology and results for the C limiter core.

No results are recorded yet.

## Metrics

- operations per second
- p50 latency per operation
- p99 latency per operation
- memory footprint
- probe length distribution
- shard contention rate

## Test Profiles

### 1. Uniform Traffic

Description:

- large key space
- low contention
- steady request rate

Purpose:

- baseline throughput

### 2. Hot-Key Burst

Description:

- a small number of keys receive most traffic
- simulates abuse or attack concentration

Purpose:

- expose contention and boundary failures

### 3. Mixed Traffic

Description:

- mostly normal keys
- some bursty clients
- variable costs per request if supported

Purpose:

- approximate real middleware load

### 4. Near-Capacity Table

Description:

- fill table close to configured key limit

Purpose:

- measure probe amplification and degradation

### 5. Rollover Stress

Description:

- frequent transitions across bucket boundaries

Purpose:

- validate expiry logic under load

## Initial Configuration Matrix

| window_ms | bucket_ms | max_keys | threads |
| --- | --- | --- | --- |
| 60000 | 1000 | 65536 | 1 |
| 60000 | 1000 | 65536 | 4 |
| 60000 | 500 | 65536 | 1 |
| 60000 | 100 | 65536 | 1 |

## Measurement Rules

- use monotonic time
- pin benchmark mode and config in output
- separate warmup from measured interval
- report compiler flags with every run
- store raw numbers before summary interpretation

## Output Template

```text
benchmark: <name>
build: <compiler + flags>
window_ms: <value>
bucket_ms: <value>
max_keys: <value>
threads: <value>
duration_s: <value>
ops_total: <value>
ops_per_sec: <value>
p50_ns: <value>
p99_ns: <value>
memory_bytes: <value>
notes: <freeform>
```

## Result Log

### Pending

- no benchmark runs completed
