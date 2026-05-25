package audit

import "time"

type Violation struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Actor     string `json:"actor"`
	Action    string `json:"action"`
	Severity  string `json:"severity"`
	Result    string `json:"result"`
}

type IncidentStore struct {
	Logs []Violation
}

func (s *IncidentStore) Log(v Violation) {
	v.Timestamp = time.Now().String()
	s.Logs = append(s.Logs, v)
}
