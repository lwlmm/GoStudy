package main

import(
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	n15 := &TreeNode {
		Val : 15,
	}
	n7 := &TreeNode {
		Val : 7,
	}
	n9 := &TreeNode {
		Val : 9,
		Left : n15,
		Right : n7,
	}
	n20 := &TreeNode {
		Val : 20,
	}
	n3 := &TreeNode {
		Val : 3,
		Left : n9,
		Right : n20,
	}
	zigzag := zigzagLevelOrder(n3)
	fmt.Println(zigzag)
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	flag := true
	stack := []*TreeNode {}
	ret := [][]int {}
	if root == nil {
		return ret
	}
	stack = append(stack, root)
	for len(stack) != 0 {
		tmpStack := []*TreeNode {}
		tmpValue := []int {}
		for i := 0; i < len(stack); i++ {
			tmpValue = append(tmpValue, stack[i].Val)
		}
		ret = append(ret, tmpValue)
		if flag {
			flag = false
			for i := len(stack)-1; i >= 0; i-- {
				if stack[i].Right != nil {
					tmpStack = append(tmpStack, stack[i].Right)
				}
				if stack[i].Left != nil {
					tmpStack = append(tmpStack, stack[i].Left)
				}
			}
			stack = append([]*TreeNode(nil), tmpStack...)
		} else {
			flag = true
			for i := len(stack)-1; i >= 0; i-- {
				if stack[i].Left != nil {
					tmpStack = append(tmpStack, stack[i].Left)
				}
				if stack[i].Right != nil {
					tmpStack = append(tmpStack, stack[i].Right)
				}
			}
			stack = append([]*TreeNode(nil), tmpStack...)
		}
	}
	return ret
}