package kvstore

import (
	"sync"
)

// Store is a thread safe key value stores

type Store struct {
	mu   sync.RWMutex
	data map[string]string
}

// Newstore creates a new k-v store
func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

// set method
func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

// fetch method
func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[key]
	return val, ok
}

// delete a key value based pair
func (s *Store) Remove(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}
