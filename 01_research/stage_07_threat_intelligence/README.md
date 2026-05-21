# Stage_07_Threat_Intelligence: Threat Intelligence

- Classification: IMPLEMENTATION
- Progression: Integrate -> Secure
- Scope: IOC/rule/feed management, MITRE mapping, enrichment
- Prerequisites: Stage_03_Networking, Stage_04_Security_Fundamentals, Stage_05_Malware_and_RE
- Next stage: Stage_08_Observability

## Source Files

| Item | Current location | Status | Type |
| --- | --- | --- | --- |
| Threat Intelligence | docs/17-threat-intelligence.md | PLANNED | Documentation |
| .gitkeep | security/honeypots/.gitkeep | IMPLEMENTATION | Coding placeholder |
| .gitkeep | security/threatintel/.gitkeep | IMPLEMENTATION | Coding placeholder |

## Topics

| Topic | Difficulty | Hours | Weeks | Kind | Risk | Priority | Blocking items |
| --- | --- | --- | --- | --- | --- | --- | --- |
| IOC storage | High | 60 | 6 | Coding/Integration | Medium | P1 | Stage gate incomplete |
| YARA/Sigma/Suricata repositories | High | 60 | 6 | Coding/Integration | Medium | P1 | Stage gate incomplete |
| Feed ingestion | High | 60 | 6 | Coding/Integration | Medium | P1 | Stage gate incomplete |
| MISP/OpenCTI | High | 60 | 6 | Coding/Integration | Medium | P1 | Stage gate incomplete |
| MITRE ATT&CK mapping | High | 60 | 6 | Coding/Integration | Medium | P1 | Stage gate incomplete |
| Enrichment rules | High | 60 | 6 | Coding/Integration | Medium | P1 | Stage gate incomplete |
| Threat hunting context | High | 60 | 6 | Coding/Integration | Medium | P1 | Stage gate incomplete |

## Gate Conditions

- All prerequisite stages complete: Stage_03_Networking, Stage_04_Security_Fundamentals, Stage_05_Malware_and_RE

- Source documents reviewed and contradictions logged

- Labs produce reproducible artifacts

- Risk items have owner and mitigation

- Outputs indexed in root reports
