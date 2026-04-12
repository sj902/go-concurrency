package main

import (
	"fmt"
	"sync"
	"time"
)

func p16() {
	var wg sync.WaitGroup
	wg.Add(3)
	c := make(chan struct{})
	for i := 0; i < 3; i++ {
		go func(c <-chan struct{}) {
			defer wg.Done()
			ticker := time.NewTicker(500 * time.Millisecond)
			defer ticker.Stop()
			for {
				select {
				case <-c:
					fmt.Println("Got close")
					return
				case <-ticker.C:
					fmt.Println("Sleeping")
				}
			}

		}(c)
	}

	time.Sleep(3000 * time.Millisecond)

	close(c)
	wg.Wait()

}
