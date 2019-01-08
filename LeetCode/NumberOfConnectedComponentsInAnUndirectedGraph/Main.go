package NumberOfConnectedComponentsInAnUndirectedGraph

func countComponents(n int, edges [][]int) (count int) {
	if n < 2 {
		return n
	}
	length := len(edges)
	if length < 2 {
		return n - length
	}
	count = n
	circles := make([]int, n+1)
	for _, pair := range edges {
		a, b := pair[0]+1, pair[1]+1
		if circles[a] == 0 {
			circles[a] = a
		}
		if circles[b] == 0 {
			circles[b] = b
		}
		for circles[a] != a {
			a = circles[a]
		}
		for circles[b] != b {
			b = circles[b]
		}
		if a != b {
			circles[b] = a
			b = a
			count--
		}
		circles[pair[0]+1] = a
		circles[pair[1]+1] = b
	}
	return
}
