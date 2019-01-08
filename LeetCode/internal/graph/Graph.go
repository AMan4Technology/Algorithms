package graph

import (
	"strconv"
	"strings"
)

func NewGraph(num int) *Graph {
	return &Graph{
		num:       num,
		relations: make([][]int, num),
	}
}

type Graph struct {
	num       int
	relations [][]int
}

func (g Graph) AddEdge(start, end int) {
	g.relations[start] = append(g.relations[start], end)
}

func (g Graph) ToPoSort() string {
	inDegrees := make([]int, g.num)
	for _, relation := range g.relations {
		for _, end := range relation {
			inDegrees[end]++
		}
	}
	var (
		queue  = make([]int, 0, 2)
		result = strings.Builder{}
	)
	for i, inDegree := range inDegrees {
		if inDegree == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) != 0 {
		i := queue[0]
		queue = queue[1:]
		result.WriteString(strconv.Itoa(i))
		result.WriteString(" -> ")
		for _, end := range g.relations[i] {
			inDegrees[end]--
			if inDegrees[end] == 0 {
				queue = append(queue, end)
			}
		}
	}
	return result.String()[:result.Len()-4]
}
