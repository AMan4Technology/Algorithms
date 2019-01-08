package chinesechess

func GetValidPosition(rowN, column int) (result [2][][]int) {
	for i := 0; i < column; i++ {
		var aPositions, bPositions []int
		for j := 0; j < rowN; j++ {
			for k := 0; k < column; k++ {
				position := j*column + k
				if k == i {
					aPositions = append(aPositions, position)
				} else {
					bPositions = append(bPositions, position)
				}
			}
		}
		// or
		// count := rowN * column
		// for position := 0; position < count; position++ {
		//     index := position % column
		//     if index == i {
		//         aPositions = append(aPositions, position)
		//     } else {
		//         bPositions = append(bPositions, position)
		//     }
		// }
		result[0] = append(result[0], aPositions)
		result[1] = append(result[1], bPositions)
	}
	return
}
