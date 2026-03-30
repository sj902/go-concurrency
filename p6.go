/*
**Problem #6 — Typed Channel Pipeline**

Write a three-function pipeline: `generate` produces integers 1–5 into a channel, `square` reads from that channel and writes squared values into a second channel, `print` reads from the second channel and prints each value.

Each function must use **directional channels** — `chan<- int` or `<-chan int` — not bidirectional `chan int`.

**Concepts tested:** `chan<- T` vs `<-chan T`, how directional channels enforce ownership and prevent misuse at compile time.

**Constraint:** 15-minute timer. Post your solution when ready.
*/

package main

import (
	"fmt"
	"sync"
)

func p6()  {
	input := make(chan int) 
	sq := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go generate(input)
	go square(input, sq)
	go print(sq, &wg)
	wg.Wait()
}

func print(sq <-chan int, wg *sync.WaitGroup)  {
	defer wg.Done()
	for{
		msg, ok := <- sq
		if(!ok){
			return
		}
		fmt.Printf("Recd. %v\n", msg)
	}
}

func square(input <-chan int, sq chan<- int)  {
	defer close(sq)
	for{
		msg, ok := <- input
		if(!ok){
			return
		}
		sq <- msg*msg
	}
}

func generate(input chan<- int){
	defer close(input)
	count := 5
	for i := 0; i < count; i++ {
		input <- i+1
	}
}

