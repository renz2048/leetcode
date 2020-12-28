package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode = nil
	next := prev
	for curr != nil {
		next = curr.Next
		curr.Next, prev, curr = prev, curr, next
	}
	return prev
}

func main() {
	l1 := new(ListNode)
	l1.Val = 1
	l2 := new(ListNode)
	l2.Val = 2
	l3 := new(ListNode)
	l3.Val = 3
	l4 := new(ListNode)
	l4.Val = 4
	l5 := new(ListNode)
	l5.Val = 5

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5

	pre := reverseList(l1)
	for pre != nil {
		fmt.Printf("%d", pre.Val)
		pre = pre.Next
	}
}
