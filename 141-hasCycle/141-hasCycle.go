package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head];ok {
			return true
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return false
}

func main() {
	var head = new(ListNode)
	head.Val = 3

	var node1 = new(ListNode)
	head.Next = node1
	node1.Val = 2

	var node2 = new(ListNode)
	node1.Next = node2
	node2.Val = 0

	var node3 = new(ListNode)
	node2.Next = node3
	node3.Val = -4
	node3.Next = nil

	isCycle := hasCycle(head)
	fmt.Println(isCycle)
}
