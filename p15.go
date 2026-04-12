package main

import (
	"fmt"
	"sync"
	"time"
)

func p15() {
	bufChan := make(chan int, 3)

	var wg sync.WaitGroup

	wg.Add(20)

	for i := 0; i < 20; i++ {
		go func(jobId int) {
			defer wg.Done()
			defer func() {
				<-bufChan
			}()
			bufChan <- jobId
			fmt.Printf("job Id: %v \n", jobId)
			time.Sleep(100 * time.Millisecond)
		}(i)
	}

	wg.Wait()
}
