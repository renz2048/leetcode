package main

type Trie struct {
	isEnd bool
	next map[int32]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{next: make(map[int32]*Trie)}
}


/** Inserts a word into the trie. */
func (t *Trie) Insert(word string)  {
	node := t
	for _, c := range word {
		if node.next[c] == nil {
			//插入时如果某个映射不存在，创建映射
			node.next[c] = new(Trie)
			node.next[c].next = make(map[int32]*Trie)
		}
		node = node.next[c]
	}
	node.isEnd = true
}


/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	node := t
	for _, c := range word{
		if node.next[c] == nil {
			return false
		}
		node = node.next[c]
	}
	return node.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for _, c := range prefix {
		if node.next[c] == nil {
			return false
		}
		node = node.next[c]
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