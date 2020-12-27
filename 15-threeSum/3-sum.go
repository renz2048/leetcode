package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first - 1] {
			continue
		}
		third := n - 1
		target := -1 * nums[first]

		for second := first + 1; second < n; second++ {
			if second > first + 1 && nums[second] == nums[second - 1] {
				continue
			}
			for second < third && nums[second] + nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

func main() {
	a := []int{-1, 0, 1, 2, -1, -4}
	ans := threeSum(a)
	for i := 0; i < len(ans); i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d, ", ans[i][j])
		}
		fmt.Printf("\n")
	}
}
