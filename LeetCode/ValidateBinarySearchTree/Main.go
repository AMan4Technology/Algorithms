package ValidateBinarySearchTree

import (
	"Algorithms"
)

func isValidBST(root *Algorithms.TreeNode) bool {
	if root == nil {
		return true
	}
	_, _, valid := validate(root)
	return valid
}

func validate(root *Algorithms.TreeNode) (min, max int, valid bool) {
	min, max = root.Val, root.Val
	if root.Left != nil {
		leftMin, leftMax, leftValid := validate(root.Left)
		if !leftValid || leftMax >= root.Val {
			return 0, 0, false
		}
		min = leftMin
	}
	if root.Right != nil {
		rightMin, rightMax, rightValid := validate(root.Right)
		if !rightValid || root.Val >= rightMin {
			return 0, 0, false
		}
		max = rightMax
	}
	return min, max, true
}
