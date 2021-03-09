package main

import (
	"context"
	"fmt"
	"sync/atomic"
)

func main() {

	var (
		count    int32 = 1
		maxCount int32 = 100
	)
	catChan, dogChan, fishChan := make(chan struct{}, 1), make(chan struct{}, 1), make(chan struct{}, 1)
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case _, ok := <-catChan:
				if count <= maxCount && ok {
					fmt.Println("cat", count)
					atomic.AddInt32(&count, 1)
					if count > maxCount {
						fmt.Println("cat return")
						return
					}
					dogChan <- struct{}{}
				}
			case <-ctx.Done():
				fmt.Println("cat exit by ctx")
				return
			}
		}
	}(ctx)

	go func(ctx context.Context) {
		for {
			select {
			case _, ok := <-dogChan:
				if count <= maxCount && ok {
					fmt.Println("dog", count)
					atomic.AddInt32(&count, 1)
					if count > maxCount {
						fmt.Println("dog return")
						return
					}
					fishChan <- struct{}{}
				}
			case <-ctx.Done():
				fmt.Println("dog exit by ctx")
				return
			}
		}
	}(ctx)

	go func(ctx context.Context) {
		for {
			select {
			case _, ok := <-fishChan:
				if count <= maxCount && ok {
					fmt.Println("fish", count)
					atomic.AddInt32(&count, 1)
					if count > maxCount {
						fmt.Println("fish return")
						return
					}
					catChan <- struct{}{}
				}
			case <-ctx.Done():
				fmt.Println("fish exit by ctx")
				return
			}
		}
	}(ctx)

	catChan <- struct{}{}

	for {
		if count > maxCount {
			cancel()
			close(catChan)
			close(dogChan)
			close(fishChan)
			return
		}
	}
}
