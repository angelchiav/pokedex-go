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

func NewCache(ttl time.Duration) *Cache {
	c := &Cache{
		data: make(map[string]cacheEntry),
		ttl:  ttl,
		done: make(chan struct{}),
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	buf := append([]byte(nil), val...)
	c.data[key] = cacheEntry{
		Val:       buf,
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

	return append([]byte(nil), item.Val...), true
}

func (c *Cache) Set(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	buf := append([]byte(nil), val...)

	c.data[key] = cacheEntry{
		Val:       buf,
		createdAt: time.Now(),
	}
}

func (c *Cache) Close() {
	if c == nil {
		return
	}
	select {
	case <-c.done:
		return
	default:
	}
	close(c.done)
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.ttl)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			now := time.Now()
			c.mu.Lock()

			for key, item := range c.data {
				if now.Sub(item.createdAt) > c.ttl {
					delete(c.data, key)
				}
			}
			c.mu.Unlock()
		case <-c.done:
			return
		}
	}
}
