// 模拟运维发布系统：目标10000物理服务器，更新应用（print hello world），发布策略最高并发20.
// 1. 协程
// 2. chan
// 3. 并发控制
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	maxGoroutineCount = 20
	machineCount      = 10000
)

func main() {
	var (
		queue = make(chan struct{}, maxGoroutineCount)
		wg    sync.WaitGroup
		count int32
	)

	go func() {
		for range queue {
			go func() {
				defer wg.Done()
				fmt.Println("hello world")
				atomic.AddInt32(&count, 1)
			}()
		}
		fmt.Println("quene close")
	}()

	for i := 0; i < machineCount; i++ {
		wg.Add(1)
		queue <- struct{}{}
	}
	wg.Wait()
	close(queue)
	fmt.Println(count)

}
