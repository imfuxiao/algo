package tree

import (
	"fmt"
	"reflect"
)

type NodeInterface interface {
	Key() int
	Value() interface{}
}

type intNode int

func (i intNode) Key() int {
	return int(i)
}

func (i intNode) Value() interface{} {
	return int(i)
}

type Node struct {
	value     NodeInterface
	leftNode  *Node
	rightNode *Node
}

func NewIntNode(value int) *Node {
	v := intNode(value)
	return &Node{
		value:     v,
		leftNode:  nil,
		rightNode: nil,
	}
}

func (n *Node) Value() interface{} {
	return n.value.Value()
}

func (n *Node) SetLeftNode(left *Node) {
	n.leftNode = left
}

func (n *Node) LeftNode() *Node {
	return n.leftNode
}

func (n *Node) SetRightNode(right *Node) {
	n.rightNode = right
}

func (n *Node) RightNode() *Node {
	return n.rightNode
}

type BinaryTree struct {
	rootNode *Node
}

func NewBinaryTree(root *Node) *BinaryTree {
	return &BinaryTree{
		rootNode: root,
	}
}

func (b *BinaryTree) Root() *Node {
	return b.rootNode
}

// 前序遍历: 先打印自己, 在左, 在右
func (b *BinaryTree) PreNode() {
	preNode(b.rootNode)
}

// 前序遍历: 先打印自己, 在左, 在右
func preNode(n *Node) {
	if n == nil {
		return
	}
	fmt.Printf("1. 先打印自己: %+v\n", n.value)
	preNode(n.leftNode)
	preNode(n.rightNode)
}

// 中序遍历: 先打印左节点, 在打印自己, 在打印右节点
func (b *BinaryTree) InNode() {
	inNode(b.rootNode)
}

// 中序遍历: 先打印左节点, 在打印自己, 在打印右节点
func inNode(n *Node) {
	if n == nil {
		return
	}
	inNode(n.leftNode)
	fmt.Printf("2. 打印自己: %+v\n", n.value)
	inNode(n.rightNode)
}

// 后续遍历: 先打印左节点, 在打印右节点, 在打印自己
func (b *BinaryTree) PostNode() {
	postNode(b.rootNode)
}

func postNode(n *Node) {
	if n == nil {
		return
	}
	postNode(n.leftNode)
	postNode(n.rightNode)
	fmt.Printf("3. 打印自己: %+v\n", n.value)
}

type BinarySearchTree struct {
	root *Node
}

func NewBinarySearchTree(n *Node) *BinarySearchTree {
	return &BinarySearchTree{
		root: n,
	}
}

func (b *BinarySearchTree) String() string {
	return fmt.Sprintf("\n%+v\n", b.Values())
}

func (b *BinarySearchTree) Values() []interface{} {
	s := make([]interface{}, 0, 100)
	return in(s, b.root)
}

// 中序遍历
func in(s []interface{}, n *Node) []interface{} {
	if n != nil {
		s = in(s, n.leftNode)
		s = append(s, n.value.Value())
		s = in(s, n.rightNode)
	}
	return s
}

func (b *BinarySearchTree) Insert(n *Node) {
	if b.root == nil {
		b.root = n
		return
	}

	p := b.root
	for {
		if n.value.Key() > p.value.Key() {
			if p.rightNode == nil {
				p.rightNode = n
				return
			}
			p = p.rightNode
		} else { // 不考虑相同元素
			if p.leftNode == nil {
				p.leftNode = n
				return
			}
			p = p.leftNode
		}
	}

}

func (b *BinarySearchTree) Search(n *Node) interface{} {
	if n == nil {
		return nil
	}
	p := b.root
	if p == nil {
		return nil
	}
	for {
		if reflect.DeepEqual(n.Value(), p.Value()) {
			return n.Value()
		}
		if n.value.Key() > p.value.Key() {
			p = p.rightNode
		} else {
			p = p.leftNode
		}
		if p == nil {
			return nil
		}
	}
}

// 删除一个元素有三种情况:
// 1. 代删除节点为为叶子节点, 则直接把父节点中的指向指针的值设置为nil
// 2. 代删除节点只有一个子节点(无论是左子节点, 还是右子节点), 只需要更新父节点, 把指针指向代删除子节点的子节点(无论是左子节点还是右子节点).
// 3. 代删除节点有两个子节点. 这个比较复杂
// 查找代删除节点的右子节点, 找到最小的那个节点, 这个节点一定没有左子节点(如果有一定不是最小节点)
// 找到这个节点后, 用它替换代删除节点
func (b *BinarySearchTree) Delete(n *Node) {
	if n == nil {
		return
	}
	if b.root == nil {
		return
	}

	// 查找代删除节点, node 为代删除节点, parent为代删除节点父节点
	node, parent, isLeft := b.root, b.root, false
	for {
		if reflect.DeepEqual(node.Value(), n.Value()) {
			break
		}
		if n.value.Key() > node.value.Key() {
			node, parent, isLeft = node.rightNode, node, false
		} else {
			node, parent, isLeft = node.leftNode, node, true
		}
		if node == nil {
			break
		}
	}

	if node == nil {
		return
	}

	// 如果代删除节点为叶子节点, 则直接删除
	if node.leftNode == nil && node.rightNode == nil {
		if isLeft {
			parent.leftNode = nil
		} else {
			parent.rightNode = nil
		}
		return
	}

	// 如果带删除节点存在两个子节点, 则需要查找右子树中最小的节点, 用这个节点代替代删除节点
	if node.leftNode != nil && node.rightNode != nil {
		// 查找右子树中最小的节点
		// minNode: 右子树最小节点
		// minNodeParent: minNode节点的父节点
		minNode, minNodeParent := node.rightNode, node
		for {
			if minNode.leftNode != nil {
				minNode, minNodeParent = minNode.leftNode, minNode
			} else {
				break
			}
		}

		// 如果最小节点存在右子节点, 则将这个节点移动至父节点的左节点上(因为我们for循环始终取的是节点的左节点)
		if minNode.rightNode != nil {
			minNodeParent.leftNode = minNode.rightNode
			minNode.rightNode = nil
		}

		minNode.leftNode = node.leftNode
		if node.rightNode != minNode {
			minNode.rightNode = node.rightNode
		}

		if isLeft {
			parent.leftNode = minNode
		} else {
			parent.rightNode = minNode
		}
		return
	}

	// 如果节点只有一个子节点(无论是左子节点, 还是右子节点)
	if node.leftNode != nil {
		if isLeft {
			parent.leftNode = node.leftNode
		} else {
			parent.rightNode = node.leftNode
		}
	} else {
		if isLeft {
			parent.leftNode = node.rightNode
		} else {
			parent.rightNode = node.rightNode
		}
	}

}
