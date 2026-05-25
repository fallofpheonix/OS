package game

import (
	"sync"
)

type ScoreState struct {
	mu                  sync.RWMutex
	Score               int      `json:"score"`
	Level               string   `json:"level"`
	Multiplier          float64  `json:"multiplier"`
	Badges              []string `json:"badges"`
	CompletedChallenges []string `json:"completed_challenges"`
}

func NewScoreState() *ScoreState {
	return &ScoreState{
		Score:               0,
		Level:               "Novice SOC Analyst",
		Multiplier:          1.0,
		Badges:              []string{},
		CompletedChallenges: []string{},
	}
}

func (s *ScoreState) GetState() ScoreState {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return ScoreState{
		Score:               s.Score,
		Level:               s.Level,
		Multiplier:          s.Multiplier,
		Badges:              append([]string(nil), s.Badges...),
		CompletedChallenges: append([]string(nil), s.CompletedChallenges...),
	}
}

func (s *ScoreState) AddPoints(pts int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	added := int(float64(pts) * s.Multiplier)
	s.Score += added
	s.updateLevel()
}

func (s *ScoreState) ApplyPenalty(pts int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Score -= pts
	if s.Score < 0 {
		s.Score = 0
	}
	s.updateLevel()
}

func (s *ScoreState) AwardBadge(badge string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, b := range s.Badges {
		if b == badge {
			return false
		}
	}
	s.Badges = append(s.Badges, badge)
	s.Multiplier += 0.1
	return true
}

func (s *ScoreState) CompleteChallenge(challenge string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, c := range s.CompletedChallenges {
		if c == challenge {
			return false
		}
	}
	s.CompletedChallenges = append(s.CompletedChallenges, challenge)
	return true
}

func (s *ScoreState) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Score = 0
	s.Level = "Novice SOC Analyst"
	s.Multiplier = 1.0
	s.Badges = []string{}
	s.CompletedChallenges = []string{}
}

func (s *ScoreState) updateLevel() {
	if s.Score < 500 {
		s.Level = "Novice SOC Analyst"
	} else if s.Score < 1500 {
		s.Level = "Warden Guardian"
	} else if s.Score < 3000 {
		s.Level = "Determinism Master"
	} else {
		s.Level = "Cybernetic Sentinel"
	}
}
