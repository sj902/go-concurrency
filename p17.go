package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func p17() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		select {
		case <-ctx.Done():
			return nil // cancelled by another task
		case <-time.After(1 * time.Second):
			fmt.Println("task done")
			return nil
		}
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
			return nil // cancelled by another task
		case <-time.After(1 * time.Second):
			fmt.Println("task done")
			return nil
		}
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
			return nil // cancelled by another task
		case <-time.After(1 * time.Second):
			return errors.New("task failed")
		}
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
			return nil // cancelled by another task
		case <-time.After(100 * time.Millisecond):
			fmt.Println("error")
			return errors.New("Error!!")
		}
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
			return nil // cancelled by another task
		case <-time.After(1 * time.Second):
			fmt.Println("task done")
			return nil
		}
	})

	e := g.Wait()
	if e != nil {
		fmt.Println("error:", e)
	}
}
