#!/bin/bash
echo "[DEPLOY] Initializing PhoenixOS Infrastructure..."
# Get root directory
ROOT_DIR=$(pwd)
mkdir -p "$ROOT_DIR/build/runtime"
cp "$ROOT_DIR/build/bin/phoenix_"* "$ROOT_DIR/build/runtime/"
echo "[DEPLOY] Runtime environment prepared in $ROOT_DIR/build/runtime/"
