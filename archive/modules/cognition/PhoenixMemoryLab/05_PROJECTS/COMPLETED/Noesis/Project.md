# Project: Noesis

## One-Liner
Noesis (Axiom Engine)

## Status
COMPLETED

## Repo
`~/engineering/workspace/archived/Noesis`

## Ports
- API: N/A
- DB: N/A

## Database
N/A

## Run Command
N/A - historical project overview

## Dependencies On Other Projects
None

## What I Deliver To Others
None

## Links
- [[03_CORE_KNOWLEDGE/ai-ml/AI]]
- [[04_ENGINEERING/architecture-patterns/Software-Engineering]]
- [[04_ENGINEERING/system-design/System Design]]
- [[03_CORE_KNOWLEDGE/ai-ml/Machine Learning]]
- [[Decisions]]
- [[Mistakes]]

## Current Blockers
None

## Last Worked On
2026-05-12

## Original Overview


**Repository:** [github.com/fallofpheonix/Noesis](https://github.com/fallofpheonix/Noesis)  
**Language:** Python | **Created:** 2026-04-17

---

## Project Summary

Axiom Engine is a modular autonomous research system combining Dreamer-style RSSM world models, latent actor-critic agents, multi-agent coordination, symbolic communication, knowledge graph-based scientific loops, and meta-learning primitives.

## Architecture

| Module | Components |
|---|---|
| **core/** | RSSM world model, Dreamer wrapper, imagination rollout, latent actor/critic, multi-agent coordinator, symbolic communication, soft rule engine |
| **science/** | Knowledge graph, hypothesis generation, experiment planning, result analysis |
| **infrastructure/** | Autonomous research loop, task workers, bounded queue, JSONL persistence |
| **optimization/** | Meta-learning genome, evolution strategy, RSI self-improvement loop |
| **src/** | Blender scene compiler vertical slice |

## Implemented Features

- RSSM world model with prior/posterior latent dynamics
- Dreamer-style imagination rollout
- Actor-critic agent in latent space
- PettingZoo-style multi-agent coordinator
- Symbolic speaker/listener module
- Soft rule engine
- Knowledge graph add/query
- Hypothesis → experiment → analyze → graph update loop
- Bounded task queue and worker abstraction
- JSONL persistence
- Meta-learning genome and RSI validation loop
- 15 passing tests

## Skills Demonstrated

`Python`, `Reinforcement Learning`, `World Models (RSSM/Dreamer)`, `Actor-Critic`, `Multi-Agent Systems`, `Knowledge Graphs`, `Symbolic AI`, `Meta-Learning`, `Autonomous Research`, `Scientific Computing`
