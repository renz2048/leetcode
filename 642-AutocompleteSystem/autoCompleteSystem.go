package trie

import "fmt"

type AutocompleteSystem struct {
	input []byte
	depository AutoComTrie
}

func AutoCompleteSystemConstructor(sentences []string, times []int) AutocompleteSystem {
	sys := AutocompleteSystem{}
	sys.depository = AutoComTrie{next: make(map[int32]*AutoComTrie)}
	for i, sentence := range sentences {
		sys.depository.Insert(sentence, times[i])
	}
	//sys.depository.BianliTrie(&sys.depository)

	return sys
}

func (a *AutocompleteSystem) Input(c byte) []string {
	a.input = append(a.input, c)
	patterns, _ := a.depository.StartWith(string(a.input))
	return patterns
}


type AutoComTrie struct {
	next map[int32]*AutoComTrie
	isEnd bool
	frequency int
	pattern string
}

func (a *AutoComTrie) Insert(sentence string, time int) {
	node := a
	for i, c := range sentence {
		if node.next[c] == nil {
			node.next[c] = &AutoComTrie{next: make(map[int32]*AutoComTrie)}
			node.next[c].pattern = sentence[:i+1]
		}
		node = node.next[c]
		node.frequency += time
	}

	node.isEnd = true
}

func (a *AutoComTrie) Search(word string) bool {
	node := a
	for _, c := range word {
		if node.next[c] == nil {
			return false
		}
		node = node.next[c]
	}
	if node.isEnd {
		return true
	}
	return false
}

func (a *AutoComTrie) StartWith(prefix string) ([]string, []int) {
	fmt.Printf("in StartWith\n")
	node := a
	patterns := make([]string, 0)
	frequencys := make([]int, 0)
	for _, c := range prefix {
		if node.next[c] == nil {
			return patterns,frequencys
		}
		node = node.next[c]
	}
	for _, value := range node.next {
		fmt.Printf("- : %#v\n", value)
		patterns, frequencys = a.SearchByFrequency(value, patterns, frequencys)
	}
	//fmt.Printf("%#v\n%#v\n", patterns, frequencys)
	return patterns, frequencys
}

func (a *AutoComTrie) SearchByFrequency(node *AutoComTrie,patterns []string, frequencys []int) ([]string, []int) {
	fmt.Printf("in SearchByFrequency\n")
	for _, value := range node.next {
		fmt.Printf("-- : value=%#v\n", value)
		if value.isEnd {
			fmt.Printf("-- : %#v\t%#v\n", patterns, frequencys)
			patterns, frequencys =a.FindPattern(value.pattern, value.frequency, patterns, frequencys)
			fmt.Printf("-- : %#v\t%#v\n", patterns, frequencys)
		}
		patterns, frequencys = a.SearchByFrequency(value, patterns, frequencys)
	}
	fmt.Printf("end SearchByFrequency\n")
	return patterns, frequencys
}

func (a *AutoComTrie) FindPattern(pattern string, frequency int, patterns []string, frequencys []int) ([]string,[]int) {
	if len(frequencys) == 0 {
		frequencys = append(frequencys, frequency)
		patterns = append(patterns, pattern)
	} else if len(frequencys) < 3 {
		if frequency > frequencys[0] {
			frequencys = append([]int{frequency}, frequencys...)
			patterns = append([]string{pattern}, patterns...)
		} else if frequency < frequencys[len(frequencys)-1] {
			frequencys = append(frequencys, frequency)
			patterns = append(patterns, pattern)
		} else {
			for i, j := range frequencys {
				if frequency > j {
					fmt.Printf("---- : slice: %#v\tword: %#v\n", patterns, pattern)
					frequencys = append(frequencys[:i], append([]int{frequency}, frequencys[i:]...)...)
					patterns = append(patterns[:i], append([]string{pattern}, patterns[i:]...)...)
					fmt.Printf("---- : slice: %#v\n", patterns)
					break
				}
			}
		}
	} else {
		if frequency >= frequencys[0] {
			frequencys = append([]int{frequency}, frequencys...)
			patterns = append([]string{pattern}, patterns...)
			frequencys = frequencys[:len(frequencys)-1]
			patterns = patterns[:len(frequencys)-1]
		} else {
			for i, j := range frequencys {
				if frequency > j {
					frequencys = append(frequencys[:i], append([]int{frequency}, frequencys[i:]...)...)
					frequencys = frequencys[:len(frequencys)-1]
					patterns = append(patterns[:i], append([]string{pattern}, patterns[i:]...)...)
					patterns = patterns[:len(frequencys)-1]
					break
				}
			}
		}
	}
	return patterns, frequencys
}

func (a *AutoComTrie) BianliTrie(node *AutoComTrie) {
	if len(node.next) == 0 {
		fmt.Printf("%#v\n", node)
	}
	for _, value := range node.next {
		fmt.Printf("%#v\n", value)
		a.BianliTrie(value)
	}
}

/**
 * Your AutocompleteSystem object will be instantiated and called as such:
 * obj := Constructor(sentences, times);
 * param_1 := obj.Input(c);
 */