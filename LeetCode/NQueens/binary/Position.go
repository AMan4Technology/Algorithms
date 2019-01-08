package normal

import (
	"strings"
)

func solveNQueens(n int) (results [][]string) {
	if n == 0 {
		return nil
	}
	positions := make([]uint, n)
	validPositions(0, uint(n), 0, 0, 0, positions, &results)
	return
}

func validPositions(row, n, col, pie, na uint, positions []uint, results *[][]string) {
	next := row + 1
	for validB := ^(col | pie | na) & (1<<n - 1); validB > 0; validB &= validB - 1 {
		positionB := validB & -validB
		positions[row] = positionB
		if next == n {
			saveResult(results, positions, n)
			continue
		}
		validPositions(next, n, col|positionB, (pie|positionB)<<1, (na|positionB)>>1, positions, results)
	}
}

func saveResult(results *[][]string, positions []uint, n uint) {
	result := make([]string, n)
	positionStr := strings.Builder{}
	for i, position := range positions {
		for k := uint(1 << (n - 1)); k != 0; k = k >> 1 {
			if position&k == 0 {
				positionStr.WriteByte('.')
				continue
			}
			positionStr.WriteByte('Q')
		}
		result[i] = positionStr.String()
		positionStr.Reset()
	}
	*results = append(*results, result)
	return
}
