# AI Agent Registry

This document records the identity, capabilities, and boundaries of all AI agents operating within PhoenixOS.

## Agent Identity: PhoenixMind (L6 Advisor)
- **Agent ID:** PM-001
- **Role:** Cybersecurity Forensic Analyzer & Explainer.
- **Input:**
  - `monitor.DriftScore` (Z-Score, EventType, Severity, Frequency)
  - `tcsScore` (Confidence index)
- **Output:**
  - Suggest Command (ISOLATE_PID, THROTTLE_NETWORK, REVOKE_UID, LOG_ONLY)
  - Confidence Score (0.0 to 1.0)
  - Reasoning (Text explaining anomaly)
- **Permissions:** Read-only access to L4 Trace DB and telemetry event feeds. No write permissions to Ledger, Warden, or raw filesystem.
- **Memory Scope:** Stateless inference per batch request (5-second aggregation window).
- **Allowed Actions:** Push JSON advice strings to log streams for manual operator inspection.
- **Escalation Path:** In case of LLM drift or anomalous command generation, the AI loop drops the advice and defaults to static Warden FSM containment rules.

## AI Alignment Mandate
AI is strictly advisory. Under **Axiom 3**, AI outputs must pass human review or warden threshold verification. AI can never autonomously escalate system state to recovery or containment.
