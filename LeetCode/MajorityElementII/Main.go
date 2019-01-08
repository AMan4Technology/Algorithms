package MajorityElementII

func majorityElement(nums []int) (results []int) {
	length := len(nums)
	if length == 0 {
		return nil
	}
	var (
		result1, result2 int
		count1, count2   int
	)
	for _, val := range nums {
		if val == result1 {
			count1++
		} else if val == result2 {
			count2++
		} else if count1 == 0 {
			result1, count1 = val, 1
		} else if count2 == 0 {
			result2, count2 = val, 1
		} else {
			count1--
			count2--
		}
	}
	count1, count2 = 0, 0
	countLine := length / 3
	for _, val := range nums {
		if val == result1 {
			count1++
		} else if val == result2 {
			count2++
		}
	}
	if count1 > countLine {
		results = append(results, result1)
	}
	if count2 > countLine {
		results = append(results, result2)
	}
	return
}
