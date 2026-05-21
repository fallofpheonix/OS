#!/bin/bash
set -e
mkdir -p artifacts/build
echo "Building agents..."
go build -o artifacts/build/agents ./src/...
go vet ./...
sha256sum artifacts/build/agents > artifacts/build/hashes.sha256
echo '{"module": "agents", "version": "0.1.0", "built": "'$(date -u +%Y-%m-%dT%H:%M:%SZ)'"}' > artifacts/build/version.json
echo "Build complete." | tee artifacts/build/build.log
