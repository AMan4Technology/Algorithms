package Sqrt

import "math"

func intSqrt(x int) (result int) {
	if x == 0 || x == 1 {
		return x
	}
	var (
		left  = 1
		right = x
		mid   int
	)
	for left <= right {
		mid = left + (right-left)/2
		other := x / mid
		if other < mid {
			right = mid - 1
		} else if mid < other {
			left = mid + 1
			result = mid
		} else {
			return mid
		}
	}
	return
}

func sqrt(y, precision float64) float64 {
	var (
		left, mid float64
		right     = math.Max(1, y)
	)
	for left <= right {
		mid = left + (right-left)/2
		if right-left < precision {
			return mid
		}
		other := y / mid
		if other < mid {
			right = mid
		} else if mid < other {
			left = mid
		} else {
			return mid
		}
	}
	return 0
}
