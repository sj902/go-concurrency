package main

import "sync"

type Cache struct {
	mutex    sync.RWMutex
	cacheMap map[string]string
}

func (c *Cache) Init() {
	cacheMap := make(map[string]string)
	c.cacheMap = cacheMap
}

func (c *Cache) Set(key, value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cacheMap[key] = value
}

func (c *Cache) Get(key string) (string, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	a, ok := c.cacheMap[key]
	return a, ok
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.cacheMap, key)
}
