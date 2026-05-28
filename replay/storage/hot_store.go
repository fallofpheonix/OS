package storage

import "sync"

type HotStore struct {
	mu   sync.RWMutex
	data map[string][]byte
}

func NewHotStore() *HotStore {
	return &HotStore{data: make(map[string][]byte)}
}

func (s *HotStore) Put(id string, val []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[id] = val
}

func (s *HotStore) Get(id string) ([]byte, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[id]
	return val, ok
}
