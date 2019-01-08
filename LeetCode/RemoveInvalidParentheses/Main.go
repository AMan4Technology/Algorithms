package RemoveInvalidParentheses

func removeInvalidParentheses(s string) (results []string) {
	length := len(s)
	if length == 0 {
		return []string{""}
	}
	var (
		positionsOfLeft  []int
		stack            = make([]int, 0, 8)
		positionsOfRight = []int{-1}
		rights           []int
	)
	for i := 0; i < length; i++ {
		switch s[i] {
		case '(':
			positionsOfLeft = append(positionsOfLeft, i)
			stack = append(stack, len(positionsOfLeft)-1)
		case ')':
			positionsOfRight = append(positionsOfRight, i)
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				rights = append(rights, len(positionsOfRight)-1)
			}
		}
	}
	var (
		resultsOfLeft []string
		head          int
		tail          = length
	)
	if length := len(rights); length > 0 {
		dFSRight(0, 0, length, positionsOfRight, rights, s, "", &results)
		head = positionsOfRight[rights[length-1]] + 1
	}
	if lengthOfLeft := len(stack); lengthOfLeft > 0 {
		positionsOfLeft = append(positionsOfLeft, length)
		dFSLeft(len(positionsOfLeft)-1, lengthOfLeft-1, positionsOfLeft, stack, s, "", &resultsOfLeft)
		tail = positionsOfLeft[stack[0]]
	}
	lenOfResults := len(results)
	if lenOfResults == 0 {
		results = append(results, "")
		lenOfResults++
	}
	for i := 0; i < lenOfResults; i++ {
		results[i] += s[head:tail]
		for _, result := range resultsOfLeft {
			results = append(results, results[i]+result)
		}
	}
	if len(results) == lenOfResults {
		return
	}
	return results[lenOfResults:]
}

func dFSRight(head, i, length int, positions, rights []int, s, result string, results *[]string) {
	if i == length {
		*results = append(*results, result+s[positions[head]+1:positions[rights[i-1]]+1])
		return
	}
	dFSRight(head+1, i+1, length, positions, rights, s, result+s[positions[head]+1:positions[head+1]], results)
	for j := head + 2; j <= rights[i]; j++ {
		if positions[j]-positions[j-1] > 1 {
			dFSRight(j, i+1, length, positions, rights, s, result+s[positions[head]+1:positions[j]], results)
		}
	}
}

func dFSLeft(tail, i int, positions, lefts []int, s, result string, results *[]string) {
	if i == -1 {
		*results = append(*results, s[positions[lefts[0]]:positions[tail]]+result)
		return
	}
	dFSLeft(tail-1, i-1, positions, lefts, s, s[positions[tail-1]+1:positions[tail]]+result, results)
	for j := tail - 2; j >= lefts[i]; j-- {
		if positions[j+1]-positions[j] > 1 {
			dFSLeft(j, i-1, positions, lefts, s, s[positions[j]+1:positions[tail]]+result, results)
		}
	}
}
