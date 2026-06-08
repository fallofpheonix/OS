# PhoenixMind Orchestration Subsystem

## Primary Responsibility
PhoenixMind is the central coordination plane for the PhoenixOS intelligence layer. It is responsible for orchestrating the 5-layer cognition stack, enforcing Guarded Autonomy, and managing the collaborative interface between the AI Oracle and the deterministic Warden substrate.

## System Architecture
The subsystem is organized into two primary packages:
1. **Intelligence:** Contains the `AIOrchestrator`, the `NexusBridge` (Oracle), and the governance/analytics hub. It executes the core `OrchestrateTick` pipeline.
2. **Memory:** Provides the context retrieval bridge that grounds AI reasoning in formal tiered episodic memory.

## Tech Stack
- Go 1.26
- gRPC / HTTP (Oracle Communication)
- internal/bus (Event Pipeline)

## AI-Specific Context
- **System Boundaries:** Central hub. Integrates all root cognitive packages (`knowledge`, `memory`, `reasoning`, `reflection`).
- **Data Flow:** Telemetry -> Monitor -> Causal Analysis -> Policy Evaluation -> AI Reasoning -> User Permission -> Warden Actuation.
