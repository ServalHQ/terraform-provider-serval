// Package cache provides a simple prefetch-based caching layer for the Serval
// Terraform provider. When PrefetchMode is enabled, all resources are loaded
// into memory during provider initialization (Configure), and subsequent
// Read/ImportState operations serve from cache only.
//
// When PrefetchMode is disabled (default), the cache is not used and all
// operations fall through to normal API calls.
package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	serval "github.com/ServalHQ/serval-go"
)

// PrefetchMode indicates whether the provider is operating in prefetch mode.
// When true, all Read/ImportState operations must be served from cache.
var PrefetchMode bool

// Store is a simple map-based cache for any resource type.
// It is populated once during provider Configure() and is read-only after that.
type Store[T any] struct {
	data map[string]*T
}

// NewStore creates a new empty store.
func NewStore[T any]() *Store[T] {
	return &Store[T]{data: make(map[string]*T)}
}

// Put adds an item to the store.
func (s *Store[T]) Put(id string, item *T) {
	s.data[id] = item
}

// Get retrieves an item from the store by ID.
func (s *Store[T]) Get(id string) (*T, bool) {
	item, ok := s.data[id]
	return item, ok
}

// Keys returns all IDs in the store.
func (s *Store[T]) Keys() []string {
	keys := make([]string, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	return keys
}

// Len returns the number of items in the store.
func (s *Store[T]) Len() int {
	return len(s.data)
}

// TryRead is the unified cache check for Read/ImportState methods.
//   - PrefetchMode off: returns (nil, false, nil) — caller falls through to normal API
//   - PrefetchMode on + hit: returns (item, true, nil)
//   - PrefetchMode on + miss: returns (nil, false, error)
func TryRead[T any](store *Store[T], id string) (*T, bool, error) {
	if !PrefetchMode || store == nil {
		return nil, false, nil
	}
	item, ok := store.Get(id)
	if !ok {
		return nil, false, fmt.Errorf(
			"resource %s not found in prefetch cache", id)
	}
	return item, true, nil
}

// IsServerError returns true if the error is a 5xx API error.
// During prefetch, server errors for individual teams/parents can be
// skipped so that prefetch succeeds for the resources it can load.
func IsServerError(err error) bool {
	var apiErr *serval.Error
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode >= 500
	}
	return false
}

// LoadFromFile populates a store from a JSON file containing an array of objects.
// The provided unmarshal function deserializes each JSON object into T, and
// getID extracts the cache key from the deserialized item.
//
// This enables future file-based cache loading where the svmeta Fetcher writes
// cache files using serval-go types, and the provider loads them here.
func LoadFromFile[T any](path string, unmarshal func([]byte, *T) error, getID func(*T) string) (*Store[T], error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read cache file: %w", err)
	}
	var raw []json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("failed to parse cache file: %w", err)
	}
	store := NewStore[T]()
	for _, r := range raw {
		var item T
		if err := unmarshal(r, &item); err != nil {
			return nil, fmt.Errorf("failed to unmarshal cache entry: %w", err)
		}
		store.Put(getID(&item), &item)
	}
	return store, nil
}
