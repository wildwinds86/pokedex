package internal

import (
	"time"
)

type Cache struct {
	interval time.Duration
	contents map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(intervalDuration time.Duration) Cache {
	return Cache{
		interval: intervalDuration,
		contents: make(map[string]cacheEntry),
	}
}

func (c *Cache) Add(key string, value []byte) {
	c.contents[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	data, exists := c.contents[key]

	if !exists {
		return nil, false
	} else {
		return data.val, true
	}
}

func (c *Cache) ReapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		timeNow := time.Now()
		for s, ce := range c.contents {
			timeDiff := timeNow.Sub(ce.createdAt)
			if timeDiff.Seconds() > 30 {
				delete(c.contents, s)
			}
		}
	}
}
