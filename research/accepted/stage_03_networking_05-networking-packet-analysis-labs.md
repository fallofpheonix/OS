# Networking And Packet Analysis Labs

## Scope

Packet-engineering curriculum for traffic analysis, IDS pipelines, VPN inspection, and zero-trust networking.

## 1. Core Networking Stack

### Topics

- OSI model.
- TCP/IP.
- Ethernet.
- ARP.
- IP.
- TCP.
- UDP.
- Routing.
- BGP.
- NAT.
- VLAN.
- DNS.
- DHCP.
- TLS.
- HTTP/2.
- HTTP/3.
- QUIC.
- WebSockets.

### Exercise

Capture and decode:

- QUIC handshake.
- TLS 1.3 session.
- DNS-over-HTTPS request.
- DHCP lease.
- ARP resolution.

### Tools

- Wireshark.
- tcpdump.
- Zeek.

### Exit Criteria

- PCAPs are saved.
- Packet flow is documented.
- Protocol fields are annotated.

## 2. Packet Inspection And DPI

### Topics

- Ethernet frames.
- IP headers.
- Fragmentation.
- TCP flags.
- Protocol options.
- Deep packet inspection.
- Signature rules.
- Heuristic analysis.
- Beaconing.
- DNS tunneling.
- TLS fingerprinting.

### Exercise

1. Write a Suricata rule for suspicious DNS patterns.
2. Feed PCAP into Zeek.
3. Generate:
   - `conn.log`
   - `dns.log`
   - `http.log`
4. Compare Suricata alerts with Zeek logs.

### Exit Criteria

- Rule has test PCAP.
- Zeek logs are parsed.
- Detection notes distinguish signature vs behavior.

## 3. Tunneling, VPNs, And Proxies

### Topics

- VPN.
- WireGuard.
- IP-in-UDP.
- Key rotation.
- Forward proxy.
- Reverse proxy.
- TLS termination.
- Load balancing.
- Cookie stickiness.
- Header-based routing.
- Zero-trust networking.
- Micro-segmentation.

### Exercise

Build:

- Two-VM WireGuard tunnel.
- Reverse proxy using nginx.
- Load balancer using HAProxy.

Capture:

- WireGuard UDP packets.
- Pre-tunnel and post-tunnel traffic.
- Proxy headers.
- TLS termination metadata.

### Exit Criteria

- `wg show` state is recorded.
- tcpdump capture exists.
- Zeek logs show proxy/load-balancer behavior.

## 4. Tools In Practice

### Wireshark

Use for:

- Protocol dissectors.
- Display filters.
- Stream reconstruction.
- Flow statistics.
- TLS/QUIC inspection where keys or metadata permit.

### tcpdump

Use for:

- Headless capture.
- Server-side packet collection.
- BPF filters.
- PCAP generation.

Example:

```sh
tcpdump -i eth0 -w traffic.pcap
```

### Suricata

Use for:

- IDS/IPS rules.
- HTTP, DNS, TLS detection.
- Offline PCAP analysis.

### Zeek

Use for:

- Behavioral logs.
- Structured metadata.
- Connection summaries.
- DNS and HTTP analysis.

### WireGuard

Use for:

- Minimal VPN lab.
- UDP encapsulation study.
- Peer-state correlation.

## 5. Combined Pipeline Exercise

Flow:

```text
tcpdump capture
  -> Wireshark inspection
  -> Suricata offline detection
  -> Zeek log extraction
  -> comparison report
```

Commands:

```sh
tcpdump -i eth0 -w traffic.pcap
suricata -r traffic.pcap -l suricata-out/
zeek -r traffic.pcap
```

Compare:

- Suricata alerts.
- Zeek `conn.log`.
- Zeek `dns.log`.
- Zeek `http.log`.
- Wireshark decoded streams.

## Suggested Repo Structure

```text
networking/
├── 01_osi_tcpip/
│   ├── README.md
│   ├── pcaps/
│   └── notes/
├── 02_routing_dhcp_dns/
│   ├── README.md
│   ├── labs/
│   └── captures/
├── 03_tls_http_quic/
│   ├── README.md
│   ├── tls13/
│   ├── http2/
│   └── quic/
├── 04_dpi_tools/
│   ├── README.md
│   ├── suricata_rules/
│   ├── zeek_logs/
│   └── pcaps/
└── 05_vpn_wireguard/
    ├── README.md
    ├── configs/
    ├── captures/
    └── analysis/
```

## 10-Week Lab Plan

| Week | Focus | Output |
|---:|---|---|
| 1 | OSI and Ethernet | annotated ARP/IP PCAP |
| 2 | TCP/UDP | connection and flow analysis |
| 3 | DHCP/DNS | bootstrap traffic captures |
| 4 | Routing/NAT/VLAN | forwarding and segmentation notes |
| 5 | TLS | TLS 1.3 handshake capture |
| 6 | HTTP/2, HTTP/3, QUIC | multiplexing comparison |
| 7 | WebSockets | long-lived flow analysis |
| 8 | Suricata | custom rule and alert output |
| 9 | Zeek | log extraction and behavior analysis |
| 10 | WireGuard/proxy | tunnel and proxy report |

## Integration With Cyber AI OS

| Networking Output | Later Use |
|---|---|
| PCAP corpus | ML feature engineering |
| Zeek logs | SOC ingestion and behavior analytics |
| Suricata rules | IDS/IPS baseline |
| WireGuard lab | secure management plane |
| Proxy/load-balancer notes | zero-trust service routing |

