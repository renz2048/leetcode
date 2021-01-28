package main

import (
	"fmt"
)

func main() {
	sentences := []string{"i love you", "island", "ironman", "i love leetcode"}
	times := []int{5,3,2,2}
	auto := AutoCompleteSystemConstructor(sentences, times)
	patterns := auto.Input('i')
	fmt.Printf("%#v\n", patterns)
}
