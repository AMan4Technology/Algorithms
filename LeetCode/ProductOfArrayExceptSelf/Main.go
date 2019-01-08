package ProductOfArrayExceptSelf

func productExceptSelf(nums []int) (result []int) {
	length := len(nums)
	result = make([]int, length)
	for left, i := 1, 0; i < length; i++ {
		result[i] = left
		left *= nums[i]
	}
	for right, i := 1, length-1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
	}
	return
}
