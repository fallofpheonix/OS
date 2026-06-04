---
Status: Planned
Implementation: 0%
Confidence: Conceptual
---
# Game Architecture — Index

This directory establishes the architecture-first specifications of the Phoenix Game Layer, before implementation code is written, to prevent architectural drift.

## Document Directory

### Vision & System Core
- [VISION.md](./VISION.md): Conceptual model of the gamified sandbox environment.
- [GAMEPLAY_LOOP.md](./GAMEPLAY_LOOP.md): Interactive cycle between agents, user, and sandbox.
- [PROGRESSION_SYSTEM.md](./PROGRESSION_SYSTEM.md): Dynamic curriculum design and level unlocking.
- [REWARD_SYSTEM.md](./REWARD_SYSTEM.md): Optimization reinforcement scoring functions.
- [LEADERBOARD_SYSTEM.md](./LEADERBOARD_SYSTEM.md): Verification and logging of high scores.
- [AI_HINT_SYSTEM.md](./AI_HINT_SYSTEM.md): Contextual hint generation using the LLM interface.
- [MULTIPLAYER_SYSTEM.md](./MULTIPLAYER_SYSTEM.md): Multiplayer state replication and sync.
- [ROADMAP.md](./ROADMAP.md): Development milestones.

### Engine Integration
- [Godot Architecture](./engine/GODOT_ARCHITECTURE.md): scene tree layout.
- [Entity System](./engine/ENTITY_SYSTEM.md): Component mapping.
- [Deterministic Simulation](./engine/DETERMINISTIC_SIMULATION.md): Time ticks.
- [Replay System](./engine/REPLAY_SYSTEM.md): Deterministic playback.
- [Physics Model](./engine/PHYSICS_MODEL.md): Dynamic triggers.

### P-Script Subsystem
- [Language Spec](./pscript/LANGUAGE_SPEC.md): Syntax and types.
- [Grammar](./pscript/GRAMMAR.md): Lexer specifications.
- [Bytecode](./pscript/BYTECODE.md): Virtual machine instructions.
- [VM Architecture](./pscript/VM_ARCHITECTURE.md): Call stack.
- [JIT Architecture](./pscript/JIT_ARCHITECTURE.md): JIT compilation.
- [LSP Architecture](./pscript/LSP_ARCHITECTURE.md): IDE integrations.

### Progression Curriculum
- [Tier 1 Deterministic](./progression/TIER_1_DETERMINISTIC.md)
- [Tier 2 Evidence](./progression/TIER_2_EVIDENCE.md)
- [Tier 3 FSM](./progression/TIER_3_FSM.md)
- [Tier 4 Thermodynamics](./progression/TIER_4_THERMODYNAMICS.md)
- [Tier 5 Byzantine](./progression/TIER_5_BYZANTINE.md)

### Game Theory Principles
- [Flow Theory](./game_theory/FLOW_THEORY.md): Balance challenge.
- [Operant Conditioning](./game_theory/OPERANT_CONDITIONING.md): Reinforcement.
- [Self Determination](./game_theory/SELF_DETERMINATION.md): Player agency.
- [Loss Aversion](./game_theory/LOSS_AVERSION.md): Risk management.
- [Social Comparison](./game_theory/SOCIAL_COMPARISON.md): Multiplayer score comparisons.
