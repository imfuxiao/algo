package linkedlist

import (
	"errors"
	"fmt"
	"reflect"
)

type element struct {
	value interface{}
	prev  *element
	next  *element
}

type List struct {
	head  *element
	trail *element
	size  int
}

func New(v ...interface{}) *List {
	// head与trail为哨兵节点, 不用做nil的判断
	head := &element{}
	trail := &element{}
	head.next = trail
	trail.prev = head
	l := &List{
		head:  head,
		trail: trail,
		size:  0,
	}
	if len(v) > 0 {
		l.Add(v...)
	}
	return l
}

func (l *List) Insert(index int, v interface{}) error {
	if !l.withinRange(index) {
		return errors.New("index is wrong")
	}

	if index == l.size {
		l.Add(v)
		return nil
	}

	e := &element{
		value: v,
	}

	for i, next := 0, l.head.next; ; i++ {
		if i != index {
			next = next.next
			continue
		}
		e.prev = next.prev
		e.next = next
		next.prev.next = e
		next.prev = e
		break
	}
	return nil
}

func (l *List) Add(v ...interface{}) {
	for _, value := range v {
		e := &element{
			value: value,
			prev:  l.trail.prev,
			next:  l.trail,
		}
		l.trail.prev.next = e
		l.trail.prev = e
		l.size++
	}
}

// 从head添加, 比如: [c, d] -> Prepend(a, b) -> [a, b, c, d]
func (l *List) Prepend(v ...interface{}) {
	for i := len(v) - 1; i >= 0; i-- {
		e := &element{
			value: v[i],
		}
		e.prev = l.head
		e.next = l.head.next
		l.head.next.prev = e
		l.head.next = e
	}
}

// 通过index获取值
func (l *List) Get(index int) (interface{}, bool) {
	if !l.withinRange(index) {
		return nil, false
	}
	var (
		value  interface{}
		result = false
	)

	for i, e := 0, l.head.next; i < l.size; i++ {
		if i == index {
			value = e.value
			result = true
			break
		}
	}

	return value, result
}

//
func (l *List) Remove(index int) error {
	if !l.withinRange(index) {
		return errors.New("index is wrong")
	}
	for i, e := 0, l.head.next; i < l.size; i++ {
		if i == index {
			l.size--
			e.prev.next = e.next
			e.next.prev = e.prev
			e.next = nil
			e.prev = nil
			e.value = nil
			break
		}
		e = e.next
	}
	return nil
}

func (l *List) Contain(v interface{}) bool {
	for i, e := 0, l.head.next; i < l.size; i++ {
		if reflect.DeepEqual(e.value, v) {
			return true
		}
		e = e.next
	}
	return false
}

// 获取值所在的index
func (l *List) IndexOf(v interface{}) int {
	index := -1
	for i, e := 0, l.head.next; i < l.size; i++ {
		if reflect.DeepEqual(e.value, v) {
			index = i
			break
		}
		e = e.next
	}
	return index
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Clean() {
	l.size = 0
	l.head.next = l.trail
	l.trail.prev = l.head
}

func (l *List) String() string {
	var values []interface{}
	for i, e := 0, l.head.next; i < l.size; i++ {
		values = append(values, e.value)
		e = e.next
	}
	return fmt.Sprintf("%+v", values)
}

// Check that the index is within bounds of the list
func (l *List) withinRange(index int) bool {
	return index >= 0 && index < l.size
}
