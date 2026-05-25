package cache

import "time"


type Cache struct {
	cache map[string]cacheEntry

}

type cacheEntry struct{
	createdAt time.Time
	val []byte
}

func NewCache() Cache {
	return Cache{
		cache: make(map[string]cacheEntry),
	}
}

func (c *Cache)Add( key string, val []byte){
	c.cache[key] = cacheEntry{
		val: val,
		createdAt: time.Now(),
	}
}

func (c *Cache)Get(key string)([]byte,bool){
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
	fiveMinAgo := time.Now().Add(-t)
	for key, entry := range c.cache {
			if (entry.createdAt.Before(fiveMinAgo)) {
				delete(c.cache, key)
			}
		}
	}
