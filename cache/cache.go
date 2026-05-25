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