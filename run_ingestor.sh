#!/bin/bash
mkdir -p debug
touch debug/panic_audit.log
while true; do
  echo "panic_detected: latency_optimization" >> debug/panic_audit.log
  sleep 2
done
