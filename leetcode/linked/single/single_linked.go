package single

import (
	"errors"
	"reflect"
)

type Node struct {
	val  interface{}
	next *Node
}

func (n *Node) Value() interface{} {
	return n.val
}

func (n *Node) Next() *Node {
	return n.next
}

type LinkedNode struct {
	length int
	head   *Node
	trail  *Node
}

func NewLinkedNode() *LinkedNode {
	trail := &Node{
		next: nil,
	}
	h := &Node{
		next: trail,
	}
	return &LinkedNode{
		length: 0,
		head:   h,
		trail:  trail,
	}
}

func (l *LinkedNode) Length() int {
	return l.length
}

// 得到链表最后一个节点, 如果链表为空, 则返回头节点
func (l *LinkedNode) getLastNode() *Node {
	if l.length == 0 {
		return l.head
	}
	lastNode := l.head
	for lastNode.next != l.trail {
		lastNode = lastNode.next
	}
	return lastNode
}

// 获取给定节点的前一个节点
func (l *LinkedNode) getPreviousNode(node *Node) *Node {
	h := l.head
	for h.next != node {
		h = h.next
	}
	return h
}

func (l *LinkedNode) Add(values ...interface{}) *LinkedNode {
	lastNode := l.getLastNode()
	for i := range values {
		n := &Node{
			val:  values[i],
			next: lastNode.next,
		}
		lastNode.next = n
		lastNode = n
		l.length++
	}
	return l
}

func (l *LinkedNode) Find(value interface{}) *Node {
	h := l.head.next
	for h != nil {
		if reflect.DeepEqual(h.Value(), value) {
			return h
		}
		h = h.next
	}
	return nil
}

func (l *LinkedNode) FindByNode(node *Node) *Node {
	h := l.head.next
	for h != nil {
		if h == node {
			return node
		}
		h = h.next
	}
	return nil
}

func (l *LinkedNode) Delete(node *Node) error {
	if node == nil {
		return nil
	}

	n := l.FindByNode(node)
	if n == nil {
		return errors.New("node not found")
	}
	previousNode := l.getPreviousNode(n)
	previousNode.next = n.next
	l.length--
	return nil
}

// 链表反转
func (l *LinkedNode) Invert() *LinkedNode {
	var prev *Node
	last := l.head.next
	h := l.head.next
	for h != l.trail {
		next := h.next
		h.next = prev
		prev = h
		h = next
	}
	l.head.next = prev
	last.next = l.trail
	return l
}

func print(l *LinkedNode) []interface{} {
	var res []interface{}
	h := l.head
	for h.next != l.trail {
		res = append(res, h.next.Value())
		h = h.next
	}
	return res
}
