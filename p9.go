package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func p9() {
	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}
	wg.Add(3)
	go a(ctx, &wg)
	go a(ctx, &wg)
	go a(ctx, &wg)
	time.Sleep(2 * time.Second)
	cancel()
	wg.Wait()
}

func a(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker exiting")
			return
		default:
			fmt.Println("working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
