# Threat Detection Stack Labs

## Scope

Build a modern SOC-style detection pipeline from raw telemetry to alerting, triage, hunting, and response orchestration.

## 1. Detection Engines

### Components

| Component | Role |
|---|---|
| IDS | Detect suspicious network events |
| IPS | Block known malicious network behavior |
| EDR | Collect and analyze endpoint behavior |
| XDR | Correlate endpoint, network, cloud, and identity events |

### Tools

- Suricata.
- Zeek.
- Wazuh.
- Security Onion.
- Osquery.
- Auditd.

### Exercise

Deploy:

- Wazuh manager.
- Wazuh Linux agent.
- Suricata sensor.
- Zeek sensor.

Generate controlled test events:

- Process spawn.
- Suspicious outbound DNS lookup.
- Suricata IDS alert.
- Endpoint file write.

Correlate:

```text
process event
  -> outbound connection
  -> IDS alert
  -> host risk increase
```

### Exit Criteria

- Endpoint and network events share host identity.
- At least one Suricata alert is ingested.
- At least one Wazuh alert is ingested.
- Correlation logic is documented.

## 2. SIEM And SOAR

### Components

| Component | Role |
|---|---|
| SIEM | Normalize, index, search, and alert on logs |
| SOAR | Orchestrate enrichment and response workflows |
| TheHive | Case management and analyst workflow |
| Cortex-style analyzers | Enrichment and analysis tasks |

### Candidate Stack

- OpenSearch or Elasticsearch.
- Wazuh.
- TheHive.
- Shuffle or n8n.
- Filebeat or Fluent Bit.

### Exercise

Build:

```text
Wazuh / Suricata / Zeek
  -> Filebeat or Fluent Bit
  -> OpenSearch
  -> alert rule
  -> TheHive case
  -> manual triage
```

### Exit Criteria

- Alerts are searchable in SIEM.
- Alert creates or updates a case.
- Case has observable fields.
- Analyst workflow is documented.

## 3. Threat-Centric Modeling

### Topics

- MITRE ATT&CK.
- IOC.
- Kill chain.
- Attack graph.
- Behavior analytics.
- Risk scoring.
- Event correlation.

### Exercise

Tag detection events with:

- ATT&CK tactic.
- ATT&CK technique.
- Asset ID.
- User ID.
- Confidence.
- Severity.

Build a minimal attack graph:

```text
user
  -> host
  -> process
  -> connection
  -> domain
  -> alert
```

Optional storage:

- Neo4j.
- SQLite edge table.
- JSON graph file.

### Exit Criteria

- At least five event types map to ATT&CK.
- Graph has nodes and edges.
- Risk score is reproducible.
- False-positive handling is documented.

## 4. Telemetry Pipelines

### Pipeline Pattern

```text
pcap / tap
  -> Suricata / Zeek
  -> logs
  -> forwarder
  -> OpenSearch
  -> enrichment
  -> SIEM alert
  -> SOAR case
```

Endpoint pattern:

```text
auditd / osquery / Wazuh agent
  -> manager
  -> indexer
  -> correlation
  -> hunt query
```

### Exercise

Implement normalized event schema:

```text
timestamp
source_tool
event_type
host_id
user_id
process_id
src_ip
dst_ip
domain
mitre_technique
severity
confidence
raw_event_ref
```

### Exit Criteria

- All tools map into the schema.
- Raw event reference is retained.
- Pipeline handles missing fields.
- Timestamps are normalized.

## 5. Threat Hunting

### Hunting Questions

- Which hosts have failed logins and suspicious outbound DNS?
- Which users spawned shells from office applications?
- Which hosts show SMB lateral movement after privilege escalation?
- Which domains have beacon-like periodicity?
- Which endpoints show abnormal process ancestry?

### Query Artifacts

Create:

- DNS tunneling hunt.
- Lateral movement hunt.
- Suspicious process tree hunt.
- Rare destination hunt.
- Privilege escalation hunt.

### Exit Criteria

- Each hunt has data requirements.
- Each hunt has expected false positives.
- Each hunt maps to ATT&CK.
- Each hunt has escalation criteria.

## Suggested Repo Structure

```text
threat_detection/
├── 01_idps/
│   ├── README.md
│   ├── suricata_rules/
│   ├── zeek_scripts/
│   └── pcaps/
├── 02_edr_xdr/
│   ├── README.md
│   ├── wazuh/
│   ├── osquery/
│   └── auditd/
├── 03_siem_soar/
│   ├── README.md
│   ├── opensearch/
│   ├── thehive/
│   ├── shuffle/
│   └── pipelines/
├── 04_threat_hunting/
│   ├── README.md
│   ├── mitre_mapping/
│   ├── queries/
│   └── risk_scoring/
└── 05_attack_graph/
    ├── README.md
    ├── schema/
    ├── examples/
    └── reports/
```

## 12-Week Lab Plan

| Week | Focus | Output |
|---:|---|---|
| 1 | Telemetry basics | event schema draft |
| 2 | Suricata | custom rule and alert ingestion |
| 3 | Zeek | structured logs and parser |
| 4 | Wazuh agent | endpoint alert ingestion |
| 5 | OpenSearch | indexed searchable events |
| 6 | SIEM rules | detection rule set |
| 7 | TheHive | case creation workflow |
| 8 | Shuffle/n8n | response workflow |
| 9 | MITRE mapping | technique-tagged alerts |
| 10 | Hunting queries | DNS + SMB correlation hunts |
| 11 | Attack graph | graph schema and example |
| 12 | Capstone | full incident timeline from telemetry to case |

## Capstone

Goal:

```text
Controlled incident replay
  -> network alert
  -> endpoint alert
  -> SIEM correlation
  -> TheHive case
  -> ATT&CK mapping
  -> attack graph
  -> response recommendation
```

Deliverables:

- PCAP.
- Endpoint logs.
- Detection rules.
- SIEM query.
- Case export.
- Attack graph.
- Incident report.

## Integration With Cyber AI OS

| Threat Detection Output | Later Use |
|---|---|
| Normalized event schema | AI feature store |
| IDS alerts | Network threat engine |
| EDR events | Host behavior model |
| SIEM searches | SOC assistant retrieval |
| SOAR playbooks | Response manager |
| Attack graph | Threat prediction and risk scoring |

