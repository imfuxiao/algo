package trie_tree

import "strings"

const (
	a      byte = 'a'
	length      = 26
)

// 我们的trie只包含 a~z 26个字母
// Trie是多叉树
type TrieNode struct {
	data      byte
	isEndChar bool // 是否为结束字符
	// 可以通过 字母 - a 快速获取其在树中的位置
	children []*TrieNode // 多个子节点
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		data:     '/',
		children: make([]*TrieNode, length),
	}
}

func getIndex(char byte) int {
	return int(char) - int(a)
}

func (n *TrieNode) Insert(str string) bool {
	str = strings.ToLower(str)
	p := n
	for i := range str {
		index := getIndex(str[i])
		// 不在字符范围内
		if index < 0 || index > 26 {
			return false
		}
		if p.children[index] == nil {
			p.children[index] = &TrieNode{
				data:     str[i],
				children: make([]*TrieNode, length),
			}
		}
		p = p.children[index]
	}
	p.isEndChar = true
	return true
}

func (n *TrieNode) Find(str string) bool {
	result := true
	str = strings.ToLower(str)
	p := n
	for i := range str {
		index := getIndex(str[i])
		// 不在字符范围内
		if (index < 0 || index > 26) || p.children[index] == nil {
			result = false
			break
		}
		p = p.children[index]
	}
	// 只要完全匹配, 且节点为结束节点才返回true
	if result && p.isEndChar {
		return true
	}
	return false
}
