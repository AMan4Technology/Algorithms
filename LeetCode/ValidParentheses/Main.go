package ValidParentheses

func isValid(s string) bool {
	if len(s) == 1 {
		return false
	}
	parenthesesMap := map[int32]int32{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]int32, 0, len(s))
	for _, c := range s {
		if leftC, ok := parenthesesMap[c]; ok {
			if len(stack) != 0 && stack[len(stack)-1] == leftC {
				stack = stack[:len(stack)-1]
				continue
			}
			return false
		}
		stack = append(stack, c)
	}
	return len(stack) == 0
}
