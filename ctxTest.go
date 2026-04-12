package main

import (
	"context"
	"fmt"
	"time"
)

func ctxTest() {
	ctx, cancel := context.WithTimeout(context.WithValue(context.Background(), "userID", 42), 1*time.Second)
	defer cancel()
	loop(ctx)

}

func loop(ctx context.Context) {
	ticker := time.NewTicker(200 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			fmt.Printf("UserId: %v", ctx.Value("userID"))
		case <-ctx.Done():
			fmt.Printf("Done %v", ctx.Err())
			return
		}
	}
}
