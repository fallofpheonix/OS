# Knowledge Transfer: Phoenix.Terminus/Arbiter

## Beginner Explanation
The **Arbiter** is the "Judge" of PhoenixOS. It looks at the evidence provided by the AI and the sensors and decides how serious a problem is. For example, if it sees a small drift, it might just log it, but if it sees a major security violation, it will decide to isolate the process.

## Intermediate Explanation
The Arbiter implements the **Strategic Layer**. It reconciles high-level AI advice with raw system metrics. Its primary job is to map a `DriftScore` and a `TCSScore` (Trust Score) into a `SystemState` (Safe, Suspicious, Critical) and an `ActuationClass` (Log, Throttle, Isolate).

## Expert Explanation
The Arbiter acts as the **Policy Reconciliation Engine**.
- **Input:** Takes `monitor.DriftScore` (statistical anomaly) and `tcsScore` (cryptographic/formal trust).
- **Logic:** Implements a state-transition matrix. Higher Z-Scores trigger more aggressive containment states.
- **Trust-Gating:** It performs a "Cross-check" where low TCS scores can downgrade or block AI-driven decisions to prevent "AI Hallucination" from triggering system-wide DoS.
- **Authorization:** It provides a signed recommendation to the Warden, which remains the final arbiter of truth.
