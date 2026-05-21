# MVP Sprint Plan: Pheonix

This sprint plan details the phased engineering roadmap for delivering the Pheonix MVP.

## Sprint 1: Architecture, Design & Governance (Milestone: Design Sign-off)
*   **Duration:** 1 week
*   **Objectives:** Establish complete architectural specifications, API interfaces, schema definitions, and threat models.
*   **Tasks:**
    1.  Detail Telemetry Schema RFC-001 (Process, Syscall, Filesystem, Network, Container, Memory).
    2.  Define telemetry_events.json JSON schema.
    3.  Detail Unified Event Bus RFC-002 (Normalizer, Broker, Publisher/Subscriber APIs).
    4.  Detail Forensic Snapshot RFC-003 (Triggers, Memory/Disk/Process capture state machine).
    5.  Detail AI Correlator & Incident Graph RFC-004 (Graph schema, anomaly rules, offline LLM interface).
    6.  Detail Containment Primitives RFC-005 (Network & Process block mechanisms).
    7.  Establish Risk Register and project execution gates.

## Sprint 2: Foundational Data Pipeline (Milestone: Event Flow Validation)
*   **Duration:** 1 week
*   **Objectives:** Implement the Telemetry Agent and Unified Event Bus with local simulation capabilities.
*   **Tasks:**
    1.  Initialize Go modular workspace under `14_experiments/`.
    2.  Implement Telemetry Agent with local macOS Simulation Mode (sending random/patterned JSON events).
    3.  Build eBPF probe scaffold (Linux structure, loaded stub for prod mode).
    4.  Implement Event Bus Broker: TCP/Unix domain socket listener, JSON event normalizer, and pub-sub router.
    5.  Validate pipeline performance (target: >50k events/sec under local simulation).

## Sprint 3: Intelligence, Incident Capture & SOC UI (Milestone: E2E MVP Release)
*   **Duration:** 1 week
*   **Objectives:** Build the correlation engine, forensic capture, and visual dashboard.
*   **Tasks:**
    1.  Implement Correlation Engine (Incident Graph builder, pattern/state rules like process-spawned-shell-with-active-network-connection).
    2.  Implement Forensic Snapshot Runtime (saves process memory dumps, file handles, and open connection lists to `08_forensics/reports/` when an anomaly alerts).
    3.  Build premium Local SOC Web Dashboard: Glassmorphic dark UI, real-time alert ticker, system event feed, incident graph visualizer, and manual forensic trigger buttons.
    4.  Run End-to-End simulation tests.
