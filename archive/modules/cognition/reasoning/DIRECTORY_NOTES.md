# Reasoning Subsystem

## Primary Responsibility
The Reasoning subsystem implements the high-level inference and audit layers for PhoenixOS. It is responsible for LLM-agnostic strategic bridging (Provider) and human-readable accountability (Explanation Layer).

## System Architecture
1. **Inference Bridge:** A decoupled interface that translates system goals and grounding context into provider-specific prompts.
2. **Explanation Layer:** A formal auditing component that generates `ReasonPaths` (step-by-step authority chains) and `Counterfactual` proofs for every autonomous decision.

## Tech Stack
- Go (Context management)
- Reasoning.Provider Interface (Agnostic Bridge)

## AI-Specific Context
- **System Boundaries:** Northbound to the Nexus Oracle (L7). Southbound to the AIOrchestrator actuation pipeline.
- **Data Flow:** Inference Request -> Explanation Generation (Pre-Action) -> Oracle Reason -> Actuation Authorization.
