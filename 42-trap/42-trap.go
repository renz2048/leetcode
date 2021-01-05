package main

import (
	"fmt"
)

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func trap(height []int) int {
	l := len(height)
	ans := 0
	for i := 1;i < l - 1; i++ {
		max_left := 0
		for j := i; j >= 0; j-- {
			max_left = max(max_left, height[j])
		}
		max_right := 0
		for j := i; j < l; j++ {
			max_right = max(max_right, height[j])
		}
		ans += min(max_left, max_right) - height[i]
	}
	return ans;
}

func main() {
	height := [...]int{0,1,0,2,1,0,1,3,2,1,2,1}
	a := trap(height[:])
	fmt.Printf("%d\n", a)
}
