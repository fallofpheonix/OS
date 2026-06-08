---\nStatus: Partial\nImplementation: 60%\nConfidence: Tested\n---\n# Merge Candidates

## pheonix-truth

Target:

truth/

Extract:

validators

scoring

resolution

consensus

Keep:

yes

Delete:

no

Status:

Extraction Complete: No assets found

---

## pheonix-validation

Target:

validation/

Extract:

runtime checks

proposal validators

observation validators

dependency validators

Keep:

yes

Delete:

no

Status:

Extraction Complete: No assets found

---

## PhoenixMind-Org

Target:

agents/

Extract:

agents (documentation only)

prompts

workflows

memory logic

planning

Keep:

partial

Delete:

after merge

---

## external

Target:

manual review

Extract:

unknown

Keep:

temporary

Delete:

blocked

---

## PhoenixOrg

Target:

detailed review

Extract:

core files (main.go, go.mod, INVARIANTS.md)
security module
observability (no functional assets found)

Keep:

temporary

Delete:

blocked

Status:

in progress: core files and security module extracted

---

## External VPN Components (Integration Target)

### OpenVPN

Target: `pheonixos/external/openvpn`

Source: `https://github.com/openvpn/openvpn.git`

Status: Cloned, pending integration strategy implementation

### ProtonVPN Go VPN Library

Target: `pheonixos/external/protonvpn-go-vpn-lib`

Source: `https://github.com/ProtonVPN/go-vpn-lib.git`

Status: Cloned, pending integration strategy implementation

### ProtonVPN Python API Core

Target: `pheonixos/external/protonvpn-python-api-core`

Source: `https://github.com/ProtonVPN/python-proton-vpn-api-core.git`

Status: Cloned, pending integration strategy implementation

### ProtonVPN Daemon

Target: `pheonixos/external/protonvpn-daemon`

Source: `https://github.com/ProtonVPN/proton-vpn-daemon.git`

Status: Cloned, pending integration strategy implementation

### ProtonVPN Local Agent (Rust)

Target: `pheonixos/external/protonvpn-local-agent-rs`

Source: `https://github.com/ProtonVPN/local-agent-rs.git`

Status: Cloned, pending integration strategy implementation