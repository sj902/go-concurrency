package main

import (
	"fmt"
	"sync"
)

func p14() {
	jobs := make(chan int)

	go func() {
		defer close(jobs)
		for i := 0; i < 100; i++ {
			jobs <- i
		}
	}()

	var wg sync.WaitGroup

	results := make(chan int)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go processP14(jobs, results, &wg)
	}

	count := 0

	go func(){
		wg.Wait()
		close(results) 
	}()

	for i := range results {
		count = count + 1
		fmt.Printf("Output: %v \t Processed: %v \n", i, count)
	}

}

func processP14(jobs <-chan int, res chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		i, ok := <-jobs
		if !ok {
			return
		}
		res <- i * 2
	}
}
