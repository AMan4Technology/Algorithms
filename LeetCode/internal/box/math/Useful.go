package math

func MaxAndMin(base int, values ...int) (max, min int) {
	max, min = base, base
	for i, length := 0, len(values); i < length; i++ {
		max = MaxOfTwo(max, values[i])
		min = MinOfTwo(min, values[i])
	}
	return
}

func Max(base int, values ...int) (max int) {
	max = base
	for i, length := 0, len(values); i < length; i++ {
		max = MaxOfTwo(max, values[i])
	}
	return
}

func Min(base int, values ...int) (min int) {
	min = base
	for i, length := 0, len(values); i < length; i++ {
		min = MinOfTwo(min, values[i])
	}
	return
}

func MaxOfTwo(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinOfTwo(a, b int) int {
	if a < b {
		return a
	}
	return b
}
