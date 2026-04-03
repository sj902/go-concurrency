package main

import "fmt"

func p7() {
	merged := make(chan string)
	a := make(chan string)
	b := make(chan string)
	go merge(a, b, merged)
	go ping1(a)
	go pong1(b)

	for {
		z, ok := <-merged
		if !ok {
			break
		}
		fmt.Printf("%v\n", z)
	}
}

func ping1(a chan<- string) {
	defer close(a)
	for i := 0; i < 5; i++ {
		a <- "ping"
	}
}
func pong1(b chan<- string) {
	defer close(b)
	for i := 0; i < 5; i++ {
		b <- "pong"
	}
}

func merge(a <-chan string, b <-chan string, c chan<- string) {
	q := true
	w := true
	for {

		select {
		case z, aok := <-a:
			if !aok {
				q = false
			} else {
				c <- z
			}

		case x, bok := <-b:
			if !bok {
				w = false
			} else {
				c <- x
			}
		}
		if !q && !w {
			close(c)
			break
		}
	}
}
