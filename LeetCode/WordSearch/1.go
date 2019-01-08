package WordSearch

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	var (
		rowNum    = len(board)
		columnNum = len(board[0])
		end       = len(word) - 1
	)
	for i, row := range board {
		for j, c := range row {
			if c == word[0] {
				board[i][j] = '#'
				if findNextC(board, word, 0, i, j, rowNum, columnNum, end) {
					return true
				}
				board[i][j] = c
			}
		}
	}
	return false
}

func findNextC(board [][]byte, word string, i, x, y, rowNum, columnNum, end int) bool {
	if i == end {
		return true
	}
	next := i + 1
	for k, dx := range dxs {
		nextX := x + dx
		if nextX == -1 || nextX == rowNum {
			continue
		}
		nextY := y + dys[k]
		if nextY == -1 || nextY == columnNum {
			continue
		}
		if board[nextX][nextY] == word[next] {
			board[nextX][nextY] = '#'
			if findNextC(board, word, next, nextX, nextY, rowNum, columnNum, end) {
				return true
			}
			board[nextX][nextY] = word[next]
		}
	}
	return false
}
