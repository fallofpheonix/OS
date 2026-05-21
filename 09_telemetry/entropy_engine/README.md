# Entropy Engine (L3)

Implementation of Shannon Entropy and KL Divergence for real-time telemetry analysis.

## Purpose
Detect high-entropy data streams (ransomware, encrypted exfiltration) in the `vfs_write` pipeline.

## Structure
- `src/`: Core Go implementation.
- `tests/`: Unit and integration tests.
- `bench/`: Performance benchmarks.
- `replay/`: Telemetry replay logic.
- `debug/`: Debugging traces.
- `artifacts/`: Build outputs.

## Performance Budget
- **Entropy Calc:** < 5us / 4KB block.
- **Memory:** O(1) space complexity (fixed 256-byte frequency table).

## Validation Gates
- [ ] Build success
- [ ] Unit test pass (100% coverage)
- [ ] Performance < 5us
- [ ] Replay deterministic output
