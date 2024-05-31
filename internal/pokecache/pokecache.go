package pokecache

import (
	// "fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mux sync.RWMutex	
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.entries[key] = cacheEntry{ 
		createdAt: time.Now(),
		val: val,
	}	
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()	
	entry, ok := c.entries[key]	
	return entry.val, ok
}

func (c *Cache) Entries() map[string]cacheEntry {
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.entries
}

func (c *Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)		
	}
}

func (c *Cache) reap(interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()	
	for key, val := range c.entries {
		if time.Since(val.createdAt) > interval {
			delete(c.entries, key)
		}
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		entries: make(map[string]cacheEntry),
	}
	go cache.ReapLoop(interval)
	return &cache
}