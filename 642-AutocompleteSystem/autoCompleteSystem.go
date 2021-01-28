package main

import "fmt"

type AutocompleteSystem struct {
	input []byte
	depository ComTrie
}

func AutoCompleteSystemConstructor(sentences []string, times []int) AutocompleteSystem {
	sys := AutocompleteSystem{}
	sys.depository = ComTrie{next: make(map[int32]*ComTrie)}
	for i, sentence := range sentences {
		sys.depository.Insert(sentence, times[i])
	}

	return sys
}

func (a *AutocompleteSystem) Input(c byte) []string {
	a.input = append(a.input, c)
	patterns, _ := a.depository.StartWith(string(a.input))
	return patterns
}


type ComTrie struct {
	next map[int32]*ComTrie
	isEnd bool
	frequency int
	pattern string
}

func (c *ComTrie) Insert(sentence string, time int) {
	node := c
	for i, c := range sentence {
		if node.next[c] == nil {
			node.next[c] = &ComTrie{next: make(map[int32]*ComTrie)}
			node.next[c].pattern = sentence[:i]
		}
		node = node.next[c]
		node.frequency += time
	}

	node.isEnd = true
}

func (c *ComTrie) StartWith(prefix string) ([]string, []int) {
	node := c
	patterns := make([]string, 0)
	frequencys := make([]int, 0)
	for _, c := range prefix {
		if node.next[c] == nil {
			return patterns,frequencys
		}
		node = node.next[c]
	}
	for _, value := range node.next {
		patterns, frequencys = c.SearchByFrequency(value, patterns, frequencys)
	}
	fmt.Printf("%#v\n%#v\n", patterns, frequencys)
	return patterns, frequencys
}

func (c *ComTrie) SearchByFrequency(node *ComTrie,patterns []string, frequencys []int) ([]string, []int) {
	for _, value := range node.next {
		if value.isEnd {
			if len(frequencys) == 0 {
				frequencys = append(frequencys, value.frequency)
				patterns = append(patterns, value.pattern)
			} else if len(frequencys) < 3 {
				for i, j := range frequencys {
					if value.frequency <= j {
						frequencys = append(append(frequencys[:i], value.frequency),frequencys[i:]...)
					}
				}
			} else {
				if value.frequency > frequencys[len(frequencys)-1] {
					frequencys = frequencys[1:]
					frequencys = append(frequencys, value.frequency)
					patterns = patterns[1:]
					patterns = append(patterns, value.pattern)
				} else if value.frequency <= frequencys[len(frequencys)-1] && value.frequency > frequencys[0] {
					for i, j := range frequencys {
						if value.frequency <= j {
							frequencys = frequencys[1:]
							frequencys = append(append(frequencys[:i], value.frequency), frequencys[i:]...)
							patterns = patterns[1:]
							patterns = append(append(patterns[:i], value.pattern), patterns[i:]...)
						}
					}
				}
			}
		}
		patterns, frequencys = c.SearchByFrequency(value, patterns, frequencys)
	}
	return patterns, frequencys
}

/**
 * Your AutocompleteSystem object will be instantiated and called as such:
 * obj := Constructor(sentences, times);
 * param_1 := obj.Input(c);
 */