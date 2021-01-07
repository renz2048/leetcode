package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	if l % 2 == 1 {
		return float64(getKthElement(nums1, nums2, l/2 + 1))
	} else {
		return float64(getKthElement(nums1, nums2, l/2) + getKthElement(nums1, nums2, l/2 + 1))/2
	}
}

func getKthElement(nums1 []int, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2 + k - 1]
		}
		if index2 == len(nums2) {
			return nums1[index1 + k - 1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		mid := k/2
		newIndex1 := min(index1 + mid, len(nums1)) - 1
		newIndex2 := min(index2 + mid, len(nums2)) - 1
		if nums1[newIndex1] <= nums2[newIndex2] {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
	
	return 0
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	nums1 := [...]int{1,3,5}
	nums2 := [...]int{1,2,3,4,5,6,7}
	a := findMedianSortedArrays(nums1[:], nums2[:])
	fmt.Printf("%v\n", a)
}
