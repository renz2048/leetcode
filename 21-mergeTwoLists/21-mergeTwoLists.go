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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	lm := new(ListNode)
	l3 := new(ListNode)
	l3 = lm
	for l1 != nil || l2 != nil {
		if l1 == nil {
			lm.Val = l2.Val
			l2 = l2.Next
		} else if l2 == nil {
			lm.Val = l1.Val
			l1 = l1.Next
		} else if l1.Val <= l2.Val {
			lm.Val = l1.Val
			l1 = l1.Next
		} else {
			lm.Val = l2.Val
			l2 = l2.Next
		}
		if l1 == nil && l2 == nil {
			break;
		}
		lm.Next = new(ListNode)
		lm = lm.Next
	}
	return l3
}

func main() {
	l1a := ListNode{Val: 1}
	l1b := ListNode{Val: 2}
	l1c := ListNode{Val: 4}
	l1a.Next = &l1b
	l1b.Next = &l1c

	l2a := ListNode{Val: 1}
	l2b := ListNode{Val: 3}
	l2c := ListNode{Val: 4}
	l2a.Next = &l2b
	l2b.Next = &l2c

	lm := mergeTwoLists(&l1a, &l2a)
	for lm != nil {
		fmt.Printf("%d ", lm.Val)
		lm = lm.Next
	}
}
