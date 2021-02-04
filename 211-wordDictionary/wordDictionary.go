package main

type WordDictionary struct {
	next map[int32]*WordDictionary
	isEnd bool
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{next: make(map[int32]*WordDictionary)}
}


func (this *WordDictionary) AddWord(word string)  {
	node := this
	for _, c := range word {
		if node.next[c] == nil {
			node.next[c] = &WordDictionary{next: make(map[int32]*WordDictionary)}
		}
		node = node.next[c]
	}
	node.isEnd = true
}


func (this *WordDictionary) Search(word string) bool {
	node := this
	for i, c := range word {
		if c == '.' {
			if len(node.next) == 0 {
				return false
			}
			for _, childnode := range node.next {
				//fmt.Printf("childnode: %#v\nword: %s\ni: %d\n", childnode, word[i+1:], i)
				if childnode.Search(word[i+1:]) {
					return true
				}
			}
			return false
		} else {
			if node.next[c] == nil {
				return false
			}
			node = node.next[c]
		}
	}
	if node.isEnd {
		return true
	}
	return false
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */