#!/bin/bash
set -e
mkdir -p artifacts/build
echo "Building __MODULE_NAME__..."
go build -o artifacts/build/__MODULE_NAME__ ./src/...
go vet ./...
sha256sum artifacts/build/__MODULE_NAME__ > artifacts/build/hashes.sha256
echo '{"module": "__MODULE_NAME__", "version": "0.1.0", "built": "'$(date -u +%Y-%m-%dT%H:%M:%SZ)'"}' > artifacts/build/version.json
echo "Build complete." | tee artifacts/build/build.log
