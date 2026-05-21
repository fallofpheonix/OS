Incident Graph Template

Purpose:
Schema and required nodes/edges for constructing causal incident graphs.

Nodes:
- Host (id, hostname, os, tags)
- Process (pid, ppid, cmdline, start_ts)
- File (path, hash, size, mtime)
- NetworkConn (src_ip, dst_ip, src_port, dst_port, proto, ts)
- Alert (id, rule, severity, ts)

Edges (examples):
- spawned_by: Process -> Process
- opened: Process -> File
- wrote: Process -> File
- connected_to: Process -> NetworkConn
- alerted_on: Process/File/NetworkConn -> Alert

Best practices:
- Include timestamps on edges
- Normalize node IDs to stable hashes
- Record evidence references (artifact paths)
