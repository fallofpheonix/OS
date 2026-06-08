---
Status: Planned
Implementation: 10%
Confidence: Conceptual
---
# PhoenixMind — Orchestration Model

Deals with the prompt layouts, token allocation, context packaging, and tool schemas.

## Context Allocation Map

```
┌────────────────────────────────────────────────────────┐
│ Context Window (1M+ Tokens)                           │
├───────────────┬───────────────────┬────────────────────┤
│ System Prompt │ Long-Term Memory  │ Interactive State  │
│ (10k tokens)  │ (500k tokens)     │ (500k tokens)      │
└───────────────┴───────────────────┴────────────────────┘
```

## Prompt Pipeline Rules
1. **System Instruction**: Enforce zero-trust rules, restricting write calls to signed actuators.
2. **Context Injection**: State Scribe embeds the top 20 most similar past memories via semantic distance metrics.
3. **Validation Loop**: Every generated command is parsed for raw bash/eval statements and blocked if security rules fail.
