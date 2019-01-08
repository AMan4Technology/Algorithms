package normal

func totalNQueens(n int) (count int) {
	if n == 0 {
		return 0
	}
	validPositionsNum(0, uint(n), 0, 0, 0, &count)
	return
}

func validPositionsNum(row, n, col, pie, na uint, count *int) {
	next := row + 1
	for validB := (^(col | pie | na)) & (1<<n - 1); validB > 0; validB &= validB - 1 {
		positionB := validB & -validB
		if next == n {
			*count++
			continue
		}
		validPositionsNum(next, n, col|positionB, (pie|positionB)<<1, (na|positionB)>>1, count)
	}
}
