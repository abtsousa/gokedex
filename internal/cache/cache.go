package cache

import (
	"sync"
	"time"
)

type Cache struct {
	mtx     *sync.Mutex
	entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{&sync.Mutex{}, map[string]cacheEntry{}}
	c.CleanEvery(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	mp := c.entries
	mp[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	entry, exists := c.entries[key]
	return entry.val, exists
}

func (c *Cache) Remove(key string) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	delete(c.entries, key)
}

func (c *Cache) CleanEvery(interval time.Duration) {
	tck := time.NewTicker(interval)
	go func() {
		for {
			select {
			case t := <-tck.C:
				c.mtx.Lock()
				for key, entry := range c.entries {
					if t.After(entry.createdAt.Add(interval)) {
						delete(c.entries, key)
					}
				}
				c.mtx.Unlock()
			}
		}
	}()
}
