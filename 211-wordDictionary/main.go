package main

import "fmt"

func main() {
	/*
	obj := Constructor()
	obj.AddWord("bad")
	obj.AddWord("dad")
	obj.AddWord("mad")
	fmt.Printf("search pad: %t\n", obj.Search("pad"))
	fmt.Printf("search bad: %t\n", obj.Search("bad"))
	fmt.Printf("search .ad: %t\n", obj.Search(".ad"))
	fmt.Printf("search b..: %t\n", obj.Search("b.."))
	 */
	/*
	obj1 := Constructor()
	obj1.AddWord("a")
	obj1.AddWord("a")
	fmt.Printf("search .: %t\n", obj1.Search("."))
	fmt.Printf("search a: %t\n", obj1.Search("a"))
	fmt.Printf("search aa: %t\n", obj1.Search("aa"))
	fmt.Printf("search .a: %t\n", obj1.Search(".a"))
	fmt.Printf("search a.: %t\n", obj1.Search("a."))
	 */
	obj1 := Constructor()
	obj1.AddWord("at")
	obj1.AddWord("and")
	obj1.AddWord("an")
	obj1.AddWord("add")
	fmt.Printf("search a: %t\n", obj1.Search("a"))
	fmt.Printf("search .at: %t\n", obj1.Search(".at"))
	obj1.AddWord("bat")
	fmt.Printf("search .at: %t\n", obj1.Search(".at"))
	fmt.Printf("search an.: %t\n", obj1.Search("an."))
	fmt.Printf("search a.d.: %t\n", obj1.Search("a.d."))
	fmt.Printf("search b.: %t\n", obj1.Search("b."))
	fmt.Printf("search a.d: %t\n", obj1.Search("a.d"))
	fmt.Printf("search .: %t\n", obj1.Search("."))
}
