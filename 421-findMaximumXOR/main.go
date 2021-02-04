package main

import (
	"fmt"
)

func main() {
	nums := []int{3,10,5,25,2,8}
	fmt.Printf("Hello, %d\n", FindMaximumXORMap(nums))

	maxXor := FindMaximumXORTrie(nums)
	fmt.Printf("maxXor: %d\n", maxXor)
}