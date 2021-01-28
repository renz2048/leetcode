package main

import "fmt"

func main() {
	trie := Constructor()
	trie.Insert("apple")
	b := trie.Search("apple")
	fmt.Println(b)

	b = trie.Search("app")
	fmt.Println(b)
	b = trie.StartsWith("app")
	fmt.Println(b)

	trie.Insert("app")
	b = trie.Search("app")
	fmt.Println(b)
}
