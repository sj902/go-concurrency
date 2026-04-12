package main

import "time"

type RateLimiter struct {
	c chan struct{}
}

func (rl *RateLimiter) start() {
	rl.c = make(chan struct{}, 5)
	for i := 0; i < 5; i++ {
		rl.c <- struct{}{}
	}

	go func(c chan struct{}) {
		ticker := time.NewTicker(200 * time.Millisecond)
		for {
			select {
			case <-ticker.C:
				{
					select {
					case rl.c <- struct{}{}: // add token
					default: // bucket full, skip
					}
				}
			}
		}
	}(rl.c)
}

func (rl *RateLimiter) consume() {
	<-rl.c
}
