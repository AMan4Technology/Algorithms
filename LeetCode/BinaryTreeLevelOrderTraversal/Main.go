package BinaryTreeLevelOrderTraversal

import (
	"Algorithms"
)

func levelOrderByBFS(root *Algorithms.TreeNode) (results [][]int) {
	if root == nil {
		return nil
	}
	queue := make([]*Algorithms.TreeNode, 0, 7)
	results = append(results, []int{root.Val})
	if !Algorithms.IsLeaf(root) {
		queue = append(queue, root)
	}
	for len(queue) != 0 {
		var levelValues []int
		length := len(queue)
		for i := 0; i < length; i++ {
			parentNode := queue[i]
			if parentNode.Left != nil {
				levelValues = append(levelValues, parentNode.Left.Val)
				if !Algorithms.IsLeaf(parentNode.Left) {
					queue = append(queue, parentNode.Left)
				}
			}
			if parentNode.Right != nil {
				levelValues = append(levelValues, parentNode.Right.Val)
				if !Algorithms.IsLeaf(parentNode.Right) {
					queue = append(queue, parentNode.Right)
				}
			}
		}
		results = append(results, levelValues)
		queue = queue[length:]
	}
	return
}

func levelOrderByDFS(root *Algorithms.TreeNode) (results [][]int) {
	if root == nil {
		return nil
	}
	var level int
	results = accessNode(root, level, results)
	return
}

func accessNode(node *Algorithms.TreeNode, level int, results [][]int) [][]int {
	if len(results) < level+1 {
		results = append(results, []int{})
	}
	results[level] = append(results[level], node.Val)
	if node.Left != nil {
		results = accessNode(node.Left, level+1, results)
	}
	if node.Right != nil {
		results = accessNode(node.Right, level+1, results)
	}
	return results
}
