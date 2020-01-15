package main

/**
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	ret = append(ret, root.Val)
	if root.Left != nil {
		l := preorderTraversal(root.Left)
		ret = append(ret, l...)
	}
	if root.Right != nil {
		r := preorderTraversal(root.Right)
		ret = append(ret, r...)
	}
	return ret
}

func preorderTraversalIter(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		size := len(stack)
		cur := stack[size-1]
		ret = append(ret, cur.Val)
		stack = stack[:size-1]
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
	return ret
}

func main() {

}
