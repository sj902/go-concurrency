/*
**Problem #4 — Parallel File Processor**

You have a slice of 5 filenames (just strings, no actual disk IO needed — simulate processing with `fmt.Println`). Process all 5 files concurrently using goroutines. Use a **Mutex to track a shared count** of how many files have been processed. Print the final count after all goroutines finish.

**Concepts tested:** `sync.WaitGroup` + `sync.Mutex` as two separate concerns — WaitGroup for lifecycle, Mutex for shared state protection.

**Constraint:** No channels. 15-minute timer. Post your solution when ready.
*/

package main

import (
	"fmt"
	"sync"
)

type count struct {
	ctr int
	m   sync.Mutex
}

func p4() {
	strings := []string{"1", "2", "3", "4", "5"}
	wg := sync.WaitGroup{}
	c := count{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go process(strings[i], &c, &wg)
	}
	wg.Wait()
	fmt.Printf("count %v \n", c.ctr)
}

func process(str string, c *count, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(str)
	c.m.Lock()
	c.ctr = c.ctr + 1
	c.m.Unlock()
}
