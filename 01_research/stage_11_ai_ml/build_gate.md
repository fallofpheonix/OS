# Phase 8 Build Gate: AI/ML Fundamentals

## Mathematical Understanding

- [ ] Perform matrix multiplication, transpose, inverse, and norm computations.
- [ ] Compute eigenvalues and eigenvectors.
- [ ] Apply SVD for dimensionality reduction.
- [ ] Explain matrix norms geometrically.
- [ ] Use covariance and Gram matrices in ML context.
- [ ] Compute conditional probabilities.
- [ ] Apply Bayes theorem.
- [ ] Conduct hypothesis testing.
- [ ] Explain confidence intervals and p-values.
- [ ] Compute gradients and partial derivatives.
- [ ] Explain gradient descent variants.
- [ ] Optimize a constrained function.

## Data Handling and EDA

- [ ] Load CSV, JSON, or database data.
- [ ] Clean missing values.
- [ ] Handle outliers.
- [ ] Normalize or standardize features.
- [ ] Visualize distributions.
- [ ] Identify anomalies.
- [ ] Encode categorical data.
- [ ] Select relevant features.
- [ ] Explain feature importance.

## Classical ML

- [ ] Train 2-3 classification algorithms.
- [ ] Achieve at least 85 percent accuracy on a suitable UCI or equivalent dataset.
- [ ] Generate precision, recall, F1, ROC-AUC, and confusion matrix.
- [ ] Implement or use isolation forest or LOF.
- [ ] Train autoencoder anomaly detector.
- [ ] Evaluate anomalies with precision, recall, and threshold trade-offs.
- [ ] Optionally implement clustering and report silhouette score.

## Deep Learning

- [ ] Build and train a PyTorch neural network.
- [ ] Implement training loop.
- [ ] Use optimizer and learning rate scheduling.
- [ ] Validate on test set.
- [ ] Plot loss and accuracy curves.
- [ ] Fine-tune a pretrained model.
- [ ] Compare against baseline.
- [ ] Export model to ONNX.

## Embeddings and RAG

- [ ] Create document embeddings.
- [ ] Build FAISS index.
- [ ] Implement semantic search.
- [ ] Evaluate retrieval quality with MRR or NDCG.
- [ ] Implement simple RAG system.
- [ ] Evaluate answer quality separately from retrieval quality.

## Security Applications

- [ ] Parse syslog, Windows event logs, or equivalent security logs.
- [ ] Extract features.
- [ ] Train malicious-log classifier.
- [ ] Achieve at least 90 percent accuracy or document why the dataset makes this invalid.
- [ ] Extract network flow features.
- [ ] Train network anomaly detector.
- [ ] Evaluate precision and recall.
- [ ] Extract static malware features.
- [ ] Train malware-family classifier.
- [ ] Achieve at least 85 percent accuracy or document dataset constraints.

## Model Optimization and Deployment

- [ ] Quantize model to INT8 or equivalent.
- [ ] Measure latency.
- [ ] Measure memory.
- [ ] Keep accuracy loss below 2 percent or justify trade-off.
- [ ] Export to ONNX.
- [ ] Benchmark on target hardware.
- [ ] Containerize inference service.
- [ ] Create REST API.
- [ ] Support batch and real-time inference.
- [ ] Implement model versioning.
- [ ] Monitor latency and drift.

## End-to-End Project

- [ ] Select KDD Cup 99, NSL-KDD, or custom log dataset.
- [ ] Document data provenance.
- [ ] Preprocess and split data.
- [ ] Train autoencoder or isolation forest.
- [ ] Evaluate precision, recall, F1, and ROC-AUC.
- [ ] Export model.
- [ ] Deploy inference API.
- [ ] Document architecture, usage, and performance.

## Exit Criteria

- [ ] Model card written for each promoted model.
- [ ] Dataset metadata written.
- [ ] Metrics reproducible.
- [ ] Failure modes documented.
- [ ] Deployment constraints documented.
- [ ] Artifacts promoted to appropriate `06_ai/` paths.
