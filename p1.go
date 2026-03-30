/*
**Problem #1 — Hello Goroutines**

Launch 5 goroutines, each printing its index (0–4). The main goroutine must wait for all of them to finish before exiting.

**Concepts tested:** `go` keyword, `sync.WaitGroup`, the loop variable capture trap.

**Constraint:** Solve it without looking anything up. Set a 15-minute timer. Post your solution when ready.

*/

package main

import (
	"fmt"
	"sync"
)

func p1() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i += 1 {
		wg.Add(1)
		go bb(i, &wg)
	}
	wg.Wait()
}

func bb(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(i)
}
