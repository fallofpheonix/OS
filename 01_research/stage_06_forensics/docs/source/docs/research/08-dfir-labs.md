# Digital Forensics And Incident Response Labs

## Scope

DFIR pipeline from evidence acquisition through artifact extraction, memory analysis, timeline reconstruction, recovery, and incident reporting.

Use only authorized systems, lab images, or intentionally prepared evidence sets.

## 1. Acquisition And Integrity

### Topics

- Disk imaging.
- Raw `.dd` images.
- Expert Witness Format `.E01`.
- Volume imaging.
- Memory dumps.
- Hashing.
- SHA-256.
- Chain of custody.
- Evidence metadata.

### Tools

- `dd`.
- `ewfacquire`.
- `sha256sum`.
- WinPMEM.
- LiME.
- FTK Imager where applicable.

### Exercise

Acquire a small test disk:

1. Create or attach lab disk.
2. Acquire raw image.
3. Compute SHA-256 before and after transfer.
4. Record acquisition metadata.
5. Store chain-of-custody log.

### Chain Of Custody Schema

```json
{
  "case_id": "CASE-0001",
  "evidence_id": "DISK-0001",
  "acquired_by": "analyst",
  "acquired_at": "YYYY-MM-DDTHH:MM:SSZ",
  "source": "lab disk",
  "tool": "dd",
  "tool_version": "TBD",
  "image_path": "evidence/disk.dd",
  "sha256": "TBD",
  "notes": "TBD"
}
```

### Exit Criteria

- Evidence image exists.
- Hash is recorded.
- Tool version is recorded.
- Chain-of-custody record exists.

## 2. Artifact Extraction And Analysis

### Topics

- File timestamps.
- System logs.
- Prefetch files.
- Jump lists.
- MFT.
- Browser history.
- Cookies.
- Downloads.
- Bookmarks.
- Cache.
- Registry hives.
- `Run` and `RunOnce`.
- USB tracking.
- Installed software.
- NTUSER artifacts.

### Tools

- Autopsy.
- Sleuth Kit.
- RegRipper where applicable.
- Browser artifact parsers.
- Custom parsers.

### Exercise

Use Autopsy against a sample disk image.

Extract:

- Browser history.
- Download history.
- Persistence entries.
- Recently modified files.
- User activity indicators.

### Exit Criteria

- Artifact report exists.
- Suspicious persistence mechanisms are listed.
- Browser activity is summarized.
- Findings include file paths and timestamps.

## 3. Timeline Reconstruction

### Topics

- Timeline aggregation.
- File modification events.
- Registry changes.
- Process start events.
- Network login events.
- Authentication logs.
- Plaso.
- Timesketch.

### Tools

- Plaso/log2timeline.
- Timesketch.
- Autopsy timeline.
- SIEM exports.

### Exercise

Build incident timeline:

1. Run Plaso over disk image.
2. Import output into Timesketch.
3. Add RAM-derived events if available.
4. Mark pivots:
   - initial access
   - execution
   - persistence
   - lateral movement
   - data staging
   - exfiltration

### Exit Criteria

- Timeline is queryable.
- Pivot points are annotated.
- Timestamp source is preserved.
- Confidence is documented per major event.

## 4. RAM Analysis

### Topics

- Memory dumps.
- Running processes.
- Loaded drivers.
- Network connections.
- Injected code.
- Hollowed processes.
- Shellcode.
- In-memory config.

### Tools

- Volatility 3.
- WinPMEM.
- LiME.
- CAPE memory dumps.

### Exercise

Analyze a memory image:

1. Identify OS/profile requirements.
2. List processes.
3. List network connections.
4. List loaded drivers/modules.
5. Search for injected memory.
6. Correlate with disk and timeline artifacts.

### Exit Criteria

- Process list is exported.
- Network connection evidence is exported.
- Suspicious memory regions are documented.
- Memory findings link to timeline events.

## 5. Filesystem Recovery And Incident Response

