package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isLeaves(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	} else if root.Left == nil {
		if isLeaves(root.Right) {
			return 0
		} else {
			return sumOfLeftLeaves(root.Right)
		}
	} else if root.Right == nil {
		if isLeaves(root.Left) {
			return root.Left.Val
		} else {
			return sumOfLeftLeaves(root.Left)
		}
	} else {
		left := 0
		if isLeaves(root.Left) {
			left = root.Left.Val
		} else {
			left = sumOfLeftLeaves(root.Left)
		}
		right := 0
		if !isLeaves(root.Right) {
			right = sumOfLeftLeaves(root.Right)
		}
		return left + right
	}
}

func makeTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: arr[0],
	}
	arr = arr[1:]
	if len(arr) > 0 {
		root.Left = makeTree(arr[:len(arr)/2])
		root.Right = makeTree(arr[len(arr)/2:])
	}
	return root
}

func printTree(node *TreeNode) {
	if node == nil {
		fmt.Print("nil ")
		return
	}
	fmt.Printf("%d ", node.Val)
	printTree(node.Left)
	printTree(node.Right)
}

func main() {
	arr := []int{3, 4, -1, 2, 2, 0, 4}
	root := makeTree(arr)
	printTree(root)
	fmt.Print(sumOfLeftLeaves(root))
}