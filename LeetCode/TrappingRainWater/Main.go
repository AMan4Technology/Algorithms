package TrappingRainWater

func trap(heights []int) (result int) {
	length := len(heights)
	if length < 3 {
		return 0
	}
	var (
		maxHeightsIndices = make([]int, 0, length/3)
		maxIndex          int
	)
	if heights[0] > heights[1] {
		maxHeightsIndices = append(maxHeightsIndices, 0)
	}
	heights = append(heights, -1)
	for i := 1; i < length; i++ {
		if heights[i-1] <= heights[i] && heights[i] >= heights[i+1] {
			for j := len(maxHeightsIndices) - 1; j > maxIndex; j-- {
				if heights[maxHeightsIndices[j]] <= heights[i] {
					maxHeightsIndices = maxHeightsIndices[:j]
				}
			}
			maxHeightsIndices = append(maxHeightsIndices, i)
			if heights[maxHeightsIndices[maxIndex]] < heights[i] {
				maxIndex++
			}
		}
	}
	for i, length := 1, len(maxHeightsIndices); i < length; i++ {
		var low int
		if heights[maxHeightsIndices[i-1]] < heights[maxHeightsIndices[i]] {
			low = heights[maxHeightsIndices[i-1]]
		} else {
			low = heights[maxHeightsIndices[i]]
		}
		for j := maxHeightsIndices[i-1] + 1; j < maxHeightsIndices[i]; j++ {
			if low > heights[j] {
				result += low - heights[j]
			}
		}
	}
	return
}
