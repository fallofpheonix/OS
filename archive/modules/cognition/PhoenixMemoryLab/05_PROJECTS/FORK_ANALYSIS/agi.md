# agi

## Purpose
Study agent control-plane, memory, tool use, and failure semantics.

## Core Architecture
- Multi-agent control-plane, memory stores, tool interfaces, decision loops.

## Interesting Systems
- Memory models, tool routing, failure modes, planning vs reactive behaviors.

## Reusable Ideas
- Agent topology patterns, memory lifecycles, execution tracing hooks, retry semantics.

## Reusable Components
- Memory interfaces, tool-adapter patterns, trace instrumentation scaffolding.

## Rendering / Runtime Concepts
- (Not rendering-focused) runtime control-plane and coordination gravity signals — compare with your locality/containment principles.

## Patterns Worth Extracting
- Orchestration metrics, coordination-gravity observability hooks, bounded-agent design patterns.

## Weaknesses
- High coordination complexity; likely heavy coupling to original control-plane assumptions.

## Integration Opportunities
- Import observational patterns into `04_ENGINEERING/control-plane-patterns/` and write comparative notes on coordination gravity.

## Delete / Archive Decision
- Classification: Strategic Analysis Source → keep as knowledge artifact only; extract patterns and convert into brain notes; avoid keeping as long-term fork unless actively evolving.
