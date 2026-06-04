---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Events — Algorithm Index

> Last verified: 2026-06-04

## 1. Signature Verification
- **Purpose**: Assures the authenticity and integrity of events emitted by authorities.
- **Inputs**: Event signature, public key, and payload hash.
- **Complexity**: $O(1)$.

## 2. Artifact Hashing and Verification
- **Purpose**: Verifies that binary payloads match the logged hashes.
- **Inputs**: File byte buffer, SHA-256 validator.
- **Complexity**: $O(B)$ where $B$ is payload byte length.
