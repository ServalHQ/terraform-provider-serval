// Package cache provides a generic caching layer for Terraform provider resources.
// It pre-fetches all resources via List endpoints on first access, dramatically
// reducing API calls from O(n) to O(1) during plan/apply operations.
//
// The cache is initialized once per provider lifecycle (per terraform plan/apply)
// and automatically resets when the provider process exits.
package cache

import (
	"sync"
)

// ResourceCache is a generic thread-safe cache for any resource type.
// It uses sync.Once to ensure the loader is called exactly once.
type ResourceCache[T any] struct {
	once   sync.Once
	mu     sync.RWMutex
	data   map[string]*T
	loaded bool
	err    error
}

// NewResourceCache creates a new empty cache.
func NewResourceCache[T any]() *ResourceCache[T] {
	return &ResourceCache[T]{
		data: make(map[string]*T),
	}
}

// KeyedCache is a cache that stores separate ResourceCache instances per key (e.g., per team_id).
// This allows efficient caching when resources are partitioned by a parent key.
type KeyedCache[T any] struct {
	mu     sync.RWMutex
	caches map[string]*ResourceCache[T]
}

// NewKeyedCache creates a new keyed cache.
func NewKeyedCache[T any]() *KeyedCache[T] {
	return &KeyedCache[T]{
		caches: make(map[string]*ResourceCache[T]),
	}
}

// GetOrCreateCache returns the cache for the given key, creating it if it doesn't exist.
func (kc *KeyedCache[T]) GetOrCreateCache(key string) *ResourceCache[T] {
	kc.mu.RLock()
	cache, exists := kc.caches[key]
	kc.mu.RUnlock()

	if exists {
		return cache
	}

	kc.mu.Lock()
	defer kc.mu.Unlock()

	// Double-check after acquiring write lock
	if cache, exists = kc.caches[key]; exists {
		return cache
	}

	cache = NewResourceCache[T]()
	kc.caches[key] = cache
	return cache
}

// FindInLoadedCaches searches all loaded caches for an item by ID.
// Returns (item, key) if found, (nil, "") if not found.
// This is useful for imports where we don't know which parent key the item belongs to.
func (kc *KeyedCache[T]) FindInLoadedCaches(id string) (*T, string) {
	kc.mu.RLock()
	defer kc.mu.RUnlock()

	for key, cache := range kc.caches {
		if item, found := cache.Get(id); found {
			return item, key
		}
	}
	return nil, ""
}

// MappingCache is a thread-safe cache for storing ID mappings (e.g., app_instance_id → team_id).
type MappingCache struct {
	mu   sync.RWMutex
	data map[string]string
}

// NewMappingCache creates a new mapping cache.
func NewMappingCache() *MappingCache {
	return &MappingCache{
		data: make(map[string]string),
	}
}

// Get retrieves a mapping by key.
func (mc *MappingCache) Get(key string) (string, bool) {
	mc.mu.RLock()
	defer mc.mu.RUnlock()
	value, exists := mc.data[key]
	return value, exists
}

// Set stores a mapping.
func (mc *MappingCache) Set(key, value string) {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	mc.data[key] = value
}

// Get retrieves an item from the cache by ID.
// Returns (item, true) if found, (nil, false) if not found or cache not loaded.
func (c *ResourceCache[T]) Get(id string) (*T, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if !c.loaded {
		return nil, false
	}
	item, exists := c.data[id]
	return item, exists
}

// IsLoaded returns whether the cache has been populated.
func (c *ResourceCache[T]) IsLoaded() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.loaded
}

// LoadError returns any error that occurred during loading.
func (c *ResourceCache[T]) LoadError() error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.err
}

// Load populates the cache using the provided loader function.
// This is thread-safe and will only execute the loader once.
func (c *ResourceCache[T]) Load(loader func() (map[string]*T, error)) error {
	c.once.Do(func() {
		data, err := loader()
		c.mu.Lock()
		defer c.mu.Unlock()
		if err != nil {
			c.err = err
			return
		}
		c.data = data
		c.loaded = true
	})
	return c.LoadError()
}

// GetOrLoad retrieves an item from the cache by ID, loading the cache first if needed.
// Returns (item, found, error). If loading fails, error is non-nil.
// If the item doesn't exist after loading, found is false.
func (c *ResourceCache[T]) GetOrLoad(id string, loader func() (map[string]*T, error)) (*T, bool, error) {
	if err := c.Load(loader); err != nil {
		return nil, false, err
	}
	item, exists := c.Get(id)
	return item, exists, nil
}
