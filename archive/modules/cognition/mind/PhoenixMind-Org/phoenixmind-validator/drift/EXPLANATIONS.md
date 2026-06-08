# EXPLANATIONS.md - Drift Detection Subsystem

## behavior_drift.py
### Beginner (What it does)
Checks if the system is starting to act differently than it usually does.
### Intermediate (How it interacts)
Will interface with system logs and telemetry to compare current behavior against a "normal" baseline.
### Expert (Architectural role)
Ensures behavioral integrity of the PhoenixMind by detecting deviations from established operational profiles.

## dependency_drift.py
### Beginner (What it does)
Watches for changes in the libraries and tools the system relies on.
### Intermediate (How it interacts)
Monitors the project's dependency tree for unauthorized or unexpected updates.
### Expert (Architectural role)
Prevents supply-chain attacks and ensure deterministic builds by validating dependency consistency.

## memory_drift.py
### Beginner (What it does)
Ensures the system's "memory" or data isn't changing in unexpected ways.
### Intermediate (How it interacts)
Validates the integrity of episodic and semantic memory stores.
### Expert (Architectural role)
Detects silent data corruption or unauthorized modifications to the knowledge base.

## proposal_drift.py
### Beginner (What it does)
Checks if the suggestions the system makes are becoming less accurate or different.
### Intermediate (How it interacts)
Monitors the quality and alignment of generated proposals over time.
### Expert (Architectural role)
Ensures that the reasoning engine remains aligned with system goals and doesn't "hallucinate" or degrade.

## runtime_drift.py
### Beginner (What it does)
Watches the system while it's running to make sure it's staying healthy.
### Intermediate (How it interacts)
Monitors runtime metrics like CPU, memory, and execution paths.
### Expert (Architectural role)
Detects execution-time anomalies that might indicate resource exhaustion or exploit attempts.

## security_drift.py
### Beginner (What it does)
Specifically looks for changes that might make the system less secure.
### Intermediate (How it interacts)
Monitors security configurations and access patterns.
### Expert (Architectural role)
Provides continuous security auditing by detecting shifts in the system's security posture.

## structure_drift.py
### Beginner (What it does)
Makes sure the basic "skeleton" or organization of the system hasn't changed.
### Intermediate (How it interacts)
Validates the structural integrity of the codebase and its architectural mappings.
### Expert (Architectural role)
Ensures that the fractal architecture remains consistent and hasn't been compromised structurally.

## telemetry_drift.py
### Beginner (What it does)
Ensures that the sensors and reports coming from the system are still accurate.
### Intermediate (How it interacts)
Monitors the health and consistency of the telemetry stream.
### Expert (Architectural role)
Validates the "eyes and ears" of the system to ensure that monitoring data is reliable.
