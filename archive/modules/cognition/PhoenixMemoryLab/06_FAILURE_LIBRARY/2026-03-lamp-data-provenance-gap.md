---
failure-id: 2026-03-lamp-data-provenance-gap
project: [[05_PROJECTS/ACTIVE/lamp]]
severity: MEDIUM
status: OPEN
date-encountered: 2026-03
tags: [failure, data, reproducibility]
---
# Failure: Undocumented data lineage for input terrain data

## What Was Tried
Using terrain/building data in LAMP's path tracing and viewshed pipelines without formal provenance documentation.

## What Happened
Input data sources (DEMs, building footprints, surface classifications) have no documented origin, version, download date, or processing history. If the data needs to be re-acquired or updated, there is no record of where it came from or how it was preprocessed.

## Root Cause
Data was treated as a given input, not as a versioned dependency. No data provenance standard was established at project start.

## What Was Learned
Scientific computing projects must treat data as a dependency with the same rigor as code dependencies. Every input dataset needs: source URL, download date, version/revision, preprocessing steps, and a checksum.

## Prevention / Resolution
- Create a datasets/contracts/ directory with provenance YAML for each input dataset
- Add checksums (sha256) for all input files
- Document preprocessing steps (reprojection, cropping, resampling) as scripts, not manual steps
- Consider DVC (Data Version Control) for large datasets

## Linked Concepts
- [[03_CORE_KNOWLEDGE/databases]] — data versioning, data lineage
- [[03_CORE_KNOWLEDGE/algorithms]] — reproducible scientific computing
