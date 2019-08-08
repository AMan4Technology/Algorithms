package GameOfLife

var (
	dx = []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy = []int{-1, 0, 1, -1, 1, -1, 0, 1}
)

func gameOfLife(board [][]int) {
	m := len(board)
	if m == 0 {
		return
	}
	var (
		n      = len(board[0])
		points = make([][]point, m)
	)
	for i := range board {
		points[i] = make([]point, n)
	}
	for i, lives := range board {
		for j, live := range lives {
			points[i][j].x, points[i][j].y = i, j
			if live == 0 {
				continue
			}
			for k, x := range dx {
				var (
					row    = i + x
					column = j + dy[k]
				)
				if row < 0 || row >= m {
					continue
				}
				if column < 0 || column >= n {
					continue
				}
				points[row][column].count++
			}
		}
	}
	queue := make([]point, 0)
	for i, row := range points {
		for j, point := range row {
			switch {
			case point.count <= 1 || point.count > 3:
				if board[i][j] == 1 {
					board[i][j] = 0
					point.live = false
					queue = append(queue, point)
				}
			case point.count == 3:
				if board[i][j] == 0 {
					board[i][j] = 1
					point.live = true
					queue = append(queue, point)
				}
			}
		}
	}
	for len(queue) != 0 {
		point := queue[0]
		queue = queue[1:]
		for k, x := range dx {
			var (
				row    = point.x + x
				column = point.y + dy[k]
			)
			if row < 0 || row >= m {
				continue
			}
			if column < 0 || column >= n {
				continue
			}
			if point.live {
				points[row][column].count++
			} else {
				points[row][column].count--
			}
		}
	}
	return
}

type point struct {
	x, y  int
	live  bool
	count int
}
