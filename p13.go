package main

import "fmt"

func p13() {
	produced := make(chan int)
	squared := make(chan int)
	filtered := make(chan int)

	go produce(produced)
	go sq(produced, squared)
	go filter(squared, filtered)

	for {
		i, ok := <-filtered

		if !ok {
			return
		}
		fmt.Printf("Val: %v \n", i)
	}

}

func produce(p chan<- int) {
	defer close(p)
	for i := 1; i <= 9; i++ {
		p <- i
	}
}

func sq(p <-chan int, squared chan<- int) {
	for {
		i, ok := <-p

		if !ok {
			close(squared)
			return
		}
		squared <- i * i
	}
}

func filter(p <-chan int, f chan<- int) {
	for {
		i, ok := <-p

		if !ok {
			close(f)
			return
		}
		if i%2 == 0 {
			f <- i
		}
	}
}
