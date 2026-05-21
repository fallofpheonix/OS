# Phase 6: Security Foundations

## Objective

Understand core security principles, threat modeling, defensive design, authentication, authorization, access control, logging, detection, and incident response.

## Security Fundamentals

| Area | Required Knowledge |
|---|---|
| CIA and extensions | Confidentiality, integrity, authenticity, non-repudiation, availability |
| Design principles | Defense in depth, least privilege, fail secure, secure defaults |
| Security mindset | Threat modeling, attack surface analysis, risk assessment |
| Trade-offs | Security versus usability, operational complexity, incident response cost |

Required distinctions:

- Vulnerability: exploitable weakness.
- Threat: actor, capability, or event that can exploit weakness.
- Risk: expected loss from likelihood and impact.
- Control: preventive, detective, corrective, or compensating mitigation.

## Threat Modeling and Risk Analysis

Required methodologies:

- STRIDE: spoofing, tampering, repudiation, information disclosure, denial of service, elevation of privilege.
- PASTA: objective-driven attack simulation and threat analysis.
- Cyber kill chain: reconnaissance, weaponization, delivery, exploitation, installation, command and control, actions on objectives.

Risk model:

```text
risk = likelihood * impact
residual_risk = inherent_risk - control_effectiveness
```

Threat model output must include:

- System boundary.
- Trust zones.
- Assets and data flows.
- Threat list.
- Likelihood and impact matrix.
- Mitigations.
- Assumptions and unresolved dependencies.

## Authentication and Identity

Required topics:

- Knowledge-based authentication: passwords and passphrases.
- Possession-based authentication: hardware tokens, TOTP, SMS.
- Biometric and location-based authentication.
- MFA and adaptive authentication.
- Password hashing versus encryption.
- Salting and stretching: bcrypt, scrypt, PBKDF2, Argon2.
- Credential stuffing and dictionary attacks.
- OAuth 2.0 flows: authorization code, client credentials, implicit legacy risk.
- OIDC identity federation.
- JWT validation, expiration, refresh, revocation, and key rotation.

Security constraints:

- Passwords are never encrypted for later recovery.
- JWTs are verified for issuer, audience, signature, expiry, not-before, and algorithm allowlist.
- Refresh tokens require rotation or replay detection.
- MFA recovery flows are part of the authentication surface.

## Authorization and Access Control

Required models:

- DAC: owner-controlled permissions; flexible but error-prone.
- MAC: centrally enforced labels and policy; strong but operationally rigid.
- RBAC: roles group permissions; scalable with hierarchy and separation of duties.
- ABAC: policy uses subject, object, action, and environment attributes.
- ACLs: explicit object-level permissions.
- Capability security: authority represented by unforgeable token or handle.

Zero trust requirements:

- Never trust network location alone.
- Continuously authenticate and authorize.
- Enforce least privilege.
- Segment services and data paths.
- Assume breach.
- Log access decisions.

## Cryptography Basics

Required topics:

- Symmetric encryption: AES, ChaCha20, modes including CBC, CTR, GCM.
- Avoid ECB.
- Prefer authenticated encryption for new designs.
- Key derivation: PBKDF2, scrypt, Argon2, HKDF.
- Public key cryptography: RSA, ECC.
- Key exchange: Diffie-Hellman, ECDHE.
- Digital signatures: RSA-PSS, ECDSA, EdDSA.
- Certificates and X.509 trust chains.
- Hashes: MD5 broken, SHA-1 deprecated, SHA-2/SHA-3 acceptable.
- HMAC for integrity and authentication.
- TLS 1.3 as default transport security target.

Failure considerations:

- Key storage dominates algorithm choice in real systems.
- IV/nonce reuse can destroy confidentiality or integrity.
- Encryption without authentication is usually incomplete.
- Crypto agility must not permit downgrade to weak algorithms.

## Logging and Monitoring

Required security logs:

- Authentication attempts and MFA events.
- Authorization decisions.
- Data access and sensitive object reads.
- Configuration changes.
- Privilege changes.
- Security control changes.
- Detection events.

Log requirements:

- Centralized.
- Tamper-evident or append-only.
- Time-synchronized through NTP.
- Retained according to incident response and compliance requirements.
- Redacted for secrets and regulated data.

Analysis topics:

- Baselines and anomaly detection.
- Thresholds and false positive control.
- Event correlation and enrichment.
- Alert fatigue.
- Batch versus real-time detection.
- On-call escalation.

## Incident Response

Lifecycle:

```text
Preparation
-> Detection and analysis
-> Containment
-> Eradication
-> Recovery
-> Post-incident review
```

Plan requirements:

- Incident classes and severity levels.
- Escalation matrix.
- Communication protocol.
- Evidence preservation and chain of custody.
- RTO and RPO targets.
- Runbooks for common incidents.
- Tabletop exercise results.

## Detection and Response Frameworks

Required systems:

- IDS/IPS: signature and anomaly detection, passive versus active response, NIDS versus HIDS.
- EDR: endpoint process, file, registry, syscall, and response telemetry.
- XDR: correlation across endpoint, network, cloud, and identity.
- SIEM: centralized log collection, alerting, retention, compliance reporting.
- SOAR: playbook automation, enrichment, and response orchestration.

Required frameworks:

- MITRE ATT&CK tactics and techniques.
- Threat hunting using hypotheses and pivots.
- CTI vocabulary: observable, IOC, campaign, TLP.

## Promotion Targets

Research outcomes promote to:

- `02_docs/threat_models/` for threat models.
- `02_docs/architecture/` for secure architecture.
- `07_security/detections/` for detection rules.
- `07_security/response/` for incident response runbooks.
- `07_security/siem/` and `07_security/soar/` for SIEM/SOAR implementation artifacts.
