# PHOENIX MATRIX SOVEREIGN ARCHITECTURE
# [STATUS]: 18-Repository Substrate Consolidated
# [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
# [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
# [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
CC := cc
CFLAGS := -std=c11 -O3 -Wall -Wextra -Wpedantic -march=native -Iinclude

.PHONY: all test bench clean

all: test bench

test: tests/test_sliding_window
	./tests/test_sliding_window

bench: bench/bench_sliding_window
	./bench/bench_sliding_window

tests/test_sliding_window: tests/test_sliding_window.c src/sliding_window.c include/sliding_window.h
	$(CC) $(CFLAGS) -o $@ tests/test_sliding_window.c src/sliding_window.c

bench/bench_sliding_window: bench/bench_sliding_window.c src/sliding_window.c include/sliding_window.h
	$(CC) $(CFLAGS) -o $@ bench/bench_sliding_window.c src/sliding_window.c

clean:
	rm -f tests/test_sliding_window bench/bench_sliding_window
