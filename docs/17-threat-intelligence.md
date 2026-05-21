# Threat Intelligence

## Feeds

- IOC database.
- CVE feeds.
- MITRE ATT&CK.
- YARA rules.
- Sigma rules.
- Vendor reports.
- Internal incident reports.

## Storage

| Store | Purpose |
|---|---|
| Graph DB | Actor, campaign, technique, infrastructure relationships |
| Threat cache | Fast local IOC lookup |
| Local indicators | Site-specific detections |
| Rule repository | YARA, Sigma, Suricata, enrichment rules |

## Functions

- IOC matching.
- Campaign correlation.
- Actor analysis.
- Risk scoring.
- Rule enrichment.
- Alert deduplication.
- Attack-chain mapping.

## Reference Frameworks

- MITRE ATT&CK.
- YARA.
- Sigma.
- CVE.
- CWE.

## Data Flow

```text
Feed import
  -> Normalization
  -> Deduplication
  -> Confidence scoring
  -> Graph update
  -> Detection-rule mapping
  -> Local cache
```

## Failure Considerations

- Expired indicators must age out.
- Low-confidence indicators must not trigger blocking alone.
- Feed poisoning must be considered.
- External feed outages must not disable local detections.

