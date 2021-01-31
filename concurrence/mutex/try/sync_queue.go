package main

import (
	"errors"
	"sync"
)

type SyncSliceQueue struct {
	mu           sync.Mutex
	elements     []interface{}
	elementCount int
	cap          int
}

func NewSyncSliceQueue(cap int) *SyncSliceQueue {
	return &SyncSliceQueue{
		elements:     make([]interface{}, cap),
		elementCount: 0,
		cap:          cap,
	}
}

func (q *SyncSliceQueue) Enqueue(element interface{}) error {
	if q.elementCount == q.cap {
		return errors.New("sync_queue is full")
	}
	q.mu.Lock()
	q.elements[q.elementCount] = element
	q.elementCount++
	q.mu.Unlock()
	return nil
}

func (q *SyncSliceQueue) Dequeue() (interface{}, error) {
	if q.elementCount == 0 {
		return nil, errors.New("sync_queue is empty")
	}
	q.mu.Lock()
	value := q.elements[q.elementCount-1]
	q.elements[q.elementCount-1] = nil
	q.elementCount--
	q.mu.Unlock()
	return value, nil
}
