package graph

import (
	"fmt"
	"testing"

	"Algorithms"
)

func TestGraph_ToPoSort(t *testing.T) {
	graph := Algorithms.NewGraph(5)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(2, 1)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)
	fmt.Println(graph.ToPoSort())
}
