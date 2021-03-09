package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var count int32 = 1
	var wg sync.WaitGroup
	catChan, dogChan, fishChan := make(chan struct{}, 1), make(chan struct{}, 1), make(chan struct{}, 1)
	go func(ctx context.Context) {
		for {
			select {
			case <-catChan:
				fmt.Println("cat")
				dogChan <- struct{}{}
			case <-ctx.Done():
				fmt.Println("cat goroutine exit")
				return
			}
		}
	}(ctx)
	go func(ctx context.Context) {
		for {
			select {
			case <-dogChan:
				fmt.Println("dog")
				fishChan <- struct{}{}
			case <-ctx.Done():
				fmt.Println("dog goroutine exit")
				return
			}
		}
	}(ctx)
	go func(ctx context.Context) {
		for {
			select {
			case <-fishChan:
				fmt.Println("fish")
				fmt.Println("count", count)
				atomic.AddInt32(&count, 1)
				wg.Done()
			case <-ctx.Done():
				fmt.Println("fish goroutine exit")
				return
			}
		}
	}(ctx)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		catChan <- struct{}{}
		wg.Wait()
	}

	cancel()
	close(catChan)
	close(dogChan)
	close(fishChan)
}
