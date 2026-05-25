# Phase 5 Build Gate: Networking

## Conceptual Knowledge

- [ ] Explain each OSI layer and protocol examples.
- [ ] Draw TCP three-way handshake and identify flags.
- [ ] Explain subnetting, CIDR, private ranges, and public ranges.
- [ ] Trace full DNS query path.
- [ ] Explain NAT rewriting, limitations, and state.
- [ ] Draw TLS handshake and explain certificate exchange and key derivation.
- [ ] Compare HTTP/1.1, HTTP/2, and HTTP/3.

## Packet Capture And Analysis

- [ ] Capture traffic with tcpdump or Wireshark.
- [ ] Filter by protocol, port, and address.
- [ ] Export and analyze capture files.
- [ ] Capture DNS query and response.
- [ ] Identify DNS record types.
- [ ] Explain DNS caching behavior.
- [ ] Capture TCP connection setup.
- [ ] Identify SYN, SYN-ACK, and ACK.
- [ ] Record sequence numbers and window sizes.
- [ ] Explain TCP options.
- [ ] Capture HTTP/HTTPS request-response.
- [ ] Decode headers and payload where plaintext is available.
- [ ] Compare HTTP/1.1 and HTTP/2 behavior.

## Network Troubleshooting

- [ ] Trace route with `traceroute` or `mtr`.
- [ ] Examine routing table with `ip route`.
- [ ] Explain default gateway.
- [ ] Diagnose routing loops.
- [ ] Query DNS with `dig` and `nslookup`.
- [ ] Identify missing DNS records.
- [ ] Test recursive vs iterative resolution.
- [ ] Test reachability with `ping`.
- [ ] Check listening ports with `netstat` or `ss`.
- [ ] Identify blocked ports with `nmap`.
- [ ] Trace packet loss and latency.

## Protocol Implementation

- [ ] Implement TCP server.
- [ ] Implement TCP client.
- [ ] Measure round-trip time.
- [ ] Implement DNS query formatter.
- [ ] Implement DNS response parser.
- [ ] Handle DNS response codes and record types.
- [ ] Inspect TLS handshake with `openssl s_client`.
- [ ] Verify certificate chain and expiration.
- [ ] Compare cipher suites.
- [ ] Implement basic HTTP/1.1 client.
- [ ] Send GET and POST.
- [ ] Parse response headers and body.
- [ ] Handle redirects and cookies.

## Advanced Analysis

- [ ] Manually parse packets from hex dump.
- [ ] Identify frame, IP, TCP, and application fields.
- [ ] Verify checksums where feasible.
- [ ] Measure latency.
- [ ] Measure throughput with `iperf`.
- [ ] Analyze retransmissions and packet loss.
- [ ] Identify network vs application bottlenecks.
- [ ] Identify unencrypted protocols.
- [ ] Detect DNS exfiltration patterns.
- [ ] Analyze TLS version and cipher strength.
- [ ] Explain MITM attack vectors.

## Pass Criteria

- Packet captures and parsed notes exist.
- TCP/DNS/HTTP implementations run against controlled tests.
- Troubleshooting command outputs are recorded.
- Security analysis identifies protocol exposure and TLS posture.

