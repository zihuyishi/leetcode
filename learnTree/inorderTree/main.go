package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}
	ret := inorderTraversal(root.Left)
	ret = append(ret, root.Val)
	ret = append(ret, inorderTraversal(root.Right)...)
	return ret
}

func inorderTraversal_iter(root *TreeNode) []int {
	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		size := len(stack)
		cur := stack[size-1]
		if cur == nil {
			continue
		}
		stack = append(stack, cur.Right)
		if cur.Left == nil {
			ret = append(ret, cur.Val)
		} else {
			stack = append(stack, cur.Left)
		}
	}
	return ret
}
