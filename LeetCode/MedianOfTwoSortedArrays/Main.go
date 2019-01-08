package MedianOfTwoSortedArrays

import . "Algorithms/LeetCode/internal/box/math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		length1 = len(nums1)
		length2 = len(nums2)
		length  = length1 + length2
	)
	if length == 0 {
		return 0
	}
	if length1 < length2 || (length1 == length2 && nums1[length1-1] <= nums2[0]) {
		length1, length2 = length2, length1
		nums1, nums2 = nums2, nums1
	}
	var (
		end        = length / 2
		prev, curr int
	)
	if length2 == 0 {
		curr = nums1[length1/2]
		if length1&1 == 0 {
			prev = nums1[length1/2-1]
		}
		goto result
	}
	if nums2[length2-1] <= nums1[0] {
		goto compute
	}
	for i, j, k := 0, 0, 0; k <= end; k++ {
		prev = curr
		if j == length2 {
			goto compute
		}
		if nums1[i] <= nums2[j] {
			curr = nums1[i]
			i++
		} else {
			curr = nums2[j]
			j++
		}
	}
	goto result
compute:
	curr = nums1[end-length2]
	if end-length2 == 0 {
		prev = nums2[length2-1]
	} else {
		prev = MaxOfTwo(nums1[end-length2-1], nums2[length2-1])
	}
result:
	if length&1 != 0 {
		return float64(curr)
	}
	return float64(prev+curr) / 2
}
