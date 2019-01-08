package tree

func IsLeaf(node *TreeNode) bool {
	if node.Left == nil && node.Right == nil {
		return true
	}
	return false
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}
