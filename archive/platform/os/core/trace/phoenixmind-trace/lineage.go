/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package trace

import (
	"time"
)

// LineageLink represents a connection between two pieces of data,
// indicating that one is derived from the other.
type LineageLink struct {
	SourceID string
	TargetID string
	Type     string // e.g., "DERIVED_FROM", "TRANSFORMED_BY", "COMPOSED_OF"
	Timestamp time.Time
}

// LineageGraph tracks the provenance and derivation of all truth-related data.
type LineageGraph struct {
	Links []LineageLink
	// Potentially a map for faster lookup by ID
}

// AddLink records a new lineage relationship.
func (lg *LineageGraph) AddLink(sourceID, targetID, linkType string) {
	lg.Links = append(lg.Links, LineageLink{
		SourceID:  sourceID,
		TargetID:  targetID,
		Type:      linkType,
		Timestamp: time.Now(),
	})
}

// GetLineageForID retrieves all direct lineage links for a given ID.
func (lg *LineageGraph) GetLineageForID(id string) []LineageLink {
	var links []LineageLink
	for _, link := range lg.Links {
		if link.SourceID == id || link.TargetID == id {
			links = append(links, link)
		}
	}
	return links
}
