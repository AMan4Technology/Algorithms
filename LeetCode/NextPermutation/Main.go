package NextPermutation

func nextPermutation(nums []int) {
	var (
		length = len(nums)
		start  = -1
	)
	for i := length - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			start = i - 1
			break
		}
	}
	if start != -1 {
		for i := length - 1; i > start; i-- {
			if nums[start] < nums[i] {
				nums[start], nums[i] = nums[i], nums[start]
				break
			}
		}
	}
	for i := start + 1; i < (length-start-1)/2+start+1; i++ {
		nums[i], nums[length-1-(i-start-1)] = nums[length-1-(i-start-1)], nums[i]
	}
}
