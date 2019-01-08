package normal

func isValidSudoku(board [][]byte) bool {
	var (
		rows    = make([][9]bool, 9)
		columns = make([][9]bool, 9)
		blocks  = make([][9]bool, 9)
	)
	for i, row := range board {
		for j, val := range row {
			if val == '.' {
				continue
			}
			hash := val - '1'
			if rows[i][hash] || columns[j][hash] || blocks[i/3*3+j/3][hash] {
				return false
			}
			rows[i][hash] = true
			columns[j][hash] = true
			blocks[i/3*3+j/3][hash] = true
		}
	}
	return true
}
