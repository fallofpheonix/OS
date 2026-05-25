package memory

import "time"

type RepairEntry struct {
	ID         string
	Failure    string
	Repair     string
	Confidence float64
	Risk       string
	Result     string
	Timestamp  time.Time
}

var History []RepairEntry

func RecordRepair(entry RepairEntry) {
	History = append(History, entry)
}
