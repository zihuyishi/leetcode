package main

import "github.com/zihuyishi/leetcode/learnTree"

func postorderTraversal(root *learnTree.TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}
	ret := postorderTraversal(root.Left)
	ret = append(ret, postorderTraversal(root.Right)...)
	ret = append(ret, root.Val)
	return ret
}

func main() {
	arr := []int{1, 2, -1, 3}
	node := learnTree.LoadTree(arr)
	ret := postorderTraversal(node)
	for _, v := range ret {
		println(v)
	}
}
