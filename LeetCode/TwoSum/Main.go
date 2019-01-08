package TwoSum

func twoSum(nums []int, target int) []int {
	indexWithNum := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := indexWithNum[target-num]; ok {
			return []int{j, i}
		}
		indexWithNum[num] = i
	}
	return nil
}
