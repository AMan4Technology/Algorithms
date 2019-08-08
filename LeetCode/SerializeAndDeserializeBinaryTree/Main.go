package SerializeAndDeserializeBinaryTree

import (
	"strconv"
	"strings"

	"Algorithms"
)

const Nil = "Nil"

func SerializeWithBFS(root *Algorithms.TreeNode) string {
	if root == nil {
		return "[]"
	}
	var (
		data  strings.Builder
		queue = []*Algorithms.TreeNode{root}
	)
	data.WriteByte('[')
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			data.WriteString(Nil + ",")
			continue
		}
		data.WriteString(strconv.Itoa(node.Val) + ",")
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}
	return data.String()[:data.Len()-1] + "]"
}

func DeserializeWithBFS(data string) (root *Algorithms.TreeNode) {
	if data == "[]" || data == "" {
		return nil
	}
	var (
		length     = len(data)
		values     = strings.Split(data[1:length-1], ",")
		rootVal, _ = strconv.Atoi(values[0])
		queue      = []*Algorithms.TreeNode{{Val: rootVal}}
	)
	root = queue[0]
	for i := 1; len(queue) != 0; i += 2 {
		node := queue[0]
		queue = queue[1:]
		if values[i] != Nil {
			val, _ := strconv.Atoi(values[i])
			node.Left = &Algorithms.TreeNode{Val: val}
			queue = append(queue, node.Left)
		}
		if values[i+1] != Nil {
			val, _ := strconv.Atoi(values[i+1])
			node.Right = &Algorithms.TreeNode{Val: val}
			queue = append(queue, node.Right)
		}
	}
	return
}

func SerializeWithDFS(root *Algorithms.TreeNode) string {
	if root == nil {
		return "[]"
	}
	var data strings.Builder
	data.WriteByte('[')
	dfsOfSerialize(&data, root)
	return data.String()[:data.Len()-1] + "]"
}

func DeserializeWithDFS(data string) (root *Algorithms.TreeNode) {
	if data == "[]" || data == "" {
		return nil
	}
	var (
		length = len(data)
		values = strings.Split(data[1:length-1], ",")
		i      int
	)
	root = new(Algorithms.TreeNode)
	dfsOfDeserialize(root, &i, length-2, values)
	return
}

func dfsOfSerialize(data *strings.Builder, node *Algorithms.TreeNode) {
	if node == nil {
		data.WriteString(Nil + ",")
		return
	}
	data.WriteString(strconv.Itoa(node.Val) + ",")
	dfsOfSerialize(data, node.Left)
	dfsOfSerialize(data, node.Right)
}

func dfsOfDeserialize(node *Algorithms.TreeNode, i *int, length int, values []string) {
	value, _ := strconv.Atoi(values[*i])
	node.Val = value
	if *i == length-1 {
		return
	}
	if *i++; values[*i] != Nil {
		node.Left = new(Algorithms.TreeNode)
		dfsOfDeserialize(node.Left, i, length, values)
	}
	if *i++; values[*i] != Nil {
		node.Right = new(Algorithms.TreeNode)
		dfsOfDeserialize(node.Right, i, length, values)
	}
}
