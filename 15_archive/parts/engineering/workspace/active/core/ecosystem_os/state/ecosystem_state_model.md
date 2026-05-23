# Ecosystem State Model

## Complex System Analysis

- **Ecosystem State**: INITIALIZED
- **Failure Boundary**: Identified at cross-runtime coupling points.
- **Stability**: MEDIUM (due to manifest fragmentation)
- **Universality**: LOW (runtime explosion in progress)

## Threshold Variables

- **N repos**: ~150
- **C dependencies**: High (coupling identified in runtime_graph.json)
- **R runtimes**: 3 active, 2 shadow
- **D domains**: 6 identified
- **A archives**: ~20
- **V validation**: Pending global score

## Detection Logic

- **coordination collapse**: Triggered if manifest synchronization fails.
- **runtime explosion**: Detected when R/N > 0.1
- **archive saturation**: Detected when A/N > 0.5
- **research drift**: Detected when research lineage is missing for > 20% of ACTIVE repos.
- **manifest fragmentation**: Detected when multiple versions of registry exist.
