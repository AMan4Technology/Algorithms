package MinimumWindowSubstring

func minWindow(s string, t string) string {
	var (
		lenOfS = len(s)
		lenOfT = len(t)
	)
	if lenOfS == 0 || lenOfT == 0 || lenOfS < lenOfT {
		return ""
	}
	var (
		min, max = -1, lenOfS
		count    = lenOfT
		numWithC = make([]int, 256)
	)
	for _, c := range t {
		numWithC[c]++
	}
	for head, tail, i := 0, 0, 0; i < lenOfS; i++ {
		tail++
		numWithC[s[i]]--
		if numWithC[s[i]] >= 0 {
			count--
		}
		for ; head < tail; head++ {
			if numWithC[s[head]] >= 0 {
				break
			}
			numWithC[s[head]]++
		}
		if count == 0 && tail-head < max-min {
			min, max = head, tail
		}
	}
	if min == -1 {
		return ""
	}
	return s[min:max]
}
