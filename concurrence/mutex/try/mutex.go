package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const (
	mutexLocked      = 1 << iota // 加锁标识的位置     001
	mutexWoken                   // 锁唤醒标识的位置   010
	mutexStaving                 // 锁饥饿标识的位置   100
	mutexWaiterShift = iota
)

type Mutex struct {
	sync.Mutex
}

func (m *Mutex) TryLock() bool {

	// fast path，如果幸运，没有其他 goroutine 争这把锁，那么，这把锁就会被这个请求的 goroutine 获取，直接返回。
	// 判断mutex是否已被加锁, 没有加锁, 即原来锁的标识位为0, 则返回true
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}

	// 如果锁已经被其他 goroutine 所持有，或者被其他唤醒的 goroutine 准备持有，那么，就直接返回 false，不再请求
	// 如果当前锁处于唤醒, 加锁或者饥饿状态任意状态, 那么这次请求就不参与竞争了
	oldStatus := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	if oldStatus&(mutexLocked|mutexWoken|mutexStaving) != 0 {
		return false
	}

	// 如果没有被持有，也没有其它唤醒的 goroutine 来竞争锁，锁也不处于饥饿状态，就尝试获取这把锁
	// 不论是否成功都将结果返回。因为，这个时候，可能还有其他的 goroutine 也在竞争这把锁
	// 所以，不能保证成功获取这把锁。
	// 尝试在竞争的状态下请求锁, oldStatus | mutexLocked 表示锁状态为已加锁
	newStatus := oldStatus | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), oldStatus, newStatus)
}

// 等待锁的所有Goroutine数量
func (m *Mutex) Count() int {
	status := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	count := status >> mutexWaiterShift
	return int(count + status&mutexLocked)
}

// 当前锁状态是否已被锁
func (m *Mutex) IsLocked() bool {
	status := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return status&mutexLocked == mutexLocked
}

// 当前锁状态是否有等待被唤醒的goroutine
func (m *Mutex) IsWoken() bool {
	status := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return status&mutexWoken == mutexWoken
}

// 当前锁是否处于饥饿模式
func (m *Mutex) IsStarving() bool {
	status := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return status&mutexStaving == mutexStaving
}

func main() {
	var mu Mutex

	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			time.Sleep(time.Duration(rand.Int63n(5)) * time.Second)
			mu.Unlock()
		}()
	}

	time.Sleep(time.Second * 1)

	ok := mu.TryLock() // 尝试获取锁
	if ok {
		fmt.Println("got the lock")
		mu.Unlock()
		return
	}

	// 在获取 state 字段的时候，并没有通过 Lock 获取这把锁，所以获取的这个 state 的值是一个瞬态的值，可能在你解析出这个字段之后，锁的状态已经发生了变化。不过没关系，因为你查看的就是调用的那一时刻的锁的状态。
	fmt.Println("could not get the lock")
	fmt.Println("wait goroutine count", mu.Count())
	fmt.Println("mutex isLock ", mu.IsLocked())
	fmt.Println("mutex isWoken  ", mu.IsWoken())
	fmt.Println("mutex isStarving  ", mu.IsStarving())
}
