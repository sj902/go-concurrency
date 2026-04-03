package main

import (
	"fmt"
	"time"
)

func p8() {
	a := make(chan string)
	go doWork(a)
	select {
	case f, ok := <-a:
		if !ok {
			fmt.Println("Done")
			return
		}
		fmt.Printf("Got %v \n", f)
		return
	case <-time.After(time.Second):
		fmt.Println("TImed out")
	}
}

func doWork(a chan<- string) {
	time.Sleep(2 * time.Second)
	close(a)
}
