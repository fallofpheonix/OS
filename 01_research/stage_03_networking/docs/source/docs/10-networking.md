# Networking

## Initial Scope

Linux-derived path:

- DHCP.
- DNS resolution.
- SSH client.
- Optional SSH server disabled by default.

Scratch path:

- NIC detection.
- Packet RX/TX.
- ARP.
- IPv4.
- ICMP.
- UDP.
- TCP later.

## Stack Layers

```text
Driver
  -> Ethernet
  -> ARP
  -> IP
  -> ICMP / UDP / TCP
  -> Sockets
  -> Userspace services
```

## Test Commands

Linux-derived image:

```sh
ip addr
ip route
ping -c 3 1.1.1.1
resolvectl status
```

## Failure Considerations

- Do not enable inbound services by default.
- Packet parsing must bounds-check every header.
- Network driver must handle RX ring exhaustion.
- DHCP failure must not block boot indefinitely.

