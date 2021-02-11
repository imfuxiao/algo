package trie

type Trie struct {
	isWorld bool
	child map[byte]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{child: make(map[byte]*Trie)}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	node := this
	child := this.child
	for _, char := range word {
		if c, exists := child[byte(char)]; exists {
			node = c
			child = c.child
		} else {
			node = &Trie{
				isWorld: false,
				child:   make(map[byte]*Trie),
			}
			child[byte(char)] = node
			child = node.child
		}
	}
	node.isWorld = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	isWorld := false
	child := this.child
	for _, char := range word {
		if c, exists := child[byte(char)]; !exists {
			isWorld = false
			break
		} else {
			isWorld = c.isWorld
			child = c.child
		}
	}
	return isWorld
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	child := this.child
	for _, char := range prefix {
		if c, exists := child[byte(char)]; !exists {
			return false
		} else {
			child = c.child
		}
	}
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */