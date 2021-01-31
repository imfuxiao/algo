package main

import (
	"fmt"
	"sync"
)

func main() {
	// Mutex 的零值是还没有 goroutine 等待的未加锁的状态，所以你不需要额外的初始化，直接声明变量（如 var mu sync.Mutex）即可。
	var mux sync.Mutex

	var count = 0

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				mux.Lock()
				count++
				mux.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
