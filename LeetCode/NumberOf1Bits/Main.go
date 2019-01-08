package NumberOf1Bits

func hammingWeight(num uint32) (count int) {
	for num != 0 {
		num &= num - 1
		count++
	}
	return
}
