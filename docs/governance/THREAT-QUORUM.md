---
Status: Draft
Implementation: 0%
Confidence: High
---
# THREAT-QUORUM: Sensor Diversity & Escalation Thresholds

> This document defines the requirements for transitioning system state to high-severity containment levels (SUSPICIOUS, CRITICAL, COMPROMISED).

## 1. The Diversity Principle
To prevent a single faulty or noisy sensor from triggering a catastrophic substrate lockdown, all high-severity state transitions MUST meet the **Diversity Threshold**.

### 1.1 Quorum Formula
A transition to state `S` is only valid if:
`ThreatScore(E) >= Threshold(S)` AND `UniqueSourceCount(E) >= Quorum(S)`

| Target State | Threat Threshold | Source Quorum | Allowed Escalation |
|--------------|------------------|---------------|-------------------|
| **WATCH**    | 0.20             | 1             | Any               |
| **SUSPICIOUS**| 0.50             | 2             | Any               |
| **CRITICAL** | 0.80             | 3             | Any               |
| **COMPROMISED**| 0.95            | 4*            | Fast-Up Only      |

*\*Special Case: A single "Authoritative Kernel Probe" (AKP) or a "Verified Invariant Breach" (VIB) from the Warden counts as a full Quorum of 4.*

## 2. Sensor Classification
Sensors are categorized by their observation domain to ensure cross-domain verification.

- **D1: Kernel/Process** (Syscall monitoring, PID isolation)
- **D2: Filesystem/I/O** (Checksum mismatch, unauthorized writes)
- **D3: Network/Distributed** (Consensus divergence, message flooding)
- **D4: Memory/Logic** (Invariant violation, heap corruption)

**Rule**: A Quorum of 2 MUST involve at least two different domains (e.g., D1 + D3).

## 3. Evidence Weighted Aggregation
The `EvidenceWeight` in an `AuthorityEscalationRequest` is the result of the weighted sum of diverse sensor inputs.

```text
EvidenceWeight = Σ (SensorConfidence * DomainWeight) / QuorumSize
```

- If `EvidenceWeight` < `Threshold(S)`, the Warden MUST reject the transition and downgrade the intent to `ClassLog`.

---
*Refer to [assurance/security/warden.go](../../assurance/security/warden.go) for implementation of EvidenceWeightInvariant.*
