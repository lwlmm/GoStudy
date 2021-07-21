/*			1
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
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//前序遍历 ：先输出root结点，再输出左子树，再输出右子树
func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println("Node val: ", root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

// 前序遍历的栈迭代法，入栈顺序：右， 左； 出栈顺序：中、左、右
func ItePreorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println("Node val: ", node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
}

// 中序遍历的迭代法：向栈中压入根节点最深的左子结点
// 每次回退一格，如果有右节点则将cur变为右节点继续压入
func IteInorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	cur := root
	stack := []*TreeNode{}
	for len(stack) != 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println("Node val: ", node.Val)
			if node.Right != nil {
				cur = node.Right
			}
	}
}

// 后序遍历的迭代：后序遍历是先遍历左子树，再右子树，最后根节点
// 利用前序遍历的栈思想：先在一个栈中压入中、左、右节点，
// 再顺次压入第二个栈中并提出
func ItePostorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	stack1 := []*TreeNode{root}
	stack2 := []*TreeNode{}
	for len(stack1) != 0 {
		node := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = append(stack2, node)
		if node.Left != nil {
			stack1 = append(stack1, node.Left)
		}
		if node.Right != nil {
			stack1 = append(stack1, node.Right)
		}
	}
	for len(stack2) != 0 {
		fmt.Println("Node val: ", stack2[len(stack2)-1].Val)
		stack2 = stack2[:len(stack2)-1]
	}
}

func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left)
	fmt.Println("Node val: ", root.Val)
	inorderTraversal(root.Right)
}

func postorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	postorderTraversal(root.Left)
	postorderTraversal(root.Right)
	fmt.Println("NOde val: ", root.Val)
}

//通过dfs深度搜索，返回某个结点开始的从上至下所有value切片
func dfsResult(root *TreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}

//递归法前序遍历结点，用append扩展slice
func dfs(root *TreeNode, result *[]int) {
	if root == nil { //递归返回条件
		return
	}
	*result = append(*result, root.Val) //加入本节点
	dfs(root.Left, result)
	dfs(root.Right, result)
}

//分治法dfs深度搜索
func dfsDivideResult(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	//分治法：先左递归，再右递归。加上本节点、所有左节点、所有右节点
	left := dfsDivideResult(root.Left)
	right := dfsDivideResult(root.Right)

	result = append(result, root.Val)
	result = append(result, left...) //不确定长度
	result = append(result, right...)

	return result
}

//BFS广度遍历
func levelOrder(root *TreeNode) [][]int {
	// 通过上一层的长度确定下一层的元素
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root) //len(queue) = 1
	for len(queue) > 0 {
		list := make([]int, 0)
		// 为什么要取length？
		// 记录当前层有多少元素（遍历当前层，再添加下一层）
		l := len(queue)
		for i := 0; i < l; i++ {

			// 出队列
			level := queue[0]
			queue = queue[1:]
			list = append(list, level.Val)

			//进队列
			if level.Left != nil {
				queue = append(queue, level.Left)
			}
			if level.Right != nil {
				queue = append(queue, level.Right)
			}
		}
		result = append(result, list)
	}
	return result
}

func main() {

	//构建一个二叉树
	root := &TreeNode{
		Val: 1,
	}

	left1 := &TreeNode{
		Val: 2,
	}
	left2 := &TreeNode{
		Val: 4,
	}
	left3 := &TreeNode{
		Val: 5,
	}

	right1 := &TreeNode{
		Val: 3,
	}
	right2 := &TreeNode{
		Val: 6,
	}
	right3 := &TreeNode{
		Val: 7,
	}

	root.Left = left1
	root.Right = right1
	left1.Left = left2
	left1.Right = left3
	right1.Left = right2
	right1.Right = right3

	/*
		preorderTraversal(root)
		fmt.Println()
		inorderTraversal(root)
		fmt.Println()
		postorderTraversal(root)
		fmt.Println()


		fmt.Println(dfsResult(right1))
		fmt.Println(dfsDivideResult(root))

		fmt.Println(levelOrder(root))

		queue := make([]*TreeNode, 0)
		queue = append(queue, root)
		fmt.Println(len(queue))
	*/
	ItePreorderTraversal(root)

	fmt.Println()

	IteInorderTraversal(root)

	fmt.Println()

	ItePostorderTraversal(root)
}
