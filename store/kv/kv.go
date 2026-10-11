package kv

import (
	"fmt"
	"redisgo/store"
	"sort"
)

type Store struct {
	data    map[string]string
	maxSize int //0 = unlimited
}

func NewStore(maxSize int) *Store {
	return &Store{
		data:    make(map[string]string),
		maxSize: maxSize,
	}
}

func (s *Store) Len() int {
	return len(s.data)
}

// keys return every key currently in the store sorted alphabetically
func (s *Store) Keys() []string {
	keys := make([]string, 0, len(s.data))

	for key := range s.data {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}

func (s *Store) Get(key string) (string, error) {
	if key == "" {
		return "", store.ErrEmptyKey
	}

	val, ok := s.data[key]
	if !ok {
		return "", fmt.Errorf("Key %s does not exist", key)
	}
	return val, nil
}

func (s *Store) Set(key, value string) error {
	if key == "" {
		return store.ErrEmptyKey
	}

	_, exists := s.data[key]

	if s.maxSize > 0 && s.Len() >= s.maxSize && !exists {
		return fmt.Errorf("Set(%q): %w", key, store.ErrStoreFull)
	}

	s.data[key] = value
	return nil
}

func (s *Store) Delete(Key string) {
	delete(s.data, Key)
}
