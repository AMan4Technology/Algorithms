package ReverseInteger

func reverse(x int) (result int) {
	for curr := x; curr != 0; curr /= 10 {
		result = result*10 + curr%10
		if result <= -1<<31 || result >= 1<<31-1 {
			return 0
		}
	}
	return
}
