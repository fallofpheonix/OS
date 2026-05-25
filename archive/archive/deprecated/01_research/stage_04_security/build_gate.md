# Phase 6 Build Gate: Security Foundations

## Threat Modeling

- [ ] Create one STRIDE or PASTA threat model.
- [ ] Define system boundary and trust zones.
- [ ] Identify assets and data flows.
- [ ] Enumerate at least 15 threats.
- [ ] Score likelihood and impact.
- [ ] Propose mitigations for each threat.
- [ ] Document assumptions and dependencies.

Minimum example systems:

- [ ] Web login system with user data and payment path.
- [ ] Microservice architecture with service-to-service communication.
- [ ] Mobile application with local storage and backend API.
- [ ] Cloud infrastructure with VPC, storage, and IAM.

## Attack Chain Analysis

- [ ] Select a real attack mapped to MITRE ATT&CK.
- [ ] Map reconnaissance through exfiltration or impact.
- [ ] Identify techniques and likely tools.
- [ ] Identify detection opportunities at each stage.
- [ ] Recommend preventive and detective controls.
- [ ] Document IOCs and signatures.
- [ ] Map each step to cyber kill chain stages.
- [ ] Identify where defenses break the chain.

## Detection and Alerting

- [ ] Build at least one SIEM or IDS detection rule.
- [ ] Define suspicious pattern or behavior.
- [ ] Include false positive mitigation.
- [ ] Test against benign and malicious data.
- [ ] Document rule logic and IOCs.
- [ ] Establish baseline behavior.
- [ ] Tune thresholds.
- [ ] Create correlation rule.
- [ ] Document escalation procedure.

Rule candidates:

- Brute force login detection.
- Unusual RDP or SSH lateral movement.
- Large outbound DNS or HTTP transfer.
- Privilege escalation through UID/GID change in process tree.

## Security Design

- [ ] Design a secure system using CIA and least privilege.
- [ ] Define authentication and authorization model.
- [ ] Design RBAC or ABAC enforcement.
- [ ] Plan encryption at rest and in transit.
- [ ] Design logging and monitoring.
- [ ] Document security architecture.
- [ ] Design zero trust microsegmentation.
- [ ] Define continuous verification checkpoints.

## Incident Response

- [ ] Define incident classes and severity levels.
- [ ] Create escalation matrix.
- [ ] Write runbooks for common incidents.
- [ ] Define communication strategy.
- [ ] Define RTO and RPO targets.
- [ ] Conduct tabletop exercise.
- [ ] Document gaps and lessons learned.

## Cryptography and Authentication

- [ ] Implement password hashing using bcrypt or Argon2.
- [ ] Implement TOTP or hardware-token MFA.
- [ ] Implement OAuth 2.0 authorization code flow.
- [ ] Implement JWT validation and refresh-token handling.
- [ ] Encrypt data at rest with AES-256 or equivalent authenticated encryption.
- [ ] Enforce TLS 1.3 for data in transit.
- [ ] Demonstrate key derivation.
- [ ] Demonstrate key rotation.

## Access Control

- [ ] Design roles and permissions.
- [ ] Implement role assignment and enforcement.
- [ ] Audit access decisions.
- [ ] Test privilege escalation mitigations.

## Exit Criteria

- [ ] Threat model reviewed and versioned under `02_docs/threat_models/`.
- [ ] At least one tested detection rule promoted to `07_security/detections/`.
- [ ] Incident response runbook promoted to `07_security/response/`.
- [ ] Access control design promoted to `02_docs/architecture/`.
