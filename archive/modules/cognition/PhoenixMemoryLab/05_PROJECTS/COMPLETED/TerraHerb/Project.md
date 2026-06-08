# Project: TerraHerb

## One-Liner
TerraHerb

## Status
COMPLETED

## Repo
`~/engineering/workspace/archived/TerraHerb`

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


**Repository:** [github.com/fallofpheonix/TerraHerb](https://github.com/fallofpheonix/TerraHerb)  
**Language:** Python | **License:** MIT | **Forks:** 1 | **Created:** 2026-02-13

---

## Project Summary

Deep learning system for plant species identification and botanical knowledge retrieval. Python-first computer vision system with both PyTorch and TensorFlow models for plant species/disease classification, plus local and remote knowledge enrichment.

## Performance Benchmarks

| Metric | MobileNetV2 (PyTorch) | EfficientNetB0 (TensorFlow) |
|---|---:|---:|
| Top-1 Accuracy | 92.8% | 97.8% |
| Top-5 Accuracy | 98.5% | 99.2% |
| Inference Latency | ~120ms | ~155ms |
| Dataset | PlantVillage (38 classes, ~54K images) | PlantVillage |

## Architecture

```
React Web UI → FastAPI Gateway → PlantPredictor → MobileNetV2 Classifier
                                → KnowledgeRetriever → UCI / GBIF / Wikipedia APIs
```

## Tech Stack

- **ML:** PyTorch (MobileNetV2), TensorFlow (EfficientNetB0)
- **API:** FastAPI + Uvicorn
- **Frontend:** Vite + React
- **Dataset:** PlantVillage via KaggleHub
- **Knowledge:** UCI Plants, GBIF API, Wikipedia API

## Skills Demonstrated

`Python`, `PyTorch`, `TensorFlow`, `Computer Vision`, `Transfer Learning`, `FastAPI`, `React`, `Image Classification`, `Knowledge Retrieval`, `Multi-Model Inference`
