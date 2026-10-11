package kv

import (
	"errors"
	"reflect"
	"testing"
)

var ErrEmptyKey = errors.New("empty key")

func TestKeys_ReturnsAllKeysSorted(t *testing.T) {
	store := NewStore(3)
	store.Set("charlie", "3")
	store.Set("alpha", "1")
	store.Set("bravo", "2")

	got := store.Keys()
	want := []string{"alpha", "bravo", "charlie"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Keys() = %v; want %v", got, want)
	}
}

func TestKeys_EmptyStore(t *testing.T) {
	store := NewStore(1)
	got := store.Keys()
	if len(got) != 0 {
		t.Errorf("Keys() = %v; want empty slice", got)
	}
}

func TestSetGet_EmptyKeys(t *testing.T) {
	store := NewStore(1)

	if _, err := store.Get(""); err == nil || !errors.Is(err, ErrEmptyKey) {
		t.Error("Get() failed")
	}
}

func TestSetGet_RoundTrip(t *testing.T) {
	store := NewStore(1)
	store.Set("hello", "world")

	if val, err := store.Get("hello"); err != nil || val != "world" {
		t.Errorf("Get() = %v, %v; want 'world', true", val, err)
	}

	if val, err := store.Get("nonexistent"); err == nil || val != "" {
		t.Errorf("Get() = %v, %v; want '', false", val, err)
	}

}

func TestDelete(t *testing.T) {
	tests := []struct {
		name      string
		setup     map[string]string
		deleteKey string
		wantLen   int
	}{

		{"deletes existing ket", map[string]string{"a": "1", "b": "2"}, "a", 1},
		{"no-op on missing key", map[string]string{"a": "1"}, "x", 1},
		{"no-op on empty store", map[string]string{}, "x", 0},
		{"deletes last key", map[string]string{"only": "value"}, "only", 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			store := NewStore(4)
			for k, v := range tc.setup {
				store.Set(k, v)
			}

			store.Delete(tc.deleteKey)

			if got := store.Len(); got != tc.wantLen {
				t.Errorf("after delete (%q): number of keys = %d, want %d", tc.deleteKey, got, tc.wantLen)
			}

			if _, err := store.Get(tc.deleteKey); err == nil {
				t.Errorf("after delete (%q): Get() did not return error for deleted key", tc.deleteKey)
			}
		})
	}
}
