/*
**Problem #5 — Thread-Safe Counter Struct**

Design a `Counter` struct that can be safely used from multiple goroutines. It must support three methods: `Increment()`, `Decrement()`, and `Value() int`. Spawn 100 goroutines — 50 call `Increment()`, 50 call `Decrement()`. Print the final value after all finish. Expected result: 0.

**Concepts tested:** `sync.Mutex` with pointer receivers, why value receivers break thread safety.

**Constraint:** No channels. 15-minute timer. Post your solution when ready.
*/

package main

import (
	"fmt"
	"sync"
)

type counter struct {
	count int
	m     sync.Mutex
}

func (c *counter) incr() {
	c.m.Lock()
	c.count = c.count + 1
	c.m.Unlock()
}

func (c *counter) decr() {
	c.m.Lock()
	c.count = c.count - 1
	c.m.Unlock()
}

func (c *counter) Value() int {
	c.m.Lock()
	defer c.m.Unlock()
	return c.count
}

func p5() {
	c := counter{}
	wg := sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.incr()
		}()
	}

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.decr()
		}()
	}

	wg.Wait()
	fmt.Printf("Count :%v\n", c.Value())
}
