# Dependency Conflicts Report

## CRITICAL
None detected.

## HIGH
None detected.

## MEDIUM
- **requests**: Version drift ['2.34.0', '2.34.1'] across ['runtime/core', 'runtime/research', 'runtime/shared']
- **tokenizers**: Version drift ['0.22.2', '0.23.1'] across ['runtime/core', 'runtime/research', 'runtime/shared']

## LOW
- **Python Version**: Core (3.13.12) vs Research/Shared (3.14.3). This is planned drift.
- **Lock Mismatch**: Research and Shared runtimes lack authoritative lock files (requirements.txt and site-packages only).
