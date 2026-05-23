# RFC Alignment & Merge Report

*   **RFC-001 (Telemetry Schema):** KEEP. Core dependency.
*   **RFC-001A (Hospital Ransomware):** MERGE into `02_docs/specifications/real_world_use_cases.md`.
*   **RFC-001B (Anomaly Logic) & RFC-001C (Phoenix Monitor):** MERGE into RFC-001 as mathematical sub-modules.
*   **RFC-001D (Strategic Containment Game):** RENAME to `RFC-010_stackelberg_containment.md`.
*   **RFC-002 (Phoenix Bus):** KEEP. Required for routing.
*   **RFC-003 (Forensic Snapshot):** DEFER. High-overhead; build process DAGs first.
*   **RFC-004 (AI Correlator):** DEFER until Graph Engine and Telemetry are fully saturated and benchmarked.
*   **RFC-005 (Phoenix Warden):** KEEP. Connects directly to PID Control.
*   **RFC-006 (Phoenix Trace):** KEEP. Critical L4 module.
*   **RFC-007 (Event Normalizer):** KEEP. Prerequisite for RFC-006.
*   **RFC-008 (Phoenix Nexus):** DEFER until R031 (PID Control) is validated.
*   **RFC-009 (Math/Physics Architecture):** KEEP as Master Arch Document.