package binary

import "sort"

func solveSudoku(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
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

func parse(board [][]byte) (valid bool, rows, columns, blocks []uint16, emptyPositions positions) {
	rows = make([]uint16, 9)
	columns = make([]uint16, 9)
	blocks = make([]uint16, 9)
	emptyPositions = NewPositions(9 * 9 / 2)
	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				emptyPositions.Add(i, j, 0)
				continue
			}
			value := uint16(1 << (c - '1'))
			if (rows[i]|columns[j]|blocks[i/3*3+j/3])&value != 0 {
				return false, nil, nil, nil, nil
			}
			rows[i] |= value
			columns[j] |= value
			blocks[i/3*3+j/3] |= value
		}
	}
	valid = true
	for _, position := range emptyPositions {
		position.update(^(rows[position.row] | columns[position.column] | blocks[position.block]) & (1<<9 - 1))
	}
	sort.Sort(emptyPositions)
	return
}

func writeVal(board [][]byte, emptyPositions positions, i, length int, rows, columns, blocks []uint16) bool {
	next := i + 1
	position := emptyPositions[i]
	for values := position.values; values != 0; values &= values - 1 {
		var (
			value uint16
			hash  int
		)
		for value = 1; values&value == 0; value = value << 1 {
			hash++
		}
		if (rows[position.row]|columns[position.column]|blocks[position.block])&value != 0 {
			continue
		}
		board[position.row][position.column] = byte('1' + hash)
		if next == length {
			return true
		}
		rows[position.row] |= value
		columns[position.column] |= value
		blocks[position.block] |= value
		if writeVal(board, emptyPositions, next, length, rows, columns, blocks) {
			return true
		}
		rows[position.row] &= ^value
		columns[position.column] &= ^value
		blocks[position.block] &= ^value
	}
	return false
}

func NewPositions(cap int) positions {
	return make(positions, 0, cap)
}

func NewPosition(row, column int, values uint16) (p *position) {
	p = &position{
		row:    row,
		column: column,
		block:  row/3*3 + column/3,
	}
	if values != 0 {
		p.update(values)
	}
	return p
}

type positions []*position

func (p *positions) Add(row, column int, values uint16) {
	*p = append(*p, NewPosition(row, column, values))
}

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
	values             uint16
	length             int
}

func (p *position) update(values uint16) {
	p.values = values
	for p.length = 0; values != 0; values &= values - 1 {
		p.length++
	}
}
