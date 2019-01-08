package Triangle

func minimumTotal(triangle [][]int) int {
	length := len(triangle)
	if length == 0 {
		return 0
	}
	distances := make([]int, length)
	copy(distances, triangle[length-1])
	for i := length - 2; i >= 0; i-- {
		for j, length := 0, len(triangle[i]); j < length; j++ {
			if distances[j] > distances[j+1] {
				distances[j] = distances[j+1]
			}
			distances[j] += triangle[i][j]
		}
	}
	return distances[0]
}
