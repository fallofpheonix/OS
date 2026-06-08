/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package resource

import "errors"

type BoundedAllocator struct {
	MaxBytes    uint64
	MaxEntries  uint64
	UsedBytes   uint64
	UsedEntries uint64
}

func NewBoundedAllocator(maxBytes uint64, maxEntries uint64) *BoundedAllocator {
	return &BoundedAllocator{
		MaxBytes:   maxBytes,
		MaxEntries: maxEntries,
	}
}

func (b *BoundedAllocator) Allocate(bytes uint64) error {
	if b.UsedBytes+bytes > b.MaxBytes || b.UsedEntries+1 > b.MaxEntries {
		return errors.New("resource limit exceeded")
	}
	b.UsedBytes += bytes
	b.UsedEntries++
	return nil
}

func (b *BoundedAllocator) Deallocate(bytes uint64) {
	if b.UsedBytes >= bytes {
		b.UsedBytes -= bytes
	}
	if b.UsedEntries > 0 {
		b.UsedEntries--
	}
}
