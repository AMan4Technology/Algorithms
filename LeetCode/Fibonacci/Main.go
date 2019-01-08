package Fibonacci

func fibonacciDP(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 1 {
		return 1
	}
	data := make([]int, n+1)
	data[0], data[1] = 1, 1
	for i := 2; i <= n; i++ {
		data[i] = data[i-1] + data[i-2]
	}
	return data[n]
}

func fibonacci(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 1 {
		return n
	}
	return multiple(pow([][]int{
		{1, 1},
		{1, 0},
	}, n-1), [][]int{{1}, {1}})[0][0]
}

func multiple(X, Y [][]int) (result [][]int) {
	if len(X[0]) != len(Y) {
		return nil
	}
	var (
		rowNum    = len(X)
		columnNum = len(Y[0])
	)
	result = make([][]int, rowNum)
	for i, row := range X {
		result[i] = make([]int, columnNum)
		for j := 0; j < columnNum; j++ {
			for k, value := range row {
				result[i][j] += value * Y[k][j]
			}
		}
	}
	return
}

func pow(K [][]int, n int) (result [][]int) {
	result = [][]int{{1, 0}, {0, 1}}
	for n != 0 {
		if n&1 == 1 {
			result = multiple(result, K)
		}
		n >>= 1
		K = multiple(K, K)
	}
	return
}
