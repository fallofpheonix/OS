/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package registry

import (
    "testing"
    "github.com/fallofpheonix/phoenixmind-validator/truth/evidence"
)

func TestTruthRegistryIntegration(t *testing.T) {
    registry := NewTruthRegistry()

    // 1. Add OBSERVED evidence
    registry.UpdateEntity(evidence.Evidence{EntityID: "test-entity", State: evidence.OBSERVED})
    entity, _ := registry.GetEntity("test-entity")
    if entity.CurrentState != evidence.OBSERVED {
        t.Errorf("Expected OBSERVED, got %s", entity.CurrentState)
    }

    // 2. Add a higher-priority WARNING
    registry.UpdateEntity(evidence.Evidence{EntityID: "test-entity", State: evidence.WARNING})
    entity, _ = registry.GetEntity("test-entity")
    if entity.CurrentState != evidence.WARNING {
        t.Errorf("Expected WARNING to override, got %s", entity.CurrentState)
    }

    // 3. Add a lower-priority VALIDATED, state should remain WARNING
    registry.UpdateEntity(evidence.Evidence{EntityID: "test-entity", State: evidence.VALIDATED})
    entity, _ = registry.GetEntity("test-entity")
    if entity.CurrentState != evidence.WARNING {
        t.Errorf("Expected WARNING to persist, got %s", entity.CurrentState)
    }
}
