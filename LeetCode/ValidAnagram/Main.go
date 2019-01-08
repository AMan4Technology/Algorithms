package ValidAnagram

func isAnagram(s string, t string) bool {
	if len(s) == 0 && len(t) == 0 {
		return true
	}
	if len(s) != len(t) {
		return false
	}
	countWithC := make([]int, 26)
	for _, c := range s {
		countWithC[c-'a']++
	}
	for _, c := range t {
		count := countWithC[c-'a']
		if count == 0 {
			return false
		}
		countWithC[c-'a'] = count - 1
	}
	return true
}
