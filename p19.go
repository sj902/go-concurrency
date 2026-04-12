package main

import "sync"

type concmap struct {
	shards []shard
}

type shard struct {
	start int
	end   int
	mutex sync.RWMutex
	m     map[int]int
}

func startCMap() concmap {
	cmap := concmap{
		shards: []shard{
			{1, 25, sync.RWMutex{}, make(map[int]int)},
			{26, 50, sync.RWMutex{}, make(map[int]int)},
			{51, 75, sync.RWMutex{}, make(map[int]int)},
			{76, 100, sync.RWMutex{}, make(map[int]int)},
		},
	}
	return cmap
}

func (c *concmap) Set(key, value int) {
	i := 3
	if key < 26 {
		i = 0

	} else if key < 51 {
		i = 1
	} else if key < 76 {
		i = 2
	}

	c.shards[i].mutex.Lock()
	defer c.shards[i].mutex.Unlock()
	c.shards[i].m[key] = value
}

func (c *concmap) Get(key int) (int, bool) {
	i := 3
	if key < 26 {
		i = 0

	} else if key < 51 {
		i = 1
	} else if key < 76 {
		i = 2
	}

	c.shards[i].mutex.RLock()
	defer c.shards[i].mutex.RUnlock()
	a, ok := c.shards[i].m[key]
	return a, ok
}
