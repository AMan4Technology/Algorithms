package LongestPalindromicSubstring

func longestPalindrome(s string) (result string) {
	length := len(s)
	if length < 2 {
		return s
	}
	for head, tail, count, i := 0, 0, 0, 0; i < length; i++ {
		if i > 0 && s[i-1] == s[i] {
			continue
		}
		for head, tail = i, i+1; tail < length; tail++ {
			if s[head] != s[tail] {
				break
			}
		}
		for head > 0 && tail < length {
			if s[head-1] != s[tail] {
				break
			}
			head--
			tail++
		}
		if count < tail-head {
			count = tail - head
			result = s[head:tail]
		}
	}
	return
}
