package MaxMinDepthOfBinaryTree

import (
	"Algorithms"
	. "Algorithms/internal/tree"
)

func maxDepth(root *Algorithms.TreeNode) (max int) {
	max, _ = maxMinDepth(root)
	return max
}

func maxMinDepth(root *Algorithms.TreeNode) (max, min int) {
	if root == nil {
		return 0, 0
	}
	queue := make([]*TreeNode, 0, 7)
	if Algorithms.IsLeaf(root) {
		max, min = 1, 1
	} else {
		queue = append(queue, root)
	}
	depth := 2
	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			parentNode := queue[i]
			if parentNode.Left != nil {
				if Algorithms.IsLeaf(parentNode.Left) {
					if min == 0 {
						min = depth
					}
					if depth > max {
						max = depth
					}
				} else {
					queue = append(queue, parentNode.Left)
				}
			}
			if parentNode.Right != nil {
				if Algorithms.IsLeaf(parentNode.Right) {
					if min == 0 {
						min = depth
					}
					if depth > max {
						max = depth
					}
				} else {
					queue = append(queue, parentNode.Right)
				}
			}
		}
		depth++
		queue = queue[length:]
	}
	return
}

func minDepth(root *Algorithms.TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*Algorithms.TreeNode, 0, 7)
	if Algorithms.IsLeaf(root) {
		return 1
	} else {
		queue = append(queue, root)
	}
	depth := 2
	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			parentNode := queue[i]
			if parentNode.Left != nil {
				if Algorithms.IsLeaf(parentNode.Left) {
					return depth
				} else {
					queue = append(queue, parentNode.Left)
				}
			}
			if parentNode.Right != nil {
				if Algorithms.IsLeaf(parentNode.Right) {
					return depth
				} else {
					queue = append(queue, parentNode.Right)
				}
			}
		}
		depth++
		queue = queue[length:]
	}
	return 0
}
