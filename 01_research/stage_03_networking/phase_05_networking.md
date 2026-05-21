# Phase 5: Networking

## Overview

Understand networking from theory to practice: OSI, TCP/IP, DNS, routing, NAT, TLS, HTTP evolution, packet capture, diagnostics, and protocol implementation.

## Classification

- Stage: `Stage_03_Networking`
- Type: `FOUNDATIONAL`
- Status: `RESEARCH_ONLY`
- Difficulty: intermediate-advanced
- Estimated duration: 6-8 weeks
- Depends on:
  - Phase 0 Computer Science Foundations
  - Phase 1 Computer Architecture
  - Phase 2 Low-Level Programming
  - Phase 3 Operating Systems

## Research

### Network Fundamentals

- OSI layer 1: physical signals and cables.
- OSI layer 2: frames, MAC addresses, switching, Ethernet, ARP.
- OSI layer 3: IP, routing, IPv4, IPv6, ICMP.
- OSI layer 4: transport.
- OSI layers 5-7: session, presentation, application.
- TCP/IP 4-layer model.
- Protocol families.
- Socket types.
- Packet encapsulation.
- Headers.

### Layer 3: Internet Protocol

- IPv4 header.
- IPv6 header.
- IP addressing.
- Subnetting.
- CIDR.
- Address classes.
- Special addresses.
- Fragmentation.
- MTU.
- TTL.
- Hop limit.
- Longest prefix match.
- Default gateway.
- Static routing.
- Dynamic routing.
- RIP.
- OSPF.
- BGP.
- NAT source and destination rewriting.
- NAT state.
- Hairpinning.
- Port forwarding.

### Layer 2: ARP And Data Link

- ARP request/reply.
- ARP cache.
- ARP TTL.
- Gratuitous ARP.
- ARP spoofing.
- Ethernet frame fields.
- MAC broadcast.
- Switch MAC address table.
- VLAN tagging.

### Layer 4: Transport Protocols

- TCP header.
- TCP flags.
- Sequence numbers.
- Acknowledgments.
- TCP state machine.
- Three-way handshake.
- Flow control.
- Congestion control:
  - Reno.
  - Cubic.
  - BBR.
- Retransmission.
- Timeouts.
- Slow start.
- TCP options.
- UDP.
- UDP checksum.
- QUIC.
- 0-RTT.
- Connection migration.
- Multiplexed streams.

### Layer 5-7: Applications And Encryption

- DNS query/response.
- DNS records:
  - A.
  - AAAA.
  - MX.
  - CNAME.
  - NS.
  - SOA.
- Recursive resolution.
- Iterative resolution.
- DNS cache and TTL.
- DoH.
- DoT.
- HTTP/1.0.
- HTTP/1.1.
- HTTP/2.
- HTTP/3.
- REST.
- WebSocket.
- TLS handshake.
- Cipher suites.
- ECDHE.
- RSA.
- Certificates.
- Trust anchors.
- TLS 1.2 vs TLS 1.3.
- Perfect forward secrecy.
- Certificate pinning.
- VPN types.
- IPSec.
- OpenVPN.
- WireGuard.

### Network Analysis And Tools

- tcpdump.
- Wireshark.
- BPF filter syntax.
- Packet dissection.
- TCP/UDP stream following.
- ping.
- traceroute.
- netstat.
- ss.
- nslookup.
- dig.
- curl.
- wget.
- iperf.
- netcat.

## Tools

Required:

- Wireshark.
- tcpdump.

Optional:

- netcat.
- telnet.
- curl.
- wget.
- nmap.
- mtr.
- iftop.
- nethogs.
- ss.
- ip.
- route.
- arp.

## Learning Outcomes

- Understand TCP/IP and OSI layers.
- Capture and analyze network packets.
- Diagnose connectivity issues.
- Understand routing, DNS, NAT, TLS, and HTTP evolution.
- Implement TCP, DNS, and HTTP clients.
- Troubleshoot production networking issues.

