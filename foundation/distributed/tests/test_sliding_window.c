/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
#include "sliding_window.h"

#include <assert.h>
#include <stdio.h>

static void test_same_bucket_accumulates(void) {
    sliding_window_t window;
    sliding_window_init(&window);

    assert(sliding_window_increment(&window, 0, 1) == 1);
    assert(sliding_window_increment(&window, 999, 2) == 3);
    assert(sliding_window_query(&window, 999) == 3);
}

static void test_bucket_rollover_keeps_live_counts(void) {
    sliding_window_t window;
    sliding_window_init(&window);

    assert(sliding_window_increment(&window, 0, 5) == 5);
    assert(sliding_window_increment(&window, 1000, 7) == 12);
    assert(sliding_window_query(&window, 1500) == 12);
}

static void test_idle_gap_expires_old_counts(void) {
    sliding_window_t window;
    sliding_window_init(&window);

    assert(sliding_window_increment(&window, 0, 4) == 4);
    assert(sliding_window_increment(&window, 59000, 6) == 10);
    assert(sliding_window_query(&window, 60000) == 6);
    assert(sliding_window_query(&window, 119000) == 0);
}

static void test_full_window_gap_resets_sum(void) {
    sliding_window_t window;
    sliding_window_init(&window);

    assert(sliding_window_increment(&window, 0, 9) == 9);
    assert(sliding_window_query(&window, 60000) == 0);
    assert(sliding_window_increment(&window, 60000, 3) == 3);
}

int main(void) {
    test_same_bucket_accumulates();
    test_bucket_rollover_keeps_live_counts();
    test_idle_gap_expires_old_counts();
    test_full_window_gap_resets_sum();
    puts("sliding_window tests passed");
    return 0;
}
