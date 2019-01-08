package CountingBits

func countBits(num int) (counts []int) {
	counts = make([]int, num+1)
	for i := 1; i <= num; i++ {
		counts[i] = counts[i&(i-1)] + 1
	}
	return
}
