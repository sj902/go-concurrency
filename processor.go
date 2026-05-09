package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

type ctxKey string

const (
	jobIDKey   ctxKey = "jobID"
	traceIDKey ctxKey = "traceID"
)

type Stats struct {
	succeeded int
	timedOut  int
	errored   int
	mu        sync.Mutex
}

func (s *Stats) IncSuccess() {
	s.mu.Lock()
	s.succeeded++
	s.mu.Unlock()
}

func (s *Stats) IncTimeout() {
	s.mu.Lock()
	s.timedOut++
	s.mu.Unlock()
}

func (s *Stats) IncError() {
	s.mu.Lock()
	s.errored++
	s.mu.Unlock()
}

func processJob(ctx context.Context, jobID int) error {
	traceID := ctx.Value(traceIDKey).(string)
	fmt.Printf("Starting job=%d traceID=%s\n", jobID, traceID)
	workTime := time.Duration(rand.IntN(100)+50) * time.Millisecond
	select {
	case <-time.After(workTime):
		fmt.Printf("Completed job=%d workTime=%v\n", jobID, workTime)
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func worker(ctx context.Context, jobs <-chan int, semaphore chan struct{}, stats *Stats) error {
	for {
		select {
		case <-ctx.Done():
			return nil // cancelled by errgroup, exit cleanly
		case jobID, ok := <-jobs:
			if !ok {
				return nil // jobs channel closed, exit cleanly
			}
			err := func() error {
				semaphore <- struct{}{}        // acquire — blocks if 2 already running
				defer func() { <-semaphore }() // release after work is done

				jobCtx, cancel := context.WithTimeout(ctx, 120*time.Millisecond)
				defer cancel()
				jobCtx = context.WithValue(jobCtx, jobIDKey, jobID)
				jobCtx = context.WithValue(jobCtx, traceIDKey, fmt.Sprintf("trace-%d", jobID))

				return processJob(jobCtx, jobID)
			}()

			if err == nil {
				stats.IncSuccess()
				continue
			}
			if errors.Is(err, context.DeadlineExceeded) {
				stats.IncTimeout()
				fmt.Printf("Timed out job=%d\n", jobID)
				if jobID == 13 {
					return fmt.Errorf("critical job 13 timed out") // triggers errgroup cancel
				}
				continue
			}
			if errors.Is(err, context.Canceled) {
				return nil // parent cancelled, exit cleanly
			}
			stats.IncError()
			fmt.Printf("Errored job=%d err=%v\n", jobID, err)
			continue
		}
	}
}

func processor() {
	jobs := make(chan int)
	stats := &Stats{}
	semaphore := make(chan struct{}, 2)

	g, ctx := errgroup.WithContext(context.Background())

	// producer
	g.Go(func() error {
		defer close(jobs)
		for i := 1; i <= 20; i++ {
			select {
			case <-ctx.Done():
				return nil
			case jobs <- i:
			}
		}
		return nil
	})

	// 4 workers
	for i := 0; i < 4; i++ {
		g.Go(func() error {
			return worker(ctx, jobs, semaphore, stats)
		})
	}

	err := g.Wait()
	if err != nil {
		fmt.Printf("\nProcessing stopped: %v\n", err)
	}

	fmt.Println("\n------ FINAL STATS ------")
	fmt.Printf("Succeeded : %d\n", stats.succeeded)
	fmt.Printf("Timed Out : %d\n", stats.timedOut)
	fmt.Printf("Errored   : %d\n", stats.errored)
}
