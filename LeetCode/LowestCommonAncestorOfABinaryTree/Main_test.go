package LowestCommonAncestorOfABinaryTree

import (
	"fmt"
	"testing"

	"Algorithms/LeetCode/internal/tree"
)

func TestLowestCommonAncestor(t *testing.T) {
	p := &tree.TreeNode{Val: 0}
	q := &tree.TreeNode{Val: 4}
	root := &tree.TreeNode{Val: 2, Left: p, Right: q}
	fmt.Println(lowestCommonAncestor(root, q, p).Val)
}
