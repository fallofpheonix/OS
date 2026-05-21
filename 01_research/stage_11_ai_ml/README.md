# Stage 11: AI/ML Fundamentals

## Purpose

Build the mathematical, classical ML, deep learning, embedding, optimization, and deployment foundations needed for security AI systems.

## Scope

- Linear algebra, probability, statistics, calculus, and optimization.
- Classical supervised and unsupervised ML.
- Feature engineering and exploratory data analysis.
- Deep learning with PyTorch.
- Embeddings, semantic search, and RAG.
- Anomaly detection and security log classification.
- Malware classification using static features.
- ONNX export, quantization, model serving, and monitoring.

## Classification

- Type: `AI_ML_FOUNDATION`
- Status: `RESEARCH_ONLY`
- Difficulty: advanced-expert
- Estimated duration: 10-12 weeks
- Upstream prerequisites:
  - Stage 00 Foundations
  - Stage 04 Security
  - Stage 09 Telemetry
- Downstream blockers:
  - Stage 12 Security AI
  - Stage 14 Automation
  - Stage 19 Production

## Research Modules

| Module | Path |
|---|---|
| Phase 8 Research Plan | `phase_08_ai_ml_fundamentals.md` |
| Phase 8 Build Gate | `build_gate.md` |

## Internal Dependency Order

```text
Math foundations
-> Data handling and EDA
-> Classical ML
-> Deep learning
-> Embeddings and retrieval
-> Security ML applications
-> Optimization and deployment
```

## Gate

Do not promote models into `06_ai/` until data provenance, metrics, failure modes, and deployment constraints are documented.
