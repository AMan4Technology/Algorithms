package SearchInRotatedSortedArray

func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	var (
		start = 0
		end   = length - 1
		split = nums[0]
		left  = target >= split
	)
	for pivot := start + (end-start)/2; start <= end; pivot = start + (end-start)/2 {
		if nums[pivot] == target {
			return pivot
		}
		if nums[pivot] < target {
			if left {
				if nums[pivot] >= split {
					start = pivot + 1
				} else {
					end = pivot - 1
				}
			} else {
				start = pivot + 1
			}
		} else {
			if left {
				end = pivot - 1
			} else {
				if nums[pivot] < split {
					end = pivot - 1
				} else {
					start = pivot + 1
				}
			}
		}
	}
	return -1
}
