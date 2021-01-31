package main

import (
	"fmt"
	"sync"
)

func main() {
	var count = 0
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				count++ // 不是一个原子操作, 它至少包含几个步骤, 比如读取变量count的当前值, 对这个值加1, 在把值保存到count中. 因为不是原子操作, 就有可能有并发的问题.

			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
