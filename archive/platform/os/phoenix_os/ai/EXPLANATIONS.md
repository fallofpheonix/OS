# Knowledge Transfer: Phoenix.Terminus/AI

## Beginner Explanation
The **AI Subsystem** in Terminus acts as the "Decision Support" engine. It takes telemetry events from the system bus and uses various "Features" (like process graphs or drift monitors) to understand if the system is behaving normally. If it detects something suspicious, it suggests a new "System State" to the Warden.

## Intermediate Explanation
This component implements the **Reasoning Layer** of the 7-layer architecture. It interacts with:
- `PhoenixCore/Bus`: For receiving real-time telemetry.
- `PhoenixGuard/Warden`: For executing state transitions.
- `PhoenixCore/Ledger`: For persisting its findings and certificates.
- `External/GitNexus`: For deep architectural awareness (via MCP).

It uses a modular `Feature` interface, allowing the OS to swap or add new detection algorithms (e.g., Graph analysis vs. Statistical drift) without changing the core orchestrator logic.

## Expert Explanation
The AI subsystem is architected as an **Asynchronous Advisory Loop**.
- **Orchestration:** The `AIOrchestrator` runs a `ProcessLoop` that drains a subscription from the `Bus`.
- **Feature Pipeline:** Each event is passed through a chain of `Features`. The `GraphFeature` builds a causal process graph, while the `MonitorFeature` calculates a statistical drift Z-Score.
- **Strategic Arbitration:** The `ArbiterFeature` reconciles these metrics against formal system invariants.
- **Actuation:** If a violation is confirmed, the `WardenFeature` issues an `AuthorityEscalationRequest` to the `PhoenixGuard` FSM.
- **Security Constraint:** The AI itself cannot directly modify system memory or kill processes; it must always go through the **Warden**, which verifies the request against a pre-compiled formal policy.
