package ttl

import (
	"fmt"
	"redisgo/store"
	"time"
)

type TTLStore struct {
	data map[string]ttlEntry
	ttl  time.Duration
}

type ttlEntry struct {
	value     string
	expiresAt time.Time
}

func NewTTLStore(ttl time.Duration) *TTLStore {
	return &TTLStore{
		data: make(map[string]ttlEntry),
		ttl:  ttl,
	}
}

func (t *TTLStore) Len() int {
	return len(t.data)
}

func (t *TTLStore) Set(key string, value string) error {
	if key == "" {
		return store.ErrEmptyKey
	}

	t.data[key] = ttlEntry{
		value:     value,
		expiresAt: time.Now().Add(t.ttl),
	}

	return nil
}

func (t *TTLStore) Get(key string) (string, error) {
	if key == "" {
		return "", store.ErrEmptyKey
	}

	entry, ok := t.data[key]

	if !ok || time.Now().After(entry.expiresAt) {
		delete(t.data, key)

		return "", fmt.Errorf("key %s does not exist", key)
	}

	return entry.value, nil

}

func (t *TTLStore) Delete(Key string) {
	delete(t.data, Key)
}
