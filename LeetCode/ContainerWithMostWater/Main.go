package ContainerWithMostWater

import . "Algorithms/internal/box/math"

func maxArea(height []int) (max int) {
	length := len(height)
	if length < 2 {
		return
	}
	for start, end := 0, length-1; start < end; {
		max = MaxOfTwo(max, MinOfTwo(height[start], height[end])*(end-start))
		if height[start] > height[end] {
			for curr := height[end]; end > start; end-- {
				if height[end] > curr {
					break
				}
			}
		} else if height[start] < height[end] {
			for curr := height[start]; start < end; start++ {
				if curr < height[start] {
					break
				}
			}
		} else {
			for curr := height[start]; start < end; start++ {
				if curr < height[start] {
					break
				}
			}
			for curr := height[end]; end > start; end-- {
				if height[end] > curr {
					break
				}
			}
		}
	}
	return
}
