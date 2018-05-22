package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type Solution struct {
	ret map[int]int
	depth int
}

func (s *Solution) travel(node *TreeNode, depth int) {
	if node == nil {
		return
	}
	_, ok := s.ret[depth]
	if !ok {
		s.ret[depth] = node.Val
	}
	if depth > s.depth {
		s.depth = depth
	}
	if node.Right != nil {
		s.travel(node.Right, depth+1)
	}
	if node.Left != nil {
		s.travel(node.Left, depth+1)
	}
}

func rightSideView(root *TreeNode) []int {
	s := &Solution{
		ret: make(map[int]int),
		depth: -1,
	}
	s.travel(root, 0)
	ret := make([]int, s.depth+1)
	for i := 0; i < s.depth+1; i++ {
		ret[i] = s.ret[i]
	}
	return ret
}