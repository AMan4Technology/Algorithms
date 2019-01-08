package RemoveKDigits

func removeKdigits(num string, k int) string {
	length := len(num)
	if k >= length {
		return "0"
	}
	var result = make([]byte, 1, length)
	result[0] = num[0]
	for i := 1; i < length; i++ {
		for j := len(result) - 1; k > 0; j-- {
			if j == -1 {
				if num[i] != '0' {
					result = append(result, num[i])
				}
				break
			}
			if result[j] <= num[i] {
				result = append(result, num[i])
				break
			}
			result = result[:j]
			k--
		}
		if k == 0 {
			result = append(result, num[i:]...)
			break
		}
	}
	var validI int
	for length := len(result); validI < length; validI++ {
		if result[validI] != '0' {
			break
		}
	}
	result = result[validI:]
	if len(result) <= k {
		return "0"
	}
	result = result[:len(result)-k]
	return string(result)
}