### Topics

- Deleted-file recovery.
- NTFS MFT.
- NTFS journal.
- ext3/ext4 metadata.
- Overwritten artifact limits.
- Incident triage.
- Containment.
- Reporting.

### Tools

- Autopsy.
- Sleuth Kit.
- extundelete where applicable.
- Plaso.
- Timesketch.

### Exercise

Simulate a ransomware-style incident in a lab image:

1. Create files.
2. Delete or modify files.
3. Capture disk image.
4. Recover deleted artifacts where possible.
5. Build timeline.
6. Write IR report.

### Exit Criteria

- Recovered files are hashed.
- Unrecoverable artifacts are documented.
- Timeline explains file deletion/modification.
- IR report answers who, what, when, where, how.

## 6. Incident Report Templates

### Initial Triage

```text
case_id:
reported_at:
reported_by:
affected_assets:
initial_observations:
evidence_collected:
immediate_risk:
next_actions:
```

### Interim Findings

```text
case_id:
scope:
timeline_summary:
confirmed_findings:
unconfirmed_findings:
containment_status:
open_questions:
```

### Final Report

```text
case_id:
executive_summary:
timeline:
root_cause:
affected_assets:
evidence:
impact:
containment:
eradication:
recovery:
lessons_learned:
recommendations:
appendices:
```

## Suggested Repo Structure

```text
dfir/
├── 01_acquisition/
│   ├── README.md
│   ├── scripts/
│   ├── chain_of_custody/
│   └── hash_logs/
├── 02_artifacts/
│   ├── README.md
│   ├── autopsy_cases/
│   ├── browser/
│   ├── registry/
│   └── reports/
├── 03_timeline/
│   ├── README.md
│   ├── plaso/
│   ├── timesketch/
│   └── pivots/
├── 04_memory/
│   ├── README.md
│   ├── volatility/
│   ├── dumps_metadata/
│   └── findings/
├── 05_recovery/
│   ├── README.md
│   ├── filesystem_notes/
│   └── recovered_metadata/
└── 06_ir_reports/
    ├── README.md
    ├── initial_triage.md
    ├── interim_findings.md
    └── final_report.md
```

## 10-Week Lab Plan

| Week | Focus | Output |
|---:|---|---|
| 1 | Evidence handling | chain-of-custody template |
| 2 | Disk imaging | raw image and hash log |
| 3 | Memory acquisition | RAM dump metadata |
| 4 | Autopsy basics | artifact report |
| 5 | Browser artifacts | browser activity report |
| 6 | Registry/system artifacts | persistence report |
| 7 | Plaso/Timesketch | unified timeline |
| 8 | Volatility | memory findings |
| 9 | Filesystem recovery | recovered artifact report |
| 10 | IR report | final incident report |

## Capstone

Goal:

```text
lab incident
  -> disk acquisition
  -> memory acquisition
  -> artifact extraction
  -> timeline reconstruction
  -> memory analysis
  -> recovery attempt
  -> incident report
```

Deliverables:

- Chain-of-custody log.
- Disk hash log.
- Memory dump metadata.
- Artifact report.
- Timeline.
- Volatility output.
- Recovery notes.
- Final IR report.

## Safety And Integrity Boundary

Required:

- Work from evidence copies.
- Preserve original hashes.
- Record tool versions.
- Keep time zone explicit.
- Separate analyst notes from evidence.
- Do not alter source evidence.

Forbidden:

- Performing acquisition without authorization.
- Modifying original evidence.
- Omitting failed analysis steps.
- Treating tool output as fact without validation.

## Integration With Cyber AI OS

| DFIR Output | Later Use |
|---|---|
| Chain-of-custody schema | Forensics mode |
| Timeline events | SOC assistant context |
| Memory findings | Malware classifier and response engine |
| Browser/registry artifacts | Threat hunting features |
| Recovery workflow | Incident response playbooks |
| Final reports | RAG knowledge base |

