/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
#ifndef SLIDING_WINDOW_H
#define SLIDING_WINDOW_H

#include <stdint.h>

#define SW_NUM_BUCKETS 60u
#define SW_BUCKET_MS 1000u
#define SW_INVALID_BUCKET UINT32_MAX

typedef struct __attribute__((aligned(64))) {
    uint32_t last_bucket;
    uint32_t bucket_ids[SW_NUM_BUCKETS];
    uint32_t counts[SW_NUM_BUCKETS];
    uint64_t rolling_sum;
} sliding_window_t;

void sliding_window_init(sliding_window_t *window);
uint64_t sliding_window_increment(sliding_window_t *window, uint64_t now_ms, uint32_t cost);
uint64_t sliding_window_query(sliding_window_t *window, uint64_t now_ms);

#endif
