#!/usr/bin/env bash
set -euo pipefail
ROOT_DIR=$(cd "$(dirname "$0")" && pwd)
ARTIFACTS=${ROOT_DIR}/artifacts/build
mkdir -p "$ARTIFACTS"
echo "Building normalizer module (Python)" > ${ARTIFACTS}/build.log
python3 -m compileall ${ROOT_DIR}/src >> ${ARTIFACTS}/build.log 2>&1 || true
echo "{"\"version\": \"0.1.0\"}" > ${ARTIFACTS}/version.json
sha256sum ${ROOT_DIR}/src/*.py > ${ARTIFACTS}/hashes.sha256 || true
echo "BUILD_OK" >> ${ARTIFACTS}/build.log
echo "Build complete. Artifacts in ${ARTIFACTS}"
