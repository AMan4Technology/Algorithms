package normal

func totalNQueens(n int) int {
	if n == 0 {
		return 0
	}
	var (
		col = make([]bool, n)
		pie = make([]bool, n-1+n-1+1)
		na  = make([]bool, n-1+n-1+1)
	)
	return validPositionsNum(0, n, 0, col, pie, na)
}

func validPositionsNum(i, n, num int, col, pie, na []bool) int {
	nextRow := i + 1
	for j := 0; j < n; j++ {
		if col[j] || pie[i+j] || na[i-j+n-1] {
			continue
		}
		if nextRow == n {
			num++
			continue
		}
		col[j], pie[i+j], na[i-j+n-1] = true, true, true
		num = validPositionsNum(nextRow, n, num, col, pie, na)
		col[j], pie[i+j], na[i-j+n-1] = false, false, false
	}
	return num
}
