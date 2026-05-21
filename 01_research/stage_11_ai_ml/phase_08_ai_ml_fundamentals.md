# Phase 8: AI/ML Fundamentals

## Objective

Build ML foundations from mathematics through deployment, then apply them to security problems such as anomaly detection, log classification, network detection, and malware classification.

## Mathematical Foundations

Required topics:

- Linear algebra: vectors, matrices, multiplication, transpose, inverse, determinant, eigenvalues, eigenvectors, SVD, norms, rank, orthogonality.
- Calculus and optimization: derivatives, partial derivatives, gradients, chain rule, backpropagation, Taylor approximation, gradient descent, SGD, Adam, RMSprop, convexity, constrained optimization.
- Probability and statistics: probability axioms, conditional probability, Bayes, distributions, expectation, variance, CLT, hypothesis testing, confidence intervals.
- Information theory: entropy, KL divergence, mutual information.

Important formulas:

| Concept | Formula |
|---|---|
| Gradient descent | `theta := theta - alpha * grad(J(theta))` |
| Cross entropy | `L = -sum(y_i * log(yhat_i))` |
| Softmax | `softmax(z_i) = exp(z_i) / sum(exp(z_j))` |
| Entropy | `H(X) = -sum(p(x) * log(p(x)))` |
| KL divergence | `D_KL(P || Q) = sum(p(x) * log(p(x) / q(x)))` |
| Attention | `softmax(QK^T / sqrt(d_k))V` |

## Data Handling and Feature Engineering

Required skills:

- Load CSV, JSON, and database-derived data.
- Clean missing values and outliers.
- Normalize and standardize features.
- Encode categorical features.
- Select features using correlation, mutual information, and model importance.
- Reduce dimensions with PCA, t-SNE, or UMAP.
- Handle imbalance with class weights, oversampling, undersampling, or SMOTE.
- Visualize distributions and relationships.

Security feature examples:

- Source and destination IP features.
- Port, protocol, byte, packet, duration, and timing features.
- Command-line tokens.
- Process ancestry.
- Authentication failure counts.
- PE imports, section entropy, strings, and metadata.

## Classical Machine Learning

Required supervised algorithms:

- Logistic regression.
- Decision trees.
- Random forests.
- Gradient boosting.
- SVM.
- Naive Bayes.
- KNN.

Required regression algorithms:

- Linear regression.
- Ridge and LASSO.
- Elastic Net.

Required unsupervised algorithms:

- k-means.
- Hierarchical clustering.
- DBSCAN.
- Gaussian mixture models.
- PCA.
- Autoencoders for nonlinear reduction.

Required anomaly detection:

- z-score and IQR.
- Isolation Forest.
- Local Outlier Factor.
- One-class SVM.
- Autoencoder-based detection.

Required metrics:

- Precision.
- Recall.
- F1.
- ROC-AUC.
- Confusion matrix.
- MSE, RMSE, MAE, R2 for regression.
- Silhouette score for clustering.

## Deep Learning

Required PyTorch knowledge:

- Tensor operations.
- Dataset and DataLoader.
- Network modules.
- Forward pass.
- Loss computation.
- Backpropagation.
- Optimizers and schedulers.
- Validation loops.
- Checkpointing.
- GPU and CPU execution differences.

Architectures:

- MLP.
- CNN.
- RNN, LSTM, GRU.
- Attention.
- Transformers.
- Autoencoders and VAEs.

Training constraints:

- Track train/validation split.
- Avoid leakage between train and test sets.
- Record seeds and deterministic settings where possible.
- Monitor overfitting.
- Store model card with intended use, data, metrics, and known failure modes.

## Embeddings and RAG

Required topics:

- Bag of words.
- TF-IDF.
- Word2Vec, GloVe, FastText.
- Sentence and document embeddings.
- CNN image embeddings.
- Contrastive learning.
- CLIP-style multimodal embeddings.
- FAISS vector indexing.
- Retrieval metrics: MRR and NDCG.
- RAG pipeline: chunking, embedding, indexing, retrieval, generation, evaluation.

Constraint:

- Retrieval quality must be measured independently from generated answer quality.

## Security Applications

Required projects:

- Security log classifier.
- Network traffic anomaly detector.
- Static malware classifier using PE or ELF metadata.
- End-to-end anomaly detection system.

Security ML failure modes:

- Data leakage.
- Label noise.
- Imbalanced classes.
- Concept drift.
- Adversarial examples.
- Evasion through mimicry.
- Overfitting to environment-specific artifacts.
- Poor calibration.

## Optimization and Deployment

Required topics:

- ONNX export.
- ONNX Runtime inference.
- INT8 quantization.
- Calibration.
- Accuracy loss measurement.
- Latency and memory benchmarking.
- Batch versus online inference.
- REST API serving.
- Model versioning.
- Drift detection.
- Canary deployment.

Deployment constraints:

- Define latency SLO.
- Define memory budget.
- Define input validation.
- Define model rollback process.
- Log predictions without leaking sensitive data.

## Promotion Targets

Research outcomes promote to:

- `06_ai/features/` for feature extractors.
- `06_ai/classifiers/` for trained classifier implementation.
- `06_ai/anomaly_detection/` for anomaly models.
- `06_ai/embeddings/` for embedding and retrieval systems.
- `06_ai/deployment/` for ONNX/runtime deployment assets.
- `04_datasets/` for dataset metadata only unless payload storage is explicitly approved.
