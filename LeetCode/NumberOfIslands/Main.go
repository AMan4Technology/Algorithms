package NumberOfIslands

func numIslands(grid [][]byte) (count int) {
	rowNum := len(grid)
	if rowNum == 0 {
		return 0
	}
	columnNum := len(grid[0])
	if columnNum == 0 {
		return 0
	}
	var (
		ids     = make([][]int, rowNum)
		islands = make([]int, 1, 5)
	)
	for i := range ids {
		ids[i] = make([]int, columnNum)
	}
	for i, row := range grid {
		for j, value := range row {
			if value == '0' {
				continue
			}
			if j > 0 && grid[i][j-1] != '0' {
				ids[i][j] = ids[i][j-1]
			}
			if i > 0 && grid[i-1][j] != '0' {
				if ids[i][j] == 0 {
					ids[i][j] = ids[i-1][j]
				} else if ids[i-1][j] != ids[i][j] {
					for islands[ids[i][j]] != ids[i][j] {
						ids[i][j] = islands[ids[i][j]]
					}
					for islands[ids[i-1][j]] != ids[i-1][j] {
						ids[i-1][j] = islands[ids[i-1][j]]
					}
					if ids[i-1][j] != ids[i][j] {
						islands[ids[i][j]] = ids[i-1][j]
						count--
					}
				}
			}
			if ids[i][j] == 0 {
				id := len(islands)
				islands = append(islands, id)
				ids[i][j] = id
				count++
			}
		}
	}
	return
}
