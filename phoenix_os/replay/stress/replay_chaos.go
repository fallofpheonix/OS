package stress

type ReplayMetrics struct {
	OverflowCount  int
	DropRate       float64
	OrderingErrors int
	ClockSkew      int64
	BurstLatency   int64
}
