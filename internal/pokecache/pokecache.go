package pokecache

import (
	"sync"
	"time"
)

// Cache
type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
	ttl     time.Duration // how long an entry stays fresh
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// New cache
func NewCache(ttl time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]cacheEntry),
		mu:      sync.Mutex{},
		ttl:     ttl,
	}

	// Start a goroutine to clean up expired entries
	go c.reapLoop()
	return c
}

// Store a new entry in the cache
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

// Get an entry from the cache if it is there
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.entries[key]
	if !exists {
		return nil, false
	}

	if time.Since(entry.createdAt) > c.ttl {
		delete(c.entries, key)
		return nil, false
	}

	return entry.val, true

}

// Continuously remove expired cache entries
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.ttl)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > c.ttl {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}

}
