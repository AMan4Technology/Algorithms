package SymmetricTree

import (
	"Algorithms"
)

func isSymmetric(root *Algorithms.TreeNode) bool {
	if root == nil {
		return true
	}
	queue := make([]*Algorithms.TreeNode, 0, 4)
	queue = append(queue, root.Left, root.Right)
	for left, right := root, root; len(queue) != 0; queue = queue[2:] {
		left, right = queue[0], queue[1]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Right, left.Right, right.Left)
	}
	return true
}
