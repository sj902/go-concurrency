package main

import (
	"fmt"
	"sync"
	"time"
)

type Config struct {
	i int
}

var (
	instance Config
	once     sync.Once
)

func initConf() {
	once.Do(func() {
		fmt.Printf("Starting.....")
		time.Sleep(100 * time.Millisecond)
		instance = Config{1}
	})
}

func p12() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go start(&wg, &once)
	}
	wg.Wait()
}

func start(wg *sync.WaitGroup, once *sync.Once) {
	defer wg.Done()
	initConf()
}
