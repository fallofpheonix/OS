#!/bin/bash
# I/O Stress script for benchmarking telemetry overhead
# Creates a large number of small writes and renames

TARGET_DIR="./stress_test_data"
mkdir -p $TARGET_DIR

echo "Starting I/O stress test in $TARGET_DIR..."

for i in {1..1000}
do
    echo "data" > "$TARGET_DIR/file_$i.txt"
    mv "$TARGET_DIR/file_$i.txt" "$TARGET_DIR/renamed_$i.txt"
done

echo "Stress test complete. Cleaning up..."
rm -rf $TARGET_DIR
