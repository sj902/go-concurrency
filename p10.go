package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func p10() {
	ctx, cancelFunc := context.WithTimeout(context.WithValue(context.Background(), "reqId", 1234), 2*time.Second)
	defer cancelFunc()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go run10(ctx, &wg)
	wg.Wait()
}

func run10(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(500 * time.Millisecond)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("exiting:", ctx.Err())
			return

		case <-ticker.C:
			fmt.Printf("ID: %v", ctx.Value("reqId"))
		}
	}

}
