package main

import(
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	head := &ListNode {
		Val : 1,
	}
	tmp := head
	for i := 2; i <= 15; i++ {
		t := &ListNode {
			Val : i,
		}
		tmp.Next = t
		tmp = tmp.Next
	}
	tmp = reverseKGroup(head, 4)
	for i := 0; i < 15; i++ {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyNode := &ListNode {
		Next : head,
	}
	stableNode := dummyNode
	for {
		probeNode := stableNode
		for i := 0; i < k; i++ {
			probeNode = probeNode.Next
			if probeNode == nil {
				return dummyNode.Next
			}
		}
		back := stableNode.Next
		fore := back.Next
		for i := 1; i < k; i++ {
			back.Next = fore.Next
			fore.Next = stableNode.Next
			stableNode.Next = fore
			fore = back.Next
		}
		for i := 0; i < k; i++ {
			stableNode = stableNode.Next
		}
	}
	return dummyNode.Next
}