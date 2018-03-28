package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	root := &TreeNode{
		Val: t1.Val + t2.Val,
	}

	if t1.Left != nil && t2.Left != nil {
		root.Left = mergeTrees(t1.Left, t2.Left)
	} else if t1.Left != nil {
		root.Left = t1.Left
	} else if t2.Left != nil {
		root.Left = t2.Left
	}

	if t1.Right != nil && t2.Right != nil {
		root.Right = mergeTrees(t1.Right, t2.Right)
	} else if t1.Right != nil {
		root.Right = t1.Right
	} else if t2.Right != nil {
		root.Right = t2.Right
	}

	return root
}
