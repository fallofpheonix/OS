package stress

type DropTracker struct {
    Seen map[int64]bool
}

func (t *DropTracker) Track(sequence int64) bool {
    if t.Seen[sequence] {
        return true
    }
    t.Seen[sequence] = true
    return false
}
