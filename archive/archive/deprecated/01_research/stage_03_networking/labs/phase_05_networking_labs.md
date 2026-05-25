# Phase 5 Labs: Networking

## Rule

Research stays in `01_research/`. Implementation evidence belongs in `14_experiments/poc/stage_03_networking/phase_05_networking/`.

## Lab 01: OSI And Encapsulation

- Map protocols to OSI layers.
- Dissect an Ethernet/IP/TCP/Application packet.
- Record encapsulation boundaries.

## Lab 02: DNS

- Capture DNS traffic.
- Parse query and response.
- Compare recursive and iterative behavior.
- Record TTL and cache behavior.

## Lab 03: TCP

- Capture handshake.
- Record flags, sequence numbers, window sizes, and options.
- Implement TCP client/server.
- Measure RTT.

## Lab 04: Routing And NAT

- Inspect routing table.
- Trace destination route.
- Model NAT state table.
- Diagnose unreachable route scenarios.

## Lab 05: HTTP And TLS

- Inspect HTTP/1.1 request/response.
- Compare HTTP/2 or HTTP/3 behavior where tooling supports it.
- Inspect TLS with `openssl s_client`.
- Record certificate chain and cipher suite.

## Lab 06: Packet Dissection

- Export capture.
- Parse selected packet manually from hex.
- Verify layer fields and checksums where feasible.

## Lab 07: Performance

- Measure latency.
- Measure throughput.
- Capture retransmissions.
- Identify bottleneck class.

## Lab 08: Security

- Identify plaintext protocols.
- Analyze TLS version/cipher strength.
- Detect DNS tunneling/exfiltration patterns in controlled data.
- Document MITM assumptions.

## Completion Record

| Lab | Status | Evidence path | Notes |
|---|---|---|---|
| OSI and encapsulation | TODO | TBD | packet layer mapping |
| DNS | TODO | TBD | query, response, cache |
| TCP | TODO | TBD | handshake, client/server |
| Routing and NAT | TODO | TBD | route and state |
| HTTP and TLS | TODO | TBD | request, certs, ciphers |
| Packet dissection | TODO | TBD | hex parsing |
| Performance | TODO | TBD | latency, throughput |
| Security | TODO | TBD | plaintext, TLS, DNS exfil |

