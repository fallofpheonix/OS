/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
#include "sliding_window.h"

#include <stdint.h>
#include <stdio.h>
#include <time.h>

static uint64_t monotonic_ns(void) {
    struct timespec ts;
    clock_gettime(CLOCK_MONOTONIC, &ts);
    return ((uint64_t)ts.tv_sec * 1000000000ull) + (uint64_t)ts.tv_nsec;
}

int main(void) {
    sliding_window_t window;
    const uint64_t ops = 50000000ull;
    const uint64_t advance_every = 256ull;
    const uint32_t cost = 1u;
    uint64_t now_ms = 0;
    uint64_t sink = 0;

    sliding_window_init(&window);

    uint64_t start_ns = monotonic_ns();
    for (uint64_t i = 0; i < ops; ++i) {
        sink ^= sliding_window_increment(&window, now_ms, cost);
        if ((i & (advance_every - 1ull)) == 0ull) {
            now_ms += 1;
        }
    }
    uint64_t end_ns = monotonic_ns();

    double seconds = (double)(end_ns - start_ns) / 1000000000.0;
    double ops_per_sec = (double)ops / seconds;

    printf("ops_total=%llu\n", (unsigned long long)ops);
    printf("duration_s=%.6f\n", seconds);
    printf("ops_per_sec=%.2f\n", ops_per_sec);
    printf("state_bytes=%zu\n", sizeof(window));
    printf("final_sum=%llu\n", (unsigned long long)window.rolling_sum);
    printf("sink=%llu\n", (unsigned long long)sink);
    return 0;
}
