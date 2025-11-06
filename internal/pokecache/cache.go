package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data map[string]cacheEntry
	mu   sync.RWMutex
	ttl  time.Duration
	done chan struct{}
}

type cacheEntry struct {
	createdAt time.Time
	Val       []byte
}

func NewCache(ttl time.Duration) Cache {
	return Cache{
		data: make(map[string]cacheEntry),
		ttl:  ttl,
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = cacheEntry{
		Val:       val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.data[key]
	if !found {
		return nil, false
	}

	return item.Val, found
}

func (c *Cache) Set(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = cacheEntry{
		Val:       []byte(key),
		createdAt: time.Now(),
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.ttl)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			now := time.Now()
			c.mu.Lock()
			defer c.mu.Unlock()

			for key, item := range c.data {
				if now.Sub(item.createdAt) > c.ttl {
					delete(c.data, key)
				}
			}
		case <-c.done:
			return
		}
	}
}
