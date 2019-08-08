package EditDistance

import (
	"Algorithms"
)

func minDistance(word1 string, word2 string) int {
	var (
		length1 = len(word1)
		length2 = len(word2)
	)
	if length1 == 0 {
		return length2
	}
	if length2 == 0 {
		return length1
	}
	distances := make([][]int, length1+1)
	for i := range distances {
		distances[i] = make([]int, length2+1)
		distances[i][0] = i
	}
	for j := range distances[0] {
		distances[0][j] = j
	}
	for i := 1; i <= length1; i++ {
		for j := 1; j <= length2; j++ {
			if word1[i-1] == word2[j-1] {
				distances[i][j] = distances[i-1][j-1]
			} else {
				distances[i][j] = 1 + Algorithms.Min(distances[i-1][j-1], distances[i][j-1], distances[i-1][j])
			}
		}
	}
	return distances[length1][length2]
}
