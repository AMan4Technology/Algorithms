package LongestSubstringWithoutRepeatingCharacters

import . "Algorithms/LeetCode/internal/box/math"

func lengthOfLongestSubstring(s string) (count int) {
	length := len(s)
	if length < 2 {
		return length
	}
	var (
		head, tail int
		nums       = make([]int, 256)
	)
	for _, c := range s {
		tail++
		nums[c]++
		if nums[c] > 1 {
			for ; head < tail; head++ {
				nums[s[head]]--
				if nums[s[head]] == 1 {
					head++
					break
				}
			}
		}
		count = MaxOfTwo(count, tail-head)
	}
	return
}
