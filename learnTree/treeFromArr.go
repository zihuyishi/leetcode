package learnTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LoadTree(arr []int) *TreeNode {
	return loadTree(arr, 0)
}

func loadTree(arr []int, pos int) *TreeNode {
	if len(arr) == 0 || pos >= len(arr) {
		return nil
	}
	if arr[pos] < 0 {
		return nil
	}

	root := &TreeNode{
		Val: arr[pos],
	}
	root.Left = loadTree(arr, pos*2+1)
	root.Right = loadTree(arr, pos*2+2)

	return root
}
