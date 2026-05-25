package stress

type OverflowDetector struct {
    Limit int
}

func (d *OverflowDetector) IsOverflow(count int) bool {
    return count > d.Limit
}
