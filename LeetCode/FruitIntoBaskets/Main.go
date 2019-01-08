package FruitIntoBaskets

import . "Algorithms/LeetCode/internal/box/math"

func totalFruit(tree []int) int {
	return totalFruitWithN(tree, 2)
}

func totalFruitWithN(tree []int, num int) (count int) {
	if num == 0 {
		return 0
	}
	length := len(tree)
	if length <= num {
		return length
	}
	numWithVal := make(map[int]int, num)
	for head, n, i := 0, num, 0; i < length; i++ {
		if _, ok := numWithVal[tree[i]]; ok {
			goto update
		}
		if n > 0 {
			n--
			goto update
		}
		for ; head < i; head++ {
			numWithVal[tree[head]]--
			if numWithVal[tree[head]] == 0 {
				delete(numWithVal, tree[head])
				head++
				break
			}
		}
	update:
		numWithVal[tree[i]]++
		count = MaxOfTwo(count, i+1-head)
	}
	return
}
