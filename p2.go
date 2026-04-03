/*
All 5 goroutines likely print `5` — because they close over the **same `i` variable**, and by the time they run, the loop has finished and `i` is 5. That's the trap.

Your solution avoids this correctly by passing `i` as an argument. Good — you understood it.

---

**Problem #2 — Ping-Pong**

Two goroutines, one channel. Goroutine A sends "ping", Goroutine B receives it and sends back "pong". Do this 5 times. Main prints each ping and pong in order.

**Concept tested:** Unbuffered channel as a rendezvous point — both sender and receiver must be ready simultaneously.

**Constraint:** No `time.Sleep`, no WaitGroup needed. 15-minute timer. Post your solution when ready.
*/

package main

import (
	"fmt"
)

func p2() {
	a := make(chan int)
	go ping(a)
	pong(a)
}

func ping(a chan int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Ping %v \n", i)
		a <- i
	}
	close(a)
}

func pong(a chan int) {
	for {
		msg, ok := <-a
		if !ok {
			break
		}
		fmt.Printf("Pong %v\n", msg)
	}
}
