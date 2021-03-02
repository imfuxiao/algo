package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count int32

	var wg sync.WaitGroup
	wg.Add(3)

	c := make(chan int, 1)
	defer close(c)

	c <- 1

	go func() {
		for atomic.LoadInt32(&count) < 100 {
			select {
			case v := <-c:
				if v == 1 && atomic.LoadInt32(&count) < 100 {
					atomic.AddInt32(&count, 1)
					fmt.Printf("count: %d,  value: %d \n", atomic.LoadInt32(&count), v)
					c <- 2
				} else {
					fmt.Printf("    get error: count: %d,  value: %d \n", atomic.LoadInt32(&count), v)
					c <- v
				}
			}
		}
		wg.Done()
	}()

	go func() {
		for atomic.LoadInt32(&count) < 100 {
			select {
			case v := <-c:
				if v == 2 && atomic.LoadInt32(&count) < 100 {
					atomic.AddInt32(&count, 1)
					fmt.Printf("count: %d,  value: %d \n", atomic.LoadInt32(&count), v)
					c <- 3
				} else {
					c <- v
				}
			}
		}
		wg.Done()
	}()
	go func() {
		for atomic.LoadInt32(&count) < 100 {
			select {
			case v := <-c:
				if v == 3 && atomic.LoadInt32(&count) < 100 {
					atomic.AddInt32(&count, 1)
					fmt.Printf("count: %d,  value: %d \n", atomic.LoadInt32(&count), v)
					c <- 1
				} else {
					c <- v
				}
			}
		}
		wg.Done()
	}()
	wg.Wait()
}
