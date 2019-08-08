package SerializeAndDeserializeBinaryTree

import (
	"fmt"
	"testing"

	"Algorithms"
)

func TestSerialize(t *testing.T) {
	root := &Algorithms.TreeNode{Val: 1}
	root.Left = &Algorithms.TreeNode{Val: 2}
	root.Right = &Algorithms.TreeNode{Val: 3}
	root.Right.Left = &Algorithms.TreeNode{Val: 4}
	root.Right.Right = &Algorithms.TreeNode{Val: 5}
	fmt.Println(SerializeWithBFS(root))
	fmt.Println(SerializeWithDFS(root))
}

func TestDeserialize(t *testing.T) {
	root := &Algorithms.TreeNode{Val: 1}
	root.Left = &Algorithms.TreeNode{Val: 2}
	root.Right = &Algorithms.TreeNode{Val: 3}
	root.Right.Left = &Algorithms.TreeNode{Val: 4}
	root.Right.Right = &Algorithms.TreeNode{Val: 5}
	newRoot := DeserializeWithBFS(SerializeWithBFS(root))
	fmt.Println(newRoot.Val)
	fmt.Println(SerializeWithBFS(newRoot))
}
