package queue

import "fmt"

var DefaultCap = 32

type ArrayQueue struct {
	elements []interface{}
	head     int
	tail     int
	cap      int
}

func New(cap int) *ArrayQueue {
	queueCap := cap
	if queueCap <= 0 {
		queueCap = DefaultCap
	}

	q := &ArrayQueue{
		elements: make([]interface{}, 0, queueCap),
		head:     0,
		tail:     0,
		cap:      queueCap,
	}

	return q
}

func (q *ArrayQueue) EnQueue(v interface{}) bool {

	// 如果队列满了, 则不能入队
	if q.head+q.tail == q.cap {
		return false
	}

	// 当队列尾指针等于队列容量时 且 队列头指针不为0时, 队列需要做数据搬迁
	if q.tail == q.cap && q.head != 0 {
		for i := q.head; i < q.tail; i++ {
			q.elements[i-q.head] = q.elements[i]
		}
		q.tail = q.tail - q.head
		q.head = 0
		q.elements = q.elements[q.head:q.tail]
	}

	q.elements = append(q.elements, v)
	q.tail++

	return true
}

func (q *ArrayQueue) DeQueue() interface{} {

	// 相等表示队列中没有数据
	if q.head == q.tail {
		return nil
	}

	e := q.elements[q.head]
	q.elements[q.head] = nil
	q.head++

	return e
}

func (q *ArrayQueue) String() string {
	return fmt.Sprintf("%+v", q.elements[q.head:q.tail])
}
