package JewelsAndStones

func numJewelsInStones(J string, S string) (count int) {
	js := make(map[rune]bool, len(J))
	for _, c := range []rune(J) {
		js[c] = true
	}
	for _, c := range []rune(S) {
		if js[c] {
			count++
		}
	}
	return
}
