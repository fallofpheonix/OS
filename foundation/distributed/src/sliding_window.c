/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
#include "sliding_window.h"

#include <stddef.h>
#include <string.h>

static inline uint32_t sw_bucket_from_ms(uint64_t now_ms) {
    return (uint32_t)(now_ms / SW_BUCKET_MS);
}

static inline uint32_t sw_index_from_bucket(uint32_t bucket_id) {
    return (uint32_t)(bucket_id % SW_NUM_BUCKETS);
}

static void sw_expire_until(sliding_window_t *window, uint32_t now_bucket) {
    if (window->last_bucket == SW_INVALID_BUCKET) {
        window->last_bucket = now_bucket;
        return;
    }

    if (now_bucket <= window->last_bucket) {
        return;
    }

    uint32_t delta = now_bucket - window->last_bucket;
    uint32_t to_clear = (delta >= SW_NUM_BUCKETS) ? SW_NUM_BUCKETS : (uint32_t)delta;

    for (uint32_t i = 1; i <= to_clear; ++i) {
        uint32_t idx = sw_index_from_bucket(window->last_bucket + i);
        window->rolling_sum -= window->counts[idx];
        window->counts[idx] = 0;
        window->bucket_ids[idx] = SW_INVALID_BUCKET;
    }

    window->last_bucket = now_bucket;
}

void sliding_window_init(sliding_window_t *window) {
    memset(window, 0, sizeof(*window));
    window->last_bucket = SW_INVALID_BUCKET;
    for (size_t i = 0; i < SW_NUM_BUCKETS; ++i) {
        window->bucket_ids[i] = SW_INVALID_BUCKET;
    }
}

uint64_t sliding_window_increment(sliding_window_t *window, uint64_t now_ms, uint32_t cost) {
    uint32_t now_bucket = sw_bucket_from_ms(now_ms);
    uint32_t idx;

    sw_expire_until(window, now_bucket);

    idx = sw_index_from_bucket(now_bucket);
    if (window->bucket_ids[idx] != now_bucket) {
        window->rolling_sum -= window->counts[idx];
        window->counts[idx] = 0;
        window->bucket_ids[idx] = now_bucket;
    }

    window->counts[idx] += cost;
    window->rolling_sum += cost;
    return window->rolling_sum;
}

uint64_t sliding_window_query(sliding_window_t *window, uint64_t now_ms) {
    sw_expire_until(window, sw_bucket_from_ms(now_ms));
    return window->rolling_sum;
}
