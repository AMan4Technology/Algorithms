package binary

func isValidSudoku(board [][]byte) bool {
	if board == nil || len(board) == 0 {
		return false
	}
	var (
		rows    = make([]uint16, 9)
		columns = make([]uint16, 9)
		block   = make([]uint16, 9)
	)
	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}
			val := uint16(1 << (c - '1'))
			if rows[i]&val != 0 || columns[j]&val != 0 || block[i/3*3+j/3]&val != 0 {
				return false
			}
			rows[i] |= val
			columns[j] |= val
			block[i/3*3+j/3] |= val
		}
	}
	return true
}
