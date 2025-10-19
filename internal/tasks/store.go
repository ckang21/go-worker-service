package tasks

import (
	"sync"
	"time"
)

type Store struct {
	mu    sync.RWMutex
	tasks map[string]Task
}

func NewStore() *Store {
	return &Store{tasks: make(map[string]Task)}
}

func (s *Store) Create(t Task) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tasks[t.ID] = t
}

func (s *Store) UpdateStatus(id string, st Status) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if t, ok := s.tasks[id]; ok {
		t.Status = st
		t.UpdatedAt = time.Now().UTC()
		s.tasks[id] = t
	}
}

func (s *Store) List() []Task {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		out = append(out, t)
	}
	return out
}
