/*															1
														/		\
												       /         \
													  2           3
													 / \          / \
													/   \        /   \
												   4     5      6     7
*/
package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//前序遍历 ：先输出root结点，再输出左子树，再输出右子树
func preorderTraversal(root *TreeNode){
	if root == nil {
		return
	}
	fmt.Println("Node val: ", root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

func inorderTraversal(root *TreeNode){
	if root == nil {
		return
	}
	inorderTraversal(root.Left)
	fmt.Println("Node val: ", root.Val)
	inorderTraversal(root.Right)
}

func postorderTraversal(root *TreeNode){
	if root == nil {
		return
	}
	postorderTraversal(root.Left)
	postorderTraversal(root.Right)
	fmt.Println("NOde val: ", root.Val)
}

func main() {

	//构建一个二叉树
	root := &TreeNode{
		Val : 1,
	}

	left1 := &TreeNode{
		Val : 2,
	}
	left2 := &TreeNode{
		Val : 4,
	}
	left3 := &TreeNode{
		Val : 5,
	}

	right1 := &TreeNode{
		Val : 3,
	}
	right2 := &TreeNode{
		Val : 6,
	}
	right3 := &TreeNode{
		Val : 7,
	}

	root.Left = left1
	root.Right = right1
	left1.Left = left2
	left2.Right = left3
	right1.Left = right2
	right1.Right = right3

	preorderTraversal(root)
	fmt.Println()
	inorderTraversal(root)
	fmt.Println()
	postorderTraversal(root)
	fmt.Println()
}