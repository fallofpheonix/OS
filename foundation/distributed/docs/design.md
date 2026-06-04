---\nStatus: Partial\nImplementation: 70%\nConfidence: Tested\n---\n# Design Notes

## Scope

This document defines the Week 1 design for the C sliding-window rate limiter core.

## Data Flow

```text
request key + timestamp
        ->
hash table lookup
        ->
per-key sliding window advance
        ->
rolling count update
        ->
allow / block result
```

## Core Components

### 1. Fixed-Size Key Table

Purpose:

- map client keys to rate limiter state
- avoid dynamic allocation during steady-state operation

Initial approach:

- open addressing
- fixed capacity
- bounded probe length instrumentation

Rejected for Week 1:

- resizing hash maps
- unbounded chaining
- cross-process shared state

### 2. Per-Key Sliding Window

Representation:

- ring of fixed-width time buckets
- rolling sum of live buckets
- last observed bucket index / timestamp

Target behavior:

- `query`: return current live count for the configured window
- `allow_and_update`: increment live count and decide whether limit is exceeded

### 3. Time Source

Requirements:

- monotonic milliseconds
- externally injectable in tests
- no wall-clock dependence in core logic

## Invariants

- expired buckets contribute zero
- rolling sum equals sum of all live buckets
- bucket advancement is monotonic per key
- state transitions are deterministic for identical `(key, now_ms, cost)` sequences
- no hot-path heap allocation after initialization

## Proposed API

```c
typedef struct limiter limiter_t;

typedef struct {
    uint32_t limit;
    uint32_t window_ms;
    uint32_t bucket_ms;
    uint32_t max_keys;
} limiter_config_t;

typedef struct {
    uint32_t current_count;
    uint32_t remaining;
    uint8_t allowed;
} limiter_result_t;

int limiter_init(limiter_t *l, const limiter_config_t *cfg);
void limiter_destroy(limiter_t *l);
limiter_result_t limiter_allow_and_update(limiter_t *l, uint64_t key, uint64_t now_ms, uint32_t cost);
uint32_t limiter_query(limiter_t *l, uint64_t key, uint64_t now_ms);
int limiter_reset(limiter_t *l, uint64_t key);
```

## Complexity Targets

- lookup/update/query: `O(1)` amortized
- memory: `O(max_keys * num_buckets)`

Where:

- `num_buckets = window_ms / bucket_ms`

## Concurrency Model

Week 1 target is correctness first, then bounded contention.

Candidate approaches:

1. shard-level locks with fixed shards
2. lock-free entry mutation on a fixed table

Preferred initial implementation:

- fixed shards
- one lock per shard
- deterministic ownership of each key by shard

Reason:

- simpler correctness proof
- easier instrumentation
- lower implementation risk for first milestone

## Risks

- hot-key contention under attack traffic
- false blocks near window boundaries if expiry logic is wrong
- probe amplification at high table occupancy
- oversized bucket arrays causing cache pressure

## Validation Plan

- unit tests for rollover and expiry
- boundary tests for window edges
- collision stress tests
- single-thread throughput benchmark
- multi-thread contention benchmark
