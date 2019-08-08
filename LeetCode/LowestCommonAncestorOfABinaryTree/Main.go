package LowestCommonAncestorOfABinaryTree

import (
	"Algorithms"
)

func lowestCommonAncestor(root, p, q *Algorithms.TreeNode) *Algorithms.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	leftRoot := lowestCommonAncestor(root.Left, p, q)
	rightRoot := lowestCommonAncestor(root.Right, p, q)
	if leftRoot != nil && rightRoot != nil {
		return root
	} else if leftRoot != nil {
		return leftRoot
	} else {
		return rightRoot
	}
}
