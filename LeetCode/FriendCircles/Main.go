package FriendCircles

func findCircleNum(M [][]int) (count int) {
	num := len(M)
	if num < 2 {
		return num
	}
	count = num
	circles := make([]int, num)
	for i := range M {
		circles[i] = i
		for j := 0; j < i; j++ {
			if M[i][j] == 0 {
				continue
			}
			if circles[i] == i {
				circles[i] = circles[j]
				count--
			} else {
				bossI, bossJ := i, j
				for circles[bossI] != bossI {
					bossI = circles[bossI]
				}
				for circles[bossJ] != bossJ {
					bossJ = circles[bossJ]
				}
				if bossI != bossJ {
					circles[bossJ] = bossI
					count--
					bossJ = bossI
				}
				circles[i] = bossI
				circles[j] = bossJ
			}
		}
	}
	return
}
