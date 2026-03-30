/*
**Problem #3 — Job Queue with Buffered Channel**

You have 10 jobs (just integers 1–10). Send all of them into a buffered channel of size 3. Have 3 worker goroutines consume from the channel and print which job they processed. Main waits for all workers to finish.

**Concept tested:** Buffered vs unbuffered channels — a buffered channel lets the sender proceed without a receiver being ready, up to the buffer size. What happens when the buffer is full (backpressure)?

**Constraint:** No `time.Sleep`. Use `close()` to signal workers to stop. 15-minute timer. Post your solution when ready.
*/
package main

import (
	"fmt"
	"sync"
)

func p3()  {
	ch := make(chan int, 3)
	wg := sync.WaitGroup{}
	wg.Add(3)
    go run(ch, &wg)
	go run(ch, &wg)
	go run(ch, &wg)
    add(ch)
	wg.Wait()
}

func add(ch chan<- int)  {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func run(ch <-chan int, wg *sync.WaitGroup){
	defer wg.Done()
	for{
		msg, ok := <- ch

		if !ok {
			break
		}
		fmt.Printf("JOB %v", msg)
	}
}