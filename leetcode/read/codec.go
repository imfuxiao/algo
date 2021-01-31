package read

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
	data []string
}

func ConstructorCodec() Codec {
	return Codec{
		data: make([]string, 0, 100),
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var (
		queue = make([]*TreeNode, 0, 100)
	)
	enqueue := func(node *TreeNode) {
		queue = append(queue, node)
	}
	dequeue := func() *TreeNode {
		node := queue[0]
		queue = queue[1:]
		return node
	}

	if root != nil {
		enqueue(root)
	}

	for len(queue) > 0 {
		for i, length := 0, len(queue); i < length; i++ {
			node := dequeue()
			if node != nil {
				this.data = append(this.data, strconv.Itoa(node.Val))
				enqueue(node.Left)
				enqueue(node.Right)
			} else {
				this.data = append(this.data, "null")
			}
		}
	}

	// 从后往前过滤null值
	for len(this.data) > 0 {
		if this.data[len(this.data)-1] == "null" {
			this.data = this.data[:len(this.data)-1]
		} else {
			break
		}
	}

	var sb strings.Builder
	sb.WriteString("[")
	for i := range this.data {
		sb.WriteString(this.data[i])
		if i+1 != len(this.data) {
			sb.WriteString(",")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var (
		//dataLength = len(data)
		sb        strings.Builder
		dataAarry []string
		queue     = make([]*TreeNode, 0, 100)
		root      *TreeNode
	)
	for i := range data {
		if data[i] == '[' || data[i] == ']' {
			continue
		}
		sb.WriteByte(data[i])
	}

	enqueue := func(node *TreeNode) {
		queue = append(queue, node)
	}
	dequeue := func() *TreeNode {
		node := queue[0]
		queue = queue[1:]
		return node
	}

	dataAarry = strings.FieldsFunc(sb.String(), func(r rune) bool {
		return r == ','
	})

	dataArrayIndex := 0
	if len(dataAarry) > 0 && dataAarry[0] != "null" && dataAarry[0] != "" {
		if value, err := strconv.Atoi(dataAarry[dataArrayIndex]); err == nil {
			root = &TreeNode{Val: value}
			enqueue(root)
			dataArrayIndex++
		}

		for len(queue) > 0 {
			for i, length := 0, len(queue); i < length; i++ {
				node := dequeue()
				if dataArrayIndex < len(dataAarry) && dataAarry[dataArrayIndex] != "null" {
					if value, err := strconv.Atoi(dataAarry[dataArrayIndex]); err == nil {
						leftNode := &TreeNode{Val: value}
						node.Left = leftNode
						enqueue(leftNode)
					}
				}
				dataArrayIndex++
				if dataArrayIndex < len(dataAarry) && dataAarry[dataArrayIndex] != "null" {
					if value, err := strconv.Atoi(dataAarry[dataArrayIndex]); err == nil {
						rightNode := &TreeNode{Val: value}
						node.Right = rightNode
						enqueue(rightNode)
					}
				}
				dataArrayIndex++
			}
		}
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
