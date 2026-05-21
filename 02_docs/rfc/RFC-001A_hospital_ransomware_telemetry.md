# RFC-001A: Hospital Ransomware Telemetry Extension

## Status
Proposed

## Context
Specific telemetry requirements for Pilot 1 (Hospital Ransomware SOC) to detect encryption bursts, lateral movement, and persistence.

## Schema Definitions

### Process Event
- `pid`: Process ID
- `ppid`: Parent Process ID
- `uid`: User ID
- `exe`: Executable path
- `cmdline`: Command line arguments
- `hash`: SHA256 of the executable
- `entropy_score`: Calculated entropy of the process memory/binary
- `start_time`: Epoch timestamp

### Filesystem Event
- `path`: File path
- `operation`: [CREATE, WRITE, RENAME, DELETE, READ]
- `extension`: File extension
- `bytes_changed`: Delta of bytes written
- `rename_rate`: Frequency of renames by PID
- `write_rate`: Bytes per second by PID

### Network Event
- `src_ip`: Source IP
- `dst_ip`: Destination IP
- `port`: Destination port
- `protocol`: [TCP, UDP, ICMP]
- `smb_session`: SMB session metadata if applicable
- `auth_type`: Authentication method used

### Memory Event
- `rss`: Resident Set Size
- `allocation_spike`: Boolean/Value indicating sudden large allocations
- `executable_pages`: Count of RX pages
- `injected_regions`: Detection of remote thread injection

### Security Event
- `privilege_change`: Transition to higher privs
- `token_use`: Impersonation tokens
- `scheduled_task`: Creation/Modification of tasks
- `service_modification`: Changes to system services

## Target Sources
- eBPF (File/Process/Network)
- Auditd/Fanotify (Filesystem)
- Procfs (Memory/Process)
- Systemd Journal (Security Events)
- SMB Parsers (Network/Lateral Movement)
