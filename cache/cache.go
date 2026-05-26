package cache

import (
	"sync"
	"time"
)


type Cache struct {
	cache map[string]cacheEntry
	mu sync.Mutex
}

type cacheEntry struct{
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	chc:=  &Cache{
		cache: make(map[string]cacheEntry),
	}
	go chc.LoopDelete(interval)  
	return chc
}

func (c *Cache)Add( key string, val []byte){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		val: val,
		createdAt: time.Now(),
	}
}

func (c *Cache)Get(key string)([]byte,bool){
	c.mu.Lock()
	defer c.mu.Unlock()
	entries, ok := c.cache[key]

	return entries.val, ok
 
}

func (c *Cache)LoopDelete(t time.Duration){
	ticker := time.NewTicker(t)
	// defer ticker.Stop()
	for range ticker.C {
			c.Delete(t)
	}
}



func (c *Cache)Delete(t time.Duration){
	c.mu.Lock()
	defer c.mu.Unlock()
	fiveMinAgo := time.Now().Add(-t)
	for key, entry := range c.cache {
			if (entry.createdAt.Before(fiveMinAgo)) {
				delete(c.cache, key)
			}
		}
	}
