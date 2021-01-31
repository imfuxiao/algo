// 641. 设计循环双端队列
// 设计实现双端队列。
// 你的实现需要支持以下操作：
//
// MyCircularDeque(k)：构造函数,双端队列的大小为k。
// insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
// insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
// deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
// deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
// getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
// getRear()：获得双端队列的最后一个元素。如果双端队列为空，返回 -1。
// isEmpty()：检查双端队列是否为空。
// isFull()：检查双端队列是否满了。
//
// 示例:
// MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
// circularDeque.insertLast(1);			        // 返回 true
// circularDeque.insertLast(2);			        // 返回 true
// circularDeque.insertFront(3);			        // 返回 true
// circularDeque.insertFront(4);			        // 已经满了，返回 false
// circularDeque.getRear();  				// 返回 2
// circularDeque.isFull();				        // 返回 true
// circularDeque.deleteLast();			        // 返回 true
// circularDeque.insertFront(4);			        // 返回 true
// circularDeque.getFront();				// 返回 4
//
// 提示:
// 所有值的范围为 [1, 1000]
// 操作次数的范围为 [1, 1000]
// 请不要使用内置的双端队列库。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/design-circular-deque
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package queue

type MyCircularDeque struct {
	elements      []int
	head          int
	trail         int
	elementsCount int
	cap           int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		elements:      make([]int, k),
		head:          0,
		trail:         0,
		elementsCount: 0,
		cap:           k,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	// enough
	if this.elementsCount == this.cap {
		return false
	}
	if this.elementsCount == 0 || this.head > 0 {
		if this.head > 0 {
			this.head--
		}
		this.elements[this.head] = value
		this.elementsCount++
		return true
	}
	// 全体向后挪
	for i := this.trail; i >= this.head; i-- {
		this.elements[i], this.elements[i+1] = this.elements[i+1], this.elements[i]
	}
	this.trail++
	this.elementsCount++
	this.elements[this.head] = value
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	// enough
	if this.elementsCount == this.cap {
		return false
	}
	if this.elementsCount == 0 {
		this.elements[this.trail] = value
		this.elementsCount++
		return true
	}

	if this.trail+1 < this.cap {
		this.trail++
		this.elements[this.trail] = value
		this.elementsCount++
		return true
	}

	// 全体向前挪
	for i := this.head; i <= this.trail; i++ {
		this.elements[i], this.elements[i-1] = this.elements[i-1], this.elements[i]
	}
	this.head--
	this.elementsCount++
	this.elements[this.trail] = value
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.elementsCount == 0 {
		return false
	}
	this.head++
	this.elementsCount--
	if this.elementsCount == 0 && this.trail <= this.head {
		this.head, this.trail = 0, 0
	}
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.elementsCount == 0 {
		return false
	}
	this.trail--
	this.elementsCount--
	if this.elementsCount == 0 && this.trail <= this.head {
		this.head, this.trail = 0, 0
	}
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.elementsCount == 0 {
		return -1
	}
	//value := this.elements[this.head]
	//this.head++
	//this.elementsCount--
	//if this.elementsCount == 0 && this.head >= this.trail {
	//	this.head, this.trail = 0, 0
	//}
	return this.elements[this.head]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.elementsCount == 0 {
		return -1
	}
	//value := this.elements[this.trail]
	//this.trail--
	//this.elementsCount--
	//if this.elementsCount == 0 && this.trail <= this.head {
	//	this.head, this.trail = 0, 0
	//}
	return this.elements[this.trail]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.elementsCount == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.elementsCount == this.cap
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
