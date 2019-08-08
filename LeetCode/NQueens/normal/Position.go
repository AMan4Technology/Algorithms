package normal

import (
	"strings"
)

func solveNQueens(n int) (results [][]string) {
	if n == 0 {
		return nil
	}
	var (
		col       = make([]bool, n)
		pie       = make([]bool, n-1+n-1+1)
		na        = make([]bool, n-1+n-1+1)
		positions = make([]int, n)
	)
	return validPositions(0, n, col, pie, na, positions, results)
}

func validPositions(i, n int, col, pie, na []bool, positions []int, results [][]string) [][]string {
	nextRow := i + 1
	for j := 0; j < n; j++ {
		if col[j] || pie[i+j] || na[i-j+n-1] {
			continue
		}
		positions[i] = j
		if nextRow == n {
			results = saveResult(results, positions, n)
			continue
		}
		col[j], pie[i+j], na[i-j+n-1] = true, true, true
		results = validPositions(nextRow, n, col, pie, na, positions, results)
		col[j], pie[i+j], na[i-j+n-1] = false, false, false
	}
	return results
}

func saveResult(results [][]string, positions []int, n int) [][]string {
	result := make([]string, n)
	var positionStr strings.Builder
	for i, position := range positions {
		for k := 0; k < position; k++ {
			positionStr.WriteByte('.')
		}
		positionStr.WriteByte('Q')
		for k := position + 1; k < n; k++ {
			positionStr.WriteByte('.')
		}
		result[i] = positionStr.String()
		positionStr.Reset()
	}
	return append(results, result)
}
