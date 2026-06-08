# Project: cognitron-game

## One-Liner
cognitron-game

## Status
COMPLETED

## Repo
`~/engineering/workspace/archived/cognitron-game`

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
- [[04_ENGINEERING/architecture-patterns/Frontend Architecture]]
- [[Decisions]]
- [[Mistakes]]

## Current Blockers
None

## Last Worked On
2026-05-12

## Original Overview


**Repository:** [github.com/fallofpheonix/cognitron-game](https://github.com/fallofpheonix/cognitron-game)  
**Language:** TypeScript | **License:** MIT | **Created:** 2025-09-09

---

## Project Summary

AI-powered Reddit game built for Reddit Hackathon 2025. Contains two systems: (1) **Cognitron** — a cognitive bias puzzle game on Reddit's Devvit platform, and (2) **GameTrend Intelligence Engine** — a 10-module automated system for discovering promising indie game ideas via community mining and NLP.

## GameTrend Intelligence Engine — 10-Module Architecture

| Module | Purpose |
|---|---|
| Module 1 — Data Collection | Reddit, Steam, YouTube, TikTok scrapers |
| Module 2 — Data Storage | Type-safe in-memory database (Map collections) |
| Module 3 — Text Processing | URL removal, sentiment scoring, mechanic detection |
| Module 4 — Idea Clustering | TF-IDF + k-means++ clustering |
| Module 5 — Trend Detection | Cross-platform trend scoring |
| Module 6 — Opportunity Gap | Community demand vs Steam catalog comparison |
| Module 7 — Concept Generator | LLM-powered game concept generation |
| Module 8 — Prototype Generator | MVP spec + Phaser 3 / Godot 4 starter code |
| Module 9 — Dashboard | React + Recharts visualization |
| Module 10 — Scheduler | Continuous pipeline execution |

## Cognitron Game Design

| Field | Value |
|---|---|
| Genre | Cognitive Roguelite |
| Platform | Reddit Devvit + Phaser 3 |
| Session length | 10–15 minutes |
| Originality | 92/100 |
| Virality | 81/100 |
| Replayability | 88/100 |

### Core Gameplay Loop
1. Belief Network (directed graph 8–20 nodes) generates procedurally
2. Each node contains a Cognitive Bias (Confirmation, Anchoring, Bandwagon, etc.)
3. Player plays Evidence Cards to neutralize biases
4. Match → node cleared (+100 pts). Mismatch → Chain Reaction via BFS
5. Corruption > 60% → run ends. Clear all → win (+500 pts)

## Skills Demonstrated

`TypeScript`, `Game Development`, `NLP`, `Sentiment Analysis`, `TF-IDF`, `k-means++`, `Data Mining`, `OpenAI API`, `Reddit Devvit`, `Phaser 3`, `Real-time Data Processing`, `Pipeline Architecture`, `Market Analysis`

## Tech Stack

- **Runtime:** TypeScript (Node.js)
- **Game Platform:** Reddit Devvit + Phaser 3
- **ML/NLP:** TF-IDF, k-means++, lexicon-based sentiment
- **LLM:** OpenAI API (with mock fallback)
- **Testing:** 81 integration tests
