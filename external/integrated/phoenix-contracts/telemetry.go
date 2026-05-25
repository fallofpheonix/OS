package contracts

// Event defines the standard structure for telemetry ingestion.
type Event interface {
	GetSeqID() int64
	GetTick() uint64
	GetWallTime() int64
	GetSource() string
	GetPID() int
	GetSeverity() float64
	GetHash() string
	Verify() bool
}
