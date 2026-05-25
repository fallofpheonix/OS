# PhoenixOS: Truth Model

`truth.TruthLedger` cryptographic hash chain.
Each `EvidenceWrapper` contains: `prev_hash` + `payload` + `sequence_id` + `sha256`.
`Verify()` recalculates from genesis.
