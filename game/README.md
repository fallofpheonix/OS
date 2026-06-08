---\nStatus: In Development\nImplementation: 5%\nConfidence: Alpha\n---\n# Phoenix Game Layer

> **Layer**: Game | **Maturity**: Alpha | **Owner**: Gameplay Team

This directory contains the Phoenix Game Layer, featuring a deterministic game loop and the `pscript` virtual machine.

## Current Status

- **pscript VM**: Initial implementation of a stack-based VM for agent logic.
- **pscript Parser**: Minimal lexer and recursive-descent parser for programmable scripts.
- **Determinism**: Integrated with `foundation/ledger` for bit-perfect fixed-point simulation.
- **WebSocket Bridge**: Real-time world state broadcasting to Godot clients.
- **Vertical Slice**: "Deterministic Proximity Probe" implemented, verified, and renderable.

## Subprojects

- `game/pkg/vm`: Core pscript virtual machine and world state management.
- `game/docs`: Architectural specifications and research.

---
*Part of the [Phoenix Master Architecture](../docs/MASTER_SYSTEM_MAP.md)*
