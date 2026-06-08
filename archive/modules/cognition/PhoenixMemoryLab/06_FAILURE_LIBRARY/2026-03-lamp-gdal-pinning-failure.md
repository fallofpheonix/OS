---
failure-id: 2026-03-lamp-gdal-pinning-failure
project: [[05_PROJECTS/ACTIVE/lamp]]
severity: HIGH
status: OPEN
date-encountered: 2026-03
tags: [failure, dependencies, geospatial]
---
# Failure: GDAL version differences cause cross-platform runtime failures

## What Was Tried
Developing LAMP on macOS and deploying/testing on Linux. GDAL installed via different package managers (Homebrew on macOS, apt on Linux).

## What Happened
GDAL behaves differently across platforms. Raster operations that work on one platform produce different results or fail entirely on another. Coordinate projection edge cases and raster I/O format support vary between GDAL versions.

## Root Cause
GDAL is a C library with complex platform-specific build configurations. Different package managers distribute different versions with different compile flags. No version pinning strategy was in place.

## What Was Learned
Geospatial dependencies (GDAL, PROJ, GEOS) must be treated as system-level infrastructure, not pip-installable libraries. Docker is the only reliable cross-platform reproducibility strategy for GDAL-dependent projects.

## Prevention / Resolution
- Pin exact GDAL version in Dockerfile (not "latest")
- Document known platform caveats (macOS GDAL from Homebrew vs Linux GDAL from apt)
- Add cross-platform CI matrix (macOS + Linux) to catch platform-specific failures
- Consider using conda-forge for local development (more consistent GDAL versions)

## Linked Concepts
- [[03_CORE_KNOWLEDGE/databases]] — geospatial data, raster processing
- [[03_CORE_KNOWLEDGE/devops]] — dependency management, reproducible environments
