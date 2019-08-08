package trie

func newNode(length int) *node {
	return &node{next: make([]*node, length)}
}

type node struct {
	next   []*node
	isWord bool
}

func (n *node) words(path string, results []string) []string {
	if n.isWord {
		results = append(results, path)
	}
	for i, next := range n.next {
		if next != nil {
			results = next.words(path+string('a'+i), results)
		}
	}
	return results
}
