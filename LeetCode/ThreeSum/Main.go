package ThreeSum

import "sort"

func threeSum(nums []int) (results [][]int) {
	length := len(nums)
	if length < 3 {
		return
	}
	sort.Ints(nums)
	for i := 0; i < length-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		j, k := i+1, length-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if sum > 0 {
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else {
				results = append(results, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return
}
