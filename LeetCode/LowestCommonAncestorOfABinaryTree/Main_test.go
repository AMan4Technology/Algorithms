package LowestCommonAncestorOfABinaryTree

import (
	"fmt"
	"testing"

	"Algorithms"
)

func TestLowestCommonAncestor(t *testing.T) {
	p := &Algorithms.TreeNode{Val: 0}
	q := &Algorithms.TreeNode{Val: 4}
	root := &Algorithms.TreeNode{Val: 2, Left: p, Right: q}
	fmt.Println(lowestCommonAncestor(root, q, p).Val)
}
