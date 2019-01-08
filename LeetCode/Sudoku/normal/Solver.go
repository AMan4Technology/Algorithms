package normal

import "sort"

func solveSudoku(board [][]byte) {
	valid, rows, columns, blocks, emptyPositions := parse(board)
	if !valid {
		return
	}
	if emptyPositions.Len() == 0 {
		return
	}
	writeVal(board, emptyPositions, 0, emptyPositions.Len(), rows, columns, blocks)
	return
}

func parse(board [][]byte) (valid bool, rows, columns, blocks [][9]bool, emptyPositions positions) {
	rows = make([][9]bool, 9)
	columns = make([][9]bool, 9)
	blocks = make([][9]bool, 9)
	emptyPositions = NewPositions(9 * 9 / 2)
	for i, row := range board {
		for j, val := range row {
			if val == '.' {
				emptyPositions = append(emptyPositions, NewPosition(i, j))
				continue
			}
			hash := val - '1'
			if rows[i][hash] || columns[j][hash] || blocks[i/3*3+j/3][hash] {
				return false, nil, nil, nil, nil
			}
			rows[i][hash] = true
			columns[j][hash] = true
			blocks[i/3*3+j/3][hash] = true
		}
	}
	valid = true
	for _, position := range emptyPositions {
		for i := 0; i < 9; i++ {
			if rows[position.row][i] || columns[position.column][i] || blocks[position.block][i] {
				continue
			}
			position.add(byte('1' + i))
		}
	}
	sort.Sort(emptyPositions)
	return
}

func writeVal(board [][]byte, emptyPositions positions, i, length int, rows, columns, blocks [][9]bool) bool {
	next := i + 1
	position := emptyPositions[i]
	for _, val := range position.values {
		hash := val - '1'
		if rows[position.row][hash] || columns[position.column][hash] || blocks[position.block][hash] {
			continue
		}
		board[position.row][position.column] = val
		if next == length {
			return true
		}
		rows[position.row][hash], columns[position.column][hash], blocks[position.block][hash] = true, true, true
		if writeVal(board, emptyPositions, next, length, rows, columns, blocks) {
			return true
		}
		board[position.row][position.column] = '.'
		rows[position.row][hash], columns[position.column][hash], blocks[position.block][hash] = false, false, false
	}
	return false
}

func NewPositions(cap int) positions {
	return make(positions, 0, cap)
}

func NewPosition(row, column int) *position {
	return &position{
		row:    row,
		column: column,
		block:  row/3*3 + column/3,
		values: make([]byte, 0, 5),
	}
}

type positions []*position

func (p positions) Len() int {
	return len(p)
}

func (p positions) Less(i, j int) bool {
	return p[i].length < p[j].length
}

func (p positions) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type position struct {
	row, column, block int
	values             []byte
	length             int
}

func (p *position) add(val byte) {
	p.values = append(p.values, val)
	p.length++
}
