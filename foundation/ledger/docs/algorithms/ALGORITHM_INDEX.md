---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Ledger — Algorithm Index

> Last verified: 2026-06-04

## 1. Reality Confidence Quorum
- **Core logic**: Collects sensor observations. Groups identical values. Computes the ratio:
  $$\text{ConfidenceScore} = \frac{\sum \text{Weights of agreeing majority}}{\sum \text{All sensor weights}}$$
- **Invariant**: Valid facts require $M$ agreeing sensors and confidence $\ge C_{min}$.

## 2. Merkle DAG Hashing (Schema V2)
- **Hash calculation**:
  $$\text{Hash} = \text{SHA256}(\text{SchemaVersion} \parallel \text{Tick} \parallel \text{EventID} \parallel \text{CauseID} \parallel \text{Heads} \parallel \text{Payload} \parallel \text{TraceHash} \parallel \text{States})$$

## 3. State Gap Verification
- **Verification Rule**:
  $$\forall E \in \text{Ledger}, \quad \forall P \in \text{Parents}(E), \quad E.\text{StateBefore} = P.\text{StateAfter}$$
