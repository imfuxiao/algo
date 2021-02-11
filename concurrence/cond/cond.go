package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})

	var ready int

	for i := 0; i < 10; i++ {
		go func(num int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

			// 加锁更改等待条件
			c.L.Lock()
			ready++
			c.L.Unlock()

			log.Printf("我是 %d 号运动员, 我准备好了.\n", num)

			// 广播通知所有等待者
			c.Broadcast()
		}(i)
	}

	c.L.Lock()
	i := 1
	//for ready != 10 { // 检查边界条件: 是否全部准备完毕
		c.Wait()
		log.Printf("裁判被唤醒 %d 次", i)
		i++
	//}
	c.L.Unlock()

	log.Println("所有运动员都准备完毕, 比赛开始: 3, 2, 1. go....")
}
