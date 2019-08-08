package tree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func (n *TreeNode) IsLeaf() bool {
	if n.Left == nil && n.Right == nil {
		return true
	}
	return false
}
